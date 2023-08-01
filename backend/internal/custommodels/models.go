package custommodels

import "github.com/sakerhetsm/ssm-wargame/internal/models"

type UserChall struct {
	models.Challenge `boil:",bind"`
	NumSolves        int    `boil:"num_solves"`
	Solved           bool   `boil:"solved"`
	Category         string `boil:"category"`
}
