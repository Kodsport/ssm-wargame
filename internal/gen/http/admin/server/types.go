// Code generated by goa v3.5.2, DO NOT EDIT.
//
// admin HTTP server types
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package server

import (
	admin "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
	adminviews "github.com/sakerhetsm/ssm-wargame/internal/gen/admin/views"
	goa "goa.design/goa/v3/pkg"
)

// CreateChallengeRequestBody is the type of the "admin" service
// "CreateChallenge" endpoint HTTP request body.
type CreateChallengeRequestBody struct {
	// A unique string that can be used in URLs
	Slug *string `form:"slug,omitempty" json:"slug,omitempty" xml:"slug,omitempty"`
	// Title displayed to user
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// A short text describing the challenge
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The number of points given to the solver
	Score *int32 `form:"score,omitempty" json:"score,omitempty" xml:"score,omitempty"`
}

// PresignChallFileUploadRequestBody is the type of the "admin" service
// "PresignChallFileUpload" endpoint HTTP request body.
type PresignChallFileUploadRequestBody struct {
	// MD5 hash of the file content
	Md5      *string `form:"md5,omitempty" json:"md5,omitempty" xml:"md5,omitempty"`
	Filename *string `form:"filename,omitempty" json:"filename,omitempty" xml:"filename,omitempty"`
}

// CreateMonthlyChallengeRequestBody is the type of the "admin" service
// "CreateMonthlyChallenge" endpoint HTTP request body.
type CreateMonthlyChallengeRequestBody struct {
	// The month(s) that the challenge is assigned for
	DisplayMonth *string `form:"display_month,omitempty" json:"display_month,omitempty" xml:"display_month,omitempty"`
	// Starting date of the monthly challenge
	StartDate *string `form:"start_date,omitempty" json:"start_date,omitempty" xml:"start_date,omitempty"`
	// Ending date of the monthly challenge
	EndDate *string `form:"end_date,omitempty" json:"end_date,omitempty" xml:"end_date,omitempty"`
	// ID of a challenge
	ChallengeID *string `form:"challengeID,omitempty" json:"challengeID,omitempty" xml:"challengeID,omitempty"`
}

// SsmChallengeResponseCollection is the type of the "admin" service
// "ListChallenges" endpoint HTTP response body.
type SsmChallengeResponseCollection []*SsmChallengeResponse

// PresignChallFileUploadResponseBody is the type of the "admin" service
// "PresignChallFileUpload" endpoint HTTP response body.
type PresignChallFileUploadResponseBody struct {
	// Signed PutObject URL
	URL string `form:"url" json:"url" xml:"url"`
}

// ListMonthlyChallengesResponseBody is the type of the "admin" service
// "ListMonthlyChallenges" endpoint HTTP response body.
type ListMonthlyChallengesResponseBody []*MonthlyChallengeMetaResponse

// ListUsersResponseBody is the type of the "admin" service "ListUsers"
// endpoint HTTP response body.
type ListUsersResponseBody []*SsmUserResponse

// ListChallengesUnauthorizedResponseBody is the type of the "admin" service
// "ListChallenges" endpoint HTTP response body for the "unauthorized" error.
type ListChallengesUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListChallengesNotFoundResponseBody is the type of the "admin" service
// "ListChallenges" endpoint HTTP response body for the "not_found" error.
type ListChallengesNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListChallengesBadRequestResponseBody is the type of the "admin" service
// "ListChallenges" endpoint HTTP response body for the "bad_request" error.
type ListChallengesBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateChallengeUnauthorizedResponseBody is the type of the "admin" service
// "CreateChallenge" endpoint HTTP response body for the "unauthorized" error.
type CreateChallengeUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateChallengeNotFoundResponseBody is the type of the "admin" service
// "CreateChallenge" endpoint HTTP response body for the "not_found" error.
type CreateChallengeNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateChallengeBadRequestResponseBody is the type of the "admin" service
// "CreateChallenge" endpoint HTTP response body for the "bad_request" error.
type CreateChallengeBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PresignChallFileUploadUnauthorizedResponseBody is the type of the "admin"
// service "PresignChallFileUpload" endpoint HTTP response body for the
// "unauthorized" error.
type PresignChallFileUploadUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PresignChallFileUploadNotFoundResponseBody is the type of the "admin"
// service "PresignChallFileUpload" endpoint HTTP response body for the
// "not_found" error.
type PresignChallFileUploadNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PresignChallFileUploadBadRequestResponseBody is the type of the "admin"
// service "PresignChallFileUpload" endpoint HTTP response body for the
// "bad_request" error.
type PresignChallFileUploadBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListMonthlyChallengesUnauthorizedResponseBody is the type of the "admin"
// service "ListMonthlyChallenges" endpoint HTTP response body for the
// "unauthorized" error.
type ListMonthlyChallengesUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListMonthlyChallengesNotFoundResponseBody is the type of the "admin" service
// "ListMonthlyChallenges" endpoint HTTP response body for the "not_found"
// error.
type ListMonthlyChallengesNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListMonthlyChallengesBadRequestResponseBody is the type of the "admin"
// service "ListMonthlyChallenges" endpoint HTTP response body for the
// "bad_request" error.
type ListMonthlyChallengesBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteMonthlyChallengeUnauthorizedResponseBody is the type of the "admin"
// service "DeleteMonthlyChallenge" endpoint HTTP response body for the
// "unauthorized" error.
type DeleteMonthlyChallengeUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteMonthlyChallengeNotFoundResponseBody is the type of the "admin"
// service "DeleteMonthlyChallenge" endpoint HTTP response body for the
// "not_found" error.
type DeleteMonthlyChallengeNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteMonthlyChallengeBadRequestResponseBody is the type of the "admin"
// service "DeleteMonthlyChallenge" endpoint HTTP response body for the
// "bad_request" error.
type DeleteMonthlyChallengeBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateMonthlyChallengeUnauthorizedResponseBody is the type of the "admin"
// service "CreateMonthlyChallenge" endpoint HTTP response body for the
// "unauthorized" error.
type CreateMonthlyChallengeUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateMonthlyChallengeNotFoundResponseBody is the type of the "admin"
// service "CreateMonthlyChallenge" endpoint HTTP response body for the
// "not_found" error.
type CreateMonthlyChallengeNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateMonthlyChallengeBadRequestResponseBody is the type of the "admin"
// service "CreateMonthlyChallenge" endpoint HTTP response body for the
// "bad_request" error.
type CreateMonthlyChallengeBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListUsersUnauthorizedResponseBody is the type of the "admin" service
// "ListUsers" endpoint HTTP response body for the "unauthorized" error.
type ListUsersUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListUsersNotFoundResponseBody is the type of the "admin" service "ListUsers"
// endpoint HTTP response body for the "not_found" error.
type ListUsersNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListUsersBadRequestResponseBody is the type of the "admin" service
// "ListUsers" endpoint HTTP response body for the "bad_request" error.
type ListUsersBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// SsmChallengeResponse is used to define fields on response body types.
type SsmChallengeResponse struct {
	ID string `form:"id" json:"id" xml:"id"`
	// A unique string that can be used in URLs
	Slug string `form:"slug" json:"slug" xml:"slug"`
	// Title displayed to user
	Title string `form:"title" json:"title" xml:"title"`
	// A short text describing the challenge
	Description string `form:"description" json:"description" xml:"description"`
	// The number of points given to the solver
	Score     int32                       `form:"score" json:"score" xml:"score"`
	Services  []*ChallengeServiceResponse `form:"services,omitempty" json:"services,omitempty" xml:"services,omitempty"`
	Files     []*ChallengeFilesResponse   `form:"files,omitempty" json:"files,omitempty" xml:"files,omitempty"`
	Published bool                        `form:"published" json:"published" xml:"published"`
	// The numer of people who solved the challenge
	Solves int64 `form:"solves" json:"solves" xml:"solves"`
}

// ChallengeServiceResponse is used to define fields on response body types.
type ChallengeServiceResponse struct {
}

// ChallengeFilesResponse is used to define fields on response body types.
type ChallengeFilesResponse struct {
}

// MonthlyChallengeMetaResponse is used to define fields on response body types.
type MonthlyChallengeMetaResponse struct {
	// The month(s) that the challenge is assigned for
	DisplayMonth string `form:"display_month" json:"display_month" xml:"display_month"`
	// Starting date of the monthly challenge
	StartDate string `form:"start_date" json:"start_date" xml:"start_date"`
	// Ending date of the monthly challenge
	EndDate string `form:"end_date" json:"end_date" xml:"end_date"`
}

// SsmUserResponse is used to define fields on response body types.
type SsmUserResponse struct {
	ID        string `form:"id" json:"id" xml:"id"`
	Email     string `form:"email" json:"email" xml:"email"`
	FirstName string `form:"first_name" json:"first_name" xml:"first_name"`
	LastName  string `form:"last_name" json:"last_name" xml:"last_name"`
}

// NewSsmChallengeResponseCollection builds the HTTP response body from the
// result of the "ListChallenges" endpoint of the "admin" service.
func NewSsmChallengeResponseCollection(res adminviews.SsmChallengeCollectionView) SsmChallengeResponseCollection {
	body := make([]*SsmChallengeResponse, len(res))
	for i, val := range res {
		body[i] = marshalAdminviewsSsmChallengeViewToSsmChallengeResponse(val)
	}
	return body
}

// NewPresignChallFileUploadResponseBody builds the HTTP response body from the
// result of the "PresignChallFileUpload" endpoint of the "admin" service.
func NewPresignChallFileUploadResponseBody(res *admin.PresignChallFileUploadResult) *PresignChallFileUploadResponseBody {
	body := &PresignChallFileUploadResponseBody{
		URL: res.URL,
	}
	return body
}

// NewListMonthlyChallengesResponseBody builds the HTTP response body from the
// result of the "ListMonthlyChallenges" endpoint of the "admin" service.
func NewListMonthlyChallengesResponseBody(res []*admin.MonthlyChallengeMeta) ListMonthlyChallengesResponseBody {
	body := make([]*MonthlyChallengeMetaResponse, len(res))
	for i, val := range res {
		body[i] = marshalAdminMonthlyChallengeMetaToMonthlyChallengeMetaResponse(val)
	}
	return body
}

// NewListUsersResponseBody builds the HTTP response body from the result of
// the "ListUsers" endpoint of the "admin" service.
func NewListUsersResponseBody(res []*admin.SsmUser) ListUsersResponseBody {
	body := make([]*SsmUserResponse, len(res))
	for i, val := range res {
		body[i] = marshalAdminSsmUserToSsmUserResponse(val)
	}
	return body
}

// NewListChallengesUnauthorizedResponseBody builds the HTTP response body from
// the result of the "ListChallenges" endpoint of the "admin" service.
func NewListChallengesUnauthorizedResponseBody(res *goa.ServiceError) *ListChallengesUnauthorizedResponseBody {
	body := &ListChallengesUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListChallengesNotFoundResponseBody builds the HTTP response body from the
// result of the "ListChallenges" endpoint of the "admin" service.
func NewListChallengesNotFoundResponseBody(res *goa.ServiceError) *ListChallengesNotFoundResponseBody {
	body := &ListChallengesNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListChallengesBadRequestResponseBody builds the HTTP response body from
// the result of the "ListChallenges" endpoint of the "admin" service.
func NewListChallengesBadRequestResponseBody(res *goa.ServiceError) *ListChallengesBadRequestResponseBody {
	body := &ListChallengesBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateChallengeUnauthorizedResponseBody builds the HTTP response body
// from the result of the "CreateChallenge" endpoint of the "admin" service.
func NewCreateChallengeUnauthorizedResponseBody(res *goa.ServiceError) *CreateChallengeUnauthorizedResponseBody {
	body := &CreateChallengeUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateChallengeNotFoundResponseBody builds the HTTP response body from
// the result of the "CreateChallenge" endpoint of the "admin" service.
func NewCreateChallengeNotFoundResponseBody(res *goa.ServiceError) *CreateChallengeNotFoundResponseBody {
	body := &CreateChallengeNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateChallengeBadRequestResponseBody builds the HTTP response body from
// the result of the "CreateChallenge" endpoint of the "admin" service.
func NewCreateChallengeBadRequestResponseBody(res *goa.ServiceError) *CreateChallengeBadRequestResponseBody {
	body := &CreateChallengeBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPresignChallFileUploadUnauthorizedResponseBody builds the HTTP response
// body from the result of the "PresignChallFileUpload" endpoint of the "admin"
// service.
func NewPresignChallFileUploadUnauthorizedResponseBody(res *goa.ServiceError) *PresignChallFileUploadUnauthorizedResponseBody {
	body := &PresignChallFileUploadUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPresignChallFileUploadNotFoundResponseBody builds the HTTP response body
// from the result of the "PresignChallFileUpload" endpoint of the "admin"
// service.
func NewPresignChallFileUploadNotFoundResponseBody(res *goa.ServiceError) *PresignChallFileUploadNotFoundResponseBody {
	body := &PresignChallFileUploadNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPresignChallFileUploadBadRequestResponseBody builds the HTTP response
// body from the result of the "PresignChallFileUpload" endpoint of the "admin"
// service.
func NewPresignChallFileUploadBadRequestResponseBody(res *goa.ServiceError) *PresignChallFileUploadBadRequestResponseBody {
	body := &PresignChallFileUploadBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListMonthlyChallengesUnauthorizedResponseBody builds the HTTP response
// body from the result of the "ListMonthlyChallenges" endpoint of the "admin"
// service.
func NewListMonthlyChallengesUnauthorizedResponseBody(res *goa.ServiceError) *ListMonthlyChallengesUnauthorizedResponseBody {
	body := &ListMonthlyChallengesUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListMonthlyChallengesNotFoundResponseBody builds the HTTP response body
// from the result of the "ListMonthlyChallenges" endpoint of the "admin"
// service.
func NewListMonthlyChallengesNotFoundResponseBody(res *goa.ServiceError) *ListMonthlyChallengesNotFoundResponseBody {
	body := &ListMonthlyChallengesNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListMonthlyChallengesBadRequestResponseBody builds the HTTP response body
// from the result of the "ListMonthlyChallenges" endpoint of the "admin"
// service.
func NewListMonthlyChallengesBadRequestResponseBody(res *goa.ServiceError) *ListMonthlyChallengesBadRequestResponseBody {
	body := &ListMonthlyChallengesBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteMonthlyChallengeUnauthorizedResponseBody builds the HTTP response
// body from the result of the "DeleteMonthlyChallenge" endpoint of the "admin"
// service.
func NewDeleteMonthlyChallengeUnauthorizedResponseBody(res *goa.ServiceError) *DeleteMonthlyChallengeUnauthorizedResponseBody {
	body := &DeleteMonthlyChallengeUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteMonthlyChallengeNotFoundResponseBody builds the HTTP response body
// from the result of the "DeleteMonthlyChallenge" endpoint of the "admin"
// service.
func NewDeleteMonthlyChallengeNotFoundResponseBody(res *goa.ServiceError) *DeleteMonthlyChallengeNotFoundResponseBody {
	body := &DeleteMonthlyChallengeNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteMonthlyChallengeBadRequestResponseBody builds the HTTP response
// body from the result of the "DeleteMonthlyChallenge" endpoint of the "admin"
// service.
func NewDeleteMonthlyChallengeBadRequestResponseBody(res *goa.ServiceError) *DeleteMonthlyChallengeBadRequestResponseBody {
	body := &DeleteMonthlyChallengeBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateMonthlyChallengeUnauthorizedResponseBody builds the HTTP response
// body from the result of the "CreateMonthlyChallenge" endpoint of the "admin"
// service.
func NewCreateMonthlyChallengeUnauthorizedResponseBody(res *goa.ServiceError) *CreateMonthlyChallengeUnauthorizedResponseBody {
	body := &CreateMonthlyChallengeUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateMonthlyChallengeNotFoundResponseBody builds the HTTP response body
// from the result of the "CreateMonthlyChallenge" endpoint of the "admin"
// service.
func NewCreateMonthlyChallengeNotFoundResponseBody(res *goa.ServiceError) *CreateMonthlyChallengeNotFoundResponseBody {
	body := &CreateMonthlyChallengeNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateMonthlyChallengeBadRequestResponseBody builds the HTTP response
// body from the result of the "CreateMonthlyChallenge" endpoint of the "admin"
// service.
func NewCreateMonthlyChallengeBadRequestResponseBody(res *goa.ServiceError) *CreateMonthlyChallengeBadRequestResponseBody {
	body := &CreateMonthlyChallengeBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListUsersUnauthorizedResponseBody builds the HTTP response body from the
// result of the "ListUsers" endpoint of the "admin" service.
func NewListUsersUnauthorizedResponseBody(res *goa.ServiceError) *ListUsersUnauthorizedResponseBody {
	body := &ListUsersUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListUsersNotFoundResponseBody builds the HTTP response body from the
// result of the "ListUsers" endpoint of the "admin" service.
func NewListUsersNotFoundResponseBody(res *goa.ServiceError) *ListUsersNotFoundResponseBody {
	body := &ListUsersNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListUsersBadRequestResponseBody builds the HTTP response body from the
// result of the "ListUsers" endpoint of the "admin" service.
func NewListUsersBadRequestResponseBody(res *goa.ServiceError) *ListUsersBadRequestResponseBody {
	body := &ListUsersBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListChallengesPayload builds a admin service ListChallenges endpoint
// payload.
func NewListChallengesPayload(token string) *admin.ListChallengesPayload {
	v := &admin.ListChallengesPayload{}
	v.Token = token

	return v
}

// NewCreateChallengePayload builds a admin service CreateChallenge endpoint
// payload.
func NewCreateChallengePayload(body *CreateChallengeRequestBody, token string) *admin.CreateChallengePayload {
	v := &admin.CreateChallengePayload{
		Slug:        *body.Slug,
		Title:       *body.Title,
		Description: *body.Description,
		Score:       *body.Score,
	}
	v.Token = token

	return v
}

// NewPresignChallFileUploadPayload builds a admin service
// PresignChallFileUpload endpoint payload.
func NewPresignChallFileUploadPayload(body *PresignChallFileUploadRequestBody, challengeID string, token string) *admin.PresignChallFileUploadPayload {
	v := &admin.PresignChallFileUploadPayload{
		Md5:      *body.Md5,
		Filename: *body.Filename,
	}
	v.ChallengeID = challengeID
	v.Token = token

	return v
}

// NewListMonthlyChallengesPayload builds a admin service ListMonthlyChallenges
// endpoint payload.
func NewListMonthlyChallengesPayload(token string) *admin.ListMonthlyChallengesPayload {
	v := &admin.ListMonthlyChallengesPayload{}
	v.Token = token

	return v
}

// NewDeleteMonthlyChallengePayload builds a admin service
// DeleteMonthlyChallenge endpoint payload.
func NewDeleteMonthlyChallengePayload(challengeID string, token string) *admin.DeleteMonthlyChallengePayload {
	v := &admin.DeleteMonthlyChallengePayload{}
	v.ChallengeID = challengeID
	v.Token = token

	return v
}

// NewCreateMonthlyChallengePayload builds a admin service
// CreateMonthlyChallenge endpoint payload.
func NewCreateMonthlyChallengePayload(body *CreateMonthlyChallengeRequestBody, token string) *admin.CreateMonthlyChallengePayload {
	v := &admin.CreateMonthlyChallengePayload{
		DisplayMonth: *body.DisplayMonth,
		StartDate:    *body.StartDate,
		EndDate:      *body.EndDate,
		ChallengeID:  *body.ChallengeID,
	}
	v.Token = token

	return v
}

// NewListUsersPayload builds a admin service ListUsers endpoint payload.
func NewListUsersPayload(token string) *admin.ListUsersPayload {
	v := &admin.ListUsersPayload{}
	v.Token = token

	return v
}

// ValidateCreateChallengeRequestBody runs the validations defined on
// CreateChallengeRequestBody
func ValidateCreateChallengeRequestBody(body *CreateChallengeRequestBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Score == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("score", "body"))
	}
	if body.Slug == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("slug", "body"))
	}
	return
}

// ValidatePresignChallFileUploadRequestBody runs the validations defined on
// PresignChallFileUploadRequestBody
func ValidatePresignChallFileUploadRequestBody(body *PresignChallFileUploadRequestBody) (err error) {
	if body.Md5 == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("md5", "body"))
	}
	if body.Filename == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("filename", "body"))
	}
	return
}

// ValidateCreateMonthlyChallengeRequestBody runs the validations defined on
// CreateMonthlyChallengeRequestBody
func ValidateCreateMonthlyChallengeRequestBody(body *CreateMonthlyChallengeRequestBody) (err error) {
	if body.StartDate == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("start_date", "body"))
	}
	if body.EndDate == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("end_date", "body"))
	}
	if body.DisplayMonth == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("display_month", "body"))
	}
	if body.ChallengeID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("challengeID", "body"))
	}
	if body.ChallengeID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.challengeID", *body.ChallengeID, goa.FormatUUID))
	}
	return
}
