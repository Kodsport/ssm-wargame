package skolverket

import (
	"context"
	_ "embed"
	"strings"

	"github.com/google/uuid"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

//go:embed unis.txt
var unisTxt string

func (i *Importer) ImportUnis() error {
	unis := strings.Split(unisTxt, "\n")

	for _, v := range unis {

		exists, err := models.Schools(
			models.SchoolWhere.Name.EQ(v),
		).Exists(context.Background(), i.db)
		if err != nil {
			return err
		}

		if exists {
			i.l.Info("school already exists", zap.String("school", v))
			continue
		}

		s := models.School{
			ID:           uuid.NewString(),
			Name:         v,
			IsUniversity: true,
		}

		err = s.Insert(context.TODO(), i.db, boil.Infer())
		if err != nil {
			i.l.Error("could not upsert", zap.Error(err), zap.String("school", v))
			return err
		}

	}
	return nil
}
