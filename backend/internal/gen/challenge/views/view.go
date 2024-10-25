// Code generated by goa v3.5.2, DO NOT EDIT.
//
// challenge views
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// SsmChallengeCollection is the viewed result type that is projected based on
// a view.
type SsmChallengeCollection struct {
	// Type to project
	Projected SsmChallengeCollectionView
	// View to render
	View string
}

// SsmUserMonthlyChallenge is the viewed result type that is projected based on
// a view.
type SsmUserMonthlyChallenge struct {
	// Type to project
	Projected *SsmUserMonthlyChallengeView
	// View to render
	View string
}

// SsmUserMonthlyChallengeCollection is the viewed result type that is
// projected based on a view.
type SsmUserMonthlyChallengeCollection struct {
	// Type to project
	Projected SsmUserMonthlyChallengeCollectionView
	// View to render
	View string
}

// SsmSchoolScoreboard is the viewed result type that is projected based on a
// view.
type SsmSchoolScoreboard struct {
	// Type to project
	Projected *SsmSchoolScoreboardView
	// View to render
	View string
}

// SsmUserScoreboard is the viewed result type that is projected based on a
// view.
type SsmUserScoreboard struct {
	// Type to project
	Projected *SsmUserScoreboardView
	// View to render
	View string
}

// SsmChallengeCollectionView is a type that runs validations on a projected
// type.
type SsmChallengeCollectionView []*SsmChallengeView

// SsmChallengeView is a type that runs validations on a projected type.
type SsmChallengeView struct {
	// ID of a file
	ID *string
	// A unique string that can be used in URLs
	Slug *string
	// Title displayed to user
	Title *string
	// A short text describing the challenge
	Description *string
	// The number of points given to the solver
	Score    *int
	Services []*ChallengeServiceView
	Files    []*ChallengeFilesView
	// The numer of people who solved the challenge
	Solves *int
	// The ID of the CTF the challenge was taken from
	CtfEventID     *string
	ChallNamespace *string
	// whether the user has solved the challenge or not
	Solved   *bool
	Category *string
	Authors  []*AuthorView
	Solvers  []*SsmSolverView
}

// ChallengeServiceView is a type that runs validations on a projected type.
type ChallengeServiceView struct {
	UserDisplay *string
	Hyperlink   *bool
}

// ChallengeFilesView is a type that runs validations on a projected type.
type ChallengeFilesView struct {
	Filename *string
	URL      *string
}

// AuthorView is a type that runs validations on a projected type.
type AuthorView struct {
	FullName    *string
	Description *string
	Sponsor     *bool
	Slug        *string
	ImageURL    *string
	Publish     *bool
	// ID of a file
	ID *string
}

// SsmSolverView is a type that runs validations on a projected type.
type SsmSolverView struct {
	ID       *string
	FullName *string
	SolvedAt *int64
}

// SsmUserMonthlyChallengeView is a type that runs validations on a projected
// type.
type SsmUserMonthlyChallengeView struct {
	ChallengeID *string
	// The month(s) that the challenge is assigned for
	DisplayMonth *string
	// Starting date of the monthly challenge
	StartDate *int64
	// Ending date of the monthly challenge
	EndDate   *int64
	Challenge *SsmChallengeView
}

// SsmUserMonthlyChallengeCollectionView is a type that runs validations on a
// projected type.
type SsmUserMonthlyChallengeCollectionView []*SsmUserMonthlyChallengeView

// SsmSchoolScoreboardView is a type that runs validations on a projected type.
type SsmSchoolScoreboardView struct {
	Scores []*SchoolScoreboardScoreView
}

// SchoolScoreboardScoreView is a type that runs validations on a projected
// type.
type SchoolScoreboardScoreView struct {
	Score        *int
	SchoolName   *string
	IsUniversity *bool
}

// SsmUserScoreboardView is a type that runs validations on a projected type.
type SsmUserScoreboardView struct {
	Scores []*UserScoreboardScoreView
}

// UserScoreboardScoreView is a type that runs validations on a projected type.
type UserScoreboardScoreView struct {
	UserID     *string
	Name       *string
	SchoolName *string
	Score      *int
}

var (
	// SsmChallengeCollectionMap is a map indexing the attribute names of
	// SsmChallengeCollection by view name.
	SsmChallengeCollectionMap = map[string][]string{
		"default": {
			"id",
			"slug",
			"title",
			"description",
			"score",
			"services",
			"files",
			"solves",
			"ctf_event_id",
			"chall_namespace",
			"solved",
			"category",
			"authors",
			"solvers",
		},
	}
	// SsmUserMonthlyChallengeMap is a map indexing the attribute names of
	// SsmUserMonthlyChallenge by view name.
	SsmUserMonthlyChallengeMap = map[string][]string{
		"default": {
			"challenge_id",
			"display_month",
			"start_date",
			"end_date",
			"challenge",
		},
	}
	// SsmUserMonthlyChallengeCollectionMap is a map indexing the attribute names
	// of SsmUserMonthlyChallengeCollection by view name.
	SsmUserMonthlyChallengeCollectionMap = map[string][]string{
		"default": {
			"challenge_id",
			"display_month",
			"start_date",
			"end_date",
			"challenge",
		},
	}
	// SsmSchoolScoreboardMap is a map indexing the attribute names of
	// SsmSchoolScoreboard by view name.
	SsmSchoolScoreboardMap = map[string][]string{
		"default": {
			"scores",
		},
	}
	// SsmUserScoreboardMap is a map indexing the attribute names of
	// SsmUserScoreboard by view name.
	SsmUserScoreboardMap = map[string][]string{
		"default": {
			"scores",
		},
	}
	// SsmChallengeMap is a map indexing the attribute names of SsmChallenge by
	// view name.
	SsmChallengeMap = map[string][]string{
		"default": {
			"id",
			"slug",
			"title",
			"description",
			"score",
			"services",
			"files",
			"solves",
			"ctf_event_id",
			"chall_namespace",
			"solved",
			"category",
			"authors",
			"solvers",
		},
	}
	// SsmSolverMap is a map indexing the attribute names of SsmSolver by view name.
	SsmSolverMap = map[string][]string{
		"default": {
			"id",
			"full_name",
			"solved_at",
		},
	}
)

// ValidateSsmChallengeCollection runs the validations defined on the viewed
// result type SsmChallengeCollection.
func ValidateSsmChallengeCollection(result SsmChallengeCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateSsmChallengeCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateSsmUserMonthlyChallenge runs the validations defined on the viewed
// result type SsmUserMonthlyChallenge.
func ValidateSsmUserMonthlyChallenge(result *SsmUserMonthlyChallenge) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateSsmUserMonthlyChallengeView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateSsmUserMonthlyChallengeCollection runs the validations defined on
// the viewed result type SsmUserMonthlyChallengeCollection.
func ValidateSsmUserMonthlyChallengeCollection(result SsmUserMonthlyChallengeCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateSsmUserMonthlyChallengeCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateSsmSchoolScoreboard runs the validations defined on the viewed
// result type SsmSchoolScoreboard.
func ValidateSsmSchoolScoreboard(result *SsmSchoolScoreboard) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateSsmSchoolScoreboardView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateSsmUserScoreboard runs the validations defined on the viewed result
// type SsmUserScoreboard.
func ValidateSsmUserScoreboard(result *SsmUserScoreboard) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateSsmUserScoreboardView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateSsmChallengeCollectionView runs the validations defined on
// SsmChallengeCollectionView using the "default" view.
func ValidateSsmChallengeCollectionView(result SsmChallengeCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateSsmChallengeView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateSsmChallengeView runs the validations defined on SsmChallengeView
// using the "default" view.
func ValidateSsmChallengeView(result *SsmChallengeView) (err error) {
	if result.Solved == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("solved", "result"))
	}
	if result.Category == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("category", "result"))
	}
	if result.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "result"))
	}
	if result.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "result"))
	}
	if result.Score == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("score", "result"))
	}
	if result.Slug == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("slug", "result"))
	}
	if result.Solves == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("solves", "result"))
	}
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.id", *result.ID, goa.FormatUUID))
	}
	for _, e := range result.Services {
		if e != nil {
			if err2 := ValidateChallengeServiceView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	for _, e := range result.Files {
		if e != nil {
			if err2 := ValidateChallengeFilesView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.CtfEventID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.ctf_event_id", *result.CtfEventID, goa.FormatUUID))
	}
	for _, e := range result.Authors {
		if e != nil {
			if err2 := ValidateAuthorView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	for _, e := range result.Solvers {
		if e != nil {
			if err2 := ValidateSsmSolverView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateChallengeServiceView runs the validations defined on
// ChallengeServiceView.
func ValidateChallengeServiceView(result *ChallengeServiceView) (err error) {
	if result.UserDisplay == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_display", "result"))
	}
	if result.Hyperlink == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("hyperlink", "result"))
	}
	return
}

// ValidateChallengeFilesView runs the validations defined on
// ChallengeFilesView.
func ValidateChallengeFilesView(result *ChallengeFilesView) (err error) {
	if result.Filename == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("filename", "result"))
	}
	if result.URL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("url", "result"))
	}
	return
}

// ValidateAuthorView runs the validations defined on AuthorView.
func ValidateAuthorView(result *AuthorView) (err error) {
	if result.Sponsor == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("sponsor", "result"))
	}
	if result.FullName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("full_name", "result"))
	}
	if result.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "result"))
	}
	if result.Slug == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("slug", "result"))
	}
	if result.Publish == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("publish", "result"))
	}
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.id", *result.ID, goa.FormatUUID))
	}
	return
}

// ValidateSsmSolverView runs the validations defined on SsmSolverView using
// the "default" view.
func ValidateSsmSolverView(result *SsmSolverView) (err error) {
	if result.SolvedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("solved_at", "result"))
	}
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.FullName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("full_name", "result"))
	}
	return
}

// ValidateSsmUserMonthlyChallengeView runs the validations defined on
// SsmUserMonthlyChallengeView using the "default" view.
func ValidateSsmUserMonthlyChallengeView(result *SsmUserMonthlyChallengeView) (err error) {
	if result.StartDate == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("start_date", "result"))
	}
	if result.EndDate == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("end_date", "result"))
	}
	if result.DisplayMonth == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("display_month", "result"))
	}
	if result.ChallengeID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("challenge_id", "result"))
	}
	if result.ChallengeID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.challenge_id", *result.ChallengeID, goa.FormatUUID))
	}
	if result.Challenge != nil {
		if err2 := ValidateSsmChallengeView(result.Challenge); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateSsmUserMonthlyChallengeCollectionView runs the validations defined
// on SsmUserMonthlyChallengeCollectionView using the "default" view.
func ValidateSsmUserMonthlyChallengeCollectionView(result SsmUserMonthlyChallengeCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateSsmUserMonthlyChallengeView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateSsmSchoolScoreboardView runs the validations defined on
// SsmSchoolScoreboardView using the "default" view.
func ValidateSsmSchoolScoreboardView(result *SsmSchoolScoreboardView) (err error) {
	if result.Scores == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("scores", "result"))
	}
	for _, e := range result.Scores {
		if e != nil {
			if err2 := ValidateSchoolScoreboardScoreView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateSchoolScoreboardScoreView runs the validations defined on
// SchoolScoreboardScoreView.
func ValidateSchoolScoreboardScoreView(result *SchoolScoreboardScoreView) (err error) {
	if result.SchoolName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("school_name", "result"))
	}
	if result.Score == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("score", "result"))
	}
	if result.IsUniversity == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("is_university", "result"))
	}
	return
}

// ValidateSsmUserScoreboardView runs the validations defined on
// SsmUserScoreboardView using the "default" view.
func ValidateSsmUserScoreboardView(result *SsmUserScoreboardView) (err error) {
	if result.Scores == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("scores", "result"))
	}
	for _, e := range result.Scores {
		if e != nil {
			if err2 := ValidateUserScoreboardScoreView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateUserScoreboardScoreView runs the validations defined on
// UserScoreboardScoreView.
func ValidateUserScoreboardScoreView(result *UserScoreboardScoreView) (err error) {
	if result.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_id", "result"))
	}
	if result.SchoolName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("school_name", "result"))
	}
	if result.Score == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("score", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	return
}
