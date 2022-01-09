package admin

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/sakerhetsm/ssm-wargame/internal/db"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
	"go.uber.org/zap"

	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
)

func (s *service) ListChallenges(ctx context.Context, req *spec.ListChallengesPayload) (spec.SsmChallengeCollection, error) {

	challs, err := db.New(s.db).ListChallengesWithSolves(ctx, true)
	if err != nil {
		s.log.Warn("could not list challs", zap.Error(err), utils.C(ctx))
		return nil, err
	}

	res := make(spec.SsmChallengeCollection, len(challs))

	for i, chall := range challs {
		res[i] = &spec.SsmChallenge{
			ID:          chall.ID.String(),
			Slug:        chall.Slug,
			Title:       chall.Title,
			Description: chall.Description,
			Score:       chall.Score,
			Solves:      chall.NumSolves,
			Published:   chall.Published,
		}

	}

	return res, nil
}

func (s *service) CreateChallenge(ctx context.Context, req *spec.CreateChallengePayload) error {

	q := db.New(s.db)

	err := q.InsertChallenge(ctx, db.InsertChallengeParams{
		ID:          uuid.New(),
		Title:       req.Title,
		Slug:        req.Slug,
		Description: req.Description,
		Score:       req.Score,
		Published:   false,
	})
	if err != nil {
		s.log.Error("inserting challenge", zap.Error(err), utils.C(ctx))
		return err
	}

	return nil
}

func (s *service) PresignChallFileUpload(ctx context.Context, req *spec.PresignChallFileUploadPayload) (*spec.PresignChallFileUploadResult, error) {

	log := s.log.With(zap.String("challengeID", req.ChallengeID), utils.C(ctx))

	challID := uuid.MustParse(req.ChallengeID) // uuid format already validated

	exists, err := db.New(s.db).ChallengeExists(ctx, challID)
	if err != nil {
		log.Error("could not exec ChallengeExists")
		return nil, err
	}

	if !exists {
		return nil, spec.MakeNotFound(errors.New("could not find challenge"))
	}

	log.Info("signing url")

	objectKey := req.ChallengeID + "/" + req.Filename

	r, _ := s.s3.PutObjectRequest(&s3.PutObjectInput{
		Bucket:     &s.cfg.S3.Bucket,
		Key:        &objectKey,
		ContentMD5: &req.Md5,
	})

	url, err := r.Presign(time.Second * 20)
	if err != nil {
		log.Error("could not sign url", zap.Error(err))
		return nil, err
	}

	fileID := uuid.New()

	err = db.New(s.db).InsertFile(ctx, db.InsertFileParams{
		ID:          fileID,
		ChallengeID: challID,
		Fname:       req.Filename,
		Bucket:      s.cfg.S3.Bucket,
		Key:         objectKey,
		Md5:         req.Md5,
	})

	if err != nil {
		log.Error("could not insert file", zap.Error(err))
		return nil, err
	}

	go func() {
		defer func() {
			err := recover()
			if err != nil {
				s.log.Error("objectWait paniced", zap.Any("err", err))
			}
		}()
		s.objectWait(context.Background(), fileID, s.cfg.S3.Bucket, objectKey)
	}()

	return &spec.PresignChallFileUploadResult{
		URL: url,
	}, nil
}

func (s *service) objectWait(ctx context.Context, fileID uuid.UUID, bucket, key string) {

	log := s.log.With(zap.String("fileID", fileID.String()))

	err := s.s3.WaitUntilObjectExistsWithContext(ctx, &s3.HeadObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}, request.WithWaiterDelay(request.ConstantWaiterDelay(time.Second*5)), request.WithWaiterMaxAttempts(5))

	if err == nil {
		err = db.New(s.db).FileMarkUploaded(ctx, fileID)
		if err != nil {
			log.Error("could not mark file as uploaded", zap.Error(err))
		}
		return
	}

	if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "ResourceNotReady" {
		log.Info("file was not uploaded", zap.String("fileID", fileID.String()))

		err = db.New(s.db).DeleteFile(ctx, fileID)
		if err != nil {
			log.Error("could not delete non-uploaded file", zap.Error(err))
		}

		return
	}

	log.Warn("could not head object", zap.Error(err), zap.String("fileID", fileID.String()))

}
