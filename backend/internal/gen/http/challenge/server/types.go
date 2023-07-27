// Code generated by goa v3.5.2, DO NOT EDIT.
//
// challenge HTTP server types
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package server

import (
	"unicode/utf8"

	challenge "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	challengeviews "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge/views"
	goa "goa.design/goa/v3/pkg"
)

// SubmitFlagRequestBody is the type of the "challenge" service "SubmitFlag"
// endpoint HTTP request body.
type SubmitFlagRequestBody struct {
	Flag *string `form:"flag,omitempty" json:"flag,omitempty" xml:"flag,omitempty"`
}

// SsmChallengeResponseCollection is the type of the "challenge" service
// "ListChallenges" endpoint HTTP response body.
type SsmChallengeResponseCollection []*SsmChallengeResponse

// ListMonthlyChallengesResponseBody is the type of the "challenge" service
// "ListMonthlyChallenges" endpoint HTTP response body.
type ListMonthlyChallengesResponseBody []*MonthlyChallengeResponse

// SubmitFlagAlreadySolvedResponseBody is the type of the "challenge" service
// "SubmitFlag" endpoint HTTP response body for the "already_solved" error.
type SubmitFlagAlreadySolvedResponseBody struct {
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

// SubmitFlagIncorrectFlagResponseBody is the type of the "challenge" service
// "SubmitFlag" endpoint HTTP response body for the "incorrect_flag" error.
type SubmitFlagIncorrectFlagResponseBody struct {
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
	Score    int                         `form:"score" json:"score" xml:"score"`
	Services []*ChallengeServiceResponse `form:"services,omitempty" json:"services,omitempty" xml:"services,omitempty"`
	Files    []*ChallengeFilesResponse   `form:"files,omitempty" json:"files,omitempty" xml:"files,omitempty"`
	// The numer of people who solved the challenge
	Solves int `form:"solves" json:"solves" xml:"solves"`
}

// ChallengeServiceResponse is used to define fields on response body types.
type ChallengeServiceResponse struct {
}

// ChallengeFilesResponse is used to define fields on response body types.
type ChallengeFilesResponse struct {
	Filename string `form:"filename" json:"filename" xml:"filename"`
	URL      string `form:"url" json:"url" xml:"url"`
}

// MonthlyChallengeResponse is used to define fields on response body types.
type MonthlyChallengeResponse struct {
	// A unique string that can be used in URLs
	Slug string `form:"slug" json:"slug" xml:"slug"`
	// Title displayed to user
	Title string `form:"title" json:"title" xml:"title"`
	// A short text describing the challenge
	Description string `form:"description" json:"description" xml:"description"`
	ChallengeID string `form:"challenge_id" json:"challenge_id" xml:"challenge_id"`
	// The month(s) that the challenge is assigned for
	DisplayMonth string `form:"display_month" json:"display_month" xml:"display_month"`
	// Starting date of the monthly challenge
	StartDate string `form:"start_date" json:"start_date" xml:"start_date"`
	// Ending date of the monthly challenge
	EndDate string `form:"end_date" json:"end_date" xml:"end_date"`
}

// NewSsmChallengeResponseCollection builds the HTTP response body from the
// result of the "ListChallenges" endpoint of the "challenge" service.
func NewSsmChallengeResponseCollection(res challengeviews.SsmChallengeCollectionView) SsmChallengeResponseCollection {
	body := make([]*SsmChallengeResponse, len(res))
	for i, val := range res {
		body[i] = marshalChallengeviewsSsmChallengeViewToSsmChallengeResponse(val)
	}
	return body
}

// NewListMonthlyChallengesResponseBody builds the HTTP response body from the
// result of the "ListMonthlyChallenges" endpoint of the "challenge" service.
func NewListMonthlyChallengesResponseBody(res []*challenge.MonthlyChallenge) ListMonthlyChallengesResponseBody {
	body := make([]*MonthlyChallengeResponse, len(res))
	for i, val := range res {
		body[i] = marshalChallengeMonthlyChallengeToMonthlyChallengeResponse(val)
	}
	return body
}

// NewSubmitFlagAlreadySolvedResponseBody builds the HTTP response body from
// the result of the "SubmitFlag" endpoint of the "challenge" service.
func NewSubmitFlagAlreadySolvedResponseBody(res *goa.ServiceError) *SubmitFlagAlreadySolvedResponseBody {
	body := &SubmitFlagAlreadySolvedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewSubmitFlagIncorrectFlagResponseBody builds the HTTP response body from
// the result of the "SubmitFlag" endpoint of the "challenge" service.
func NewSubmitFlagIncorrectFlagResponseBody(res *goa.ServiceError) *SubmitFlagIncorrectFlagResponseBody {
	body := &SubmitFlagIncorrectFlagResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListChallengesPayload builds a challenge service ListChallenges endpoint
// payload.
func NewListChallengesPayload(token *string) *challenge.ListChallengesPayload {
	v := &challenge.ListChallengesPayload{}
	v.Token = token

	return v
}

// NewListMonthlyChallengesPayload builds a challenge service
// ListMonthlyChallenges endpoint payload.
func NewListMonthlyChallengesPayload(token *string) *challenge.ListMonthlyChallengesPayload {
	v := &challenge.ListMonthlyChallengesPayload{}
	v.Token = token

	return v
}

// NewSubmitFlagPayload builds a challenge service SubmitFlag endpoint payload.
func NewSubmitFlagPayload(body *SubmitFlagRequestBody, challengeID string, token string) *challenge.SubmitFlagPayload {
	v := &challenge.SubmitFlagPayload{
		Flag: *body.Flag,
	}
	v.ChallengeID = challengeID
	v.Token = token

	return v
}

// ValidateSubmitFlagRequestBody runs the validations defined on
// SubmitFlagRequestBody
func ValidateSubmitFlagRequestBody(body *SubmitFlagRequestBody) (err error) {
	if body.Flag == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("flag", "body"))
	}
	if body.Flag != nil {
		if utf8.RuneCountInString(*body.Flag) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.flag", *body.Flag, utf8.RuneCountInString(*body.Flag), 200, false))
		}
	}
	return
}
