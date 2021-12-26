package auth

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/sakerhetsm/ssm-wargame/internal/config"
	"github.com/sakerhetsm/ssm-wargame/internal/db"
	auth_spec "github.com/sakerhetsm/ssm-wargame/internal/gen/auth"
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"github.com/bwmarrin/discordgo"
)

type service struct {
	db        *pgx.Conn
	q         *db.Queries
	log       *zap.Logger
	config    *oauth2.Config
	jwtSecret []byte
}

func NewService(conn *pgx.Conn, log *zap.Logger, cfg *config.Config) auth_spec.Service {

	config := &oauth2.Config{
		ClientID:     cfg.OAuth.Discord.ClientID,
		ClientSecret: cfg.OAuth.Discord.ClientSecret,
		RedirectURL:  cfg.OAuth.Discord.RedirectURL,
		Scopes:       []string{"email", "identify"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://discord.com/api/oauth2/authorize",
			TokenURL: "https://discord.com/api/oauth2/token",
		},
	}

	return &service{
		db:        conn,
		q:         db.New(conn),
		log:       log.Named("auth"),
		config:    config,
		jwtSecret: []byte(cfg.JWTSecret),
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

	token, err := s.config.Exchange(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	dc, err := discordgo.New(token.Type() + " " + token.AccessToken)
	if err != nil {
		return nil, err
	}

	dcUser, err := dc.User("@me")
	if err != nil {
		s.log.Warn("could not get discord user", zap.Error(err))
		return nil, err
	}

	// Checks if the user already exists, insert them if not
	userID, err := s.q.UserIDByDiscordID(ctx, dcUser.ID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		userID = uuid.New()
		err = s.q.InsertUserDiscord(ctx, db.InsertUserDiscordParams{
			ID:        userID,
			DiscordID: dcUser.ID,
			Email:     dcUser.Email,
		})
		if err != nil {
			s.log.Error("could not insert new user", zap.Error(err), zap.String("discordID", dcUser.ID))
			return nil, err
		}
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		// Id: uuid.NewString(), // don't store, but might be useful for log context
		Subject:   userID.String(),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 4).Unix(),
		NotBefore: time.Now().Unix(),
	})

	jwtStr, err := jwtToken.SignedString(s.jwtSecret)
	if err != nil {
		return nil, err
	}

	return &auth_spec.ExchangeDiscordResult{
		JWT: jwtStr,
	}, nil
}
