package ctf

import (
	"context"
	"database/sql"
	"errors"
	"math"
	"math/rand"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sakerhetsm/ssm-wargame/internal/custommodels"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/ctf"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Get(ctx context.Context, req *spec.GetPayload) (*spec.CTFInfo, error) {
	ctf, err := models.CTFS(models.CTFWhere.Slug.EQ(req.Slug)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}

	// manuel query för att få theme
	var theme sql.NullString
	err = s.db.QueryRowContext(ctx, "SELECT theme FROM ctfs WHERE slug = $1", req.Slug).Scan(&theme)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var themePtr *string
	if theme.Valid {
		themePtr = &theme.String
	}

	return &spec.CTFInfo{
		ID:          ctf.ID,
		Name:        ctf.Name,
		Description: ctf.Description,
		StartTime:   ctf.StartTime.Format(time.RFC3339),
		EndTime:     ctf.EndTime.Format(time.RFC3339),
		Slug:        ctf.Slug,
		TeamBased:   ctf.TeamBased,
		Theme:       themePtr,
	}, nil
}

func (s *Service) RegisterUser(ctx context.Context, req *spec.RegisterUserPayload) (*spec.CTFUser, error) {
	ctf, err := models.CTFS(models.CTFWhere.Slug.EQ(req.Slug)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}

	if ctf.TeamBased {
		// Team-based registration: req must include TeamCode
		if req.TeamCode == nil || len(*req.TeamCode) == 0 {
			return nil, errors.New("team code required for this CTF")
		}
		var teamID, teamName string
		err = s.db.QueryRowContext(ctx, `SELECT id, teamname FROM ctf_teams WHERE ctf_id = $1 AND password = $2`, ctf.ID, *req.TeamCode).Scan(&teamID, &teamName)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("invalid team code")
			}
			return nil, err
		}
		// Username must be unique within the team
		var exists bool
		err = s.db.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM ctf_users WHERE ctf_id = $1 AND team_id = $2 AND username = $3)`, ctf.ID, teamID, req.Username).Scan(&exists)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, errors.New("username already taken for this team")
		}
		var password string
		for {
			password, err = generatePassphrase(4)
			if err != nil {
				return nil, err
			}
			err = s.db.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM ctf_users WHERE ctf_id = $1 AND password = $2)`, ctf.ID, password).Scan(&exists)
			if err != nil {
				return nil, err
			}
			if !exists {
				break
			}
		}
		id := uuid.New().String()
		user := &models.CTFUser{
			ID:       id,
			CTFID:    ctf.ID,
			Username: req.Username,
			Password: password,
			TeamID:   null.StringFrom(teamID),
		}
		err = user.Insert(ctx, s.db, boil.Infer())
		if err != nil {
			return nil, err
		}
		return &spec.CTFUser{
			ID:       user.ID,
			CtfID:    user.CTFID,
			Username: user.Username,
			Password: user.Password,
			TeamID:   user.TeamID.Ptr(),
			Teamname: &teamName,
		}, nil
	} else {
		// Individual registration
		if len(req.Username) < 3 || len(req.Username) > 30 {
			return nil, errors.New("username must be between 3 and 30 characters")
		}
		var exists bool
		err = s.db.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM ctf_users WHERE ctf_id = $1 AND username = $2)`, ctf.ID, req.Username).Scan(&exists)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, errors.New("username already taken for this ctf")
		}
		var password string
		for {
			password, err = generatePassphrase(4)
			if err != nil {
				return nil, err
			}
			err = s.db.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM ctf_users WHERE ctf_id = $1 AND password = $2)`, ctf.ID, password).Scan(&exists)
			if err != nil {
				return nil, err
			}
			if !exists {
				break
			}
		}
		id := uuid.New().String()
		user := &models.CTFUser{
			ID:       id,
			CTFID:    ctf.ID,
			Username: req.Username,
			Password: password,
		}
		err = user.Insert(ctx, s.db, boil.Infer())
		if err != nil {
			return nil, err
		}
		var teamname *string
		if user.TeamID.Valid {
			var tname string
			err = s.db.QueryRowContext(ctx, `SELECT teamname FROM ctf_teams WHERE id = $1`, user.TeamID.String).Scan(&tname)
			if err == nil {
				teamname = &tname
			}
		}
		return &spec.CTFUser{
			ID:       user.ID,
			CtfID:    user.CTFID,
			Username: user.Username,
			Password: user.Password,
			TeamID:   user.TeamID.Ptr(),
			Teamname: teamname,
		}, nil
	}
}

func (s *Service) GetUser(ctx context.Context, req *spec.GetUserPayload) (*spec.CTFUser, error) {
	ctf, err := models.CTFS(models.CTFWhere.Slug.EQ(req.Slug)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}
	user, err := models.CTFUsers(models.CTFUserWhere.Password.EQ(req.Password), models.CTFUserWhere.CTFID.EQ(ctf.ID)).One(ctx, s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found for this ctf")
		}
		return nil, err
	}

	var teamname *string
	if user.TeamID.Valid {
		var tname string
		err = s.db.QueryRowContext(ctx, `SELECT teamname FROM ctf_teams WHERE id = $1`, user.TeamID.String).Scan(&tname)
		if err == nil {
			teamname = &tname
		}
	}
	return &spec.CTFUser{
		ID:       user.ID,
		CtfID:    user.CTFID,
		Username: user.Username,
		Password: user.Password,
		TeamID:   user.TeamID.Ptr(),
		Teamname: teamname,
	}, nil
}

func (s *Service) GetUserSolves(ctx context.Context, req *spec.GetUserSolvesPayload) (*spec.CTFUserSolves, error) {
	ctf, err := models.CTFS(models.CTFWhere.Slug.EQ(req.Slug)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}
	var res []*spec.CTFUserSolve
	var username string
	if ctf.TeamBased {
		// req.ID is the team ID
		team, err := models.CTFTeams(models.CTFTeamWhere.ID.EQ(req.ID), models.CTFTeamWhere.CTFID.EQ(ctf.ID)).One(ctx, s.db)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("team not found for this ctf")
			}
			return nil, err
		}
		username = team.Teamname
		rows, err := s.db.QueryContext(ctx, `
					   SELECT DISTINCT ON (s.challenge_id) s.challenge_id, s.created_at
					   FROM ctf_solves s
					   WHERE s.ctf_id = $1 AND s.team_id = $2
					   ORDER BY s.challenge_id, s.created_at ASC
			   `, ctf.ID, req.ID)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var challengeID string
			var solvedAt time.Time
			if err := rows.Scan(&challengeID, &solvedAt); err == nil {
				res = append(res, &spec.CTFUserSolve{
					ChallengeID: challengeID,
					SolvedAt:    solvedAt.Format(time.RFC3339),
				})
			}
		}
	} else {
		// req.ID is the user ID
		user, err := models.CTFUsers(models.CTFUserWhere.ID.EQ(req.ID), models.CTFUserWhere.CTFID.EQ(ctf.ID)).One(ctx, s.db)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("user not found for this ctf")
			}
			return nil, err
		}
		username = user.Username
		solves, err := models.CTFSolves(models.CTFSolfWhere.UserID.EQ(user.ID), models.CTFSolfWhere.CTFID.EQ(ctf.ID)).All(ctx, s.db)
		if err != nil {
			return nil, err
		}
		res = make([]*spec.CTFUserSolve, len(solves))
		for i, solve := range solves {
			res[i] = &spec.CTFUserSolve{
				ChallengeID: solve.ChallengeID,
				SolvedAt:    solve.CreatedAt.Format(time.RFC3339),
			}
		}
	}
	return &spec.CTFUserSolves{
		Username: username,
		Solves:   res,
	}, nil
}

func (s *Service) ListChallenges(ctx context.Context, req *spec.ListChallengesPayload) (spec.SsmCtfChallengeCollection, error) {
	ctf, err := models.CTFS(models.CTFWhere.Slug.EQ(req.Slug)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}

	if ctf.StartTime.After(time.Now()) {
		return nil, errors.New("ctf not started yet")
	}

	rows, err := s.db.QueryContext(ctx, `SELECT challenge_id, custom_score, display_order FROM ctf_challenges WHERE ctf_id = $1`, ctf.ID)
	if err != nil {
		return nil, err
	}

	customScores := make(map[string]*int)
	displayOrders := make(map[string]int)
	for rows.Next() {
		var cid string
		var customScore sql.NullInt64
		var displayOrder int
		if err := rows.Scan(&cid, &customScore, &displayOrder); err != nil {
			return nil, err
		}
		if customScore.Valid {
			v := int(customScore.Int64)
			customScores[cid] = &v
		} else {
			customScores[cid] = nil
		}
		displayOrders[cid] = displayOrder
	}
	rows.Close()

	var challengeIDs []string
	for cid := range customScores {
		challengeIDs = append(challengeIDs, cid)
	}

	var user *models.CTFUser
	var teamID *string
	if req.Password != nil {
		user, err = models.CTFUsers(models.CTFUserWhere.Password.EQ(*req.Password), models.CTFUserWhere.CTFID.EQ(ctf.ID)).One(ctx, s.db)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("user not found for this ctf")
			}
			return nil, err
		}
		if user.TeamID.Valid {
			teamID = &user.TeamID.String
		}
	}
	// precalculate 1 query
	challengeIDList := make([]string, len(challengeIDs))
	for i, id := range challengeIDs {
		challengeIDList[i] = "'" + id + "'"
	}
	challengeIDsSQL := strings.Join(challengeIDList, ",")

	solveCounts := make(map[string]int)
	teamSolveCounts := make(map[string]int)
	userTeamSolveCounts := make(map[string]int)
	userSolved := make(map[string]bool)
	userTeamSolved := make(map[string]bool)

	if len(challengeIDs) > 0 {
		solveCountQuery := `
			SELECT 
				challenge_id,
				COUNT(1) as total_solves,
				COUNT(DISTINCT CASE WHEN team_id IS NOT NULL THEN team_id END) as team_solves
			FROM ctf_solves 
			WHERE ctf_id = $1 AND challenge_id IN (` + challengeIDsSQL + `)
			GROUP BY challenge_id`

		rows, err := s.db.QueryContext(ctx, solveCountQuery, ctf.ID)
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				var challengeID string
				var totalSolves, teamSolves int
				if err := rows.Scan(&challengeID, &totalSolves, &teamSolves); err == nil {
					solveCounts[challengeID] = totalSolves
					teamSolveCounts[challengeID] = teamSolves
				}
			}
		}

		if teamID != nil && *teamID != "" {
			teamSolveQuery := `
				SELECT 
					challenge_id,
					COUNT(DISTINCT user_id) as team_user_solves,
					COUNT(1) > 0 as team_solved
				FROM ctf_solves 
				WHERE ctf_id = $1 AND challenge_id IN (` + challengeIDsSQL + `) AND team_id = $2
				GROUP BY challenge_id`

			rows, err := s.db.QueryContext(ctx, teamSolveQuery, ctf.ID, *teamID)
			if err == nil {
				defer rows.Close()
				for rows.Next() {
					var challengeID string
					var teamUserSolves int
					var teamSolved bool
					if err := rows.Scan(&challengeID, &teamUserSolves, &teamSolved); err == nil {
						userTeamSolveCounts[challengeID] = teamUserSolves
						userTeamSolved[challengeID] = teamSolved
					}
				}
			}
		}

		if user != nil {
			userSolveQuery := `
				SELECT DISTINCT challenge_id
				FROM ctf_solves 
				WHERE ctf_id = $1 AND challenge_id IN (` + challengeIDsSQL + `) AND user_id = $2`

			rows, err := s.db.QueryContext(ctx, userSolveQuery, ctf.ID, user.ID)
			if err == nil {
				defer rows.Close()
				for rows.Next() {
					var challengeID string
					if err := rows.Scan(&challengeID); err == nil {
						userSolved[challengeID] = true
					}
				}
			}
		}
	}

	// Now build the main query without expensive subqueries
	challs := make([]*custommodels.UserChall, 0)
	q := models.NewQuery(
		qm.Select("challenges.*, categories.name as category"),
		qm.From(models.TableNames.Challenges),
		qm.InnerJoin("categories ON categories.id = challenges.category_id"),
		qm.Load(models.ChallengeRels.ChallengeFiles),
		qm.Load(models.ChallengeRels.ChallengeServices),
		qm.Load(models.ChallengeRels.Authors),
	)

	models.ChallengeWhere.ID.IN(challengeIDs).Apply(q)

	err = q.Bind(ctx, s.db, &challs)
	if err != nil {
		return nil, err
	}

	allSolvers := make(map[string][]*spec.SsmSolver)
	allTeamSolvers := make(map[string][]*spec.SsmSolver)
	allSolversInTeam := make(map[string][]*spec.SsmSolver)

	challengeIDList = make([]string, len(challengeIDs))
	for i, id := range challengeIDs {
		challengeIDList[i] = "'" + id + "'"
	}
	challengeIDsSQL = strings.Join(challengeIDList, ",")

	if len(challengeIDs) > 0 {
		solversQuery := `
			WITH ranked_solves AS (
				SELECT s.challenge_id, u.id, u.username, s.created_at,
					   ROW_NUMBER() OVER (PARTITION BY s.challenge_id ORDER BY s.created_at ASC) as rn
				FROM ctf_solves s
				INNER JOIN ctf_users u ON u.id = s.user_id
				WHERE s.ctf_id = $1 AND s.challenge_id IN (` + challengeIDsSQL + `)
			)
			SELECT challenge_id, id, username, created_at
			FROM ranked_solves
			WHERE rn <= 5
			ORDER BY challenge_id, created_at ASC`

		rows, err := s.db.QueryContext(ctx, solversQuery, ctf.ID)
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				var challengeID, id, username string
				var solvedAt time.Time
				if err := rows.Scan(&challengeID, &id, &username, &solvedAt); err == nil {
					if allSolvers[challengeID] == nil {
						allSolvers[challengeID] = make([]*spec.SsmSolver, 0)
					}
					allSolvers[challengeID] = append(allSolvers[challengeID], &spec.SsmSolver{
						ID:       id,
						FullName: username,
						SolvedAt: solvedAt.Unix(),
					})
				}
			}
		}

		// Batch query for team solvers (first 5 teams per challenge)
		if ctf.TeamBased {
			teamSolversQuery := `
				WITH first_team_solves AS (
					SELECT s.challenge_id, s.team_id, MIN(s.created_at) as first_solve_time
					FROM ctf_solves s
					WHERE s.ctf_id = $1 AND s.challenge_id IN (` + challengeIDsSQL + `) AND s.team_id IS NOT NULL
					GROUP BY s.challenge_id, s.team_id
				),
				ranked_team_solves AS (
					SELECT fts.challenge_id, fts.team_id, fts.first_solve_time,
						   ROW_NUMBER() OVER (PARTITION BY fts.challenge_id ORDER BY fts.first_solve_time ASC) as rn
					FROM first_team_solves fts
				)
				SELECT rts.challenge_id, t.id, t.teamname, rts.first_solve_time
				FROM ranked_team_solves rts
				INNER JOIN ctf_teams t ON t.id = rts.team_id
				WHERE rts.rn <= 5
				ORDER BY rts.challenge_id, rts.first_solve_time ASC`

			rows, err = s.db.QueryContext(ctx, teamSolversQuery, ctf.ID)
			if err == nil {
				defer rows.Close()
				for rows.Next() {
					var challengeID, teamID, teamname string
					var solvedAt time.Time
					if err := rows.Scan(&challengeID, &teamID, &teamname, &solvedAt); err == nil {
						if allTeamSolvers[challengeID] == nil {
							allTeamSolvers[challengeID] = make([]*spec.SsmSolver, 0)
						}
						allTeamSolvers[challengeID] = append(allTeamSolvers[challengeID], &spec.SsmSolver{
							ID:       teamID,
							FullName: teamname,
							SolvedAt: solvedAt.Unix(),
						})
					}
				}
			}
		}

		// Batch query for solvers in user's team
		if ctf.TeamBased && user != nil && user.TeamID.Valid {
			teamSolversQuery := `
				WITH ranked_team_member_solves AS (
					SELECT s.challenge_id, u.id, u.username, s.created_at,
						   ROW_NUMBER() OVER (PARTITION BY s.challenge_id ORDER BY s.created_at ASC) as rn
					FROM ctf_solves s
					INNER JOIN ctf_users u ON u.id = s.user_id
					WHERE s.ctf_id = $1 AND s.challenge_id IN (` + challengeIDsSQL + `) AND s.team_id = $2
				)
				SELECT challenge_id, id, username, created_at
				FROM ranked_team_member_solves
				WHERE rn <= 5
				ORDER BY challenge_id, created_at ASC`

			rows, err = s.db.QueryContext(ctx, teamSolversQuery, ctf.ID, user.TeamID.String)
			if err == nil {
				defer rows.Close()
				for rows.Next() {
					var challengeID, id, username string
					var solvedAt time.Time
					if err := rows.Scan(&challengeID, &id, &username, &solvedAt); err == nil {
						if allSolversInTeam[challengeID] == nil {
							allSolversInTeam[challengeID] = make([]*spec.SsmSolver, 0)
						}
						allSolversInTeam[challengeID] = append(allSolversInTeam[challengeID], &spec.SsmSolver{
							ID:       id,
							FullName: username,
							SolvedAt: solvedAt.Unix(),
						})
					}
				}
			}
		}
	}

	res := make(spec.SsmCtfChallengeCollection, len(challs))

	for i, chall := range challs {
		// Get pre-calculated solve counts
		numSolves := solveCounts[chall.ID]
		numTeamSolves := teamSolveCounts[chall.ID]
		numSolvesInTeam := userTeamSolveCounts[chall.ID]
		solvedInTeam := userTeamSolved[chall.ID]
		solved := userSolved[chall.ID]

		score := chall.StaticScore.Int
		if cs, ok := customScores[chall.ID]; ok && cs != nil {
			score = *cs
		} else if !chall.StaticScore.Valid {
			score = dynamicScore(500, 100, float64(numSolves))
		}

		solvers := allSolvers[chall.ID]
		teamSolvers := allTeamSolvers[chall.ID]
		solversInTeam := allSolversInTeam[chall.ID]
		
		// json saker
		if solvers == nil {
			solvers = make([]*spec.SsmSolver, 0)
		}
		if teamSolvers == nil {
			teamSolvers = make([]*spec.SsmSolver, 0)
		}
		if solversInTeam == nil {
			solversInTeam = make([]*spec.SsmSolver, 0)
		}

		res[i] = &spec.SsmCtfChallenge{
			ID:              chall.ID,
			Slug:            chall.Slug,
			Title:           chall.Title,
			Description:     chall.Description,
			Score:           score,
			Solves:          numSolves,
			NumTeamSolves:   &numTeamSolves,
			NumSolvesInTeam: &numSolvesInTeam,
			SolvedInTeam:    &solvedInTeam,
			Solved:          solved,
			Category:        chall.Category,
			CtfEventID:      chall.CTFEventID.Ptr(),
			ChallNamespace:  chall.ChallNamespace.Ptr(),
			DisplayOrder:    displayOrders[chall.ID],
			Solvers:         solvers,
			TeamSolvers:     teamSolvers,
			SolversInTeam:   solversInTeam,
		}

		res[i].Files = make([]*spec.ChallengeFiles, len(chall.R.ChallengeFiles))
		for i2, file := range chall.R.ChallengeFiles {
			res[i].Files[i2] = &spec.ChallengeFiles{
				Filename: file.FriendlyName,
				URL:      file.URL,
			}
		}

		res[i].Authors = make([]*spec.Author, len(chall.R.Authors))
		for i2, v := range chall.R.Authors {
			res[i].Authors[i2] = &spec.Author{
				ID:          v.ID,
				FullName:    v.FullName,
				Description: v.Description,
				Sponsor:     v.Sponsor,
				Slug:        v.Slug,
				ImageURL:    v.ImageURL.Ptr(),
				Publish:     v.Publish,
			}
		}

		res[i].Services = make([]*spec.ChallengeService, len(chall.R.ChallengeServices))
		for i2, v := range chall.R.ChallengeServices {
			res[i].Services[i2] = &spec.ChallengeService{
				UserDisplay: v.UserDisplay,
				Hyperlink:   v.Hyperlink,
			}
		}
	}

	return res, nil
}

func dynamicScore(init, min, solvers float64) int {
	if solvers == 0 {
		return int(init)
	}
	decay := 25.0
	return int(init + ((min-init)/math.Pow(decay, 2))*math.Pow(solvers-1, 2))
}

func (s *Service) Scoreboard(ctx context.Context, req *spec.ScoreboardPayload) ([]*spec.CTFScore, error) {
	ctf, err := models.CTFS(models.CTFWhere.Slug.EQ(req.Slug)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}

	if ctf.TeamBased {
		// Team scoreboard: sum team scores, only count each challenge once per team
		teams, err := s.db.QueryContext(ctx, `SELECT id, teamname FROM ctf_teams WHERE ctf_id = $1`, ctf.ID)
		if err != nil {
			return nil, err
		}
		defer teams.Close()
		var scores []*spec.CTFScore
		for teams.Next() {
			var teamID, teamName string
			if err := teams.Scan(&teamID, &teamName); err != nil {
				return nil, err
			}
			// Get unique challenge_ids solved by any member of the team
			rows, err := s.db.QueryContext(ctx, `SELECT DISTINCT challenge_id FROM ctf_solves WHERE ctf_id = $1 AND team_id = $2`, ctf.ID, teamID)
			if err != nil {
				return nil, err
			}
			var challengeIDs []string
			for rows.Next() {
				var challID string
				if err := rows.Scan(&challID); err != nil {
					rows.Close()
					return nil, err
				}
				challengeIDs = append(challengeIDs, challID)
			}
			rows.Close()
			var score int64
			var solves []string
			for _, challID := range challengeIDs {
				// Get the first solver for this challenge in this team
				var solver string
				err := s.db.QueryRowContext(ctx, `SELECT u.username FROM ctf_solves s INNER JOIN ctf_users u ON u.id = s.user_id WHERE s.ctf_id = $1 AND s.team_id = $2 AND s.challenge_id = $3 ORDER BY s.created_at ASC LIMIT 1`, ctf.ID, teamID, challID).Scan(&solver)
				if err != nil {
					return nil, err
				}
				solves = append(solves, challID+":"+solver)
				// Get the score for this challenge (custom or static)
				var challScore sql.NullInt64
				err = s.db.QueryRowContext(ctx, `SELECT COALESCE(cc.custom_score, ch.static_score) FROM challenges ch LEFT JOIN ctf_challenges cc ON cc.ctf_id = $1 AND cc.challenge_id = ch.id WHERE ch.id = $2`, ctf.ID, challID).Scan(&challScore)
				if err != nil {
					return nil, err
				}
				if challScore.Valid {
					score += challScore.Int64
				}
			}
			scores = append(scores, &spec.CTFScore{
				ID:       teamID,
				Username: teamName,
				Score:    score,
				Solves:   solves,
			})
		}
		// Sort by score desc, then teamname asc
		sort.Slice(scores, func(i, j int) bool {
			if scores[i].Score == scores[j].Score {
				return scores[i].Username < scores[j].Username
			}
			return scores[i].Score > scores[j].Score
		})
		return scores, nil
	} else {
		// Individual scoreboard
		rows, err := s.db.QueryContext(ctx, `
			SELECT u.id, u.username, COALESCE(SUM(COALESCE(cc.custom_score, ch.static_score)), 0) as score,
				MAX(s.created_at) as last_solve
			FROM ctf_users u
			LEFT JOIN (
				SELECT DISTINCT ON (user_id, challenge_id) *
				FROM ctf_solves
				WHERE ctf_id = $1
				ORDER BY user_id, challenge_id, created_at ASC
			) s ON u.id = s.user_id AND s.ctf_id = $1
			LEFT JOIN challenges ch ON s.challenge_id = ch.id
			LEFT JOIN ctf_challenges cc ON cc.ctf_id = $1 AND cc.challenge_id = ch.id
			WHERE u.ctf_id = $1
			GROUP BY u.id, u.username
			ORDER BY score DESC, last_solve ASC NULLS LAST, u.username ASC
		`, ctf.ID)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var scores []*spec.CTFScore
		for rows.Next() {
			var userID string
			var username string
			var score int64
			var lastSolve sql.NullTime
			if err := rows.Scan(&userID, &username, &score, &lastSolve); err != nil {
				return nil, err
			}
			solvesRows, err := s.db.QueryContext(ctx, `
				SELECT DISTINCT s.challenge_id
				FROM ctf_solves s
				INNER JOIN ctf_users u ON u.id = s.user_id
				WHERE s.ctf_id = $1 AND u.username = $2
			`, ctf.ID, username)
			if err != nil {
				return nil, err
			}
			var solves []string
			for solvesRows.Next() {
				var challID string
				if err := solvesRows.Scan(&challID); err != nil {
					solvesRows.Close()
					return nil, err
				}
				solves = append(solves, challID)
			}
			solvesRows.Close()
			scores = append(scores, &spec.CTFScore{
				ID:       userID,
				Username: username,
				Score:    score,
				Solves:   solves,
			})
		}
		return scores, nil
	}
}

func (s *Service) SubmitFlag(ctx context.Context, req *spec.SubmitFlagPayload) (*spec.CTFSolve, error) {
	ctf, err := models.CTFS(models.CTFWhere.Slug.EQ(req.Slug)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}

	if ctf.StartTime.After(time.Now()) || ctf.EndTime.Before(time.Now()) {
		return nil, errors.New("ctf not active")
	}

	var userID, teamID string
	if ctf.TeamBased {
		err = s.db.QueryRowContext(ctx, `SELECT id, team_id FROM ctf_users WHERE ctf_id = $1 AND password = $2`, ctf.ID, req.Password).Scan(&userID, &teamID)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("user not registered for this ctf")
			}
			return nil, err
		}
	} else {
		err = s.db.QueryRowContext(ctx, `SELECT id FROM ctf_users WHERE ctf_id = $1 AND password = $2`, ctf.ID, req.Password).Scan(&userID)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("user not registered for this ctf")
			}
			return nil, err
		}
	}
	var challengeID string
	err = s.db.QueryRowContext(ctx, `SELECT challenge_id FROM ctf_challenges WHERE ctf_id = $1 AND challenge_id = $2`, ctf.ID, req.ChallengeID).Scan(&challengeID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("challenge not part of this ctf")
		}
		return nil, err
	}
	var alreadySolved bool
	err = s.db.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM ctf_solves WHERE ctf_id = $1 AND user_id = $2 AND challenge_id = $3)`, ctf.ID, userID, req.ChallengeID).Scan(&alreadySolved)
	if err != nil {
		return nil, err
	}
	if alreadySolved {
		return nil, errors.New("challenge already solved by this user")
	}

	flags, err := models.Flags(
		models.FlagWhere.ChallengeID.EQ(challengeID),
	).All(ctx, s.db)
	if err != nil {
		return nil, err
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
	if !flagCorrect {
		return nil, errors.New("incorrect flag")
	}

	id := uuid.New().String()
	now := time.Now().UTC()
	solve := &models.CTFSolf{
		ID:          id,
		UserID:      userID,
		CTFID:       ctf.ID,
		ChallengeID: req.ChallengeID,
		CreatedAt:   now,
	}
	if ctf.TeamBased {
		solve.TeamID = null.StringFrom(teamID)
	}
	err = solve.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &spec.CTFSolve{
		ID:          id,
		CtfID:       ctf.ID,
		UserID:      userID,
		ChallengeID: req.ChallengeID,
		SolvedAt:    now.Format(time.RFC3339),
	}, nil
}

func generatePassphrase(numWords int) (string, error) {
	rand.Seed(time.Now().UnixNano())
	passWords := make([]string, numWords)

	for i := 0; i < numWords; i++ {
		lineNum := rand.Intn(len(words) - 1)
		passWords[i] = words[lineNum]
	}

	return strings.Join(passWords, "-"), nil
}
