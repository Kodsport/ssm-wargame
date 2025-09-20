package custommodels

import (
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/volatiletech/null/v8"
)

type UserChall struct {
	models.Challenge `boil:",bind"`
	NumSolves        int  `boil:"num_solves"`
	NumSolvesInTeam  int  `boil:"num_solves_in_team"`
	NumTeamSolves    int  `boil:"num_team_solves"`
	SolvedInTeam     bool `boil:"solved_in_team"`

	Solved   bool   `boil:"solved"`
	Category string `boil:"category"`
}

type UserMonthlyChall struct {
	models.MonthlyChallenge `boil:",bind"`
	NumSolves               int    `boil:"num_solves"`
	Solved                  bool   `boil:"solved"`
	Category                string `boil:"category"`
}

type Course struct {
	models.Course `boil:",bind"`
	Finished      null.Bool `boil:"finished"`
}
