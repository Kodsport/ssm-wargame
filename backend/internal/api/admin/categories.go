package admin

import (
	"context"

	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"go.uber.org/zap"
)

func (s *service) ListCategories(ctx context.Context, req *spec.ListCategoriesPayload) ([]*spec.Category, error) {
	categories, err := models.Categories().All(ctx, s.db)
	if err != nil {
		s.log.Error("could not get categories", zap.Error(err))
		return nil, err
	}

	res := make([]*spec.Category, len(categories))
	for i, v := range categories {
		res[i] = &spec.Category{
			ID:   v.ID,
			Name: v.Name,
		}
	}

	return res, nil
}
