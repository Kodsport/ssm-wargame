package admin

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/sakerhetsm/ssm-wargame/internal/auth"
	"go.uber.org/zap"
	"goa.design/goa/v3/security"

	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
)

type service struct {
	auther spec.Auther
	db     *pgx.Conn
	log    *zap.Logger
}

func NewService(conn *pgx.Conn, log *zap.Logger, auther spec.Auther) spec.Service {
	return &service{
		auther: auther,
		db:     conn,
		log:    log,
	}
}

func (s *service) JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {
	ctx, err := s.auther.JWTAuth(ctx, token, schema)
	if err != nil {
		return ctx, err
	}

	if !auth.HasRole(ctx, "author", "admin") {
		return ctx, spec.MakeUnauthorized(errors.New("user is not admin or author"))
	}

	return ctx, nil

}
