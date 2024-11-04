package challenge

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

func (s *service) KnackKodenSubmitFlag(ctx context.Context, req *spec.KnackKodenSubmitFlagPayload) error {

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback() //nolint:errcheck

	challID := uuid.MustParse(req.ChallengeID)

	exits, err := models.Challenges(
		models.ChallengeWhere.ID.EQ(req.ChallengeID),
		models.ChallengeWhere.ChallNamespace.EQ(null.StringFrom("knackkoden")),
	).Exists(ctx, tx)

	if err != nil {
		s.log.Error("could not check chall exists", zap.Error(err))
		return err
	}

	if !exits {
		return errors.New("challenge not found")
	}

	team, err := models.KnackKodenTeams(
		models.KnackKodenTeamWhere.Password.EQ(req.Password),
	).One(ctx, tx)

	if err != nil {
		return err
	}

	solved, err := models.KnackKodenSolves(
		models.KnackKodenSolfWhere.KnackKodenTeamID.EQ(team.ID),
		models.KnackKodenSolfWhere.ChallengeID.EQ(challID.String()),
	).Exists(ctx, tx)

	if err != nil {
		return err
	}

	if solved {
		return spec.MakeAlreadySolved(errors.New("already solved"))
	}

	flags, err := models.Flags(
		models.FlagWhere.ChallengeID.EQ(challID.String()),
	).All(ctx, tx)
	if err != nil {
		s.log.Error("could not get flags", zap.Error(err), utils.C(ctx))
		return err
	}

	flagCorrect := false
	for _, flag := range flags {

		inp := strings.TrimPrefix(strings.TrimSuffix(req.Flag, flag.FlagSuffix), flag.FlagPrefix)

		if flag.Type == "regex" {
			if regexp.MustCompile(flag.Flag).Match([]byte(inp)) {
				flagCorrect = true
				break
			}
			continue
		}

		if flag.Flag == inp {
			flagCorrect = true
			break
		}
	}

	if flagCorrect {
		solve := models.KnackKodenSolf{
			KnackKodenTeamID: team.ID,
			ChallengeID:      challID.String(),
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

func (s *service) KnackKodenScoreboard(ctx context.Context, req *spec.KnackKodenScoreboardPayload) (*spec.SsmSchoolScoreboard, error) {

	query := `
	select 
		knack_koden_teams.id AS team_id,
		knack_koden_teams.school_name AS school_name,
		knack_koden_teams.class_name AS class_name,
		SUM(challenges.static_score) AS score
	from knack_koden_solves 
	JOIN challenges ON challenges.id = knack_koden_solves.challenge_id AND challenges.chall_namespace = 'knackkoden'
	JOIN knack_koden_teams ON knack_koden_teams.id = knack_koden_solves.knack_koden_team_id
	GROUP BY knack_koden_teams.id
	ORDER BY score DESC;
	`

	type Row struct {
		ChallengeID string `boil:"team_id"`
		SchoolName  string `boil:"school_name"`
		ClassName   string `boil:"class_name"`
		Score       int    `boil:"score"`
	}

	rows := []*Row{}
	err := models.NewQuery(
		qm.SQL(query),
	).Bind(ctx, s.db, &rows)

	if err != nil {
		return nil, err
	}

	res := &spec.SsmSchoolScoreboard{Scores: make([]*spec.SchoolScoreboardScore, len(rows))}

	for i, r := range rows {
		res.Scores[i] = &spec.SchoolScoreboardScore{
			Score:      r.Score,
			SchoolName: fmt.Sprintf("%s (%s)", r.SchoolName, r.ClassName),
		}

	}
	return res, nil
}

func (s *service) KnackKodenRegisterClass(ctx context.Context, req *spec.KnackKodenRegisterClassPayload) (*spec.KnackKodenRegisterClassResult, error) {

	httpReq := ctx.Value("ssm_http_req").(*http.Request)

	password := randpasswd(10)

	team := &models.KnackKodenTeam{
		ID:              uuid.NewString(),
		TeacherFullName: req.TeacherFullName,
		TeacherEmail:    req.TeacherEmail,
		TeacherPhonenr:  req.TeacherPhonenr,
		SchoolName:      req.SchoolName,
		ClassName:       req.ClassName,
		PostalCode:      req.PostalCode,
		Password:        password,
		UserAgent:       httpReq.Header.Get("user-agent"),
		IPAddr:          httpReq.RemoteAddr,
	}

	err := team.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		s.log.Error("could not insert knack koden team", zap.Error(err))
		return nil, err
	}

	return &spec.KnackKodenRegisterClassResult{
		Password: password,
	}, nil

}

func (s *service) KnackKodenGetClass(ctx context.Context, req *spec.KnackKodenGetClassPayload) (*spec.KnackKodenGetClassResult, error) {

	team, err := models.KnackKodenTeams(
		models.KnackKodenTeamWhere.Password.EQ(req.Password),
		qm.Load(models.KnackKodenTeamRels.KnackKodenSolves),
	).One(ctx, s.db)

	if err != nil {
		return nil, err
	}

	solves := []string{}
	for _, s := range team.R.KnackKodenSolves {
		solves = append(solves, s.ChallengeID)

	}

	return &spec.KnackKodenGetClassResult{
		ClassName: fmt.Sprintf("%s (%s)", team.SchoolName, team.ClassName),
		Solves:    solves,
	}, nil

}

var letterRunes = []rune("23456789ABCDEFGHIJKLMNOPQRSTUVWXYZ1")

func randpasswd(n int) string {
	rand.Seed(time.Now().Unix()) // wtf, fick cykel efter 62 lösen
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
