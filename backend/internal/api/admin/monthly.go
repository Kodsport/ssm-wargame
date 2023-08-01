package admin

import (
	"context"
	"errors"
	"time"

	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

func (s *service) CreateMonthlyChallenge(ctx context.Context, req *spec.CreateMonthlyChallengePayload) error {

	mc := models.MonthlyChallenge{
		ChallengeID:  req.ChallengeID,
		StartDate:    time.Unix(req.StartDate, 0),
		EndDate:      time.Unix(req.EndDate, 0),
		DisplayMonth: req.DisplayMonth,
	}

	err := mc.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		s.log.Warn("could not insert monthly challs", zap.Error(err), utils.C(ctx))
		return err
	}

	return nil
}

func (s *service) ListMonthlyChallenges(ctx context.Context, req *spec.ListMonthlyChallengesPayload) ([]*spec.MonthlyChallenge, error) {
	challs, err := models.MonthlyChallenges().All(ctx, s.db)
	if err != nil {
		s.log.Warn("could not list monthly challs", zap.Error(err), utils.C(ctx))
		return nil, err
	}

	res := make([]*spec.MonthlyChallenge, len(challs))

	for i, chall := range challs {
		res[i] = &spec.MonthlyChallenge{
			DisplayMonth: chall.DisplayMonth,
			StartDate:    chall.StartDate.Unix(),
			EndDate:      chall.EndDate.Unix(),
			ChallengeID:  chall.ChallengeID,
		}
	}

	return res, nil
}

func (s *service) DeleteMonthlyChallenge(ctx context.Context, req *spec.DeleteMonthlyChallengePayload) error {
	n, err := models.MonthlyChallenges(
		models.MonthlyChallengeWhere.ChallengeID.EQ(req.ChallengeID),
	).DeleteAll(ctx, s.db)

	if err != nil {
		s.log.Warn("could not delete monthly chall", zap.Error(err), utils.C(ctx))
		return err
	}
	if n == 0 {
		return spec.MakeNotFound(errors.New("monthly challenge not found"))
	}

	return nil
}
