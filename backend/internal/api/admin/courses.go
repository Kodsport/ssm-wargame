package admin

import (
	"context"

	"github.com/google/uuid"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *service) ListCourses(ctx context.Context, req *spec.ListCoursesPayload) (spec.SsmAdminCourseCollection, error) {
	courses, err := models.Courses(
		qm.Load(models.CourseRels.Authors),
	).All(ctx, s.db)
	if err != nil {
		return nil, err
	}

	res := make([]*spec.SsmAdminCourse, len(courses))
	for i, v := range courses {
		res[i] = &spec.SsmAdminCourse{
			ID:          v.ID,
			Title:       v.Title,
			Slug:        v.Slug,
			Category:    v.Category,
			Difficulty:  v.Difficulty,
			Description: v.Description,
			Publish:     v.Publish,
			AuthorIds:   make([]string, len(v.R.Authors)),
		}

		for i2, v2 := range v.R.Authors {
			res[i].AuthorIds[i2] = v2.ID
		}
	}

	return res, nil
}

func (s *service) CreateCourse(ctx context.Context, req *spec.CreateCoursePayload) error {
	c := models.Course{
		ID:          uuid.NewString(),
		Title:       req.Title,
		Slug:        req.Slug,
		Category:    req.Category,
		Difficulty:  req.Difficulty,
		Description: req.Description,
		Publish:     req.Publish,
	}

	err := c.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return err
	}

	authors := make([]*models.Author, len(req.AuthorIds))
	for i, v := range req.AuthorIds {
		authors[i] = &models.Author{ID: v}
	}
	err = c.AddAuthors(ctx, s.db, true, authors...)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateCourse(ctx context.Context, req *spec.UpdateCoursePayload) error {

	_, err := models.Courses(
		models.CourseWhere.ID.EQ(req.ID),
	).UpdateAll(ctx, s.db, models.M{
		models.CourseColumns.Category:    req.Category,
		models.CourseColumns.Title:       req.Title,
		models.CourseColumns.Slug:        req.Slug,
		models.CourseColumns.Difficulty:  req.Difficulty,
		models.CourseColumns.Description: req.Description,
		models.CourseColumns.Publish:     req.Publish,
	})
	if err != nil {
		return err
	}

	return nil
}
