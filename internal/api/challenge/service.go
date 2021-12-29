package challenge

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/sakerhetsm/ssm-wargame/internal/auth"
	"github.com/sakerhetsm/ssm-wargame/internal/db"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
	"go.uber.org/zap"
)

type service struct {
	spec.Auther
	db  *pgx.Conn
	log *zap.Logger
}

func NewService(conn *pgx.Conn, log *zap.Logger, auther spec.Auther) spec.Service {
	return &service{
		Auther: auther,
		db:     conn,
		log:    log,
	}
}

func (s *service) ListMonthlyChallenges(ctx context.Context, req *spec.ListMonthlyChallengesPayload) (spec.SsmMonthlyChallengeCollection, error) {
	return nil, nil
}

func (s *service) ListChallenges(ctx context.Context, req *spec.ListChallengesPayload) (spec.SsmChallengeCollection, error) {

	challs, err := db.New(s.db).ListChallengesWithSolves(ctx, false)
	if err != nil {
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
			UserID:      uuid.MustParse(""),
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
