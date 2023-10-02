package challenge

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/sakerhetsm/ssm-wargame/internal/auth"
	"github.com/sakerhetsm/ssm-wargame/internal/custommodels"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *service) EnrollCourse(ctx context.Context, req *spec.EnrollCoursePayload) error {
	enroll := models.CourseEnrollment{
		ID:       uuid.NewString(),
		CourseID: req.ID,
		UserID:   auth.GetUser(ctx).ID,
		Finished: false,
	}
	return enroll.Insert(ctx, s.db, boil.Infer())
}

func (s *service) CompleteCourse(ctx context.Context, req *spec.CompleteCoursePayload) error {

	n, err := models.CourseEnrollments(
		models.CourseEnrollmentWhere.UserID.EQ(auth.GetUser(ctx).ID),
		models.CourseEnrollmentWhere.CourseID.EQ(req.ID),
		models.CourseEnrollmentWhere.Finished.EQ(false),
	).UpdateAll(ctx, s.db, models.M{
		models.CourseEnrollmentColumns.Finished:   true,
		models.CourseEnrollmentColumns.FinishedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	if n == 0 {
		return spec.MakeNotFound(errors.New("not valid enrollment"))
	}

	return nil
}

func (s *service) ListCourses(ctx context.Context, req *spec.ListCoursesPayload) ([]*spec.Course, error) {

	q := models.NewQuery(
		qm.Select("courses.*"),
		models.CourseWhere.Publish.EQ(true),
		qm.Load(models.CourseRels.Authors),
		qm.Load(models.CourseRels.CourseItems),
		qm.From(models.TableNames.Courses),
	)

	if auth.IsAuthed(ctx) {
		qm.Select("(SELECT finished FROM course_enrollments WHERE course_id = courses.id AND user_id = '" + auth.GetUser(ctx).ID + "') AS finished").Apply(q)

	}

	courses := []*custommodels.Course{}
	err := q.Bind(ctx, s.db, &courses)
	if err != nil {
		return nil, err
	}

	res := make([]*spec.Course, len(courses))
	for i, v := range courses {
		res[i] = &spec.Course{
			ID:          v.ID,
			Title:       v.Title,
			Slug:        v.Slug,
			Category:    v.Category,
			Difficulty:  v.Difficulty,
			Description: v.Description,
			Enrolled:    v.Finished.Valid,
			Completed:   v.Finished.Bool,
			Authors:     make([]*spec.Author, len(v.R.Authors)),
			CourseItems: make([]*spec.CourseItem, len(v.R.CourseItems)),
		}

		for i2, author := range v.R.Authors {
			res[i].Authors[i2] = &spec.Author{
				ID:          author.ID,
				FullName:    author.FullName,
				Description: author.Description,
				Sponsor:     author.Sponsor,
				Slug:        author.Slug,
				Publish:     author.Publish,
				ImageURL:    author.ImageURL.Ptr(),
			}
		}

		for i2, ci := range v.R.CourseItems {
			res[i].CourseItems[i2] = &spec.CourseItem{
				ID:          ci.ID,
				ChallengeID: ci.ChallengeID,
				Position:    ci.Position,
			}
		}
	}

	return res, nil
}
