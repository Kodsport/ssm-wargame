package challenge

import (
	"context"

	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"go.uber.org/zap"
)

func (s *service) ListEvents(ctx context.Context, req *spec.ListEventsPayload) ([]*spec.CTFEvent, error) {

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
