package challenge

import (
	"context"
	"database/sql"
	"errors"
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

func (s *service) GetCurrentMonthlyChallenge(ctx context.Context, req *spec.GetCurrentMonthlyChallengePayload) (*spec.SsmUserMonthlyChallenge, error) {

	q := s.monthlyQuery(ctx)
	models.MonthlyChallengeWhere.EndDate.GT(time.Now()).Apply(q)

	chall := &custommodels.UserMonthlyChall{}
	err := q.Bind(ctx, s.db, chall)
	if err == sql.ErrNoRows {
		return nil, spec.MakeNotFound(errors.New("no current monthly chall"))
	}
	if err != nil {
		s.log.Error("could not list monthly challs", zap.Error(err), utils.C(ctx))
		return nil, err
	}

	return convertMonthly(chall), nil
}

func (s *service) ListMonthlyChallenges(ctx context.Context, req *spec.ListMonthlyChallengesPayload) (spec.SsmUserMonthlyChallengeCollection, error) {

	q := s.monthlyQuery(ctx)
	challs := make([]*custommodels.UserMonthlyChall, 0)
	err := q.Bind(ctx, s.db, &challs)
	if err != nil {
		s.log.Error("could not list monthly challs", zap.Error(err), utils.C(ctx))
		return nil, err
	}

	res := make(spec.SsmUserMonthlyChallengeCollection, len(challs))
	for i, chall := range challs {
		res[i] = convertMonthly(chall)
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
		qm.Where("publish_at IS NULL OR publish_at < NOW()"),
		qm.Load(models.ChallengeRels.ChallengeFiles),
		qm.Load(models.ChallengeRels.ChallengeServices),
		qm.Load(models.ChallengeRels.Authors),
		qm.Load(qm.Rels(models.ChallengeRels.UserSolves, models.UserSolfRels.User), qm.Limit(5), qm.OrderBy(models.UserSolfColumns.CreatedAt)),
	)

	if auth.IsAuthed(ctx) {
		qm.Select(
			"EXISTS(SELECT 1 FROM user_solves us2 WHERE us2.challenge_id = c.id AND us2.user_id = '" + auth.GetUser(ctx).ID + "') AS solved",
		).Apply(q)
	}

	if req.Slug != nil {
		qm.Where("slug = ?", req.Slug).Apply(q)
	}

	if req.AuthorSlug != nil {
		// qm.Where("author_id = ?", req.AuthorSlug).Apply(q)
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
			CtfEventID:  chall.CTFEventID.Ptr(),
		}

		res[i].Files = make([]*spec.ChallengeFiles, len(chall.R.ChallengeFiles))
		for i2, file := range chall.R.ChallengeFiles {
			res[i].Files[i2] = &spec.ChallengeFiles{
				Filename: file.FriendlyName,
				URL:      file.URL,
			}
		}

		res[i].Authors = make([]*spec.Author, len(chall.R.Authors))
		for i2, v := range chall.R.Authors {
			res[i].Authors[i2] = &spec.Author{
				ID:          v.ID,
				FullName:    v.FullName,
				Description: v.Description,
				Sponsor:     v.Sponsor,
				Slug:        v.Slug,
				ImageURL:    v.ImageURL.Ptr(),
				Publish:     v.Publish,
			}
		}

		res[i].Solvers = make([]*spec.SsmSolver, len(chall.R.UserSolves))
		for i2, v := range chall.R.UserSolves {
			res[i].Solvers[i2] = &spec.SsmSolver{
				ID:       v.UserID,
				FullName: v.R.User.FullName,
				SolvedAt: v.CreatedAt.Unix(),
			}
		}

		res[i].Services = make([]*spec.ChallengeService, len(chall.R.ChallengeServices))
		for i2, v := range chall.R.ChallengeServices {
			res[i].Services[i2] = &spec.ChallengeService{
				UserDisplay: v.UserDisplay,
				Hyperlink:   v.Hyperlink,
			}
		}
	}

	return res, nil
}

func (s *service) ListAuthors(ctx context.Context, req *spec.ListAuthorsPayload) ([]*spec.Author, error) {
	authors, err := models.Authors(
		models.AuthorWhere.Publish.EQ(true),
	).All(ctx, s.db)
	if err != nil {
		return nil, err
	}

	res := make([]*spec.Author, len(authors))
	for i, v := range authors {
		res[i] = &spec.Author{
			ID:          v.ID,
			FullName:    v.FullName,
			Description: v.Description,
			Sponsor:     v.Sponsor,
			Slug:        v.Slug,
			Publish:     v.Publish,
			ImageURL:    v.ImageURL.Ptr(),
		}
	}

	return res, nil
}
