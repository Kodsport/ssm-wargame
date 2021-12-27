package auth

import (
	"context"

	"github.com/sakerhetsm/ssm-wargame/internal/db"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"goa.design/goa/v3/middleware"
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

func LogCtx(ctx context.Context) []zap.Field {
	// might exists better solutions for logging context values https://github.com/uber-go/zap/issues/654

	fields := []zapcore.Field{}

	user, ok := ctx.Value(UserKey).(*db.User)
	if ok {
		fields = append(fields, zap.String("userID", user.ID.String()), zap.String("role", user.Role))
	}

	reqID, ok := ctx.Value(middleware.RequestIDKey).(string)
	if ok {
		fields = append(fields, zap.String("requestID", reqID))
	}

	return fields
}
