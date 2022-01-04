package admin

import (
	"context"

	"github.com/google/uuid"
	"github.com/sakerhetsm/ssm-wargame/internal/db"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
	"go.uber.org/zap"

	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
)

func (s *service) ListChallenges(ctx context.Context, req *spec.ListChallengesPayload) (spec.SsmChallengeCollection, error) {

	challs, err := db.New(s.db).ListChallengesWithSolves(ctx, true)
	if err != nil {
		s.log.Warn("could not list challs", zap.Error(err), utils.C(ctx))
		return nil, err
	}

	res := make(spec.SsmChallengeCollection, len(challs))

	for i, chall := range challs {
		res[i] = &spec.SsmChallenge{
			ID:          chall.ID.String(),
			Slug:        chall.Slug,
			Title:       chall.Title,
			Description: chall.Description,
			Score:       chall.Score,
			Solves:      chall.NumSolves,
			Published:   chall.Published,
		}

	}

	return res, nil
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
