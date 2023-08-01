package challenge

import (
	"context"
	"database/sql"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sakerhetsm/ssm-wargame/internal/auth"
	"github.com/sakerhetsm/ssm-wargame/internal/custommodels"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

type service struct {
	spec.Auther
	db  *sql.DB
	log *zap.Logger
	s3  *s3.S3
}

func NewService(conn *sql.DB, log *zap.Logger, auther spec.Auther, s3c *s3.S3) spec.Service {
	return &service{
		Auther: auther,
		db:     conn,
		log:    log,
		s3:     s3c,
	}
}

func (s *service) ListMonthlyChallenges(ctx context.Context, req *spec.ListMonthlyChallengesPayload) (spec.SsmUsermonthlychallengesCollection, error) {
	challs := make([]*custommodels.UserMonthlyChall, 0)
	q := models.NewQuery(
		qm.Select("monthly_challenges.*"),
		qm.Select("(SELECT COUNT(1) FROM user_solves WHERE challenge_id = monthly_challenges.challenge_id) num_solves"),
		qm.From("monthly_challenges"),
		qm.Load(qm.Rels(models.MonthlyChallengeRels.Challenge, models.ChallengeRels.ChallengeFiles)),
		qm.Load(qm.Rels(models.MonthlyChallengeRels.Challenge, models.ChallengeRels.Category)),
		models.MonthlyChallengeWhere.StartDate.LT(time.Now()),
	)

	if auth.IsAuthed(ctx) {
		qm.Select(
			"EXISTS(SELECT 1 FROM user_solves us2 WHERE us2.challenge_id = monthly_challenges.challenge_id AND us2.user_id = '" + auth.GetUser(ctx).ID + "') AS solved",
		).Apply(q)
	}

	err := q.Bind(ctx, s.db, &challs)
	if err != nil {
		s.log.Error("could not list monthly challs", zap.Error(err), utils.C(ctx))
		return nil, err
	}

	res := make(spec.SsmUsermonthlychallengesCollection, len(challs))
	for i, chall := range challs {
		res[i] = &spec.SsmUsermonthlychallenges{
			ChallengeID:  chall.ChallengeID,
			DisplayMonth: chall.DisplayMonth,
			StartDate:    chall.StartDate.Unix(),
			EndDate:      chall.EndDate.Unix(),
		}

		res[i].Challenge = &spec.SsmChallenge{
			ID:          chall.R.Challenge.ID,
			Title:       chall.R.Challenge.Title,
			Description: chall.R.Challenge.Description,
			Slug:        chall.R.Challenge.Slug,
			Category:    chall.R.Challenge.R.Category.Name,
			Solves:      chall.NumSolves,
			Solved:      chall.Solved,
		}

		res[i].Challenge.Files = make([]*spec.ChallengeFiles, len(chall.R.Challenge.R.ChallengeFiles))
		for i2, file := range chall.R.Challenge.R.ChallengeFiles {
			res[i].Challenge.Files[i2] = s.signFile(ctx, file)
		}
	}

	return res, nil
}

func (s *service) ListChallenges(ctx context.Context, req *spec.ListChallengesPayload) (spec.SsmChallengeCollection, error) {

	challs := make([]*custommodels.UserChall, 0)
	q := models.NewQuery(
		qm.Select("c.*, categories.name as category"),
		qm.Select("(SELECT COUNT(1) FROM user_solves WHERE challenge_id = c.id) num_solves"),
		qm.From("challenges c"),
		qm.InnerJoin("categories ON categories.id = c.category_id"),
		qm.Load(models.ChallengeRels.ChallengeFiles),
		qm.Where("publish_at IS NULL OR publish_at < NOW()"),
	)

	if auth.IsAuthed(ctx) {
		qm.Select(
			"EXISTS(SELECT 1 FROM user_solves us2 WHERE us2.challenge_id = c.id AND us2.user_id = '" + auth.GetUser(ctx).ID + "') AS solved",
		).Apply(q)
	}

	err := q.Bind(ctx, s.db, &challs)

	if err != nil {
		s.log.Error("could not list challs", zap.Error(err), utils.C(ctx))
		return nil, err
	}
	res := make(spec.SsmChallengeCollection, len(challs))

	for i, chall := range challs {
		res[i] = &spec.SsmChallenge{
			ID:          chall.ID,
			Slug:        chall.Slug,
			Title:       chall.Title,
			Description: chall.Description,
			Score:       chall.Score,
			Solves:      chall.NumSolves,
			Solved:      chall.Solved,
			Category:    chall.Category,
		}

		res[i].Files = make([]*spec.ChallengeFiles, len(chall.R.ChallengeFiles))
		for i2, file := range chall.R.ChallengeFiles {
			res[i].Files[i2] = s.signFile(ctx, file)
		}
	}

	return res, nil
}

func (s *service) signFile(ctx context.Context, file *models.ChallengeFile) *spec.ChallengeFiles {

	req, _ := s.s3.GetObjectRequest(&s3.GetObjectInput{
		Bucket: &file.Bucket,
		Key:    &file.Key,
	})

	url, err := req.Presign(time.Hour * 1)
	if err != nil {
		s.log.Warn("could not sign url", zap.Error(err), utils.C(ctx))
	}

	return &spec.ChallengeFiles{
		Filename: file.FriendlyName,
		URL:      url,
	}
}
