// Code generated by goa v3.5.2, DO NOT EDIT.
//
// admin service
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package admin

import (
	"context"

	adminviews "github.com/sakerhetsm/ssm-wargame/internal/gen/admin/views"
	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Service is the admin service interface.
type Service interface {
	// ListChallenges implements ListChallenges.
	ListChallenges(context.Context, *ListChallengesPayload) (res SsmAdminChallengeCollection, err error)
	// CreateChallenge implements CreateChallenge.
	CreateChallenge(context.Context, *CreateChallengePayload) (err error)
	// UpdateChallenge implements UpdateChallenge.
	UpdateChallenge(context.Context, *UpdateChallengePayload) (err error)
	// PresignChallFileUpload implements PresignChallFileUpload.
	PresignChallFileUpload(context.Context, *PresignChallFileUploadPayload) (res *PresignChallFileUploadResult, err error)
	// ListMonthlyChallenges implements ListMonthlyChallenges.
	ListMonthlyChallenges(context.Context, *ListMonthlyChallengesPayload) (res []*MonthlyChallenge, err error)
	// DeleteMonthlyChallenge implements DeleteMonthlyChallenge.
	DeleteMonthlyChallenge(context.Context, *DeleteMonthlyChallengePayload) (err error)
	// DeleteFile implements DeleteFile.
	DeleteFile(context.Context, *DeleteFilePayload) (err error)
	// CreateMonthlyChallenge implements CreateMonthlyChallenge.
	CreateMonthlyChallenge(context.Context, *CreateMonthlyChallengePayload) (err error)
	// ListUsers implements ListUsers.
	ListUsers(context.Context, *ListUsersPayload) (res []*SsmUser, err error)
	// AddFlag implements AddFlag.
	AddFlag(context.Context, *AddFlagPayload) (err error)
	// DeleteFlag implements DeleteFlag.
	DeleteFlag(context.Context, *DeleteFlagPayload) (err error)
	// ListCategories implements ListCategories.
	ListCategories(context.Context, *ListCategoriesPayload) (res []*Category, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "admin"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [12]string{"ListChallenges", "CreateChallenge", "UpdateChallenge", "PresignChallFileUpload", "ListMonthlyChallenges", "DeleteMonthlyChallenge", "DeleteFile", "CreateMonthlyChallenge", "ListUsers", "AddFlag", "DeleteFlag", "ListCategories"}

// ListChallengesPayload is the payload type of the admin service
// ListChallenges method.
type ListChallengesPayload struct {
	Token string
}

// SsmAdminChallengeCollection is the result type of the admin service
// ListChallenges method.
type SsmAdminChallengeCollection []*SsmAdminChallenge

// CreateChallengePayload is the payload type of the admin service
// CreateChallenge method.
type CreateChallengePayload struct {
	// A unique string that can be used in URLs
	Slug string
	// Title displayed to user
	Title string
	// A short text describing the challenge
	Description string
	// The number of points given to the solver
	Score int
	// unix timestamp
	PublishAt *int64
	// The ID of the CTF the challenge was taken from
	CtfEventID *string
	CategoryID string
	Token      string
}

// UpdateChallengePayload is the payload type of the admin service
// UpdateChallenge method.
type UpdateChallengePayload struct {
	// A unique string that can be used in URLs
	Slug string
	// Title displayed to user
	Title string
	// A short text describing the challenge
	Description string
	// The number of points given to the solver
	Score int
	// unix timestamp
	PublishAt *int64
	// The ID of the CTF the challenge was taken from
	CtfEventID *string
	CategoryID string
	Token      string
	// ID of a challenge
	ChallengeID string
}

// PresignChallFileUploadPayload is the payload type of the admin service
// PresignChallFileUpload method.
type PresignChallFileUploadPayload struct {
	// MD5 hash of the file content in base64
	Md5      string
	Filename string
	// the files number of bytes
	Size  int64
	Token string
	// ID of a challenge
	ChallengeID string
}

// PresignChallFileUploadResult is the result type of the admin service
// PresignChallFileUpload method.
type PresignChallFileUploadResult struct {
	// Signed PutObject URL
	URL string
}

// ListMonthlyChallengesPayload is the payload type of the admin service
// ListMonthlyChallenges method.
type ListMonthlyChallengesPayload struct {
	Token string
}

// DeleteMonthlyChallengePayload is the payload type of the admin service
// DeleteMonthlyChallenge method.
type DeleteMonthlyChallengePayload struct {
	Token string
	// ID of a challenge
	ChallengeID string
}

// DeleteFilePayload is the payload type of the admin service DeleteFile method.
type DeleteFilePayload struct {
	Token string
	// ID of a file
	FileID string
}

// CreateMonthlyChallengePayload is the payload type of the admin service
// CreateMonthlyChallenge method.
type CreateMonthlyChallengePayload struct {
	Token string
	// A unique string that can be used in URLs
	Slug string
	// Title displayed to user
	Title string
	// A short text describing the challenge
	Description string
	ChallengeID string
	// The month(s) that the challenge is assigned for
	DisplayMonth string
	// Starting date of the monthly challenge
	StartDate string
	// Ending date of the monthly challenge
	EndDate string
}

// ListUsersPayload is the payload type of the admin service ListUsers method.
type ListUsersPayload struct {
	Token string
}

// AddFlagPayload is the payload type of the admin service AddFlag method.
type AddFlagPayload struct {
	Flag  string
	Token string
	// ID of a challenge
	ChallengeID string
}

// DeleteFlagPayload is the payload type of the admin service DeleteFlag method.
type DeleteFlagPayload struct {
	FlagID string
	Token  string
	// ID of a challenge
	ChallengeID string
}

// ListCategoriesPayload is the payload type of the admin service
// ListCategories method.
type ListCategoriesPayload struct {
	Token string
}

// A Wargame challenge
type SsmAdminChallenge struct {
	ID string
	// A unique string that can be used in URLs
	Slug string
	// Title displayed to user
	Title string
	// A short text describing the challenge
	Description string
	// The number of points given to the solver
	Score    int
	Services []*ChallengeService
	Files    []*AdminChallengeFiles
	// unix timestamp
	PublishAt *int64
	// The numer of people who solved the challenge
	Solves     int
	Flags      []*AdminChallengeFlag
	CategoryID string
}

type ChallengeService struct {
}

type AdminChallengeFiles struct {
	ID       string
	Filename string
	URL      string
	Bucket   string
	Key      string
	Size     int64
	// MD5 hash of the file content in base64
	Md5 string
}

type AdminChallengeFlag struct {
	ID   string
	Flag string
}

type MonthlyChallenge struct {
	// A unique string that can be used in URLs
	Slug string
	// Title displayed to user
	Title string
	// A short text describing the challenge
	Description string
	ChallengeID string
	// The month(s) that the challenge is assigned for
	DisplayMonth string
	// Starting date of the monthly challenge
	StartDate string
	// Ending date of the monthly challenge
	EndDate string
}

type SsmUser struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
	Role      string
	SchoolID  *int
}

type Category struct {
	ID   string
	Name string
}

// MakeUnauthorized builds a goa.ServiceError from an error.
func MakeUnauthorized(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "unauthorized",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not_found",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "bad_request",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewSsmAdminChallengeCollection initializes result type
// SsmAdminChallengeCollection from viewed result type
// SsmAdminChallengeCollection.
func NewSsmAdminChallengeCollection(vres adminviews.SsmAdminChallengeCollection) SsmAdminChallengeCollection {
	return newSsmAdminChallengeCollection(vres.Projected)
}

// NewViewedSsmAdminChallengeCollection initializes viewed result type
// SsmAdminChallengeCollection from result type SsmAdminChallengeCollection
// using the given view.
func NewViewedSsmAdminChallengeCollection(res SsmAdminChallengeCollection, view string) adminviews.SsmAdminChallengeCollection {
	p := newSsmAdminChallengeCollectionView(res)
	return adminviews.SsmAdminChallengeCollection{Projected: p, View: "default"}
}

// newSsmAdminChallengeCollection converts projected type
// SsmAdminChallengeCollection to service type SsmAdminChallengeCollection.
func newSsmAdminChallengeCollection(vres adminviews.SsmAdminChallengeCollectionView) SsmAdminChallengeCollection {
	res := make(SsmAdminChallengeCollection, len(vres))
	for i, n := range vres {
		res[i] = newSsmAdminChallenge(n)
	}
	return res
}

// newSsmAdminChallengeCollectionView projects result type
// SsmAdminChallengeCollection to projected type
// SsmAdminChallengeCollectionView using the "default" view.
func newSsmAdminChallengeCollectionView(res SsmAdminChallengeCollection) adminviews.SsmAdminChallengeCollectionView {
	vres := make(adminviews.SsmAdminChallengeCollectionView, len(res))
	for i, n := range res {
		vres[i] = newSsmAdminChallengeView(n)
	}
	return vres
}

// newSsmAdminChallenge converts projected type SsmAdminChallenge to service
// type SsmAdminChallenge.
func newSsmAdminChallenge(vres *adminviews.SsmAdminChallengeView) *SsmAdminChallenge {
	res := &SsmAdminChallenge{
		PublishAt: vres.PublishAt,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Slug != nil {
		res.Slug = *vres.Slug
	}
	if vres.Title != nil {
		res.Title = *vres.Title
	}
	if vres.Description != nil {
		res.Description = *vres.Description
	}
	if vres.Score != nil {
		res.Score = *vres.Score
	}
	if vres.Solves != nil {
		res.Solves = *vres.Solves
	}
	if vres.CategoryID != nil {
		res.CategoryID = *vres.CategoryID
	}
	if vres.Services != nil {
		res.Services = make([]*ChallengeService, len(vres.Services))
		for i, val := range vres.Services {
			res.Services[i] = transformAdminviewsChallengeServiceViewToChallengeService(val)
		}
	}
	if vres.Files != nil {
		res.Files = make([]*AdminChallengeFiles, len(vres.Files))
		for i, val := range vres.Files {
			res.Files[i] = transformAdminviewsAdminChallengeFilesViewToAdminChallengeFiles(val)
		}
	}
	if vres.Flags != nil {
		res.Flags = make([]*AdminChallengeFlag, len(vres.Flags))
		for i, val := range vres.Flags {
			res.Flags[i] = transformAdminviewsAdminChallengeFlagViewToAdminChallengeFlag(val)
		}
	}
	return res
}

// newSsmAdminChallengeView projects result type SsmAdminChallenge to projected
// type SsmAdminChallengeView using the "default" view.
func newSsmAdminChallengeView(res *SsmAdminChallenge) *adminviews.SsmAdminChallengeView {
	vres := &adminviews.SsmAdminChallengeView{
		ID:          &res.ID,
		Slug:        &res.Slug,
		Title:       &res.Title,
		Description: &res.Description,
		Score:       &res.Score,
		PublishAt:   res.PublishAt,
		Solves:      &res.Solves,
		CategoryID:  &res.CategoryID,
	}
	if res.Services != nil {
		vres.Services = make([]*adminviews.ChallengeServiceView, len(res.Services))
		for i, val := range res.Services {
			vres.Services[i] = transformChallengeServiceToAdminviewsChallengeServiceView(val)
		}
	}
	if res.Files != nil {
		vres.Files = make([]*adminviews.AdminChallengeFilesView, len(res.Files))
		for i, val := range res.Files {
			vres.Files[i] = transformAdminChallengeFilesToAdminviewsAdminChallengeFilesView(val)
		}
	}
	if res.Flags != nil {
		vres.Flags = make([]*adminviews.AdminChallengeFlagView, len(res.Flags))
		for i, val := range res.Flags {
			vres.Flags[i] = transformAdminChallengeFlagToAdminviewsAdminChallengeFlagView(val)
		}
	}
	return vres
}

// transformAdminviewsChallengeServiceViewToChallengeService builds a value of
// type *ChallengeService from a value of type *adminviews.ChallengeServiceView.
func transformAdminviewsChallengeServiceViewToChallengeService(v *adminviews.ChallengeServiceView) *ChallengeService {
	if v == nil {
		return nil
	}
	res := &ChallengeService{}

	return res
}

// transformAdminviewsAdminChallengeFilesViewToAdminChallengeFiles builds a
// value of type *AdminChallengeFiles from a value of type
// *adminviews.AdminChallengeFilesView.
func transformAdminviewsAdminChallengeFilesViewToAdminChallengeFiles(v *adminviews.AdminChallengeFilesView) *AdminChallengeFiles {
	if v == nil {
		return nil
	}
	res := &AdminChallengeFiles{
		ID:       *v.ID,
		Filename: *v.Filename,
		URL:      *v.URL,
		Bucket:   *v.Bucket,
		Key:      *v.Key,
		Size:     *v.Size,
		Md5:      *v.Md5,
	}

	return res
}

// transformAdminviewsAdminChallengeFlagViewToAdminChallengeFlag builds a value
// of type *AdminChallengeFlag from a value of type
// *adminviews.AdminChallengeFlagView.
func transformAdminviewsAdminChallengeFlagViewToAdminChallengeFlag(v *adminviews.AdminChallengeFlagView) *AdminChallengeFlag {
	if v == nil {
		return nil
	}
	res := &AdminChallengeFlag{
		ID:   *v.ID,
		Flag: *v.Flag,
	}

	return res
}

// transformChallengeServiceToAdminviewsChallengeServiceView builds a value of
// type *adminviews.ChallengeServiceView from a value of type *ChallengeService.
func transformChallengeServiceToAdminviewsChallengeServiceView(v *ChallengeService) *adminviews.ChallengeServiceView {
	if v == nil {
		return nil
	}
	res := &adminviews.ChallengeServiceView{}

	return res
}

// transformAdminChallengeFilesToAdminviewsAdminChallengeFilesView builds a
// value of type *adminviews.AdminChallengeFilesView from a value of type
// *AdminChallengeFiles.
func transformAdminChallengeFilesToAdminviewsAdminChallengeFilesView(v *AdminChallengeFiles) *adminviews.AdminChallengeFilesView {
	res := &adminviews.AdminChallengeFilesView{
		ID:       &v.ID,
		Filename: &v.Filename,
		URL:      &v.URL,
		Bucket:   &v.Bucket,
		Key:      &v.Key,
		Size:     &v.Size,
		Md5:      &v.Md5,
	}

	return res
}

// transformAdminChallengeFlagToAdminviewsAdminChallengeFlagView builds a value
// of type *adminviews.AdminChallengeFlagView from a value of type
// *AdminChallengeFlag.
func transformAdminChallengeFlagToAdminviewsAdminChallengeFlagView(v *AdminChallengeFlag) *adminviews.AdminChallengeFlagView {
	if v == nil {
		return nil
	}
	res := &adminviews.AdminChallengeFlagView{
		ID:   &v.ID,
		Flag: &v.Flag,
	}

	return res
}
