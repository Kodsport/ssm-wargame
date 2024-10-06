package challenge

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"strings"

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

	exits, err := models.Challenges(
		models.ChallengeWhere.ID.EQ(req.ChallengeID),
		models.ChallengeWhere.ChallNamespace.IsNull(),
	).Exists(ctx, tx)

	if err != nil {
		s.log.Error("could not check chall exists", zap.Error(err))
		return err
	}

	if !exits {
		return errors.New("challenge not found")
	}

	flags, err := models.Flags(
		models.FlagWhere.ChallengeID.EQ(challID.String()),
	).All(ctx, tx)
	if err != nil {
		s.log.Error("could not get flags", zap.Error(err), utils.C(ctx))
		return err
	}

	flagCorrect := false
	for _, flag := range flags {

		inp := strings.TrimPrefix(strings.TrimSuffix(req.Flag, flag.FlagSuffix), flag.FlagPrefix)

		if flag.Type == "regex" {
			if regexp.MustCompile(flag.Flag).Match([]byte(inp)) {
				flagCorrect = true
				break
			}
			continue
		}

		if flag.Flag == inp {
			flagCorrect = true
			break
		}
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

		if user.SchoolID.Valid {
			schoolSolve := models.SchoolSolf{
				SchoolID:    user.SchoolID.String,
				UserID:      user.ID,
				ChallengeID: challID.String(),
			}
			err = schoolSolve.Insert(ctx, tx, boil.Infer())
			if err != nil {
				return err
			}
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
