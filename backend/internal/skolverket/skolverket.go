package skolverket

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

type Importer struct {
	l  *zap.Logger
	db *sql.DB
}

func New(db *sql.DB, l *zap.Logger) *Importer {
	return &Importer{
		l:  l,
		db: db,
	}
}

func (i *Importer) Import() error {

	i.l.Info("starting importing job")

	resp, err := http.DefaultClient.Get("https://api.skolverket.se/skolenhetsregistret/v1/skolenhet")
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		i.l.Warn("could not get units", zap.String("status", resp.Status))
		return errors.New("not 200")
	}

	var units MiniSchoolUnitResp
	err = json.NewDecoder(resp.Body).Decode(&units)
	if err != nil {
		i.l.Error("could not parse mini units", zap.Error(err))
		return err
	}

	for _, mu := range units.SchoolUnits {
		time.Sleep(time.Second) // be kind to skolverkets api

		id, err := strconv.Atoi(mu.SchoolUnitCode)
		if err != nil {
			i.l.Warn("not int", zap.String("id", mu.SchoolUnitCode))
			continue
		}

		l := i.l.With(zap.Int("school_id", id))

		l.Info("importing school")
		resp, err := http.DefaultClient.Get("https://api.skolverket.se/skolenhetsregistret/v1/skolenhet/" + mu.SchoolUnitCode)
		if err != nil {
			l.Warn("could not get school data", zap.Error(err))
			continue
		}

		if resp.StatusCode != 200 {
			l.Warn("could not get school data", zap.String("status", resp.Status))
			continue
		}

		var unit Unit
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			l.Warn("could not read data", zap.Error(err))
			continue
		}
		err = json.Unmarshal(data, &unit)
		if err != nil {
			l.Warn("could not unmarshal data", zap.Error(err))
			continue
		}

		longStr := unit.SkolenhetInfo.Besoksadress.GeoData.KoordinatWGS84Lng
		long, err := strconv.ParseFloat(longStr, 32)
		if err != nil {
			l.Warn("could not parse long", zap.Error(err), zap.String("long", longStr))
			continue
		}

		latStr := unit.SkolenhetInfo.Besoksadress.GeoData.KoordinatWGS84Lat
		lat, err := strconv.ParseFloat(latStr, 32)
		if err != nil {
			l.Warn("could not parse lat", zap.Error(err), zap.String("lat", latStr))
			continue
		}

		/*
			school types

			Central
			Forskoleklass
			Fritidshem
			Grundsameskola
			Grundsarskola
			Grundskola
			Gymnasiesarskola
			Gymnasieskola
			Komvux
			Oppenfritidsverksamhet
			Sarvux
			Sfi
			Specialskola
		*/

		isHighSchool := false
		isElementarySchool := false
		for _, v := range unit.SkolenhetInfo.Skolformer {
			if v.Type == "Grundskola" || v.Type == "Grundsarskola" || v.Type == "Grundsameskola" {
				isElementarySchool = true
			}
			if v.Type == "Gymnasieskola" || v.Type == "Gymnasiesarskola" {
				isHighSchool = true
			}

		}

		school := models.School{
			ID:                   id,
			Name:                 unit.SkolenhetInfo.Namn,
			GeographicalAreaCode: unit.SkolenhetInfo.Besoksadress.Postnr,
			Status:               unit.SkolenhetInfo.Status,
			Latitude:             lat,
			Longitude:            long,
			IsHighSchool:         isHighSchool,
			IsElementarySchool:   isElementarySchool,
			RawSkolverketData:    data,
		}

		err = school.Upsert(context.Background(), i.db, true, []string{"id"}, boil.Infer(), boil.Infer())
		if err != nil {
			l.Error("could not upsert school", zap.Error(err), zap.Any("school", school))
			continue
		}
	}

	return nil
}
