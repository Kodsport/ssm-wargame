package admin

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (s *service) CreateCTF(ctx context.Context, req *spec.CreateCTFPayload) (*spec.CTF, error) {
	id := uuid.New().String()

	theme := ""
	if req.Theme != nil {
		theme = *req.Theme
	}

	_, err := s.db.ExecContext(ctx, `
		INSERT INTO ctfs (id, name, description, slug, private, start_time, end_time, team_based, theme, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW())
	`, id, req.Name, req.Description, req.Slug, false, time.Unix(req.StartTime, 0), time.Unix(req.EndTime, 0), req.TeamBased, theme)
	if err != nil {
		return nil, err
	}

	for _, chall := range req.Challenges {
		if chall.CustomScore != nil {
			_, err = s.db.ExecContext(ctx, `INSERT INTO ctf_challenges (ctf_id, challenge_id, custom_score, display_order) VALUES ($1, $2, $3, $4)`, id, chall.ID, *chall.CustomScore, chall.DisplayOrder)
		} else {
			_, err = s.db.ExecContext(ctx, `INSERT INTO ctf_challenges (ctf_id, challenge_id, display_order) VALUES ($1, $2, $3)`, id, chall.ID, chall.DisplayOrder)
		}
		if err != nil {
			return nil, err
		}
	}
	return &spec.CTF{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Slug:        req.Slug,
		StartTime:   time.Unix(req.StartTime, 0).Format(time.RFC3339),
		EndTime:     time.Unix(req.EndTime, 0).Format(time.RFC3339),
		Challenges:  req.Challenges,
		TeamBased:   req.TeamBased,
		Theme:       req.Theme,
	}, nil
}

func (s *service) ListCTFs(ctx context.Context, req *spec.ListCTFsPayload) ([]*spec.CTF, error) {
	rows, err := s.db.QueryContext(ctx, `
		SELECT id, name, description, start_time, end_time, slug, team_based, COALESCE(theme, '') as theme
		FROM ctfs 
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*spec.CTF
	for rows.Next() {
		var ctf struct {
			ID          string
			Name        string
			Description string
			StartTime   time.Time
			EndTime     time.Time
			Slug        string
			TeamBased   bool
			Theme       string
		}

		if err := rows.Scan(&ctf.ID, &ctf.Name, &ctf.Description, &ctf.StartTime, &ctf.EndTime, &ctf.Slug, &ctf.TeamBased, &ctf.Theme); err != nil {
			return nil, err
		}

		challengeRows, err := s.db.QueryContext(ctx, `SELECT challenge_id, custom_score, display_order FROM ctf_challenges WHERE ctf_id = $1`, ctf.ID)
		if err != nil {
			return nil, err
		}

		var challenges []*spec.CTFChallenge
		for challengeRows.Next() {
			var cid string
			var customScore sql.NullInt64
			var displayOrder int
			if err := challengeRows.Scan(&cid, &customScore, &displayOrder); err != nil {
				challengeRows.Close()
				return nil, err
			}
			var csPtr *int = nil
			if customScore.Valid {
				v := int(customScore.Int64)
				csPtr = &v
			}
			challenges = append(challenges, &spec.CTFChallenge{
				ID:           cid,
				CustomScore:  csPtr,
				DisplayOrder: displayOrder,
			})
		}
		challengeRows.Close()

		var themePtr *string
		if ctf.Theme != "" {
			themePtr = &ctf.Theme
		}

		result = append(result, &spec.CTF{
			ID:          ctf.ID,
			Name:        ctf.Name,
			Description: ctf.Description,
			StartTime:   ctf.StartTime.Format(time.RFC3339),
			EndTime:     ctf.EndTime.Format(time.RFC3339),
			Slug:        ctf.Slug,
			Challenges:  challenges,
			TeamBased:   ctf.TeamBased,
			Theme:       themePtr,
		})
	}
	return result, nil
}

func (s *service) UpdateCTF(ctx context.Context, req *spec.UpdateCTFPayload) (*spec.CTF, error) {
	ctf, err := models.FindCTF(ctx, s.db, req.ID)
	if err != nil {
		return nil, err
	}

	// Check if slug is unique
	existingCTF, err := models.CTFS(models.CTFWhere.Slug.EQ(req.Slug)).One(ctx, s.db)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if existingCTF != nil && existingCTF.ID != ctf.ID {
		return nil, errors.New("slug already exists")
	}

	teamBased := false
	if req.TeamBased != nil {
		teamBased = *req.TeamBased
	}

	theme := ""
	if req.Theme != nil {
		theme = *req.Theme
	}
	// kör direkta sql querys istället för modeller
	_, err = s.db.ExecContext(ctx, `
		UPDATE ctfs 
		SET name = $1, description = $2, start_time = $3, end_time = $4, slug = $5, team_based = $6, theme = $7, updated_at = NOW()
		WHERE id = $8
	`, req.Name, req.Description, time.Unix(req.StartTime, 0), time.Unix(req.EndTime, 0), req.Slug, teamBased, theme, req.ID)
	if err != nil {
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, `DELETE FROM ctf_challenges WHERE ctf_id = $1`, ctf.ID)
	if err != nil {
		return nil, err
	}
	for _, chall := range req.Challenges {
		if chall.CustomScore != nil {
			_, err = s.db.ExecContext(ctx, `INSERT INTO ctf_challenges (ctf_id, challenge_id, custom_score, display_order) VALUES ($1, $2, $3, $4)`, ctf.ID, chall.ID, *chall.CustomScore, chall.DisplayOrder)
		} else {
			_, err = s.db.ExecContext(ctx, `INSERT INTO ctf_challenges (ctf_id, challenge_id, display_order) VALUES ($1, $2, $3)`, ctf.ID, chall.ID, chall.DisplayOrder)
		}
		if err != nil {
			return nil, err
		}
	}

	return &spec.CTF{
		ID:          ctf.ID,
		Name:        req.Name,
		Description: req.Description,
		StartTime:   time.Unix(req.StartTime, 0).Format(time.RFC3339),
		EndTime:     time.Unix(req.EndTime, 0).Format(time.RFC3339),
		Slug:        req.Slug,
		Challenges:  req.Challenges,
		TeamBased:   teamBased,
		Theme:       req.Theme,
	}, nil
}

func (s *service) DeleteCTF(ctx context.Context, req *spec.DeleteCTFPayload) error {
	_, err := models.CTFS(models.CTFWhere.ID.EQ(req.ID)).DeleteAll(ctx, s.db)
	return err
}

func (s *service) CreateChallengeGroup(ctx context.Context, req *spec.CreateChallengeGroupPayload) (*spec.ChallengeGroup, error) {
	id := uuid.New().String()
	group := &models.ChallengeGroup{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
	}
	err := group.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	for _, chall := range req.Challenges {
		if chall.CustomScore != nil {
			_, err = s.db.ExecContext(ctx, `INSERT INTO challenge_group_challenges (challenge_group_id, challenge_id, custom_score, display_order) VALUES ($1, $2, $3, $4)`, id, chall.ID, *chall.CustomScore, chall.DisplayOrder)
		} else {
			_, err = s.db.ExecContext(ctx, `INSERT INTO challenge_group_challenges (challenge_group_id, challenge_id, display_order) VALUES ($1, $2, $3)`, id, chall.ID, chall.DisplayOrder)
		}
		if err != nil {
			return nil, err
		}
	}
	return &spec.ChallengeGroup{
		ID:          group.ID,
		Name:        group.Name,
		Description: group.Description,
		Challenges:  req.Challenges,
	}, nil
}

func (s *service) UpdateChallengeGroup(ctx context.Context, req *spec.UpdateChallengeGroupPayload) (*spec.ChallengeGroup, error) {
	group, err := models.FindChallengeGroup(ctx, s.db, req.ID)
	if err != nil {
		return nil, err
	}
	group.Name = req.Name
	group.Description = req.Description
	_, err = group.Update(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	_, err = s.db.ExecContext(ctx, `DELETE FROM challenge_group_challenges WHERE challenge_group_id = $1`, group.ID)
	if err != nil {
		return nil, err
	}
	for _, chall := range req.Challenges {
		if chall.CustomScore != nil {
			_, err = s.db.ExecContext(ctx, `INSERT INTO challenge_group_challenges (challenge_group_id, challenge_id, custom_score, display_order) VALUES ($1, $2, $3, $4)`, group.ID, chall.ID, *chall.CustomScore, chall.DisplayOrder)
		} else {
			_, err = s.db.ExecContext(ctx, `INSERT INTO challenge_group_challenges (challenge_group_id, challenge_id, display_order) VALUES ($1, $2, $3)`, group.ID, chall.ID, chall.DisplayOrder)
		}
		if err != nil {
			return nil, err
		}
	}
	return &spec.ChallengeGroup{
		ID:          group.ID,
		Name:        group.Name,
		Description: group.Description,
		Challenges:  req.Challenges,
	}, nil
}

func (s *service) DeleteChallengeGroup(ctx context.Context, req *spec.DeleteChallengeGroupPayload) error {
	_, err := models.ChallengeGroups(models.ChallengeGroupWhere.ID.EQ(req.ID)).DeleteAll(ctx, s.db)
	return err
}

func (s *service) ListChallengeGroups(ctx context.Context, req *spec.ListChallengeGroupsPayload) ([]*spec.ChallengeGroup, error) {
	groups, err := models.ChallengeGroups().All(ctx, s.db)
	if err != nil {
		return nil, err
	}
	var result []*spec.ChallengeGroup
	for _, group := range groups {
		rows, err := s.db.QueryContext(ctx, `SELECT challenge_id, custom_score, display_order FROM challenge_group_challenges WHERE challenge_group_id = $1`, group.ID)
		if err != nil {
			return nil, err
		}
		var challenges []*spec.CTFChallenge
		for rows.Next() {
			var cid string
			var customScore sql.NullInt64
			var displayOrder int
			if err := rows.Scan(&cid, &customScore, &displayOrder); err != nil {
				rows.Close()
				return nil, err
			}
			var csPtr *int = nil
			if customScore.Valid {
				v := int(customScore.Int64)
				csPtr = &v
			}
			challenges = append(challenges, &spec.CTFChallenge{
				ID:           cid,
				CustomScore:  csPtr,
				DisplayOrder: displayOrder,
			})
		}
		rows.Close()
		result = append(result, &spec.ChallengeGroup{
			ID:          group.ID,
			Name:        group.Name,
			Description: group.Description,
			Challenges:  challenges,
		})
	}
	return result, nil
}

func (s *service) ListCTFUsers(ctx context.Context, req *spec.ListCTFUsersPayload) ([]*spec.CTFUser, error) {
	ctf, err := models.CTFS(models.CTFWhere.ID.EQ(req.CtfID)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}
	rows, err := s.db.QueryContext(ctx, `SELECT id, username, password FROM ctf_users WHERE ctf_id = $1 ORDER BY created_at`, ctf.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*spec.CTFUser
	for rows.Next() {
		var userID string
		var username string
		var password string
		if err := rows.Scan(&userID, &username, &password); err != nil {
			return nil, err
		}
		result = append(result, &spec.CTFUser{
			ID:       userID,
			CtfID:    ctf.ID,
			Username: username,
			Password: password,
		})
	}
	return result, nil
}

func (s *service) DeleteCTFUser(ctx context.Context, req *spec.DeleteCTFUserPayload) error {
	ctf, err := models.CTFS(models.CTFWhere.ID.EQ(req.CtfID)).One(ctx, s.db)
	if err != nil {
		return err
	}
	_, err = models.CTFUsers(models.CTFUserWhere.ID.EQ(req.UserID), models.CTFUserWhere.CTFID.EQ(ctf.ID)).DeleteAll(ctx, s.db)
	return err
}

func (s *service) UpdateCTFUser(ctx context.Context, req *spec.UpdateCTFUserPayload) (*spec.CTFUser, error) {
	ctf, err := models.CTFS(models.CTFWhere.ID.EQ(req.CtfID)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}
	user, err := models.CTFUsers(models.CTFUserWhere.ID.EQ(req.UserID), models.CTFUserWhere.CTFID.EQ(ctf.ID)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}

	// Check if username already exists for this ctf
	var exists bool
	err = s.db.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM ctf_users WHERE ctf_id = $1 AND username = $2)`, ctf.ID, req.Username).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("username already taken for this ctf")
	}

	user.Username = req.Username

	_, err = user.Update(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return &spec.CTFUser{
		ID:       user.ID,
		CtfID:    ctf.ID,
		Username: user.Username,
	}, nil
}

func (s *service) ListCTFTeams(ctx context.Context, req *spec.ListCTFTeamsPayload) ([]*spec.CTFTeam, error) {
	ctf, err := models.CTFS(models.CTFWhere.ID.EQ(req.CtfID)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}
	teams, err := models.CTFTeams(models.CTFTeamWhere.CTFID.EQ(ctf.ID)).All(ctx, s.db)
	if err != nil {
		return nil, err
	}
	var result []*spec.CTFTeam
	for _, team := range teams {
		result = append(result, &spec.CTFTeam{
			ID:       team.ID,
			CtfID:    team.CTFID,
			Teamname: team.Teamname,
			Password: team.Password,
		})
	}
	return result, nil
}

func (s *service) CreateCTFTeam(ctx context.Context, req *spec.CreateCTFTeamPayload) (*spec.CTFTeam, error) {
	ctf, err := models.CTFS(models.CTFWhere.ID.EQ(req.CtfID)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}

	// Check if team name already exists for this ctf
	var exists bool
	err = s.db.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM ctf_teams WHERE ctf_id = $1 AND teamname = $2)`, ctf.ID, req.Teamname).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("team name already taken for this ctf")
	}

	// Generate a unique team password (team code), check for collisions
	var teamPassword string
	for {
		teamPassword = uuid.New().String()[:8]
		var exists bool
		err = s.db.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM ctf_teams WHERE ctf_id = $1 AND password = $2)`, ctf.ID, teamPassword).Scan(&exists)
		if err != nil {
			return nil, err
		}
		if !exists {
			break
		}
	}

	id := uuid.New().String()
	team := &models.CTFTeam{
		ID:       id,
		CTFID:    ctf.ID,
		Teamname: req.Teamname,
		Password: teamPassword,
	}
	err = team.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &spec.CTFTeam{
		ID:       team.ID,
		CtfID:    team.CTFID,
		Teamname: team.Teamname,
		Password: team.Password,
	}, nil
}

func (s *service) DeleteCTFTeam(ctx context.Context, req *spec.DeleteCTFTeamPayload) error {
	ctf, err := models.CTFS(models.CTFWhere.ID.EQ(req.CtfID)).One(ctx, s.db)
	if err != nil {
		return err
	}
	_, err = models.CTFTeams(models.CTFTeamWhere.ID.EQ(req.TeamID), models.CTFTeamWhere.CTFID.EQ(ctf.ID)).DeleteAll(ctx, s.db)
	return err
}

func (s *service) UpdateCTFTeam(ctx context.Context, req *spec.UpdateCTFTeamPayload) (*spec.CTFTeam, error) {
	ctf, err := models.CTFS(models.CTFWhere.ID.EQ(req.CtfID)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}
	team, err := models.CTFTeams(models.CTFTeamWhere.ID.EQ(req.TeamID), models.CTFTeamWhere.CTFID.EQ(ctf.ID)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}

	// Check if team name already exists for this ctf (excluding current team)
	var exists bool
	err = s.db.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM ctf_teams WHERE ctf_id = $1 AND teamname = $2 AND id != $3)`, ctf.ID, req.Teamname, team.ID).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("team name already taken for this ctf")
	}

	team.Teamname = req.Teamname

	_, err = team.Update(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return &spec.CTFTeam{
		ID:       team.ID,
		CtfID:    team.CTFID,
		Teamname: team.Teamname,
		Password: team.Password,
	}, nil
}
