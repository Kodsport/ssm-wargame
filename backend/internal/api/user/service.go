package user

import (
	"context"
	"database/sql"
	"errors"
	"regexp"

	"github.com/sakerhetsm/ssm-wargame/internal/auth"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/user"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

type service struct {
	spec.Auther
	db  *sql.DB
	log *zap.Logger
}

func NewService(conn *sql.DB, log *zap.Logger, auther spec.Auther) spec.Service {
	return &service{
		Auther: auther,
		db:     conn,
		log:    log,
	}
}

func (s *service) GetSelf(ctx context.Context, req *spec.GetSelfPayload) (*spec.GetSelfResult, error) {
	user := auth.GetUser(ctx)

	res := &spec.GetSelfResult{
		ID:             user.ID,
		Email:          user.Email,
		FullName:       user.FullName,
		Role:           user.Role,
		SchoolID:       user.SchoolID.Ptr(),
		OnboardingDone: user.OnboardingDone,
	}

	if user.SchoolID.Valid {
		err := user.L.LoadSchool(ctx, s.db, true, user, qm.Select(models.SchoolColumns.Name))
		if err != nil {
			s.log.Error("could not get school", zap.Error(err))
			return nil, err
		}
		res.SchoolName = &user.R.School.Name
	}

	return res, nil
}

func (s *service) CompleteOnboarding(ctx context.Context, req *spec.CompleteOnboardingPayload) error {
	user := auth.GetUser(ctx)

	if user.OnboardingDone {
		return errors.New("already done")
	}

	user.OnboardingDone = true

	_, err := user.Update(ctx, s.db, boil.Whitelist(models.UserColumns.OnboardingDone))
	if err != nil {
		s.log.Error("could not update user", zap.Error(err), utils.C(ctx))
		return err
	}

	return nil
}

func (s *service) UpdateSelf(ctx context.Context, req *spec.UpdateSelfPayload) error {
	re, _ := regexp.Compile(`[\t\n\f\r]`)
	if re.FindAllString(req.FullName, -1) != nil || len(req.FullName) > 30 || len(req.FullName) < 3 {
		return spec.MakeInvalidUsername(errors.New("fullname must be less than 30 characters and not contain newlines"))
	}
	_, err := models.Users(
		models.UserWhere.ID.EQ(auth.GetUser(ctx).ID),
	).UpdateAll(ctx, s.db, models.M{
		models.UserColumns.FullName: req.FullName,
	})

	if err != nil {
		return err
	}

	return nil
}
