package admin

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/sakerhetsm/ssm-wargame/internal/auth"
	"github.com/sakerhetsm/ssm-wargame/internal/db"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
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

func (s *service) ListChallenges(ctx context.Context, req *spec.ListChallengesPayload) (spec.SsmChallengeCollection, error) {

	return nil, nil
}

func (s *service) CreateChallenge(ctx context.Context, req *spec.CreateChallengePayload) error {

	q := db.New(s.db)

	err := q.InsertChallenge(ctx, db.InsertChallengeParams{
		ID:          uuid.New(),
		Title:       req.Title,
		Slug:        req.Slug,
		Description: req.Description,
		Score:       req.Score,
		Published:   false,
	})
	if err != nil {
		s.log.Error("inserting challenge", zap.Error(err), utils.C(ctx))
		return err
	}

	return nil
}
