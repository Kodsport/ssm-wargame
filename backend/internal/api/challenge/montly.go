package challenge

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
	"github.com/google/uuid"
	"github.com/sakerhetsm/ssm-wargame/internal/auth"
	"github.com/sakerhetsm/ssm-wargame/internal/custommodels"
	spec "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	"github.com/sakerhetsm/ssm-wargame/internal/models"
	"github.com/sakerhetsm/ssm-wargame/internal/utils"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

const (
	discordGuildID       = "407149490307334144"
	discordMonthlyRoleID = "1452409956992155749"
)

type DiscordFooter struct {
	Text string `json:"text"`
}

type DiscordEmbed struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Color       int           `json:"color"`
	Footer      DiscordFooter `json:"footer"`
}

type WebhookPayload struct {
	Embeds    []DiscordEmbed `json:"embeds"`
	Username  string         `json:"username"`
	AvatarURL string         `json:"avatar_url"`
}

func (s *service) monthlyQuery(ctx context.Context) *queries.Query {
	q := models.NewQuery(
		qm.Select("monthly_challenges.*"),
		qm.Select("(SELECT COUNT(1) FROM user_solves WHERE challenge_id = monthly_challenges.challenge_id) num_solves"),
		qm.From("monthly_challenges"),
		qm.Load(qm.Rels(models.MonthlyChallengeRels.Challenge, models.ChallengeRels.ChallengeFiles)),
		qm.Load(qm.Rels(models.MonthlyChallengeRels.Challenge, models.ChallengeRels.ChallengeServices)),
		qm.Load(qm.Rels(models.MonthlyChallengeRels.Challenge, models.ChallengeRels.Category)),
		qm.Load(qm.Rels(models.MonthlyChallengeRels.Challenge, models.ChallengeRels.Authors)),
		models.MonthlyChallengeWhere.StartDate.LT(time.Now()),
	)

	if auth.IsAuthed(ctx) {
		qm.Select(
			"EXISTS(SELECT 1 FROM user_solves us2 WHERE us2.challenge_id = monthly_challenges.challenge_id AND us2.user_id = '"+auth.GetUser(ctx).ID+"') AS solved",
		).Apply(q)
	}
	return q
}

func convertMonthly(chall *custommodels.UserMonthlyChall) *spec.SsmUserMonthlyChallenge {
	res := &spec.SsmUserMonthlyChallenge{
		ChallengeID:  chall.ChallengeID,
		DisplayMonth: chall.DisplayMonth,
		StartDate:    chall.StartDate.Unix(),
		EndDate:      chall.EndDate.Unix(),
	}

	res.Challenge = &spec.SsmChallenge{
		ID:          chall.R.Challenge.ID,
		Title:       chall.R.Challenge.Title,
		Description: chall.R.Challenge.Description,
		Slug:        chall.R.Challenge.Slug,
		Category:    chall.R.Challenge.R.Category.Name,
		Solves:      chall.NumSolves,
		Solved:      chall.Solved,
	}

	res.Challenge.Files = make([]*spec.ChallengeFiles, len(chall.R.Challenge.R.ChallengeFiles))
	for i, file := range chall.R.Challenge.R.ChallengeFiles {
		res.Challenge.Files[i] = &spec.ChallengeFiles{
			Filename: file.FriendlyName,
			URL:      file.URL,
		}
	}

	res.Challenge.Authors = make([]*spec.Author, len(chall.R.Challenge.R.Authors))
	for i, v := range chall.R.Challenge.R.Authors {
		res.Challenge.Authors[i] = &spec.Author{
			ID:          v.ID,
			FullName:    v.FullName,
			Description: v.Description,
			Sponsor:     v.Sponsor,
			Slug:        v.Slug,
			ImageURL:    v.ImageURL.Ptr(),
			Publish:     v.Publish,
		}
	}
	res.Challenge.Services = make([]*spec.ChallengeService, len(chall.R.Challenge.R.ChallengeServices))
	for i, v := range chall.R.Challenge.R.ChallengeServices {
		res.Challenge.Services[i] = &spec.ChallengeService{
			UserDisplay: v.UserDisplay,
			Hyperlink:   v.Hyperlink,
		}
	}

	return res
}

// helper funcs
func (s *service) getDiscordSession() (*discordgo.Session, error) {
	botToken := os.Getenv("DISCORD_BOT_TOKEN")
	if botToken == "" {
		return nil, fmt.Errorf("DISCORD_BOT_TOKEN not set rookie")
	}
	return discordgo.New("Bot " + botToken)
}

func (s *service) checkUserInGuild(session *discordgo.Session, discordID string) bool {
	_, err := session.GuildMember(discordGuildID, discordID)
	return err == nil
}

func (s *service) addDiscordRole(session *discordgo.Session, discordID, roleID string) error {
	return session.GuildMemberRoleAdd(discordGuildID, discordID, roleID)
}

func (s *service) removeDiscordRole(session *discordgo.Session, discordID, roleID string) error {
	return session.GuildMemberRoleRemove(discordGuildID, discordID, roleID)
}

func (s *service) HandleMonthlySolve(ctx context.Context, user *models.User, challID uuid.UUID) error {
	now := time.Now()
	isActiveMonthly, err := models.MonthlyChallenges(
		models.MonthlyChallengeWhere.ChallengeID.EQ(challID.String()),
		models.MonthlyChallengeWhere.StartDate.LTE(now),
		models.MonthlyChallengeWhere.EndDate.GTE(now),
		qm.Load(models.MonthlyChallengeRels.Challenge),
	).Exists(ctx, s.db)
	if err != nil {
		return err
	}
	
	if !isActiveMonthly {
		return nil
	}

	monthlyChall, err := models.MonthlyChallenges(
		models.MonthlyChallengeWhere.ChallengeID.EQ(challID.String()),
		qm.Load(models.MonthlyChallengeRels.Challenge),
	).One(ctx, s.db)
	if err != nil {
		return fmt.Errorf("failed to get monthly challenge: %w", err)
	}
	
	monthlyName := monthlyChall.R.Challenge.Title

	if user.DiscordID.Valid {
		s.log.Info("user solved current monthly", 
			zap.String("user_id", user.ID), 
			zap.String("challenge_id", challID.String()), 
			zap.String("discord_id", user.DiscordID.String), 
			utils.C(ctx))
		
		go func() {
			session, err := s.getDiscordSession()
			if err != nil {
				s.log.Warn("failed to get discord session", zap.Error(err))
				return
			}
			
			userInDiscord := s.checkUserInGuild(session, user.DiscordID.String)
			
			if userInDiscord {
				if err := s.addDiscordRole(session, user.DiscordID.String, discordMonthlyRoleID); err != nil {
					s.log.Warn("failed to add discord role", zap.Error(err))
				}
			}
			
			s.sendDiscordWebhook(user.FullName, user.DiscordID.String, monthlyName, userInDiscord, challID)
		}()
	} else {
		s.log.Info("user solved current monthly but has no discord id", 
			zap.String("user_id", user.ID), 
			zap.String("challenge_id", challID.String()), 
			utils.C(ctx))
		go s.sendDiscordWebhook(user.FullName, "", monthlyName, false, challID)
	}
	
	return nil
}

func (s *service) sendDiscordWebhook(userName, discordID, monthlyName string, userInDiscord bool, challID uuid.UUID) {
	webhookURL := os.Getenv("DISCORD_WEBHOOK")
	if webhookURL == "" {
		return
	}

	ctx := context.Background()
	rand.Seed(time.Now().UnixNano())
	solveCount, err := models.UserSolves(qm.Where("challenge_id = ?", challID.String())).Count(ctx, s.db)
	if err != nil {
		s.log.Warn("failed to get solve count", zap.Error(err))
		return
	}

	monthly, err := models.MonthlyChallenges(models.MonthlyChallengeWhere.ChallengeID.EQ(challID.String())).One(ctx, s.db)
	if err != nil {
		s.log.Warn("failed to get monthly challenge", zap.Error(err))
		return
	}

	var mention string
	mention = userName // discord privacy smh
	/*
	if userInDiscord && discordID != "" {
		mention = fmt.Sprintf("<@%s>", discordID)
	} else {
		mention = userName
	}
		*/

	var title, description string
	var color int
	if solveCount == 1 {
		title = ":drop_of_blood: First Blood!"
		description = fmt.Sprintf("Grattis till **%s** som var först med att lösa **%s**!", mention, monthlyName)
		if 1 == 2 { // TODO: rand.Intn(10) == 0 efter första monthly webhooken.
			description = fmt.Sprintf("**%s** har fått första blodet på **%s**!", mention, monthlyName)
			title = ":drop_of_blood: Första Blodet!" // :P
		}
		color = 0xba2014
	} else {
		title = ":fire: Ny Månadslösning!"
		if rand.Intn(10) == 0 {
			title = ":wolf: Ny Månadslösning!" // :P
		}
		description = fmt.Sprintf("Grattis till **%s** som löst **%s**!", mention, monthlyName)
		color = 0x002c36
	}

	payload := WebhookPayload{
		Embeds: []DiscordEmbed{
			{
				Title:       title,
				Description: description,
				Color:       color,
				Footer: DiscordFooter{
					Text: fmt.Sprintf("%s Solve (%s)", humanize.Ordinal(int(solveCount)), monthly.DisplayMonth),
				},
			},
		},
		Username:  "Månadens problem",
		AvatarURL: "https://sakerhetssm.se/kodsport.png",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		s.log.Warn("failed to marshal webhook data", zap.Error(err))
		return
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		s.log.Warn("failed to send discord webhook", zap.Error(err))
		return
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body := make([]byte, 1024)
		n, _ := resp.Body.Read(body)
		s.log.Warn("discord webhook failed", 
			zap.Int("status", resp.StatusCode), 
			zap.String("body", string(body[:n])))
	}
}

func (s *service) ClearMonthlyRolesForChallenge(ctx context.Context, challengeID string) error {
	solvers, err := models.Users(
		qm.InnerJoin("user_solves us ON us.user_id = users.id"),
		qm.Where("us.challenge_id = ?", challengeID),
		models.UserWhere.DiscordID.IsNotNull(),
	).All(ctx, s.db)
	if err != nil {
		return fmt.Errorf("failed to get solvers for challenge %s: %w", challengeID, err)
	}

	session, err := s.getDiscordSession()
	if err != nil {
		return fmt.Errorf("failed to get discord session: %w", err)
	}

	for _, user := range solvers {
		if err := s.removeDiscordRole(session, user.DiscordID.String, discordMonthlyRoleID); err != nil {
			s.log.Warn("failed to remove discord role", 
				zap.Error(err), 
				zap.String("discord_id", user.DiscordID.String), 
				utils.C(ctx))
		} else {
			s.log.Info("discord role removed", 
				zap.String("discord_id", user.DiscordID.String), 
				zap.String("role_id", discordMonthlyRoleID), 
				utils.C(ctx))
		}
		time.Sleep(100 * time.Millisecond)
	}
	
	return nil
}

func (s *service) ClearMonthlyRolesForPrevious(ctx context.Context, previousChallengeID, currentChallengeID string) error {
	previousSolvers, err := models.Users(
		qm.InnerJoin("user_solves us ON us.user_id = users.id"),
		qm.Where("us.challenge_id = ?", previousChallengeID),
		models.UserWhere.DiscordID.IsNotNull(),
	).All(ctx, s.db)
	if err != nil {
		return fmt.Errorf("failed to get previous solvers: %w", err)
	}

	currentSolvers, err := models.Users(
		qm.InnerJoin("user_solves us ON us.user_id = users.id"),
		qm.Where("us.challenge_id = ?", currentChallengeID),
		models.UserWhere.DiscordID.IsNotNull(),
	).All(ctx, s.db)
	if err != nil {
		return fmt.Errorf("failed to get current solvers: %w", err)
	}

	currentSolverDiscordIDs := make(map[string]bool)
	for _, user := range currentSolvers {
		currentSolverDiscordIDs[user.DiscordID.String] = true
	}

	session, err := s.getDiscordSession()
	if err != nil {
		return fmt.Errorf("failed to get discord session: %w", err)
	}

	g, gctx := errgroup.WithContext(ctx)
	for _, user := range previousSolvers {
		if currentSolverDiscordIDs[user.DiscordID.String] {
			continue
		}
		
		user := user
		g.Go(func() error {
			if err := s.removeDiscordRole(session, user.DiscordID.String, discordMonthlyRoleID); err != nil {
				s.log.Info("discord role removal issue", 
					zap.String("discord_id", user.DiscordID.String), 
					zap.Error(err),
					utils.C(gctx))
			} else {
				s.log.Info("discord role removed", 
					zap.String("discord_id", user.DiscordID.String), 
					utils.C(gctx))
			}
			return nil
		})
	}

	return g.Wait()
}