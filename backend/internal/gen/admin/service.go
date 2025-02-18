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
	// GetChallengeMeta implements GetChallengeMeta.
	GetChallengeMeta(context.Context, *GetChallengeMetaPayload) (res *ChallengeMeta, err error)
	// CreateChallenge implements CreateChallenge.
	CreateChallenge(context.Context, *CreateChallengePayload) (err error)
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
	// ListAuthors implements ListAuthors.
	ListAuthors(context.Context, *ListAuthorsPayload) (res []*Author, err error)
	// UpdateAuthor implements UpdateAuthor.
	UpdateAuthor(context.Context, *UpdateAuthorPayload) (err error)
	// CreateAuthor implements CreateAuthor.
	CreateAuthor(context.Context, *CreateAuthorPayload) (err error)
	// DeleteAuthor implements DeleteAuthor.
	DeleteAuthor(context.Context, *DeleteAuthorPayload) (err error)
	// AddFlag implements AddFlag.
	AddFlag(context.Context, *AddFlagPayload) (err error)
	// DeleteFlag implements DeleteFlag.
	DeleteFlag(context.Context, *DeleteFlagPayload) (err error)
	// ListCategories implements ListCategories.
	ListCategories(context.Context, *ListCategoriesPayload) (res []*Category, err error)
	// ChalltoolsImport implements ChalltoolsImport.
	ChalltoolsImport(context.Context, *ChalltoolsImportPayload) (err error)
	// ListCTFEvents implements ListCTFEvents.
	ListCTFEvents(context.Context, *ListCTFEventsPayload) (res []*CTFEvent, err error)
	// CreateCTFEvent implements CreateCTFEvent.
	CreateCTFEvent(context.Context, *CreateCTFEventPayload) (err error)
	// DeleteCTFEvent implements DeleteCTFEvent.
	DeleteCTFEvent(context.Context, *DeleteCTFEventPayload) (err error)
	// CreateCTFEventImportToken implements CreateCTFEventImportToken.
	CreateCTFEventImportToken(context.Context, *CreateCTFEventImportTokenPayload) (res *CreateCTFEventImportTokenResult, err error)
	// ListCourses implements ListCourses.
	ListCourses(context.Context, *ListCoursesPayload) (res SsmAdminCourseCollection, err error)
	// CreateCourse implements CreateCourse.
	CreateCourse(context.Context, *CreateCoursePayload) (err error)
	// UpdateCourse implements UpdateCourse.
	UpdateCourse(context.Context, *UpdateCoursePayload) (err error)
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
var MethodNames = [24]string{"ListChallenges", "GetChallengeMeta", "CreateChallenge", "PresignChallFileUpload", "ListMonthlyChallenges", "DeleteMonthlyChallenge", "DeleteFile", "CreateMonthlyChallenge", "ListUsers", "ListAuthors", "UpdateAuthor", "CreateAuthor", "DeleteAuthor", "AddFlag", "DeleteFlag", "ListCategories", "ChalltoolsImport", "ListCTFEvents", "CreateCTFEvent", "DeleteCTFEvent", "CreateCTFEventImportToken", "ListCourses", "CreateCourse", "UpdateCourse"}

// ListChallengesPayload is the payload type of the admin service
// ListChallenges method.
type ListChallengesPayload struct {
	Token string
}

// SsmAdminChallengeCollection is the result type of the admin service
// ListChallenges method.
type SsmAdminChallengeCollection []*SsmAdminChallenge

// GetChallengeMetaPayload is the payload type of the admin service
// GetChallengeMeta method.
type GetChallengeMetaPayload struct {
	Token string
	// ID of a challenge
	ChallengeID string
}

// ChallengeMeta is the result type of the admin service GetChallengeMeta
// method.
type ChallengeMeta struct {
	Solvers     []*ChallengeSolver
	Submissions []*ChallengeSubmission
}

// CreateChallengePayload is the payload type of the admin service
// CreateChallenge method.
type CreateChallengePayload struct {
	// A unique string that can be used in URLs
	Slug string
	// Title displayed to user
	Title string
	// A short text describing the challenge
	Description string
	// unix timestamp
	PublishAt *int64
	// The ID of the CTF the challenge was taken from
	CtfEventID  *string
	Hide        bool
	StaticScore *int
	CategoryID  string
	Authors     []string
	Token       string
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
	Token       string
	ChallengeID string
	// The month(s) that the challenge is assigned for
	DisplayMonth string
	// Starting date of the monthly challenge
	StartDate int64
	// Ending date of the monthly challenge
	EndDate int64
}

// ListUsersPayload is the payload type of the admin service ListUsers method.
type ListUsersPayload struct {
	Token string
}

// ListAuthorsPayload is the payload type of the admin service ListAuthors
// method.
type ListAuthorsPayload struct {
	Token string
}

// UpdateAuthorPayload is the payload type of the admin service UpdateAuthor
// method.
type UpdateAuthorPayload struct {
	Token       string
	FullName    string
	Description string
	Sponsor     bool
	Slug        string
	ImageURL    *string
	Publish     bool
	// ID of a file
	ID string
}

// CreateAuthorPayload is the payload type of the admin service CreateAuthor
// method.
type CreateAuthorPayload struct {
	Token       string
	FullName    string
	Description string
	Sponsor     bool
	Slug        string
	ImageURL    *string
	Publish     bool
}

// DeleteAuthorPayload is the payload type of the admin service DeleteAuthor
// method.
type DeleteAuthorPayload struct {
	ID    string
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

// ChalltoolsImportPayload is the payload type of the admin service
// ChalltoolsImport method.
type ChalltoolsImportPayload struct {
	ImportToken      string
	Title            string
	Description      string
	Authors          []string
	Categories       []string
	Score            *int
	ChallengeID      string
	FlagFormatPrefix *string
	FlagFormatSuffix *string
	FileUrls         []string
	Flags            []*ImportChallFlag
	Order            *int
	Services         []*ImportChallService
	HumanMetadata    *ImportChallHumanMetadata
	Custom           *ImportChallCustom
}

// ListCTFEventsPayload is the payload type of the admin service ListCTFEvents
// method.
type ListCTFEventsPayload struct {
	Token string
}

// CreateCTFEventPayload is the payload type of the admin service
// CreateCTFEvent method.
type CreateCTFEventPayload struct {
	Name  string
	Token string
}

// DeleteCTFEventPayload is the payload type of the admin service
// DeleteCTFEvent method.
type DeleteCTFEventPayload struct {
	// ID of a file
	ID    string
	Token string
}

// CreateCTFEventImportTokenPayload is the payload type of the admin service
// CreateCTFEventImportToken method.
type CreateCTFEventImportTokenPayload struct {
	Name      string
	ExpiresIn string
	Token     string
}

// CreateCTFEventImportTokenResult is the result type of the admin service
// CreateCTFEventImportToken method.
type CreateCTFEventImportTokenResult struct {
	Token string
}

// ListCoursesPayload is the payload type of the admin service ListCourses
// method.
type ListCoursesPayload struct {
	Token string
}

// SsmAdminCourseCollection is the result type of the admin service ListCourses
// method.
type SsmAdminCourseCollection []*SsmAdminCourse

// CreateCoursePayload is the payload type of the admin service CreateCourse
// method.
type CreateCoursePayload struct {
	Token       string
	Title       string
	Slug        string
	Category    string
	Difficulty  string
	Description string
	Publish     bool
	AuthorIds   []string
}

// UpdateCoursePayload is the payload type of the admin service UpdateCourse
// method.
type UpdateCoursePayload struct {
	Token       string
	Title       string
	Slug        string
	Category    string
	Difficulty  string
	Description string
	Publish     bool
	AuthorIds   []string
	// ID of a file
	ID string
}

// A Wargame challenge
type SsmAdminChallenge struct {
	// ID of a file
	ID string
	// A unique string that can be used in URLs
	Slug string
	// Title displayed to user
	Title string
	// A short text describing the challenge
	Description string
	Services    []*ChallengeService
	Files       []*AdminChallengeFiles
	// unix timestamp
	PublishAt *int64
	// The numer of people who solved the challenge
	Solves int
	Hide   bool
	// The ID of the CTF the challenge was taken from
	CtfEventID  *string
	Flags       []*AdminChallengeFlag
	StaticScore *int
	CategoryID  string
	Authors     []string
}

type ChallengeService struct {
	UserDisplay string
	Hyperlink   bool
}

type AdminChallengeFiles struct {
	Filename string
	URL      string
	// ID of a file
	ID string
}

type AdminChallengeFlag struct {
	Flag string
	// ID of a file
	ID string
}

type ChallengeSolver struct {
	UserID   string
	SolvedAt int64
}

type ChallengeSubmission struct {
	Input       string
	Successful  bool
	UserID      string
	SubmittedAt int64
	// ID of a file
	ID string
}

type MonthlyChallenge struct {
	ChallengeID string
	// The month(s) that the challenge is assigned for
	DisplayMonth string
	// Starting date of the monthly challenge
	StartDate int64
	// Ending date of the monthly challenge
	EndDate int64
}

type SsmUser struct {
	ID       string
	Email    string
	FullName string
	Role     string
	SchoolID *string
}

type Author struct {
	FullName    string
	Description string
	Sponsor     bool
	Slug        string
	ImageURL    *string
	Publish     bool
	// ID of a file
	ID string
}

type Category struct {
	Name string
	// ID of a file
	ID string
}

type ImportChallFlag struct {
	Type string
	Flag string
}

type ImportChallService struct {
	UserDisplay string
	Hyperlink   bool
}

type ImportChallHumanMetadata struct {
	EventName *string
}

type ImportChallCustom struct {
	Publish        *bool
	PublishAt      *string
	Slug           *string
	ChallNamespace *string
}

type CTFEvent struct {
	Name string
	// ID of a file
	ID string
}

type SsmAdminCourse struct {
	// ID of a file
	ID          string
	Title       string
	Slug        string
	Category    string
	Difficulty  string
	Description string
	Publish     bool
	AuthorIds   []string
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

// NewSsmAdminCourseCollection initializes result type SsmAdminCourseCollection
// from viewed result type SsmAdminCourseCollection.
func NewSsmAdminCourseCollection(vres adminviews.SsmAdminCourseCollection) SsmAdminCourseCollection {
	return newSsmAdminCourseCollection(vres.Projected)
}

// NewViewedSsmAdminCourseCollection initializes viewed result type
// SsmAdminCourseCollection from result type SsmAdminCourseCollection using the
// given view.
func NewViewedSsmAdminCourseCollection(res SsmAdminCourseCollection, view string) adminviews.SsmAdminCourseCollection {
	p := newSsmAdminCourseCollectionView(res)
	return adminviews.SsmAdminCourseCollection{Projected: p, View: "default"}
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
		PublishAt:   vres.PublishAt,
		CtfEventID:  vres.CtfEventID,
		StaticScore: vres.StaticScore,
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
	if vres.Solves != nil {
		res.Solves = *vres.Solves
	}
	if vres.Hide != nil {
		res.Hide = *vres.Hide
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
	if vres.Authors != nil {
		res.Authors = make([]string, len(vres.Authors))
		for i, val := range vres.Authors {
			res.Authors[i] = val
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
		PublishAt:   res.PublishAt,
		Solves:      &res.Solves,
		Hide:        &res.Hide,
		CtfEventID:  res.CtfEventID,
		StaticScore: res.StaticScore,
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
	if res.Authors != nil {
		vres.Authors = make([]string, len(res.Authors))
		for i, val := range res.Authors {
			vres.Authors[i] = val
		}
	}
	return vres
}

// newSsmAdminCourseCollection converts projected type SsmAdminCourseCollection
// to service type SsmAdminCourseCollection.
func newSsmAdminCourseCollection(vres adminviews.SsmAdminCourseCollectionView) SsmAdminCourseCollection {
	res := make(SsmAdminCourseCollection, len(vres))
	for i, n := range vres {
		res[i] = newSsmAdminCourse(n)
	}
	return res
}

// newSsmAdminCourseCollectionView projects result type
// SsmAdminCourseCollection to projected type SsmAdminCourseCollectionView
// using the "default" view.
func newSsmAdminCourseCollectionView(res SsmAdminCourseCollection) adminviews.SsmAdminCourseCollectionView {
	vres := make(adminviews.SsmAdminCourseCollectionView, len(res))
	for i, n := range res {
		vres[i] = newSsmAdminCourseView(n)
	}
	return vres
}

// newSsmAdminCourse converts projected type SsmAdminCourse to service type
// SsmAdminCourse.
func newSsmAdminCourse(vres *adminviews.SsmAdminCourseView) *SsmAdminCourse {
	res := &SsmAdminCourse{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Title != nil {
		res.Title = *vres.Title
	}
	if vres.Slug != nil {
		res.Slug = *vres.Slug
	}
	if vres.Category != nil {
		res.Category = *vres.Category
	}
	if vres.Difficulty != nil {
		res.Difficulty = *vres.Difficulty
	}
	if vres.Description != nil {
		res.Description = *vres.Description
	}
	if vres.Publish != nil {
		res.Publish = *vres.Publish
	}
	if vres.AuthorIds != nil {
		res.AuthorIds = make([]string, len(vres.AuthorIds))
		for i, val := range vres.AuthorIds {
			res.AuthorIds[i] = val
		}
	}
	return res
}

// newSsmAdminCourseView projects result type SsmAdminCourse to projected type
// SsmAdminCourseView using the "default" view.
func newSsmAdminCourseView(res *SsmAdminCourse) *adminviews.SsmAdminCourseView {
	vres := &adminviews.SsmAdminCourseView{
		ID:          &res.ID,
		Title:       &res.Title,
		Slug:        &res.Slug,
		Category:    &res.Category,
		Difficulty:  &res.Difficulty,
		Description: &res.Description,
		Publish:     &res.Publish,
	}
	if res.AuthorIds != nil {
		vres.AuthorIds = make([]string, len(res.AuthorIds))
		for i, val := range res.AuthorIds {
			vres.AuthorIds[i] = val
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
	res := &ChallengeService{
		UserDisplay: *v.UserDisplay,
		Hyperlink:   *v.Hyperlink,
	}

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
		Filename: *v.Filename,
		URL:      *v.URL,
		ID:       *v.ID,
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
		Flag: *v.Flag,
		ID:   *v.ID,
	}

	return res
}

// transformChallengeServiceToAdminviewsChallengeServiceView builds a value of
// type *adminviews.ChallengeServiceView from a value of type *ChallengeService.
func transformChallengeServiceToAdminviewsChallengeServiceView(v *ChallengeService) *adminviews.ChallengeServiceView {
	if v == nil {
		return nil
	}
	res := &adminviews.ChallengeServiceView{
		UserDisplay: &v.UserDisplay,
		Hyperlink:   &v.Hyperlink,
	}

	return res
}

// transformAdminChallengeFilesToAdminviewsAdminChallengeFilesView builds a
// value of type *adminviews.AdminChallengeFilesView from a value of type
// *AdminChallengeFiles.
func transformAdminChallengeFilesToAdminviewsAdminChallengeFilesView(v *AdminChallengeFiles) *adminviews.AdminChallengeFilesView {
	res := &adminviews.AdminChallengeFilesView{
		Filename: &v.Filename,
		URL:      &v.URL,
		ID:       &v.ID,
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
		Flag: &v.Flag,
		ID:   &v.ID,
	}

	return res
}
