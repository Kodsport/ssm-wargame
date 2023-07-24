package admin

import (
	"context"

	"github.com/sakerhetsm/ssm-wargame/internal/db"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
	"go.uber.org/zap"
)

func (s *service) ListUsers(ctx context.Context, req *spec.ListUsersPayload) ([]*spec.SsmUser, error) {
	users, err := db.New(s.db).Users(ctx)
	if err != nil {
		s.log.Warn("could not list users", zap.Error(err), utils.C(ctx))
		return nil, err
	}

	res := make([]*spec.SsmUser, len(users))
	for i, u := range users {
		res[i] = &spec.SsmUser{
			ID:        u.ID.String(),
			Email:     u.Email,
			FirstName: u.FirstName.String,
			LastName:  u.LastName.String,
			Role:      u.Role,
		}
	}

	return res, nil
}
