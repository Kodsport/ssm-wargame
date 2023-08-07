package auth

import (
	"context"
	"database/sql"
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/sakerhetsm/ssm-wargame/internal/config"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"go.uber.org/zap"
	"goa.design/goa/v3/security"
)

type ContextKey int

const (
	UserKey ContextKey = iota
)

type Auther struct {
	jwtSecret []byte
	log       *zap.Logger
	db        *sql.DB
}

func New(cfg *config.Config, log *zap.Logger, db *sql.DB) *Auther {
	return &Auther{
		jwtSecret: []byte(cfg.JWTSecret),
		log:       log.Named("auth-middleware"),
		db:        db,
	}
}

func (s *Auther) JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {
	if token == "" {
		return ctx, nil
	}

	claims := jwt.StandardClaims{}
	t, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		// limit alg?
		return s.jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, errors.New("token invalid")
	}

	user, err := models.Users(
		models.UserWhere.ID.EQ(claims.Subject),
	).One(ctx, s.db)
	if err == sql.ErrNoRows {
		s.log.Warn("subject from valid jwt doesn't exist in DB, cross-env token usage or security issue?", zap.String("subject", user.ID))
		return nil, errors.New("user not found")
	}
	if err != nil {
		s.log.Warn("db err when getting user", zap.String("subject", claims.Subject))
		return nil, err
	}

	ctx = context.WithValue(ctx, UserKey, user)

	return ctx, nil
}
