package admin

import (
	"context"
	"errors"
	"time"

	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

func (s *service) CreateMonthlyChallenge(ctx context.Context, req *spec.CreateMonthlyChallengePayload) error {
	// time format/resolution TBD
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return err
	}
	endDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return err
	}

	mc := models.MonthlyChallenge{
		ChallengeID:  req.ChallengeID,
		StartDate:    startDate,
		EndDate:      endDate,
		DisplayMonth: req.DisplayMonth,
	}
	err = mc.Insert(ctx, s.db, boil.Infer())

	if err != nil {
		s.log.Warn("could not insert monthly challs", zap.Error(err), utils.C(ctx))
		return err
	}

	return nil
}

func (s *service) ListMonthlyChallenges(ctx context.Context, req *spec.ListMonthlyChallengesPayload) ([]*spec.MonthlyChallenge, error) {
	challs, err := models.MonthlyChallenges(
		qm.Load(models.MonthlyChallengeRels.Challenge),
	).All(ctx, s.db)
	if err != nil {
		s.log.Warn("could not list monthly challs", zap.Error(err), utils.C(ctx))
		return nil, err
	}

	res := make([]*spec.MonthlyChallenge, len(challs))

	for i, chall := range challs {
		res[i] = &spec.MonthlyChallenge{
			DisplayMonth: chall.DisplayMonth,
			StartDate:    chall.StartDate.Format("2006-01-02"),
			EndDate:      chall.EndDate.Format("2006-01-02"),
			Slug:         chall.R.Challenge.Slug,
			Title:        chall.R.Challenge.Title,
			Description:  chall.R.Challenge.Description,
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
