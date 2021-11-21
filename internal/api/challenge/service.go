package challenge

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/sakerhetsm/ssm-wargame/internal/db"
	challenge_spec "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	"go.uber.org/zap"
)

type service struct {
	db  *pgx.Conn
	log *zap.Logger
}

func NewService(conn *pgx.Conn, log *zap.Logger) challenge_spec.Service {
	return &service{
		db:  conn,
		log: log.Named("challenge"),
	}
}

func (s *service) ListChallenges(ctx context.Context) (challenge_spec.SsmChallengeCollection, error) {

	challs, err := db.New(s.db).ListChallengesWithSolves(ctx, true)
	if err != nil {
		return nil, err
	}

	res := make(challenge_spec.SsmChallengeCollection, len(challs))
	for i, c := range challs {
		res[i] = &challenge_spec.SsmChallenge{
			ID:          c.ID.String(),
			Title:       c.Title,
			Description: c.Description,
			Score:       uint(c.Score),
			Published:   c.Published,
			Solves:      uint(c.NumSolves),
		}
	}

	return res, nil
}

func (s *service) SubmitFlag(ctx context.Context, req *challenge_spec.SubmitFlagPayload) error {

	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	txq := (&db.Queries{}).WithTx(tx)

	userID := uuid.MustParse("")
	challID := uuid.MustParse(req.ChallengeID)

	solved, err := txq.UserHasSolved(ctx, db.UserHasSolvedParams{
		ChallengeID: challID,
		UserID:      userID,
	})

	if err != nil {
		return err
	}

	if solved {
		return challenge_spec.MakeAlreadySolved(errors.New("already solved"))
	}

	exists, err := txq.FlagExists(ctx, db.FlagExistsParams{
		ChallengeID: challID,
		Flag:        req.Flag,
	})

	if err != nil {
		s.log.Error("FlagExists failed", zap.Error(err))
		return err
	}

	err = txq.InsertAttempt(ctx, db.InsertAttemptParams{
		UserID:      userID,
		ChallengeID: challID,
		Successful:  exists,
		Input:       req.Flag,
	})
	if err != nil {
		return err
	}

	if exists {
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
		return err
	}

	if !exists {
		return challenge_spec.MakeIncorrectFlag(errors.New("incorrect flag"))
	}

	return nil

}
