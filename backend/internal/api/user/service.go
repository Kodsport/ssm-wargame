package user

import (
	"context"
	"database/sql"

	"github.com/sakerhetsm/ssm-wargame/internal/auth"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/user"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
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
		ID:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
		Role:     user.Role,
		SchoolID: user.SchoolID.Ptr(),
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
