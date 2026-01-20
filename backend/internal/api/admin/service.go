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
	ctfapi "github.com/sakerhetsm/ssm-wargame/internal/api/ctf"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
)

type service struct {
	auther     spec.Auther
	db         *sql.DB
	log        *zap.Logger
	s3         *s3.S3
	cfg        *config.Config
	ctfService *ctfapi.Service
}

func NewService(conn *sql.DB, log *zap.Logger, auther spec.Auther, s3c *s3.S3, cfg *config.Config, ctfService *ctfapi.Service) spec.Service {
	return &service{
		auther:     auther,
		db:         conn,
		log:        log,
		s3:         s3c,
		cfg:        cfg,
		ctfService: ctfService,
	}
}

func (s *service) JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {
	ctx, err := s.auther.JWTAuth(ctx, token, schema)
	if err != nil {
		return ctx, err
	}

	if !auth.HasRole(ctx, "author", "org", "admin") {
		return ctx, spec.MakeUnauthorized(errors.New("user is not admin or author"))
	}

	return ctx, nil
}
