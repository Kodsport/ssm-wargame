package admin

import (
	"context"
	"database/sql"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/lib/pq"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

var slugRegex = regexp.MustCompile(`[^a-zA-Z\-0-9]`)

func (s *service) ChalltoolsImport(ctx context.Context, req *spec.ChalltoolsImportPayload) error {

	s.log.Info("beginning import", zap.String("id", req.ChallengeID))
	tx, err := s.db.Begin()
	if err != nil {
		s.log.Error("could not begin tx", zap.Error(err))
		return err
	}
	defer tx.Rollback()

	// category

	categoryID := ""
	{
		x := []string{}
		for _, v := range req.Categories {
			x = append(x, "%"+v+"%")
		}
		cat, err := models.Categories(
			qm.Where("name ILIKE ANY (?)", pq.Array(x)),
		).One(ctx, tx)

		if err == sql.ErrNoRows {
			categoryID = uuid.NewString()
			cat = &models.Category{
				ID:   categoryID,
				Name: req.Categories[0],
			}
			err = cat.Insert(ctx, tx, boil.Infer())
		}
		if err != nil {
			s.log.Error("could not something", zap.Error(err))
			return err
		}

		categoryID = cat.ID
	}

	score := 0
	if req.Score != nil {
		score = *req.Score
	}

	slug := strings.ReplaceAll(req.Title, " ", "_")
	slug = slugRegex.ReplaceAllString(slug, "")
	chall := models.Challenge{
		ID:          req.ChallengeID,
		Slug:        slug,
		Title:       req.Title,
		Description: req.Description,
		Score:       score,
		CategoryID:  categoryID,
	}

	err = chall.Upsert(ctx, tx, true, []string{}, boil.Infer(), boil.Infer())
	if err != nil {
		s.log.Error("could not upsert chall", zap.Error(err))
		return err
	}

	// files
	{

		_, err = models.ChallengeFiles(
			models.ChallengeFileWhere.ChallengeID.EQ(null.StringFrom(req.ChallengeID)),
		).DeleteAll(ctx, tx)
		if err != nil {
			s.log.Error("could delete flags", zap.Error(err))
			return err
		}

		for _, v := range req.FileUrls {
			parts := strings.Split(v, "/")
			file := parts[len(parts)-1]
			f := models.ChallengeFile{
				ID:           uuid.NewString(),
				ChallengeID:  null.StringFrom(req.ChallengeID),
				FriendlyName: file,
				URL:          v,
			}
			err = f.Insert(ctx, tx, boil.Infer())
			if err != nil {
				s.log.Error("could not insert file", zap.Error(err))
				return err
			}
		}
	}

	// flags
	{
		_, err = models.Flags(
			models.FlagWhere.ChallengeID.EQ(req.ChallengeID),
		).DeleteAll(ctx, tx)
		if err != nil {
			s.log.Error("could delete flags", zap.Error(err))
			return err
		}

		for _, v := range req.Flags {

			f := models.Flag{
				ID:          uuid.NewString(),
				ChallengeID: req.ChallengeID,
				Flag:        v.Flag,
				FlagPrefix:  strDefault(req.FlagFormatPrefix),
				FlagSuffix:  strDefault(req.FlagFormatSuffix),
			}
			err := f.Insert(ctx, tx, boil.Infer())
			if err != nil {
				s.log.Error("could not insert flag", zap.Error(err))
				return err
			}
		}
	}

	// authors
	{
		_, err = tx.ExecContext(ctx, "DELETE FROM challenge_authors WHERE challenge_id = $1", req.ChallengeID)
		if err != nil {
			s.log.Error("could delete authors", zap.Error(err))
			return err
		}

		otherAuthors := []string{}
		for _, v := range req.Authors {
			user, err := models.Users(
				qm.Where("full_name ILIKE ?", "%"+v+"%"),
			).One(ctx, tx)
			if err != nil {
				otherAuthors = append(otherAuthors, v)
				continue
			}
			_, err = tx.ExecContext(ctx, "INSERT INTO challenge_authors (challenge_id, user_id) VALUES ($1, $2)", req.ChallengeID, user.ID)
			if err != nil {
				s.log.Error("could not insert chall author", zap.Error(err))
				return err
			}
		}

		if len(otherAuthors) != 0 {
			_, err = models.Challenges(
				qm.Where("id = ?", req.ChallengeID),
			).UpdateAll(ctx, tx, models.M{
				models.ChallengeColumns.OtherAuthors: pq.Array(otherAuthors),
			})
			if err != nil {
				s.log.Error("could not add other authors", zap.Error(err))
				return err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		s.log.Error("could not commit", zap.Error(err))
		return err
	}

	return nil
}

func strDefault(x *string) string {
	if x == nil {
		return ""
	}
	return *x
}
