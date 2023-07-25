// Code generated by goa v3.12.3, DO NOT EDIT.
//
// HTTP request path constructors for the challenge service.
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package server

import (
	"fmt"
)

// ListChallengesChallengePath returns the URL path to the challenge service ListChallenges HTTP endpoint.
func ListChallengesChallengePath() string {
	return "/challenges"
}

// ListMonthlyChallengesChallengePath returns the URL path to the challenge service ListMonthlyChallenges HTTP endpoint.
func ListMonthlyChallengesChallengePath() string {
	return "/monthly_challenges"
}

// SubmitFlagChallengePath returns the URL path to the challenge service SubmitFlag HTTP endpoint.
func SubmitFlagChallengePath(challengeID string) string {
	return fmt.Sprintf("/challenges/%v/attempt", challengeID)
}
