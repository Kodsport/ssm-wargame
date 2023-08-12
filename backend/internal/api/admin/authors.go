package admin

import (
	"context"

	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
	"go.uber.org/zap"
)

func (s *service) ListAuthors(ctx context.Context, req *spec.ListAuthorsPayload) ([]*spec.SsmAuthor, error) {
	users, err := models.Authors().All(ctx, s.db)
	if err != nil {
		s.log.Warn("could not list users", zap.Error(err), utils.C(ctx))
		return nil, err
	}

	res := make([]*spec.SsmAuthor, len(users))
	for i, u := range users {

		res[i] = &spec.SsmAuthor{
			ID:          u.ID,
			FullName:    u.FullName,
			Description: u.Description,
			Sponsor:     u.Sponsor,
			Slug:        u.Slug,
			ImageURL:    u.ImageURL.Ptr(),
			Publish:     u.Publish,
		}
	}

	return res, nil
}

func (s *service) UpdateAuthor(ctx context.Context, req *spec.UpdateAuthorPayload) error {
	_, err := models.Authors(
		models.AuthorWhere.ID.EQ(req.ID),
	).UpdateAll(ctx, s.db, models.M{
		models.AuthorColumns.Slug:        req.Slug,
		models.AuthorColumns.Description: req.Description,
		models.AuthorColumns.FullName:    req.FullName,
		models.AuthorColumns.ImageURL:    req.ImageURL,
		models.AuthorColumns.Publish:     req.Publish,
		models.AuthorColumns.Sponsor:     req.Sponsor,
	})
	if err != nil {
		return err
	}

	return nil
}
