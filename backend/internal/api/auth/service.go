package auth

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/sakerhetsm/ssm-wargame/internal/config"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/auth"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"github.com/bwmarrin/discordgo"
)

type service struct {
	db        *sql.DB
	log       *zap.Logger
	config    *oauth2.Config
	jwtSecret []byte
}

func NewService(conn *sql.DB, log *zap.Logger, cfg *config.Config) spec.Service {
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
		log:       log,
		config:    config,
		jwtSecret: []byte(cfg.JWTSecret),
	}
}

func (s *service) GenerateDiscordAuthURL(ctx context.Context) (*spec.GenerateDiscordAuthURLResult, error) {
	url := s.config.AuthCodeURL("state_todo")
	return &spec.GenerateDiscordAuthURLResult{
		URL: url,
	}, nil
}

func (s *service) ExchangeDiscord(ctx context.Context, req *spec.ExchangeDiscordPayload) (*spec.ExchangeDiscordResult, error) {
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
		s.log.Warn("could not get discord user", zap.Error(err), utils.C(ctx))
		return nil, err
	}

	// Checks if the user already exists, insert them if not
	user, err := models.Users(
		models.UserWhere.DiscordID.EQ(null.StringFrom(dcUser.ID)),
		qm.Select(models.UserColumns.ID),
	).One(ctx, s.db)
	if err != nil && err != sql.ErrNoRows {
		s.log.Error("db err", zap.Error(err), zap.String("discordID", dcUser.ID), utils.C(ctx))
		return nil, err
	}

	if err == sql.ErrNoRows {
		user = &models.User{
			ID:        uuid.New().String(),
			DiscordID: null.StringFrom(dcUser.ID),
			Email:     dcUser.Email,
			Role:      "solver",
			FullName:  dcUser.Username,
		}
		err = user.Insert(ctx, s.db, boil.Infer())

		if err != nil {
			s.log.Error("could not insert new user", zap.Error(err), zap.String("discordID", dcUser.ID), utils.C(ctx))
			return nil, err
		}
	}

	jwtStr, err := s.genJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return &spec.ExchangeDiscordResult{
		JWT: jwtStr,
	}, nil
}

func (s *service) genJWT(userID string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		// Id: uuid.NewString(), // don't store, but might be useful for log context
		Subject:   userID,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 4).Unix(),
		NotBefore: time.Now().Unix(),
	})

	return jwtToken.SignedString(s.jwtSecret)
}
