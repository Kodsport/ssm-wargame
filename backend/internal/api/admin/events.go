package admin

import (
	"context"
	"encoding/hex"
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
)

func (s *service) ListCTFEvents(ctx context.Context, req *spec.ListCTFEventsPayload) ([]*spec.CTFEvent, error) {

	events, err := models.CTFEvents().All(ctx, s.db)
	if err != nil {
		s.log.Error("could not get events", zap.Error(err))
		return nil, err
	}

	res := make([]*spec.CTFEvent, len(events))
	for i, v := range events {
		res[i] = &spec.CTFEvent{
			ID:   v.ID,
			Name: v.Name,
		}
	}

	return res, nil
}

func (s *service) CreateCTFEvent(ctx context.Context, req *spec.CreateCTFEventPayload) error {
	e := models.CTFEvent{
		ID:   uuid.NewString(),
		Name: req.Name,
	}

	err := e.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		s.log.Error("could not insert event", zap.Error(err))
		return err
	}

	return nil
}

func (s *service) DeleteCTFEvent(ctx context.Context, req *spec.DeleteCTFEventPayload) error {
	n, err := models.CTFEvents(
		models.CTFEventWhere.ID.EQ(req.ID),
	).DeleteAll(ctx, s.db)
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("not found")
	}

	return nil

}

func (s *service) CreateCTFEventImportToken(ctx context.Context, req *spec.CreateCTFEventImportTokenPayload) (*spec.CreateCTFEventImportTokenResult, error) {

	r := make([]byte, 8)
	_, err := rand.Read(r)
	if err != nil {
		return nil, err
	}

	id := uuid.New()
	token := "ctfimp_" + hex.EncodeToString(id[:]) + "_" + hex.EncodeToString(r)

	hash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	t := models.ChalltoolsImportToken{
		ID:        id.String(),
		Token:     string(hash),
		ExpiresAt: time.Now().Add(time.Hour),
	}
	err = t.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &spec.CreateCTFEventImportTokenResult{
		Token: token,
	}, nil
}
