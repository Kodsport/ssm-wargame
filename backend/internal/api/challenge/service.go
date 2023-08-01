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

func (s *service) ListMonthlyChallenges(ctx context.Context, req *spec.ListMonthlyChallengesPayload) ([]*spec.MonthlyChallenge, error) {
	challs, err := models.MonthlyChallenges(
		qm.Load(models.MonthlyChallengeRels.Challenge),
	).All(ctx, s.db)
	if err != nil {
		s.log.Error("could not list monthly challs", zap.Error(err), utils.C(ctx))
		return nil, err
	}

	res := make([]*spec.MonthlyChallenge, len(challs))
	for i, chall := range challs {
		res[i] = &spec.MonthlyChallenge{
			ChallengeID:  chall.ChallengeID,
			DisplayMonth: chall.DisplayMonth,
			StartDate:    chall.StartDate.Format("2006-01-02"),
			EndDate:      chall.EndDate.Format("2006-01-02"),
			Slug:         chall.R.Challenge.Slug,
			Title:        chall.R.Challenge.Title,
			Description:  chall.R.Challenge.Description,
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

			req, _ := s.s3.GetObjectRequest(&s3.GetObjectInput{
				Bucket: &file.Bucket,
				Key:    &file.Key,
			})

			url, err := req.Presign(time.Hour * 1)
			if err != nil {
				s.log.Warn("could not sign url", zap.Error(err), utils.C(ctx))
			}

			res[i].Files[i2] = &spec.ChallengeFiles{
				Filename: file.FriendlyName,
				URL:      url,
			}
		}
	}

	return res, nil
}
