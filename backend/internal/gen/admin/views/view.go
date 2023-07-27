// Code generated by goa v3.5.2, DO NOT EDIT.
//
// admin views
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// SsmAdminChallengeCollection is the viewed result type that is projected
// based on a view.
type SsmAdminChallengeCollection struct {
	// Type to project
	Projected SsmAdminChallengeCollectionView
	// View to render
	View string
}

// SsmAdminChallengeCollectionView is a type that runs validations on a
// projected type.
type SsmAdminChallengeCollectionView []*SsmAdminChallengeView

// SsmAdminChallengeView is a type that runs validations on a projected type.
type SsmAdminChallengeView struct {
	ID *string
	// A unique string that can be used in URLs
	Slug *string
	// Title displayed to user
	Title *string
	// A short text describing the challenge
	Description *string
	// The number of points given to the solver
	Score     *int32
	Services  []*ChallengeServiceView
	Files     []*AdminChallengeFilesView
	Published *bool
	// The numer of people who solved the challenge
	Solves *int64
	Flags  []*AdminChallengeFlagView
}

// ChallengeServiceView is a type that runs validations on a projected type.
type ChallengeServiceView struct {
}

// AdminChallengeFilesView is a type that runs validations on a projected type.
type AdminChallengeFilesView struct {
	ID       *string
	Filename *string
	URL      *string
	Bucket   *string
	Key      *string
	Size     *int64
	// MD5 hash of the file content in base64
	Md5 *string
}

// AdminChallengeFlagView is a type that runs validations on a projected type.
type AdminChallengeFlagView struct {
	ID   *string
	Flag *string
}

var (
	// SsmAdminChallengeCollectionMap is a map indexing the attribute names of
	// SsmAdminChallengeCollection by view name.
	SsmAdminChallengeCollectionMap = map[string][]string{
		"default": {
			"id",
			"slug",
			"title",
			"description",
			"score",
			"services",
			"files",
			"published",
			"solves",
			"flags",
		},
	}
	// SsmAdminChallengeMap is a map indexing the attribute names of
	// SsmAdminChallenge by view name.
	SsmAdminChallengeMap = map[string][]string{
		"default": {
			"id",
			"slug",
			"title",
			"description",
			"score",
			"services",
			"files",
			"published",
			"solves",
			"flags",
		},
	}
)

// ValidateSsmAdminChallengeCollection runs the validations defined on the
// viewed result type SsmAdminChallengeCollection.
func ValidateSsmAdminChallengeCollection(result SsmAdminChallengeCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateSsmAdminChallengeCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateSsmAdminChallengeCollectionView runs the validations defined on
// SsmAdminChallengeCollectionView using the "default" view.
func ValidateSsmAdminChallengeCollectionView(result SsmAdminChallengeCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateSsmAdminChallengeView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateSsmAdminChallengeView runs the validations defined on
// SsmAdminChallengeView using the "default" view.
func ValidateSsmAdminChallengeView(result *SsmAdminChallengeView) (err error) {
	if result.Files == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("files", "result"))
	}
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
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
	if result.Published == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("published", "result"))
	}
	if result.Slug == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("slug", "result"))
	}
	if result.Solves == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("solves", "result"))
	}
	for _, e := range result.Files {
		if e != nil {
			if err2 := ValidateAdminChallengeFilesView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	for _, e := range result.Flags {
		if e != nil {
			if err2 := ValidateAdminChallengeFlagView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateChallengeServiceView runs the validations defined on
// ChallengeServiceView.
func ValidateChallengeServiceView(result *ChallengeServiceView) (err error) {

	return
}

// ValidateAdminChallengeFilesView runs the validations defined on
// AdminChallengeFilesView.
func ValidateAdminChallengeFilesView(result *AdminChallengeFilesView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Filename == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("filename", "result"))
	}
	if result.URL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("url", "result"))
	}
	if result.Key == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("key", "result"))
	}
	if result.Bucket == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("bucket", "result"))
	}
	if result.Size == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("size", "result"))
	}
	if result.Md5 == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("md5", "result"))
	}
	return
}

// ValidateAdminChallengeFlagView runs the validations defined on
// AdminChallengeFlagView.
func ValidateAdminChallengeFlagView(result *AdminChallengeFlagView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Flag == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("flag", "result"))
	}
	return
}
