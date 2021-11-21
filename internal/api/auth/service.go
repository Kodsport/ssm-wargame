package auth

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/sakerhetsm/ssm-wargame/internal/db"
	auth_spec "github.com/sakerhetsm/ssm-wargame/internal/gen/auth"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type service struct {
	db     *pgx.Conn
	q      *db.Queries
	log    *zap.Logger
	config *oauth2.Config
}

func NewService(conn *pgx.Conn, log *zap.Logger) auth_spec.Service {

	config := &oauth2.Config{
		ClientID:     "todo",
		ClientSecret: "todo",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://discord.com/api/oauth2/authorize",
			TokenURL: "https://discord.com/api/oauth2/token",
		},
		Scopes: []string{"email", "identify"},
	}

	return &service{
		db:     conn,
		q:      db.New(conn),
		log:    log.Named("auth"),
		config: config,
	}
}

func (s *service) GenerateDiscordAuthURL(ctx context.Context) (*auth_spec.GenerateDiscordAuthURLResult, error) {
	url := s.config.AuthCodeURL("state_todo")
	return &auth_spec.GenerateDiscordAuthURLResult{
		URL: url,
	}, nil
}

func (s *service) ExchangeDiscord(ctx context.Context, req *auth_spec.ExchangeDiscordPayload) (*auth_spec.ExchangeDiscordResult, error) {

	if req.State != "state_todo" {
		return nil, errors.New("invalid state")
	}

	_, err := s.config.Exchange(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	// create user if new etc
	// generate jwt

	return &auth_spec.ExchangeDiscordResult{}, nil
}
