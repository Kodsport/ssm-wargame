package admin

import (
	"context"
	"errors"

	"github.com/google/uuid"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

func (s *service) CreateAuthor(ctx context.Context, req *spec.CreateAuthorPayload) error {
	a := models.Author{
		ID:          uuid.NewString(),
		Slug:        req.Slug,
		FullName:    req.FullName,
		Sponsor:     req.Sponsor,
		Description: req.Description,
		ImageURL:    null.StringFromPtr(req.ImageURL),
		Publish:     req.Publish,
	}

	return a.Insert(ctx, s.db, boil.Infer())
}

func (s *service) ListAuthors(ctx context.Context, req *spec.ListAuthorsPayload) ([]*spec.Author, error) {
	users, err := models.Authors().All(ctx, s.db)
	if err != nil {
		s.log.Warn("could not list users", zap.Error(err), utils.C(ctx))
		return nil, err
	}

	res := make([]*spec.Author, len(users))
	for i, u := range users {

		res[i] = &spec.Author{
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

func (s *service) DeleteAuthor(ctx context.Context, req *spec.DeleteAuthorPayload) error {
	n, err := models.Authors(
		models.AuthorWhere.ID.EQ(req.ID),
	).DeleteAll(ctx, s.db)

	if n == 0 {
		return spec.MakeNotFound(errors.New("no such author"))
	}
	if err != nil {
		return err
	}

	return nil
}
