package ctf

import (
	"bufio"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sakerhetsm/ssm-wargame/internal/custommodels"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/ctf"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
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
	return &spec.CTFInfo{
		ID:          ctf.ID,
		Name:        ctf.Name,
		Description: ctf.Description,
		StartTime:   ctf.StartTime.Format(time.RFC3339),
		EndTime:     ctf.EndTime.Format(time.RFC3339),
		Slug:        ctf.Slug,
	}, nil
}

func (s *Service) RegisterUser(ctx context.Context, req *spec.RegisterUserPayload) (*spec.CTFUser, error) {

	if len(req.Username) < 3 || len(req.Username) > 30 {
		return nil, errors.New("username must be between 3 and 30 characters")
	}

	ctf, err := models.CTFS(models.CTFWhere.Slug.EQ(req.Slug)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}

	// Check if username exists for this ctf
	var exists bool
	err = s.db.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM ctf_users WHERE ctf_id = $1 AND username = $2)`, ctf.ID, req.Username).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("username already taken for this ctf")
	}

	// Generate a passphrase for the user and check that it is unique so it
	// can be used to login
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

	return &spec.CTFUser{
		ID:       user.ID,
		CtfID:    user.CTFID,
		Username: user.Username,
		Password: user.Password,
	}, nil
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
	return &spec.CTFUser{
		ID:       user.ID,
		CtfID:    user.CTFID,
		Username: user.Username,
		Password: user.Password,
	}, nil
}

func (s *Service) GetUserSolves(ctx context.Context, req *spec.GetUserSolvesPayload) (*spec.CTFUserSolves, error) {
	ctf, err := models.CTFS(models.CTFWhere.Slug.EQ(req.Slug)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}
	user, err := models.CTFUsers(models.CTFUserWhere.ID.EQ(req.ID), models.CTFUserWhere.CTFID.EQ(ctf.ID)).One(ctx, s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found for this ctf")
		}
		return nil, err
	}
	solves, err := models.CTFSolves(models.CTFSolfWhere.UserID.EQ(user.ID), models.CTFSolfWhere.CTFID.EQ(ctf.ID)).All(ctx, s.db)
	if err != nil {
		return nil, err
	}
	res := make([]*spec.CTFUserSolve, len(solves))
	for i, solve := range solves {
		res[i] = &spec.CTFUserSolve{
			ChallengeID: solve.ChallengeID,
			SolvedAt:    solve.CreatedAt.Format(time.RFC3339),
		}
	}
	return &spec.CTFUserSolves{
		Username: user.Username,
		Solves:   res,
	}, nil
}

func (s *Service) ListChallenges(ctx context.Context, req *spec.ListChallengesPayload) (spec.SsmChallengeCollection, error) {
	ctf, err := models.CTFS(models.CTFWhere.Slug.EQ(req.Slug)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}

	if ctf.StartTime.After(time.Now()) {
		return nil, errors.New("ctf not started yet")
	}

	rows, err := s.db.QueryContext(ctx, `SELECT challenge_id, custom_score FROM ctf_challenges WHERE ctf_id = $1`, ctf.ID)
	if err != nil {
		return nil, err
	}

	customScores := make(map[string]*int)
	for rows.Next() {
		var cid string
		var customScore sql.NullInt64
		if err := rows.Scan(&cid, &customScore); err != nil {
			return nil, err
		}
		if customScore.Valid {
			v := int(customScore.Int64)
			customScores[cid] = &v
		} else {
			customScores[cid] = nil
		}
	}
	rows.Close()

	var challengeIDs []string
	for cid := range customScores {
		challengeIDs = append(challengeIDs, cid)
	}

	challs := make([]*custommodels.UserChall, 0)
	q := models.NewQuery(
		qm.Select("challenges.*, categories.name as category"),
		qm.Select("(SELECT COUNT(1) FROM ctf_solves WHERE challenge_id = challenges.id AND ctf_id = '"+ctf.ID+"') num_solves"),
		qm.From(models.TableNames.Challenges),
		qm.InnerJoin("categories ON categories.id = challenges.category_id"),
		qm.Load(models.ChallengeRels.ChallengeFiles),
		qm.Load(models.ChallengeRels.ChallengeServices),
		qm.Load(models.ChallengeRels.Authors),
		qm.Load(qm.Rels(models.ChallengeRels.UserSolves, models.UserSolfRels.User)),
	)

	models.ChallengeWhere.ID.IN(challengeIDs).Apply(q)

	if req.Password != nil {
		user, err := models.CTFUsers(models.CTFUserWhere.Password.EQ(*req.Password), models.CTFUserWhere.CTFID.EQ(ctf.ID)).One(ctx, s.db)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("user not found for this ctf")
			}
			return nil, err
		}
		qm.Select(
			"EXISTS(SELECT 1 FROM ctf_solves us2 WHERE us2.challenge_id = challenges.id AND ctf_id = '" + ctf.ID + "' AND us2.user_id = '" + user.ID + "') AS solved",
		).Apply(q)
	}

	err = q.Bind(ctx, s.db, &challs)
	if err != nil {
		return nil, err
	}
	res := make(spec.SsmChallengeCollection, len(challs))

	for i, chall := range challs {
		score := chall.StaticScore.Int
		if cs, ok := customScores[chall.ID]; ok && cs != nil {
			score = *cs
		} else if !chall.StaticScore.Valid {
			score = dynamicScore(500, 100, float64(chall.NumSolves))
		}

		res[i] = &spec.SsmChallenge{
			ID:          chall.ID,
			Slug:        chall.Slug,
			Title:       chall.Title,
			Description: chall.Description,
			Score:       score,
			Solves:      chall.NumSolves,

			Solved:         chall.Solved,
			Category:       chall.Category,
			CtfEventID:     chall.CTFEventID.Ptr(),
			ChallNamespace: chall.ChallNamespace.Ptr(),
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

		sort.SliceStable(chall.R.CTFSolves, func(x, j int) bool {
			return chall.R.CTFSolves[x].CreatedAt.Before(chall.R.CTFSolves[j].CreatedAt)
		})

		res[i].Solvers = make([]*spec.SsmSolver, 0, 5)
		for i2, v := range chall.R.CTFSolves {
			if i2 == 5 {
				break
			}
			res[i].Solvers = append(res[i].Solvers, &spec.SsmSolver{
				ID:       v.UserID,
				FullName: v.R.User.Username,
				SolvedAt: v.CreatedAt.Unix(),
			})
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

	rows, err := s.db.QueryContext(ctx, `
		SELECT u.id, u.username, COALESCE(SUM(COALESCE(cc.custom_score, ch.static_score)), 0) as score,
			MIN(s.created_at) as first_solve
		FROM ctf_users u
		LEFT JOIN ctf_solves s ON u.id = s.user_id AND s.ctf_id = $1
		LEFT JOIN challenges ch ON s.challenge_id = ch.id
		LEFT JOIN ctf_challenges cc ON cc.ctf_id = $1 AND cc.challenge_id = ch.id
		WHERE u.ctf_id = $1
		GROUP BY u.id, u.username
		ORDER BY score DESC, first_solve ASC NULLS LAST, u.username ASC
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
		var firstSolve sql.NullTime
		if err := rows.Scan(&userID, &username, &score, &firstSolve); err != nil {
			return nil, err
		}
		solvesRows, err := s.db.QueryContext(ctx, `
			SELECT s.challenge_id
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

func (s *Service) SubmitFlag(ctx context.Context, req *spec.SubmitFlagPayload) (*spec.CTFSolve, error) {
	ctf, err := models.CTFS(models.CTFWhere.Slug.EQ(req.Slug)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}

	if ctf.StartTime.After(time.Now()) || ctf.EndTime.Before(time.Now()) {
		return nil, errors.New("ctf not active")
	}

	var userID string
	err = s.db.QueryRowContext(ctx, `SELECT id FROM ctf_users WHERE ctf_id = $1 AND password = $2`, ctf.ID, req.Password).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not registered for this ctf")
		}
		return nil, err
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

func getLine(filename string, lineNum int) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentLine := 0
	for scanner.Scan() {
		if currentLine == lineNum {
			return scanner.Text(), nil
		}
		currentLine++
	}
	return "", fmt.Errorf("line %d not found", lineNum)
}

func generatePassphrase(numWords int) (string, error) {
	rand.Seed(time.Now().UnixNano())

	const totalWords = 7776
	words := make([]string, numWords)

	for i := 0; i < numWords; i++ {
		lineNum := rand.Intn(totalWords)
		word, err := getLine("eff-long.txt", lineNum)
		if err != nil {
			return "", err
		}
		words[i] = word
	}

	return strings.Join(words, "-"), nil
}
