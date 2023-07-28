package auth

import (
	"context"

	"github.com/sakerhetsm/ssm-wargame/internal/models"
)

func HasRole(ctx context.Context, roles ...string) bool {
	user := GetUser(ctx)

	for _, role := range roles {
		if role == user.Role {
			return true
		}
	}
	return false
}

func GetUser(ctx context.Context) *models.User {
	return ctx.Value(UserKey).(*models.User)
}

func IsAuthed(ctx context.Context) bool {
	_, ok := ctx.Value(UserKey).(*models.User)
	return ok
}
