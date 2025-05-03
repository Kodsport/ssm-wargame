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
	ctf := &models.CTF{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Slug:        req.Slug,
		Private:     false,
		StartTime:   time.Unix(req.StartTime, 0),
		EndTime:     time.Unix(req.EndTime, 0),
	}
	err := ctf.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	for _, chall := range req.Challenges {
		if chall.CustomScore != nil {
			_, err = s.db.ExecContext(ctx, `INSERT INTO ctf_challenges (ctf_id, challenge_id, custom_score) VALUES ($1, $2, $3)`, id, chall.ID, *chall.CustomScore)
		} else {
			_, err = s.db.ExecContext(ctx, `INSERT INTO ctf_challenges (ctf_id, challenge_id) VALUES ($1, $2)`, id, chall.ID)
		}
		if err != nil {
			return nil, err
		}
	}
	return &spec.CTF{
		ID:          ctf.ID,
		Name:        ctf.Name,
		Description: ctf.Description,
		Slug:        ctf.Slug,
		StartTime:   ctf.StartTime.Format(time.RFC3339),
		EndTime:     ctf.EndTime.Format(time.RFC3339),
		Challenges:  req.Challenges,
	}, nil
}

func (s *service) ListCTFs(ctx context.Context, req *spec.ListCTFsPayload) ([]*spec.CTF, error) {
	ctfs, err := models.CTFS().All(ctx, s.db)
	if err != nil {
		return nil, err
	}
	var result []*spec.CTF
	for _, ctf := range ctfs {
		rows, err := s.db.QueryContext(ctx, `SELECT challenge_id, custom_score FROM ctf_challenges WHERE ctf_id = $1`, ctf.ID)
		if err != nil {
			return nil, err
		}
		var challenges []*spec.CTFChallenge
		for rows.Next() {
			var cid string
			var customScore sql.NullInt64
			if err := rows.Scan(&cid, &customScore); err != nil {
				rows.Close()
				return nil, err
			}
			var csPtr *int = nil
			if customScore.Valid {
				v := int(customScore.Int64)
				csPtr = &v
			}
			challenges = append(challenges, &spec.CTFChallenge{
				ID:          cid,
				CustomScore: csPtr,
			})
		}
		rows.Close()
		result = append(result, &spec.CTF{
			ID:          ctf.ID,
			Name:        ctf.Name,
			Description: ctf.Description,
			StartTime:   ctf.StartTime.Format(time.RFC3339),
			EndTime:     ctf.EndTime.Format(time.RFC3339),
			Slug:        ctf.Slug,
			Challenges:  challenges,
		})
	}
	return result, nil
}

func (s *service) UpdateCTF(ctx context.Context, req *spec.UpdateCTFPayload) (*spec.CTF, error) {
	ctf, err := models.FindCTF(ctx, s.db, req.ID)
	if err != nil {
		return nil, err
	}
	ctf.Name = req.Name
	ctf.Description = req.Description
	ctf.StartTime = time.Unix(req.StartTime, 0)
	ctf.EndTime = time.Unix(req.EndTime, 0)

	// Check if slug is unique
	existingCTF, err := models.CTFS(models.CTFWhere.Slug.EQ(req.Slug)).One(ctx, s.db)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if existingCTF != nil && existingCTF.ID != ctf.ID {
		return nil, errors.New("slug already exists")
	}
	ctf.Slug = req.Slug

	_, err = ctf.Update(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	_, err = s.db.ExecContext(ctx, `DELETE FROM ctf_challenges WHERE ctf_id = $1`, ctf.ID)
	if err != nil {
		return nil, err
	}
	for _, chall := range req.Challenges {
		if chall.CustomScore != nil {
			_, err = s.db.ExecContext(ctx, `INSERT INTO ctf_challenges (ctf_id, challenge_id, custom_score) VALUES ($1, $2, $3)`, ctf.ID, chall.ID, *chall.CustomScore)
		} else {
			_, err = s.db.ExecContext(ctx, `INSERT INTO ctf_challenges (ctf_id, challenge_id) VALUES ($1, $2)`, ctf.ID, chall.ID)
		}
		if err != nil {
			return nil, err
		}
	}
	return &spec.CTF{
		ID:          ctf.ID,
		Name:        ctf.Name,
		Description: ctf.Description,
		StartTime:   ctf.StartTime.Format(time.RFC3339),
		EndTime:     ctf.EndTime.Format(time.RFC3339),
		Slug:        ctf.Slug,
		Challenges:  req.Challenges,
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
			_, err = s.db.ExecContext(ctx, `INSERT INTO challenge_group_challenges (challenge_group_id, challenge_id, custom_score) VALUES ($1, $2, $3)`, id, chall.ID, *chall.CustomScore)
		} else {
			_, err = s.db.ExecContext(ctx, `INSERT INTO challenge_group_challenges (challenge_group_id, challenge_id) VALUES ($1, $2)`, id, chall.ID)
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
			_, err = s.db.ExecContext(ctx, `INSERT INTO challenge_group_challenges (challenge_group_id, challenge_id, custom_score) VALUES ($1, $2, $3)`, group.ID, chall.ID, *chall.CustomScore)
		} else {
			_, err = s.db.ExecContext(ctx, `INSERT INTO challenge_group_challenges (challenge_group_id, challenge_id) VALUES ($1, $2)`, group.ID, chall.ID)
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
		rows, err := s.db.QueryContext(ctx, `SELECT challenge_id, custom_score FROM challenge_group_challenges WHERE challenge_group_id = $1`, group.ID)
		if err != nil {
			return nil, err
		}
		var challenges []*spec.CTFChallenge
		for rows.Next() {
			var cid string
			var customScore sql.NullInt64
			if err := rows.Scan(&cid, &customScore); err != nil {
				rows.Close()
				return nil, err
			}
			var csPtr *int = nil
			if customScore.Valid {
				v := int(customScore.Int64)
				csPtr = &v
			}
			challenges = append(challenges, &spec.CTFChallenge{
				ID:          cid,
				CustomScore: csPtr,
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
