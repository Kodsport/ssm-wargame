package admin

import (
	"context"
	"database/sql"
	"errors"

	"github.com/sakerhetsm/ssm-wargame/internal/auth"
	"github.com/sakerhetsm/ssm-wargame/internal/config"
	"go.uber.org/zap"
	"goa.design/goa/v3/security"

	"github.com/aws/aws-sdk-go/service/s3"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
)

type service struct {
	auther spec.Auther
	db     *sql.DB
	log    *zap.Logger
	s3     *s3.S3
	cfg    *config.Config
}

func NewService(conn *sql.DB, log *zap.Logger, auther spec.Auther, s3c *s3.S3, cfg *config.Config) spec.Service {
	return &service{
		auther: auther,
		db:     conn,
		log:    log,
		s3:     s3c,
		cfg:    cfg,
	}
}

func (s *service) APIKeyAuth(ctx context.Context, token string, schema *security.APIKeyScheme) (context.Context, error) {
	return s.auther.APIKeyAuth(ctx, token, schema)
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
