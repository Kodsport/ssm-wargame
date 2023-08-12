package admin

import (
	"context"

	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
	"go.uber.org/zap"
)

func (s *service) ListUsers(ctx context.Context, req *spec.ListUsersPayload) ([]*spec.SsmUser, error) {
	users, err := models.Users().All(ctx, s.db)
	if err != nil {
		s.log.Warn("could not list users", zap.Error(err), utils.C(ctx))
		return nil, err
	}

	res := make([]*spec.SsmUser, len(users))
	for i, u := range users {
		res[i] = &spec.SsmUser{
			ID:       u.ID,
			Email:    u.Email,
			FullName: u.FullName,
			Role:     u.Role,
		}
	}

	return res, nil
}
