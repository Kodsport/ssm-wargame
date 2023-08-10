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

// GetChallengeMetaAdminPath returns the URL path to the admin service GetChallengeMeta HTTP endpoint.
func GetChallengeMetaAdminPath(challengeID string) string {
	return fmt.Sprintf("/admin/challenges/%v", challengeID)
}

// CreateChallengeAdminPath returns the URL path to the admin service CreateChallenge HTTP endpoint.
func CreateChallengeAdminPath() string {
	return "/admin/challenges"
}

// UpdateChallengeAdminPath returns the URL path to the admin service UpdateChallenge HTTP endpoint.
func UpdateChallengeAdminPath(challengeID string) string {
	return fmt.Sprintf("/admin/challenges/%v", challengeID)
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

// ListAuthorsAdminPath returns the URL path to the admin service ListAuthors HTTP endpoint.
func ListAuthorsAdminPath() string {
	return "/admin/authors"
}

// UpdateAuthorAdminPath returns the URL path to the admin service UpdateAuthor HTTP endpoint.
func UpdateAuthorAdminPath(id string) string {
	return fmt.Sprintf("/admin/authors/%v", id)
}

// AddFlagAdminPath returns the URL path to the admin service AddFlag HTTP endpoint.
func AddFlagAdminPath(challengeID string) string {
	return fmt.Sprintf("/admin/challenges/%v/flags", challengeID)
}

// DeleteFlagAdminPath returns the URL path to the admin service DeleteFlag HTTP endpoint.
func DeleteFlagAdminPath(challengeID string, flagID string) string {
	return fmt.Sprintf("/admin/challenges/%v/flags/%v", challengeID, flagID)
}

// ListCategoriesAdminPath returns the URL path to the admin service ListCategories HTTP endpoint.
func ListCategoriesAdminPath() string {
	return "/admin/categories"
}

// ChalltoolsImportAdminPath returns the URL path to the admin service ChalltoolsImport HTTP endpoint.
func ChalltoolsImportAdminPath() string {
	return "/admin/push_challenge"
}

// ListCTFEventsAdminPath returns the URL path to the admin service ListCTFEvents HTTP endpoint.
func ListCTFEventsAdminPath() string {
	return "/admin/events"
}

// CreateCTFEventAdminPath returns the URL path to the admin service CreateCTFEvent HTTP endpoint.
func CreateCTFEventAdminPath() string {
	return "/admin/events"
}

// DeleteCTFEventAdminPath returns the URL path to the admin service DeleteCTFEvent HTTP endpoint.
func DeleteCTFEventAdminPath(id string) string {
	return fmt.Sprintf("/admin/events/%v", id)
}

// CreateCTFEventImportTokenAdminPath returns the URL path to the admin service CreateCTFEventImportToken HTTP endpoint.
func CreateCTFEventImportTokenAdminPath() string {
	return "/admin/import_token"
}
