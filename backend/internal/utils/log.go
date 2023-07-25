package utils

import (
	"github.com/sakerhetsm/ssm-wargame/internal/auth"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"goa.design/goa/v3/middleware"
)

type valuer interface {
	Value(key interface{}) interface{}
}

func C(value valuer) zap.Field {
	return zap.Object("ctx", valuerI{value})
}

type valuerI struct {
	valuer
}

func (v valuerI) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	user, ok := v.Value(auth.UserKey).(*models.User)
	if ok {
		enc.AddString("userID", user.ID)
		enc.AddString("role", user.Role)
	}

	reqID, ok := v.Value(middleware.RequestIDKey).(string)
	if ok {
		enc.AddString("requestID", reqID)
	}

	return nil
}
