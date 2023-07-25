package admin

import (
	"context"
	"errors"

	"github.com/google/uuid"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

func (s *service) DeleteFlag(ctx context.Context, req *spec.DeleteFlagPayload) error {

	n, err := models.Flags(
		models.FlagWhere.ID.EQ(req.FlagID),
		models.FlagWhere.ChallengeID.EQ(req.ChallengeID),
	).DeleteAll(ctx, s.db)
	if err != nil {
		s.log.Error("could not delete flag", zap.Error(err), zap.String("flagID", req.FlagID))
		return err
	}

	if n == 0 {
		return spec.MakeNotFound(errors.New("flag not found"))
	}

	return nil
}
func (s *service) AddFlag(ctx context.Context, req *spec.AddFlagPayload) error {

	flag := models.Flag{
		ID:          uuid.NewString(),
		ChallengeID: req.ChallengeID,
		Flag:        req.Flag,
	}

	err := flag.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		s.log.Error("could not insert flag", zap.Error(err), zap.Any("flag", flag))
		return err
	}

	return nil
}
