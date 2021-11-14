package api

import (
	"context"

	"github.com/sakerhetsm/ssm-wargame/internal/db"
	challenge_spec "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
)

type service struct {
	db *db.Queries
}

func NewService(db *db.Queries) challenge_spec.Service {
	return &service{
		db: db,
	}
}

func (s *service) ListChallenges(ctx context.Context) (challenge_spec.SsmChallengeCollection, error) {

	challs, err := s.db.ListChallenges(ctx)
	if err != nil {
		return nil, err
	}

	res := make(challenge_spec.SsmChallengeCollection, len(challs))
	for i, c := range challs {
		res[i] = &challenge_spec.SsmChallenge{
			ID:          c.ID.String(),
			Title:       c.Title,
			Description: c.Description,
			Score:       int(c.Score),
			Published:   c.Published,
		}
	}

	return res, nil
}
