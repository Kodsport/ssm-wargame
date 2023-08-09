package challenge

import (
	"context"

	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

func (s *service) SchoolScoreboard(ctx context.Context, req *spec.SchoolScoreboardPayload) (*spec.SsmSchoolScoreboard, error) {

	type SchoolScoreboardScore struct {
		Score      int    `boil:"score"`
		SchoolName string `boil:"name"`
	}

	scores := make([]*SchoolScoreboardScore, 0, 50)
	err := models.NewQuery(
		qm.Select("s.name name, SUM(score) score"),
		qm.From("user_solves us"),
		qm.InnerJoin("users u ON u.id = us.user_id AND u.school_id IS NOT NULL"),
		qm.InnerJoin("challenges c ON c.id = us.challenge_id"),
		qm.InnerJoin("schools s ON s.id = u.school_id"),
		qm.GroupBy("s.name"),
		qm.OrderBy("score DESC"),
		qm.Limit(10),
	).Bind(ctx, s.db, &scores)

	if err != nil {
		s.log.Error("could not get scoreboard", zap.Error(err))
		return nil, err
	}

	res := make([]*spec.SchoolScoreboardScore, len(scores))
	for i, sss := range scores {
		res[i] = &spec.SchoolScoreboardScore{
			Score:      sss.Score,
			SchoolName: sss.SchoolName,
		}
	}

	return &spec.SsmSchoolScoreboard{
		Scores: res,
	}, nil
}

func (s *service) UserScoreboard(ctx context.Context, req *spec.UserScoreboardPayload) (*spec.SsmUserScoreboard, error) {

	type SchoolScoreboardScore struct {
		UserID string `boild:"user_id"`
		Score  int    `boil:"score"`
		Name   string `boil:"name"`
	}

	scores := make([]*SchoolScoreboardScore, 0, 50)
	err := models.NewQuery(
		qm.Select("users.id AS user_id, users.full_name AS name, SUM(score) AS score"),
		qm.From("user_solves"),
		qm.LeftOuterJoin("users ON users.id = user_id"),
		qm.InnerJoin("challenges ON challenges.id = challenge_id"),
		qm.GroupBy("users.id"),
		qm.OrderBy("score DESC"),
		qm.Limit(10),
	).Bind(ctx, s.db, &scores)

	if err != nil {
		s.log.Error("could not get scoreboard", zap.Error(err))
		return nil, err
	}

	res := make([]*spec.UserScoreboardScore, len(scores))
	for i, sss := range scores {
		res[i] = &spec.UserScoreboardScore{
			UserID: sss.UserID,
			Score:  sss.Score,
			Name:   sss.Name,
		}
	}

	return &spec.SsmUserScoreboard{
		Scores: res,
	}, nil
}
