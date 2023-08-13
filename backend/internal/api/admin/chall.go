package admin

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/sakerhetsm/ssm-wargame/internal/custommodels"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"

	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
)

func (s *service) ListChallenges(ctx context.Context, req *spec.ListChallengesPayload) (spec.SsmAdminChallengeCollection, error) {

	challs := make([]*custommodels.UserChall, 0)
	err := models.NewQuery(
		qm.Select("c.*"),
		qm.Select("(SELECT COUNT(1) FROM user_solves WHERE challenge_id = c.id) num_solves"),
		qm.From("challenges c"),
		qm.Load(models.ChallengeRels.Flags),
		qm.Load(models.ChallengeRels.ChallengeFiles),
		qm.Load(models.ChallengeRels.Authors),
	).Bind(ctx, s.db, &challs)
	if err != nil {
		s.log.Warn("could not list challs", zap.Error(err), utils.C(ctx))
		return nil, err
	}

	res := make(spec.SsmAdminChallengeCollection, len(challs))

	for i, chall := range challs {
		res[i] = &spec.SsmAdminChallenge{
			ID:          chall.ID,
			Slug:        chall.Slug,
			Title:       chall.Title,
			Description: chall.Description,
			StaticScore: chall.StaticScore.Ptr(),
			Solves:      chall.NumSolves,
			PublishAt:   utils.NullTimeToUnix(chall.PublishAt),
			CategoryID:  chall.CategoryID,
			CtfEventID:  chall.CTFEventID.Ptr(),
		}

		res[i].Flags = make([]*spec.AdminChallengeFlag, len(chall.R.Flags))
		for i2, v := range chall.R.Flags {
			res[i].Flags[i2] = &spec.AdminChallengeFlag{
				ID:   v.ID,
				Flag: v.Flag,
			}
		}

		res[i].Files = make([]*spec.AdminChallengeFiles, len(chall.R.ChallengeFiles))
		for i2, file := range chall.R.ChallengeFiles {
			res[i].Files[i2] = &spec.AdminChallengeFiles{
				ID:       file.ID,
				Filename: file.FriendlyName,
				URL:      file.URL,
			}
		}

		res[i].Authors = make([]string, len(chall.R.Authors))
		for i2, v := range chall.R.Authors {
			res[i].Authors[i2] = v.ID
		}
	}

	return res, nil
}

func (s *service) CreateChallenge(ctx context.Context, req *spec.CreateChallengePayload) error {

	var pubAt null.Time
	if req.PublishAt != nil {
		pubAt = null.TimeFrom(time.Unix(*req.PublishAt, 0))
	}

	chall := models.Challenge{
		ID:          uuid.New().String(),
		Title:       req.Title,
		Slug:        req.Slug,
		Description: req.Description,
		StaticScore: null.IntFromPtr(req.StaticScore),
		PublishAt:   pubAt,
		CTFEventID:  null.StringFromPtr(req.CtfEventID),
		CategoryID:  req.CategoryID,
	}

	err := chall.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		s.log.Error("inserting challenge", zap.Error(err), utils.C(ctx))
		return err
	}

	for _, v := range req.Authors {
		_, err = s.db.ExecContext(ctx, "INSERT INTO challenge_authors (challenge_id, user_id) VALUES ($1,$2)", chall.ID, v)
		if err != nil {
			s.log.Error("could not insert author", zap.Error(err))
		}
	}

	return nil
}

func (s *service) UpdateChallenge(ctx context.Context, req *spec.UpdateChallengePayload) error {

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var pubAt null.Time
	if req.PublishAt != nil {
		pubAt = null.TimeFrom(time.Unix(*req.PublishAt, 0))
	}

	n, err := models.Challenges(
		models.ChallengeWhere.ID.EQ(req.ChallengeID),
	).UpdateAll(ctx, tx, models.M{
		models.ChallengeColumns.Title:       req.Title,
		models.ChallengeColumns.StaticScore: null.IntFromPtr(req.StaticScore),
		models.ChallengeColumns.Slug:        req.Slug,
		models.ChallengeColumns.Description: req.Description,
		models.ChallengeColumns.PublishAt:   pubAt,
		models.ChallengeColumns.CTFEventID:  null.StringFromPtr(req.CtfEventID),
		models.ChallengeColumns.CategoryID:  req.CategoryID,
	})
	if err != nil {
		s.log.Error("could not update chall", zap.Error(err))
		return err
	}
	if n == 0 {
		return spec.MakeNotFound(errors.New("chall not found"))
	}

	// hacky
	_, err = tx.ExecContext(ctx, "DELETE FROM challenge_authors WHERE challenge_id = $1", req.ChallengeID)
	if err != nil {
		s.log.Error("could not delete old authors", zap.Error(err))
		return err
	}
	for _, v := range req.Authors {
		_, err = tx.ExecContext(ctx, "INSERT INTO challenge_authors (challenge_id, author_id) VALUES ($1,$2)", req.ChallengeID, v)
		if err != nil {
			s.log.Error("could not insert new author", zap.Error(err))
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		s.log.Error("could not commit", zap.Error(err))
		return err
	}

	return nil
}

func (s *service) PresignChallFileUpload(ctx context.Context, req *spec.PresignChallFileUploadPayload) (*spec.PresignChallFileUploadResult, error) {
	log := s.log.With(zap.String("challengeID", req.ChallengeID), utils.C(ctx))

	exists, err := models.ChallengeExists(ctx, s.db, req.ChallengeID)
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
		Bucket:        &s.cfg.S3.Bucket,
		Key:           &objectKey,
		ContentMD5:    &req.Md5,
		ContentLength: &req.Size,
	})

	url, err := r.Presign(time.Second * 20)
	if err != nil {
		log.Error("could not sign url", zap.Error(err))
		return nil, err
	}

	fileID := uuid.New()

	file := models.ChallengeFile{
		ID:           fileID.String(),
		ChallengeID:  null.StringFrom(req.ChallengeID),
		FriendlyName: req.Filename,
		URL:          fmt.Sprintf("%s/%s/%s", s.cfg.S3.Endpoint, s.cfg.S3.Bucket, objectKey),
	}
	err = file.Insert(ctx, s.db, boil.Infer())

	if err != nil {
		log.Error("could not insert file", zap.Error(err))
		return nil, err
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				s.log.Error("objectWait paniced", zap.Any("err", err))
			}
		}()
		s.checkUploaded(context.Background(), fileID, s.cfg.S3.Bucket, objectKey)
	}()

	return &spec.PresignChallFileUploadResult{
		URL: url,
	}, nil
}

func (s *service) checkUploaded(ctx context.Context, fileID uuid.UUID, bucket, key string) {
	log := s.log.With(zap.String("fileID", fileID.String()))

	err := s.s3.WaitUntilObjectExistsWithContext(ctx, &s3.HeadObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}, request.WithWaiterDelay(request.ConstantWaiterDelay(time.Second*5)), request.WithWaiterMaxAttempts(5))

	if err == nil {
		return
	}

	if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "ResourceNotReady" {
		log.Info("file was not uploaded", zap.String("fileID", fileID.String()))

		_, err = models.ChallengeFiles(
			models.ChallengeFileWhere.ID.EQ(fileID.String()),
		).DeleteAll(ctx, s.db)
		if err != nil {
			log.Error("could not delete non-uploaded file", zap.Error(err))
		}

		return
	}

	log.Warn("could not head object", zap.Error(err), zap.String("fileID", fileID.String()))
}

func (s *service) DeleteFile(ctx context.Context, req *spec.DeleteFilePayload) error {
	fileID := uuid.MustParse(req.FileID)
	log := s.log.With(zap.String("fileID", req.FileID), utils.C(ctx))
	log.Info("deleting file")

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.Error("could not open tx", zap.Error(err))
		return err
	}
	defer tx.Rollback() //nolint:errcheck

	file, err := models.ChallengeFiles(
		models.ChallengeFileWhere.ID.EQ(fileID.String()),
	).One(ctx, tx)
	if err == sql.ErrNoRows {
		return spec.MakeNotFound(errors.New("could not find file"))
	}
	if err != nil {
		log.Error("could not get file", zap.Error(err))
		return err
	}

	_, err = file.Delete(ctx, tx)
	if err != nil {
		log.Error("could not delete file", zap.Error(err))
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Error("could not commit file deletion", zap.Error(err))
		return err
	}

	url, err := url.Parse(file.URL)
	if err != nil {
		log.Error("bad file url", zap.Error(err), zap.String("url", file.URL))
		return err
	}

	fake := "fake"
	out, err := s.s3.DeleteObjectWithContext(ctx, &s3.DeleteObjectInput{
		Bucket: &fake,
		Key:    &fake,
	}, func(r *request.Request) {
		r.HTTPRequest.URL = url // hack since challtools won't let us know specifics
	})
	if err != nil {
		log.Error("delete object errored", zap.Error(err))
		return err
	}

	if out.DeleteMarker != nil && !*out.DeleteMarker {
		log.Warn("s3 file was not deleted")
	}

	return nil
}

func (s *service) GetChallengeMeta(ctx context.Context, req *spec.GetChallengeMetaPayload) (*spec.ChallengeMeta, error) {

	subs, err := models.Submissions(
		models.SubmissionWhere.ChallengeID.EQ(req.ChallengeID),
	).All(ctx, s.db)
	if err != nil {
		return nil, err
	}

	solves, err := models.UserSolves(
		models.UserSolfWhere.ChallengeID.EQ(req.ChallengeID),
	).All(ctx, s.db)

	res := &spec.ChallengeMeta{
		Solvers:     make([]*spec.ChallengeSolver, len(solves)),
		Submissions: make([]*spec.ChallengeSubmission, len(subs)),
	}

	for i, v := range subs {
		res.Submissions[i] = &spec.ChallengeSubmission{
			ID:          v.ID,
			UserID:      v.UserID,
			Input:       v.Input,
			Successful:  v.Successful,
			SubmittedAt: v.CreatedAt.Unix(),
		}
	}

	for i, v := range solves {
		res.Solvers[i] = &spec.ChallengeSolver{
			UserID:   v.UserID,
			SolvedAt: v.CreatedAt.Unix(),
		}
	}

	return res, nil
}
