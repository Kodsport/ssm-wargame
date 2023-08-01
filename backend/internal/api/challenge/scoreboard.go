package challenge

import (
	"context"

	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

func (s *service) SchoolScoreboard(ctx context.Context, req *spec.SchoolScoreboardPayload) (*spec.SsmShoolscoreboard, error) {

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

	return &spec.SsmShoolscoreboard{
		Scores: res,
	}, nil
}
