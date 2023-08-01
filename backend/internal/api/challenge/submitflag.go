package challenge

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/sakerhetsm/ssm-wargame/internal/auth"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

func (s *service) SubmitFlag(ctx context.Context, req *spec.SubmitFlagPayload) error {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback() //nolint:errcheck

	user := auth.GetUser(ctx)
	challID := uuid.MustParse(req.ChallengeID)

	solved, err := models.UserSolves(
		models.UserSolfWhere.UserID.EQ(user.ID),
		models.UserSolfWhere.ChallengeID.EQ(challID.String()),
	).Exists(ctx, tx)

	if err != nil {
		return err
	}

	if solved {
		return spec.MakeAlreadySolved(errors.New("already solved"))
	}

	flagCorrect, err := models.Flags(
		models.FlagWhere.ChallengeID.EQ(challID.String()),
		models.FlagWhere.Flag.EQ(req.Flag),
	).Exists(ctx, tx)

	if err != nil {
		s.log.Error("FlagExists failed", zap.Error(err), utils.C(ctx))
		return err
	}

	attempt := models.Submission{
		ID:          uuid.New().String(),
		UserID:      user.ID,
		ChallengeID: challID.String(),
		Successful:  flagCorrect,
		Input:       req.Flag,
	}
	err = attempt.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return err
	}

	if flagCorrect {
		solve := models.UserSolf{
			UserID:      user.ID,
			ChallengeID: challID.String(),
		}
		err = solve.Insert(ctx, tx, boil.Infer())
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		s.log.Error("could not commit", zap.Error(err), utils.C(ctx))
		return err
	}

	if !flagCorrect {
		return spec.MakeIncorrectFlag(errors.New("incorrect flag"))
	}

	return nil
}
