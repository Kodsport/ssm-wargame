package challenge

import (
	"context"
	"math"

	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

var maxScores = 10

func (s *service) challScores(ctx context.Context) (map[string]float64, error) {
	type ChallSolves struct {
		ChallengeID string   `boil:"challenge_id"`
		Solves      int      `boil:"solves"`
		StaticScore null.Int `boil:"static_score"`
	}

	challUserSolves := []*ChallSolves{}
	err := models.NewQuery(
		qm.SQL("SELECT challenge_id, count(1) as solves, (SELECT static_score FROM challenges WHERE id = challenge_id) FROM user_solves GROUP BY challenge_id"),
	).Bind(ctx, s.db, &challUserSolves)
	if err != nil {
		return nil, err
	}

	challScores := make(map[string]float64)

	for _, r := range challUserSolves {
		if r.StaticScore.Valid {
			challScores[r.ChallengeID] = float64(r.StaticScore.Int)
		} else {
			challScores[r.ChallengeID] = float64(dynamicScore(500, 100, float64(r.Solves)))
		}
	}

	return challScores, nil
}

func (s *service) SchoolScoreboard(ctx context.Context, req *spec.SchoolScoreboardPayload) (*spec.SsmSchoolScoreboard, error) {

	type SchoolSolve struct {
		ChallengeID  string `boil:"challenge_id"`
		SchoolID     string `boil:"school_id"`
		SchoolName   string `boil:"name"`
		Solves       int    `boil:"solves"`
		IsUniversity bool   `boil:"is_university"`
	}

	challScores, err := s.challScores(ctx)
	if err != nil {
		return nil, err
	}

	schoolSolveCount := []*SchoolSolve{}
	err = models.NewQuery(
		qm.SQL("SELECT challenge_id, schools.id, count(1) as solves, name, is_university FROM school_solves INNER JOIN schools ON id = school_id GROUP BY schools.id, challenge_id"),
	).Bind(ctx, s.db, &schoolSolveCount)
	if err != nil {
		return nil, err
	}

	schoolScore := func(solvers, challScore float64) float64 {
		return challScore * (1.5 - math.Pow(0.5, solvers-1)/2)
	}

	schoolScores := make(map[string]*spec.SchoolScoreboardScore)
	for _, r := range schoolSolveCount {
		if _, ok := schoolScores[r.SchoolID]; !ok {
			schoolScores[r.SchoolID] = &spec.SchoolScoreboardScore{
				SchoolName:   r.SchoolName,
				Score:        0,
				IsUniversity: r.IsUniversity,
			}
		}

		schoolScores[r.SchoolID].Score += int(schoolScore(float64(r.Solves), float64(challScores[r.ChallengeID])))
	}

	res := make([]*spec.SchoolScoreboardScore, 0, len(schoolScores))
	for _, score := range schoolScores {
		res = append(res, score)
	}

	return &spec.SsmSchoolScoreboard{
		Scores: res,
	}, nil
}

func (s *service) UserScoreboard(ctx context.Context, req *spec.UserScoreboardPayload) (*spec.SsmUserScoreboard, error) {

	challScores, err := s.challScores(ctx)
	if err != nil {
		return nil, err
	}

	type SchoolScoreboardScore struct {
		UserID string `boild:"user_id"`
		Score  int    `boil:"score"`
		Name   string `boil:"name"`
	}

	users, err := models.Users(
		qm.Load(models.UserRels.UserSolves),
	).All(ctx, s.db)

	if err != nil {
		s.log.Error("could not get scoreboard", zap.Error(err))
		return nil, err
	}

	res := make([]*spec.UserScoreboardScore, len(users))
	for i, v := range users {
		res[i] = &spec.UserScoreboardScore{
			UserID: v.ID,
			Name:   v.FullName,
			Score:  0,
		}

		for _, v2 := range v.R.UserSolves {
			res[i].Score += int(challScores[v2.ChallengeID])
		}
	}

	return &spec.SsmUserScoreboard{
		Scores: res,
	}, nil
}
