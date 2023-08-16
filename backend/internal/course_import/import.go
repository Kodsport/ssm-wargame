package courseimport

import (
	"context"
	"database/sql"
	"io/fs"
	"os"

	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Importer struct {
	db *sql.DB
	l  *zap.Logger
}

func New(db *sql.DB, l *zap.Logger) *Importer {
	return &Importer{
		db: db,
		l:  l,
	}
}

func (i *Importer) ParseCourse(repoPath string) error {
	ctx := context.Background()

	files, err := fs.Glob(os.DirFS(repoPath), "*/*/course.yml")
	if err != nil {
		return err
	}

	if len(files) == 0 {
		i.l.Warn("no files")
		return nil
	}

	for _, v := range files {

		i.l.Info("reading course file", zap.String("file", v))
		f, err := os.OpenFile(repoPath+v, 0o600, fs.FileMode(os.O_RDONLY))
		if err != nil {
			return err
		}

		var c course
		err = yaml.NewDecoder(f).Decode(&c)
		if err != nil {
			return err
		}

		tx, err := i.db.Begin()
		if err != nil {
			return err
		}

		course := models.Course{
			ID:          c.ID,
			Slug:        c.Slug,
			Category:    c.Category,
			Difficulty:  c.Difficulty,
			Description: c.Description,
			Publish:     c.Publish,
			Title:       c.Title,
		}
		err = course.Upsert(ctx, tx, true, []string{models.CourseColumns.ID}, boil.Infer(), boil.Infer())
		if err != nil {
			return err
		}

		ciIds := []string{}
		for _, ci := range c.CoruseItems {
			ciIds = append(ciIds, ci.ID)

		}
		_, err = models.CourseItems(
			models.CourseItemWhere.CourseID.EQ(c.ID),
			models.CourseItemWhere.ID.NIN(ciIds),
		).DeleteAll(ctx, tx)

		for i, ci := range c.CoruseItems {
			x := models.CourseItem{
				ID:          ci.ID,
				Position:    i,
				CourseID:    c.ID,
				ChallengeID: ci.ChallengeID,
			}

			err = x.Upsert(ctx, tx, true, []string{"id"}, boil.Infer(), boil.Infer())
			if err != nil {
				return err
			}
		}

		err = tx.Commit()
		if err != nil {
			return err
		}

	}

	return nil
}
