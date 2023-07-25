package custommodels

import "github.com/sakerhetsm/ssm-wargame/internal/models"

type ChallWithSovles struct {
	models.Challenge `boil:",bind"`
	NumSolves        int `boil:"num_solves"`
}
