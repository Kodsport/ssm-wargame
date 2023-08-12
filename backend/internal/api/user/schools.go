package user

import (
	"context"
	"errors"

	"github.com/sakerhetsm/ssm-wargame/internal/auth"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/user"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

func (s *service) SearchSchools(ctx context.Context, req *spec.SearchSchoolsPayload) ([]*spec.School, error) {
	q := models.Schools(
		qm.Where("name ILIKE ?", "%"+req.Q+"%"),
		qm.Limit(10),
	)

	if req.University != nil {
		qm.Where("is_university = ?", req.University).Apply(q.Query)
	}

	schools, err := q.All(ctx, s.db)

	if err != nil {
		s.log.Error("could not search schools", zap.Error(err), zap.String("q", req.Q))
		return nil, err
	}

	res := make([]*spec.School, len(schools))
	for i, v := range schools {
		res[i] = &spec.School{
			ID:               v.ID,
			Name:             v.Name,
			MunicipalityName: v.MunicipalityName,
			IsUniversity:     v.IsUniversity,
		}
	}

	return res, nil
}

func (s *service) JoinSchool(ctx context.Context, req *spec.JoinSchoolPayload) error {

	// maybe notify others in the school etc, therefore it's own endpoint

	user := auth.GetUser(ctx)

	if user.SchoolID.Valid {
		return errors.New("already in a school")
	}

	_, err := models.Users(
		models.UserWhere.ID.EQ(user.ID),
	).UpdateAll(ctx, s.db, models.M{
		models.UserColumns.SchoolID: req.SchoolID,
	})

	if err != nil {
		s.log.Error("could not join school", zap.Error(err))
		return err
	}

	return nil
}

func (s *service) LeaveSchool(ctx context.Context, req *spec.LeaveSchoolPayload) error {

	user := auth.GetUser(ctx)

	if !user.SchoolID.Valid {
		return errors.New("no school to leave")
	}

	_, err := models.Users(
		models.UserWhere.ID.EQ(user.ID),
	).UpdateAll(ctx, s.db, models.M{
		models.UserColumns.SchoolID: nil,
	})

	if err != nil {
		s.log.Error("could not leave school", zap.Error(err))
		return err
	}

	return nil
}
