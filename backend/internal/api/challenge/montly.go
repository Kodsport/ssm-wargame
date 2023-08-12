package challenge

import (
	"context"
	"time"

	"github.com/sakerhetsm/ssm-wargame/internal/auth"
	"github.com/sakerhetsm/ssm-wargame/internal/custommodels"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *service) monthlyQuery(ctx context.Context) *queries.Query {
	q := models.NewQuery(
		qm.Select("monthly_challenges.*"),
		qm.Select("(SELECT COUNT(1) FROM user_solves WHERE challenge_id = monthly_challenges.challenge_id) num_solves"),
		qm.From("monthly_challenges"),
		qm.Load(qm.Rels(models.MonthlyChallengeRels.Challenge, models.ChallengeRels.ChallengeFiles)),
		qm.Load(qm.Rels(models.MonthlyChallengeRels.Challenge, models.ChallengeRels.Category)),
		qm.Load(qm.Rels(models.MonthlyChallengeRels.Challenge, models.ChallengeRels.Authors)),
		models.MonthlyChallengeWhere.StartDate.LT(time.Now()),
	)

	if auth.IsAuthed(ctx) {
		qm.Select(
			"EXISTS(SELECT 1 FROM user_solves us2 WHERE us2.challenge_id = monthly_challenges.challenge_id AND us2.user_id = '" + auth.GetUser(ctx).ID + "') AS solved",
		).Apply(q)
	}
	return q
}

func convertMonthly(chall *custommodels.UserMonthlyChall) *spec.SsmUserMonthlyChallenge {
	res := &spec.SsmUserMonthlyChallenge{
		ChallengeID:  chall.ChallengeID,
		DisplayMonth: chall.DisplayMonth,
		StartDate:    chall.StartDate.Unix(),
		EndDate:      chall.EndDate.Unix(),
	}

	res.Challenge = &spec.SsmChallenge{
		ID:          chall.R.Challenge.ID,
		Title:       chall.R.Challenge.Title,
		Description: chall.R.Challenge.Description,
		Slug:        chall.R.Challenge.Slug,
		Category:    chall.R.Challenge.R.Category.Name,
		Solves:      chall.NumSolves,
		Solved:      chall.Solved,
	}

	res.Challenge.Files = make([]*spec.ChallengeFiles, len(chall.R.Challenge.R.ChallengeFiles))
	for i, file := range chall.R.Challenge.R.ChallengeFiles {
		res.Challenge.Files[i] = &spec.ChallengeFiles{
			Filename: file.FriendlyName,
			URL:      file.URL,
		}
	}

	res.Challenge.Authors = make([]*spec.Author, len(chall.R.Challenge.R.Authors))
	for i, v := range chall.R.Challenge.R.Authors {
		res.Challenge.Authors[i] = &spec.Author{
			ID:          v.ID,
			FullName:    v.FullName,
			Description: v.Description,
			Sponsor:     v.Sponsor,
			Slug:        v.Slug,
			ImageURL:    v.ImageURL.Ptr(),
			Publish:     v.Publish,
		}
	}
	return res
}
