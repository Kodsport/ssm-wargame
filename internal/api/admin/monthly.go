package admin

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/sakerhetsm/ssm-wargame/internal/db"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
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
	displayDate, err := time.Parse("2006-01-02", req.DisplayMonth)
	if err != nil {
		return err
	}

	err = db.New(s.db).InsertMonthlyChallenge(ctx, db.InsertMonthlyChallengeParams{
		ChallengeID:  uuid.MustParse(req.ChallengeID),
		StartDate:    startDate,
		EndDate:      endDate,
		DisplayMonth: displayDate,
	})

	if err != nil {
		s.log.Warn("could not insert monthly challs", zap.Error(err), utils.C(ctx))
		return err
	}

	return nil
}

func (s *service) ListMonthlyChallenges(ctx context.Context, req *spec.ListMonthlyChallengesPayload) ([]*spec.MonthlyChallengeMeta, error) {

	challs, err := db.New(s.db).ListMonthlyChallenges(ctx)
	if err != nil {
		s.log.Warn("could not list monthly challs", zap.Error(err), utils.C(ctx))
		return nil, err
	}

	res := make([]*spec.MonthlyChallengeMeta, len(challs))

	for i, chall := range challs {
		res[i] = &spec.MonthlyChallengeMeta{
			DisplayMonth: chall.DisplayMonth.Format("2006-01-02"), // detta kanske ska vara fritext
			StartDate:    chall.StartDate.Format("2006-01-02"),
			EndDate:      chall.EndDate.Format("2006-01-02"),
		}
	}

	return res, nil
}

func (s *service) DeleteMonthlyChallenge(ctx context.Context, req *spec.DeleteMonthlyChallengePayload) error {

	err := db.New(s.db).DeleteMonthlyChallenge(ctx, uuid.MustParse(req.ChallengeID))
	if err == pgx.ErrNoRows {
		return spec.MakeNotFound(errors.New("monthly challenge not found"))
	}
	if err != nil {
		s.log.Warn("could not delete monthly chall", zap.Error(err), utils.C(ctx))
		return err
	}

	return nil
}
