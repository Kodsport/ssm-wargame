package api

import (
	"context"
	"errors"

	challenge_spec "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
)

type service struct {
}

func NewService() challenge_spec.Service {
	return &service{}
}

func (s *service) ListChallenges(context.Context) (res challenge_spec.SsmChallengeCollection, err error) {
	return nil, errors.New("not implemented")
}
