// Code generated by goa v3.5.2, DO NOT EDIT.
//
// HTTP request path constructors for the admin service.
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package server

import (
	"fmt"
)

// ListChallengesAdminPath returns the URL path to the admin service ListChallenges HTTP endpoint.
func ListChallengesAdminPath() string {
	return "/admin/challenges"
}

// CreateChallengeAdminPath returns the URL path to the admin service CreateChallenge HTTP endpoint.
func CreateChallengeAdminPath() string {
	return "/admin/challenges"
}

// ListMonthlyChallengesAdminPath returns the URL path to the admin service ListMonthlyChallenges HTTP endpoint.
func ListMonthlyChallengesAdminPath() string {
	return "/admin/monthly_challenges"
}

// DeleteMonthlyChallengeAdminPath returns the URL path to the admin service DeleteMonthlyChallenge HTTP endpoint.
func DeleteMonthlyChallengeAdminPath(monthlyChallengeID string) string {
	return fmt.Sprintf("/admin/monthly_challenges/%v", monthlyChallengeID)
}

// CreateMonthlyChallengeAdminPath returns the URL path to the admin service CreateMonthlyChallenge HTTP endpoint.
func CreateMonthlyChallengeAdminPath() string {
	return "/admin/monthly_challenges"
}
