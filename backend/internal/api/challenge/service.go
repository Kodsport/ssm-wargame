package challenge

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/sakerhetsm/ssm-wargame/internal/auth"
	"github.com/sakerhetsm/ssm-wargame/internal/custommodels"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
	// todo(mosu) code duplication with admin
	challs := make([]*custommodels.ChallWithSovles, 0)
	err := models.NewQuery(
		qm.Select("c.*, COUNT(us.user_id) num_solves"),
		qm.From("challenges c"),
		qm.LeftOuterJoin("user_solves us ON us.challenge_id = c.id"),
		qm.GroupBy("c.id"),
		qm.Load(models.ChallengeRels.ChallengeFiles),
		qm.Where("publish_at IS NULL OR publish_at < NOW()"),
	).Bind(ctx, s.db, &challs)

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

func (s *service) SubmitFlag(ctx context.Context, req *spec.SubmitFlagPayload) error {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback() //nolint:errcheck

	user := auth.GetUser(ctx)
	challID := uuid.MustParse(req.ChallengeID)

	solved, err := models.UserSolves(
		models.UserSolfWhere.UserID.EQ(user.ID),
		models.UserSolfWhere.ChallengeID.EQ(challID.String()),
	).Exists(ctx, tx)

	if err != nil {
		return err
	}

	if solved {
		return spec.MakeAlreadySolved(errors.New("already solved"))
	}

	flagCorrect, err := models.Flags(
		models.FlagWhere.ChallengeID.EQ(challID.String()),
		models.FlagWhere.Flag.EQ(req.Flag),
	).Exists(ctx, tx)

	if err != nil {
		s.log.Error("FlagExists failed", zap.Error(err), utils.C(ctx))
		return err
	}

	attempt := models.Submission{
		ID:          uuid.New().String(),
		UserID:      user.ID,
		ChallengeID: challID.String(),
		Successful:  flagCorrect,
		Input:       req.Flag,
	}
	err = attempt.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return err
	}

	if flagCorrect {
		solve := models.UserSolf{
			UserID:      user.ID,
			ChallengeID: challID.String(),
		}
		err = solve.Insert(ctx, tx, boil.Infer())
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		s.log.Error("could not commit", zap.Error(err), utils.C(ctx))
		return err
	}

	if !flagCorrect {
		return spec.MakeIncorrectFlag(errors.New("incorrect flag"))
	}

	return nil
}
