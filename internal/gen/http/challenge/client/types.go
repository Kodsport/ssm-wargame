// Code generated by goa v3.5.2, DO NOT EDIT.
//
// challenge HTTP client types
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package client

import (
	challenge "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	challengeviews "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge/views"
	goa "goa.design/goa/v3/pkg"
)

// ListChallengesRequestBody is the type of the "challenge" service
// "ListChallenges" endpoint HTTP request body.
type ListChallengesRequestBody struct {
	View string `form:"view" json:"view" xml:"view"`
}

// SubmitFlagRequestBody is the type of the "challenge" service "SubmitFlag"
// endpoint HTTP request body.
type SubmitFlagRequestBody struct {
	Flag string `form:"flag" json:"flag" xml:"flag"`
}

// ListChallengesResponseBody is the type of the "challenge" service
// "ListChallenges" endpoint HTTP response body.
type ListChallengesResponseBody []*SsmChallengeResponse

// SubmitFlagAlreadySolvedResponseBody is the type of the "challenge" service
// "SubmitFlag" endpoint HTTP response body for the "already_solved" error.
type SubmitFlagAlreadySolvedResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// SubmitFlagIncorrectFlagResponseBody is the type of the "challenge" service
// "SubmitFlag" endpoint HTTP response body for the "incorrect_flag" error.
type SubmitFlagIncorrectFlagResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// SsmChallengeResponse is used to define fields on response body types.
type SsmChallengeResponse struct {
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// A unique string that can be used in URLs
	Slug *string `form:"slug,omitempty" json:"slug,omitempty" xml:"slug,omitempty"`
	// Title displayed to user
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// A short text describing the challenge
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The number of points given to the solver
	Score     *uint                       `form:"score,omitempty" json:"score,omitempty" xml:"score,omitempty"`
	Published *bool                       `form:"published,omitempty" json:"published,omitempty" xml:"published,omitempty"`
	Services  []*ChallengeServiceResponse `form:"services,omitempty" json:"services,omitempty" xml:"services,omitempty"`
	Files     []*ChallengeFilesResponse   `form:"files,omitempty" json:"files,omitempty" xml:"files,omitempty"`
	// The numer of people who solved the challenge
	Solves *uint `form:"solves,omitempty" json:"solves,omitempty" xml:"solves,omitempty"`
}

// ChallengeServiceResponse is used to define fields on response body types.
type ChallengeServiceResponse struct {
}

// ChallengeFilesResponse is used to define fields on response body types.
type ChallengeFilesResponse struct {
}

// NewListChallengesRequestBody builds the HTTP request body from the payload
// of the "ListChallenges" endpoint of the "challenge" service.
func NewListChallengesRequestBody(p *challenge.ListChallengesPayload) *ListChallengesRequestBody {
	body := &ListChallengesRequestBody{
		View: p.View,
	}
	return body
}

// NewSubmitFlagRequestBody builds the HTTP request body from the payload of
// the "SubmitFlag" endpoint of the "challenge" service.
func NewSubmitFlagRequestBody(p *challenge.SubmitFlagPayload) *SubmitFlagRequestBody {
	body := &SubmitFlagRequestBody{
		Flag: p.Flag,
	}
	return body
}

// NewListChallengesSsmChallengeCollectionOK builds a "challenge" service
// "ListChallenges" endpoint result from a HTTP "OK" response.
func NewListChallengesSsmChallengeCollectionOK(body ListChallengesResponseBody) challengeviews.SsmChallengeCollectionView {
	v := make([]*challengeviews.SsmChallengeView, len(body))
	for i, val := range body {
		v[i] = unmarshalSsmChallengeResponseToChallengeviewsSsmChallengeView(val)
	}

	return v
}

// NewSubmitFlagAlreadySolved builds a challenge service SubmitFlag endpoint
// already_solved error.
func NewSubmitFlagAlreadySolved(body *SubmitFlagAlreadySolvedResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewSubmitFlagIncorrectFlag builds a challenge service SubmitFlag endpoint
// incorrect_flag error.
func NewSubmitFlagIncorrectFlag(body *SubmitFlagIncorrectFlagResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateSubmitFlagAlreadySolvedResponseBody runs the validations defined on
// SubmitFlag_already_solved_Response_Body
func ValidateSubmitFlagAlreadySolvedResponseBody(body *SubmitFlagAlreadySolvedResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateSubmitFlagIncorrectFlagResponseBody runs the validations defined on
// SubmitFlag_incorrect_flag_Response_Body
func ValidateSubmitFlagIncorrectFlagResponseBody(body *SubmitFlagIncorrectFlagResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateSsmChallengeResponse runs the validations defined on
// SsmChallengeResponse
func ValidateSsmChallengeResponse(body *SsmChallengeResponse) (err error) {
	if body.Solves == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("solves", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Score == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("score", "body"))
	}
	if body.Published == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("published", "body"))
	}
	return
}
