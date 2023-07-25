// Code generated by goa v3.12.3, DO NOT EDIT.
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

// PresignChallFileUploadAdminPath returns the URL path to the admin service PresignChallFileUpload HTTP endpoint.
func PresignChallFileUploadAdminPath(challengeID string) string {
	return fmt.Sprintf("/admin/challenges/%v/file_url", challengeID)
}

// ListMonthlyChallengesAdminPath returns the URL path to the admin service ListMonthlyChallenges HTTP endpoint.
func ListMonthlyChallengesAdminPath() string {
	return "/admin/monthly_challenges"
}

// DeleteMonthlyChallengeAdminPath returns the URL path to the admin service DeleteMonthlyChallenge HTTP endpoint.
func DeleteMonthlyChallengeAdminPath(challengeID string) string {
	return fmt.Sprintf("/admin/monthly_challenges/%v", challengeID)
}

// DeleteFileAdminPath returns the URL path to the admin service DeleteFile HTTP endpoint.
func DeleteFileAdminPath(fileID string) string {
	return fmt.Sprintf("/admin/files/%v", fileID)
}

// CreateMonthlyChallengeAdminPath returns the URL path to the admin service CreateMonthlyChallenge HTTP endpoint.
func CreateMonthlyChallengeAdminPath() string {
	return "/admin/monthly_challenges"
}

// ListUsersAdminPath returns the URL path to the admin service ListUsers HTTP endpoint.
func ListUsersAdminPath() string {
	return "/admin/users"
}

// AddFlagAdminPath returns the URL path to the admin service AddFlag HTTP endpoint.
func AddFlagAdminPath(challengeID string) string {
	return fmt.Sprintf("/admin/challenges/%v/flags", challengeID)
}

// DeleteFlagAdminPath returns the URL path to the admin service DeleteFlag HTTP endpoint.
func DeleteFlagAdminPath(challengeID string, flagID string) string {
	return fmt.Sprintf("/admin/challenges/%v/flags/%v", challengeID, flagID)
}
