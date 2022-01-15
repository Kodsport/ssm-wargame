package challenge

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sakerhetsm/ssm-wargame/internal/auth"
	"github.com/sakerhetsm/ssm-wargame/internal/db"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
	"go.uber.org/zap"
)

type service struct {
	spec.Auther
	db  *pgxpool.Pool
	log *zap.Logger
	s3  *s3.S3
}

func NewService(conn *pgxpool.Pool, log *zap.Logger, auther spec.Auther, s3c *s3.S3) spec.Service {
	return &service{
		Auther: auther,
		db:     conn,
		log:    log,
		s3:     s3c,
	}
}

func (s *service) ListMonthlyChallenges(ctx context.Context, req *spec.ListMonthlyChallengesPayload) (spec.SsmMonthlyChallengeCollection, error) {
	return nil, nil
}

func (s *service) ListChallenges(ctx context.Context, req *spec.ListChallengesPayload) (spec.SsmChallengeCollection, error) {

	db := db.New(s.db)

	challs, err := db.ListChallengesWithSolves(ctx, false)
	if err != nil {
		s.log.Error("could not list challs", zap.Error(err), utils.C(ctx))
		return nil, err
	}

	ids := make([]uuid.UUID, len(challs))
	for i, chall := range challs {
		ids[i] = chall.ID
	}

	files, err := db.GetChallFiles(ctx, ids)
	if err != nil {
		s.log.Error("could not list files", zap.Error(err), utils.C(ctx))
		return nil, err
	}

	res := make(spec.SsmChallengeCollection, len(challs))
	for i, c := range challs {
		res[i] = &spec.SsmChallenge{
			ID:          c.ID.String(),
			Slug:        c.Slug,
			Title:       c.Title,
			Description: c.Description,
			Score:       c.Score,
			Published:   c.Published,
			Solves:      c.NumSolves,
		}
	}

	for _, file := range files {
		for i, r := range res {
			if r.ID != file.ChallengeID.UUID.String() {
				continue
			}

			if res[i].Files == nil {
				res[i].Files = make([]*spec.ChallengeFiles, 0, 1)
			}

			req, _ := s.s3.GetObjectRequest(&s3.GetObjectInput{
				Bucket: &file.Bucket,
				Key:    &file.Key,
			})

			url, err := req.Presign(time.Hour * 4)

			if err != nil {
				s.log.Warn("could not sign url", zap.Error(err))
			}

			res[i].Files = append(res[i].Files, &spec.ChallengeFiles{
				Filename: file.FriendlyName,
				URL:      url, // TODO: this should be restructured, cache signed url in db?
			})
		}
	}

	return res, nil
}

func (s *service) SubmitFlag(ctx context.Context, req *spec.SubmitFlagPayload) error {

	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	txq := (&db.Queries{}).WithTx(tx)

	user := auth.GetUser(ctx)
	challID := uuid.MustParse(req.ChallengeID)

	solved, err := txq.UserHasSolved(ctx, db.UserHasSolvedParams{
		ChallengeID: challID,
		UserID:      user.ID,
	})
	if err != nil {
		return err
	}

	if solved {
		return spec.MakeAlreadySolved(errors.New("already solved"))
	}

	flagCorrect, err := txq.FlagExists(ctx, db.FlagExistsParams{
		ChallengeID: challID,
		Flag:        req.Flag,
	})
	if err != nil {
		s.log.Error("FlagExists failed", zap.Error(err), utils.C(ctx))
		return err
	}

	err = txq.InsertAttempt(ctx, db.InsertAttemptParams{
		ID:          uuid.New(),
		UserID:      user.ID,
		ChallengeID: challID,
		Successful:  flagCorrect,
		Input:       req.Flag,
	})
	if err != nil {
		return err
	}

	if flagCorrect {
		err = txq.InsertSolve(ctx, db.InsertSolveParams{
			UserID:      user.ID,
			ChallengeID: challID,
		})
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		s.log.Error("could not commit", zap.Error(err), utils.C(ctx))
		return err
	}

	if !flagCorrect {
		return spec.MakeIncorrectFlag(errors.New("incorrect flag"))
	}

	return nil

}
