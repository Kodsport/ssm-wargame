package admin

import (
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) authImport(ctx context.Context, token string) (*models.ChalltoolsImportToken, error) {

	id := strings.Split(token, "_")[1]
	idBytes, err := hex.DecodeString(id)
	if err != nil {
		return nil, err
	}
	uuidID, err := uuid.FromBytes(idBytes)
	if err != nil {
		return nil, err
	}
	t, err := models.ChalltoolsImportTokens(
		models.ChalltoolsImportTokenWhere.ID.EQ(uuidID.String()),
		models.ChalltoolsImportTokenWhere.ExpiresAt.GT(time.Now()),
	).One(ctx, s.db)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(t.Token), []byte(token)); err != nil {
		return nil, err
	}
	return t, nil
}

func (s *service) ChalltoolsImport(ctx context.Context, req *spec.ChalltoolsImportPayload) error {

	token, err := s.authImport(ctx, req.ImportToken)
	if err != nil {
		fmt.Println(err.Error())

		return err
	}

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
		} else if err != nil {
			s.log.Error("could not something", zap.Error(err))
			return err
		}

		categoryID = cat.ID
	}

	score := 0
	if req.Score != nil {
		score = *req.Score
	}

	chall := models.Challenge{
		ID:          req.ChallengeID,
		Slug:        utils.Slugify(req.Title),
		Title:       req.Title,
		Description: req.Description,
		Score:       score,
		CategoryID:  categoryID,
		CTFEventID:  token.CTFEventID,
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
				Type:        v.Type,
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

		for _, v := range req.Authors {
			author, err := models.Authors(
				qm.Where("full_name ILIKE ?", "%"+v+"%"),
			).One(ctx, tx)

			if err == sql.ErrNoRows {
				author = &models.Author{
					ID:       uuid.NewString(),
					Slug:     utils.Slugify(v),
					FullName: v,
					Sponsor:  false,
					Publish:  false,
				}

				err = author.Insert(ctx, tx, boil.Infer())
			}

			if err != nil {
				return err
			}

			_, err = tx.ExecContext(ctx, "INSERT INTO challenge_authors (challenge_id, author_id) VALUES ($1, $2)", req.ChallengeID, author.ID)
			if err != nil {
				s.log.Error("could not insert chall author", zap.Error(err))
				return err
			}
		}

	}

	// services
	{
		_, err = models.ChallengeServices(
			models.ChallengeServiceWhere.ChallengeID.EQ(req.ChallengeID),
		).DeleteAll(ctx, tx)
		if err != nil {
			s.log.Error("could delete services", zap.Error(err))
			return err
		}

		for _, v := range req.Services {
			f := models.ChallengeService{
				ID:          uuid.NewString(),
				ChallengeID: req.ChallengeID,
				UserDisplay: v.UserDisplay,
				Hyperlink:   v.Hyperlink,
			}
			err = f.Insert(ctx, tx, boil.Infer())
			if err != nil {
				s.log.Error("could not insert service", zap.Error(err))
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
