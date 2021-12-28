package admin

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/sakerhetsm/ssm-wargame/internal/db"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
	"go.uber.org/zap"

	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
)

type service struct {
	spec.Auther
	db  *pgx.Conn
	log *zap.Logger
}

func NewService(conn *pgx.Conn, log *zap.Logger, auther spec.Auther) spec.Service {
	return &service{
		Auther: auther,
		db:     conn,
		log:    log.Named("admin"),
	}
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
