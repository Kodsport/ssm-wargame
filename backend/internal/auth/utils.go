package auth

import (
	"context"

	"github.com/sakerhetsm/ssm-wargame/internal/db"
)

func HasRole(ctx context.Context, roles ...string) bool {
	user := ctx.Value(UserKey).(*db.User)

	for _, role := range roles {
		if role == user.Role {
			return true
		}
	}
	return false
}

func GetUser(ctx context.Context) *db.User {
	return ctx.Value(UserKey).(*db.User)
}

func IsAuthed(ctx context.Context) bool {
	_, ok := ctx.Value(UserKey).(*db.User)
	return ok
}
