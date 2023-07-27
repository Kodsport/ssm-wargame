package user

import (
	"context"
	"database/sql"

	"github.com/sakerhetsm/ssm-wargame/internal/auth"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/user"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
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
	user := ctx.Value(auth.UserKey).(*models.User)

	return &spec.GetSelfResult{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName.String,
		LastName:  user.LastName.String,
		Role:      user.Role,
	}, nil
}
