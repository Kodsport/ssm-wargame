// Code generated by goa v3.5.2, DO NOT EDIT.
//
// challenge HTTP server types
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design

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

// KnackKodenSubmitFlagRequestBody is the type of the "challenge" service
// "KnackKodenSubmitFlag" endpoint HTTP request body.
type KnackKodenSubmitFlagRequestBody struct {
	Flag     *string `form:"flag,omitempty" json:"flag,omitempty" xml:"flag,omitempty"`
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// KnackKodenRegisterClassRequestBody is the type of the "challenge" service
// "KnackKodenRegisterClass" endpoint HTTP request body.
type KnackKodenRegisterClassRequestBody struct {
	TeacherFullName *string `form:"teacher_full_name,omitempty" json:"teacher_full_name,omitempty" xml:"teacher_full_name,omitempty"`
	TeacherEmail    *string `form:"teacher_email,omitempty" json:"teacher_email,omitempty" xml:"teacher_email,omitempty"`
	TeacherPhonenr  *string `form:"teacher_phonenr,omitempty" json:"teacher_phonenr,omitempty" xml:"teacher_phonenr,omitempty"`
	SchoolName      *string `form:"school_name,omitempty" json:"school_name,omitempty" xml:"school_name,omitempty"`
	ClassName       *string `form:"class_name,omitempty" json:"class_name,omitempty" xml:"class_name,omitempty"`
	PostalCode      *string `form:"postal_code,omitempty" json:"postal_code,omitempty" xml:"postal_code,omitempty"`
}

// KnackKodenGetClassRequestBody is the type of the "challenge" service
// "KnackKodenGetClass" endpoint HTTP request body.
type KnackKodenGetClassRequestBody struct {
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// SsmChallengeResponseCollection is the type of the "challenge" service
// "ListChallenges" endpoint HTTP response body.
type SsmChallengeResponseCollection []*SsmChallengeResponse

// ListEventsResponseBody is the type of the "challenge" service "ListEvents"
// endpoint HTTP response body.
type ListEventsResponseBody []*CTFEventResponse

// GetCurrentMonthlyChallengeResponseBody is the type of the "challenge"
// service "GetCurrentMonthlyChallenge" endpoint HTTP response body.
type GetCurrentMonthlyChallengeResponseBody struct {
	ChallengeID string `form:"challenge_id" json:"challenge_id" xml:"challenge_id"`
	// The month(s) that the challenge is assigned for
	DisplayMonth string `form:"display_month" json:"display_month" xml:"display_month"`
	// Starting date of the monthly challenge
	StartDate int64 `form:"start_date" json:"start_date" xml:"start_date"`
	// Ending date of the monthly challenge
	EndDate   int64                     `form:"end_date" json:"end_date" xml:"end_date"`
	Challenge *SsmChallengeResponseBody `form:"challenge" json:"challenge" xml:"challenge"`
}

// SsmUserMonthlyChallengeResponseCollection is the type of the "challenge"
// service "ListMonthlyChallenges" endpoint HTTP response body.
type SsmUserMonthlyChallengeResponseCollection []*SsmUserMonthlyChallengeResponse

// SchoolScoreboardResponseBody is the type of the "challenge" service
// "SchoolScoreboard" endpoint HTTP response body.
type SchoolScoreboardResponseBody struct {
	Scores []*SchoolScoreboardScoreResponseBody `form:"scores" json:"scores" xml:"scores"`
}

// UserScoreboardResponseBody is the type of the "challenge" service
// "UserScoreboard" endpoint HTTP response body.
type UserScoreboardResponseBody struct {
	Scores []*UserScoreboardScoreResponseBody `form:"scores" json:"scores" xml:"scores"`
}

// ListAuthorsResponseBody is the type of the "challenge" service "ListAuthors"
// endpoint HTTP response body.
type ListAuthorsResponseBody []*AuthorResponse

// ListCoursesResponseBody is the type of the "challenge" service "ListCourses"
// endpoint HTTP response body.
type ListCoursesResponseBody []*CourseResponse

// KnackKodenScoreboardResponseBody is the type of the "challenge" service
// "KnackKodenScoreboard" endpoint HTTP response body.
type KnackKodenScoreboardResponseBody struct {
	Scores []*SchoolScoreboardScoreResponseBody `form:"scores" json:"scores" xml:"scores"`
}

// KnackKodenRegisterClassResponseBody is the type of the "challenge" service
// "KnackKodenRegisterClass" endpoint HTTP response body.
type KnackKodenRegisterClassResponseBody struct {
	Password string `form:"password" json:"password" xml:"password"`
}

// KnackKodenGetClassResponseBody is the type of the "challenge" service
// "KnackKodenGetClass" endpoint HTTP response body.
type KnackKodenGetClassResponseBody struct {
	ClassName string   `form:"class_name" json:"class_name" xml:"class_name"`
	Solves    []string `form:"solves" json:"solves" xml:"solves"`
}

// GetCurrentMonthlyChallengeNotFoundResponseBody is the type of the
// "challenge" service "GetCurrentMonthlyChallenge" endpoint HTTP response body
// for the "not_found" error.
type GetCurrentMonthlyChallengeNotFoundResponseBody struct {
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

// KnackKodenSubmitFlagAlreadySolvedResponseBody is the type of the "challenge"
// service "KnackKodenSubmitFlag" endpoint HTTP response body for the
// "already_solved" error.
type KnackKodenSubmitFlagAlreadySolvedResponseBody struct {
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

// KnackKodenSubmitFlagIncorrectFlagResponseBody is the type of the "challenge"
// service "KnackKodenSubmitFlag" endpoint HTTP response body for the
// "incorrect_flag" error.
type KnackKodenSubmitFlagIncorrectFlagResponseBody struct {
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
	// ID of a file
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
	// The ID of the CTF the challenge was taken from
	CtfEventID     *string `form:"ctf_event_id,omitempty" json:"ctf_event_id,omitempty" xml:"ctf_event_id,omitempty"`
	ChallNamespace *string `form:"chall_namespace,omitempty" json:"chall_namespace,omitempty" xml:"chall_namespace,omitempty"`
	// whether the user has solved the challenge or not
	Solved   bool                 `form:"solved" json:"solved" xml:"solved"`
	Category string               `form:"category" json:"category" xml:"category"`
	Authors  []*AuthorResponse    `form:"authors,omitempty" json:"authors,omitempty" xml:"authors,omitempty"`
	Solvers  []*SsmSolverResponse `form:"solvers,omitempty" json:"solvers,omitempty" xml:"solvers,omitempty"`
}

// ChallengeServiceResponse is used to define fields on response body types.
type ChallengeServiceResponse struct {
	UserDisplay string `form:"user_display" json:"user_display" xml:"user_display"`
	Hyperlink   bool   `form:"hyperlink" json:"hyperlink" xml:"hyperlink"`
}

// ChallengeFilesResponse is used to define fields on response body types.
type ChallengeFilesResponse struct {
	Filename string `form:"filename" json:"filename" xml:"filename"`
	URL      string `form:"url" json:"url" xml:"url"`
}

// AuthorResponse is used to define fields on response body types.
type AuthorResponse struct {
	FullName    string  `form:"full_name" json:"full_name" xml:"full_name"`
	Description string  `form:"description" json:"description" xml:"description"`
	Sponsor     bool    `form:"sponsor" json:"sponsor" xml:"sponsor"`
	Slug        string  `form:"slug" json:"slug" xml:"slug"`
	ImageURL    *string `form:"image_url,omitempty" json:"image_url,omitempty" xml:"image_url,omitempty"`
	Publish     bool    `form:"publish" json:"publish" xml:"publish"`
	// ID of a file
	ID string `form:"id" json:"id" xml:"id"`
}

// SsmSolverResponse is used to define fields on response body types.
type SsmSolverResponse struct {
	ID       string `form:"id" json:"id" xml:"id"`
	FullName string `form:"full_name" json:"full_name" xml:"full_name"`
	SolvedAt int64  `form:"solved_at" json:"solved_at" xml:"solved_at"`
}

// CTFEventResponse is used to define fields on response body types.
type CTFEventResponse struct {
	Name string `form:"name" json:"name" xml:"name"`
	// ID of a file
	ID string `form:"id" json:"id" xml:"id"`
}

// SsmChallengeResponseBody is used to define fields on response body types.
type SsmChallengeResponseBody struct {
	// ID of a file
	ID string `form:"id" json:"id" xml:"id"`
	// A unique string that can be used in URLs
	Slug string `form:"slug" json:"slug" xml:"slug"`
	// Title displayed to user
	Title string `form:"title" json:"title" xml:"title"`
	// A short text describing the challenge
	Description string `form:"description" json:"description" xml:"description"`
	// The number of points given to the solver
	Score    int                             `form:"score" json:"score" xml:"score"`
	Services []*ChallengeServiceResponseBody `form:"services,omitempty" json:"services,omitempty" xml:"services,omitempty"`
	Files    []*ChallengeFilesResponseBody   `form:"files,omitempty" json:"files,omitempty" xml:"files,omitempty"`
	// The numer of people who solved the challenge
	Solves int `form:"solves" json:"solves" xml:"solves"`
	// The ID of the CTF the challenge was taken from
	CtfEventID     *string `form:"ctf_event_id,omitempty" json:"ctf_event_id,omitempty" xml:"ctf_event_id,omitempty"`
	ChallNamespace *string `form:"chall_namespace,omitempty" json:"chall_namespace,omitempty" xml:"chall_namespace,omitempty"`
	// whether the user has solved the challenge or not
	Solved   bool                     `form:"solved" json:"solved" xml:"solved"`
	Category string                   `form:"category" json:"category" xml:"category"`
	Authors  []*AuthorResponseBody    `form:"authors,omitempty" json:"authors,omitempty" xml:"authors,omitempty"`
	Solvers  []*SsmSolverResponseBody `form:"solvers,omitempty" json:"solvers,omitempty" xml:"solvers,omitempty"`
}

// ChallengeServiceResponseBody is used to define fields on response body types.
type ChallengeServiceResponseBody struct {
	UserDisplay string `form:"user_display" json:"user_display" xml:"user_display"`
	Hyperlink   bool   `form:"hyperlink" json:"hyperlink" xml:"hyperlink"`
}

// ChallengeFilesResponseBody is used to define fields on response body types.
type ChallengeFilesResponseBody struct {
	Filename string `form:"filename" json:"filename" xml:"filename"`
	URL      string `form:"url" json:"url" xml:"url"`
}

// AuthorResponseBody is used to define fields on response body types.
type AuthorResponseBody struct {
	FullName    string  `form:"full_name" json:"full_name" xml:"full_name"`
	Description string  `form:"description" json:"description" xml:"description"`
	Sponsor     bool    `form:"sponsor" json:"sponsor" xml:"sponsor"`
	Slug        string  `form:"slug" json:"slug" xml:"slug"`
	ImageURL    *string `form:"image_url,omitempty" json:"image_url,omitempty" xml:"image_url,omitempty"`
	Publish     bool    `form:"publish" json:"publish" xml:"publish"`
	// ID of a file
	ID string `form:"id" json:"id" xml:"id"`
}

// SsmSolverResponseBody is used to define fields on response body types.
type SsmSolverResponseBody struct {
	ID       string `form:"id" json:"id" xml:"id"`
	FullName string `form:"full_name" json:"full_name" xml:"full_name"`
	SolvedAt int64  `form:"solved_at" json:"solved_at" xml:"solved_at"`
}

// SsmUserMonthlyChallengeResponse is used to define fields on response body
// types.
type SsmUserMonthlyChallengeResponse struct {
	ChallengeID string `form:"challenge_id" json:"challenge_id" xml:"challenge_id"`
	// The month(s) that the challenge is assigned for
	DisplayMonth string `form:"display_month" json:"display_month" xml:"display_month"`
	// Starting date of the monthly challenge
	StartDate int64 `form:"start_date" json:"start_date" xml:"start_date"`
	// Ending date of the monthly challenge
	EndDate   int64                 `form:"end_date" json:"end_date" xml:"end_date"`
	Challenge *SsmChallengeResponse `form:"challenge" json:"challenge" xml:"challenge"`
}

// SchoolScoreboardScoreResponseBody is used to define fields on response body
// types.
type SchoolScoreboardScoreResponseBody struct {
	Score        int    `form:"score" json:"score" xml:"score"`
	SchoolName   string `form:"school_name" json:"school_name" xml:"school_name"`
	IsUniversity bool   `form:"is_university" json:"is_university" xml:"is_university"`
}

// UserScoreboardScoreResponseBody is used to define fields on response body
// types.
type UserScoreboardScoreResponseBody struct {
	UserID     string `form:"user_id" json:"user_id" xml:"user_id"`
	Name       string `form:"name" json:"name" xml:"name"`
	SchoolName string `form:"school_name" json:"school_name" xml:"school_name"`
	Score      int    `form:"score" json:"score" xml:"score"`
}

// CourseResponse is used to define fields on response body types.
type CourseResponse struct {
	Title       string                `form:"title" json:"title" xml:"title"`
	Slug        string                `form:"slug" json:"slug" xml:"slug"`
	Category    string                `form:"category" json:"category" xml:"category"`
	Difficulty  string                `form:"difficulty" json:"difficulty" xml:"difficulty"`
	Description string                `form:"description" json:"description" xml:"description"`
	Enrolled    bool                  `form:"enrolled" json:"enrolled" xml:"enrolled"`
	Publish     bool                  `form:"publish" json:"publish" xml:"publish"`
	Completed   bool                  `form:"completed" json:"completed" xml:"completed"`
	Authors     []*AuthorResponse     `form:"authors,omitempty" json:"authors,omitempty" xml:"authors,omitempty"`
	CourseItems []*CourseItemResponse `form:"course_items,omitempty" json:"course_items,omitempty" xml:"course_items,omitempty"`
	// ID of a file
	ID string `form:"id" json:"id" xml:"id"`
}

// CourseItemResponse is used to define fields on response body types.
type CourseItemResponse struct {
	// to sort after
	Position int `form:"position" json:"position" xml:"position"`
	// ID of a file
	ID string `form:"id" json:"id" xml:"id"`
	// ID of a challenge
	ChallengeID string `form:"challenge_id" json:"challenge_id" xml:"challenge_id"`
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

// NewListEventsResponseBody builds the HTTP response body from the result of
// the "ListEvents" endpoint of the "challenge" service.
func NewListEventsResponseBody(res []*challenge.CTFEvent) ListEventsResponseBody {
	body := make([]*CTFEventResponse, len(res))
	for i, val := range res {
		body[i] = marshalChallengeCTFEventToCTFEventResponse(val)
	}
	return body
}

// NewGetCurrentMonthlyChallengeResponseBody builds the HTTP response body from
// the result of the "GetCurrentMonthlyChallenge" endpoint of the "challenge"
// service.
func NewGetCurrentMonthlyChallengeResponseBody(res *challengeviews.SsmUserMonthlyChallengeView) *GetCurrentMonthlyChallengeResponseBody {
	body := &GetCurrentMonthlyChallengeResponseBody{
		ChallengeID:  *res.ChallengeID,
		DisplayMonth: *res.DisplayMonth,
		StartDate:    *res.StartDate,
		EndDate:      *res.EndDate,
	}
	if res.Challenge != nil {
		body.Challenge = marshalChallengeviewsSsmChallengeViewToSsmChallengeResponseBody(res.Challenge)
	}
	return body
}

// NewSsmUserMonthlyChallengeResponseCollection builds the HTTP response body
// from the result of the "ListMonthlyChallenges" endpoint of the "challenge"
// service.
func NewSsmUserMonthlyChallengeResponseCollection(res challengeviews.SsmUserMonthlyChallengeCollectionView) SsmUserMonthlyChallengeResponseCollection {
	body := make([]*SsmUserMonthlyChallengeResponse, len(res))
	for i, val := range res {
		body[i] = marshalChallengeviewsSsmUserMonthlyChallengeViewToSsmUserMonthlyChallengeResponse(val)
	}
	return body
}

// NewSchoolScoreboardResponseBody builds the HTTP response body from the
// result of the "SchoolScoreboard" endpoint of the "challenge" service.
func NewSchoolScoreboardResponseBody(res *challengeviews.SsmSchoolScoreboardView) *SchoolScoreboardResponseBody {
	body := &SchoolScoreboardResponseBody{}
	if res.Scores != nil {
		body.Scores = make([]*SchoolScoreboardScoreResponseBody, len(res.Scores))
		for i, val := range res.Scores {
			body.Scores[i] = marshalChallengeviewsSchoolScoreboardScoreViewToSchoolScoreboardScoreResponseBody(val)
		}
	}
	return body
}

// NewUserScoreboardResponseBody builds the HTTP response body from the result
// of the "UserScoreboard" endpoint of the "challenge" service.
func NewUserScoreboardResponseBody(res *challengeviews.SsmUserScoreboardView) *UserScoreboardResponseBody {
	body := &UserScoreboardResponseBody{}
	if res.Scores != nil {
		body.Scores = make([]*UserScoreboardScoreResponseBody, len(res.Scores))
		for i, val := range res.Scores {
			body.Scores[i] = marshalChallengeviewsUserScoreboardScoreViewToUserScoreboardScoreResponseBody(val)
		}
	}
	return body
}

// NewListAuthorsResponseBody builds the HTTP response body from the result of
// the "ListAuthors" endpoint of the "challenge" service.
func NewListAuthorsResponseBody(res []*challenge.Author) ListAuthorsResponseBody {
	body := make([]*AuthorResponse, len(res))
	for i, val := range res {
		body[i] = marshalChallengeAuthorToAuthorResponse(val)
	}
	return body
}

// NewListCoursesResponseBody builds the HTTP response body from the result of
// the "ListCourses" endpoint of the "challenge" service.
func NewListCoursesResponseBody(res []*challenge.Course) ListCoursesResponseBody {
	body := make([]*CourseResponse, len(res))
	for i, val := range res {
		body[i] = marshalChallengeCourseToCourseResponse(val)
	}
	return body
}

// NewKnackKodenScoreboardResponseBody builds the HTTP response body from the
// result of the "KnackKodenScoreboard" endpoint of the "challenge" service.
func NewKnackKodenScoreboardResponseBody(res *challengeviews.SsmSchoolScoreboardView) *KnackKodenScoreboardResponseBody {
	body := &KnackKodenScoreboardResponseBody{}
	if res.Scores != nil {
		body.Scores = make([]*SchoolScoreboardScoreResponseBody, len(res.Scores))
		for i, val := range res.Scores {
			body.Scores[i] = marshalChallengeviewsSchoolScoreboardScoreViewToSchoolScoreboardScoreResponseBody(val)
		}
	}
	return body
}

// NewKnackKodenRegisterClassResponseBody builds the HTTP response body from
// the result of the "KnackKodenRegisterClass" endpoint of the "challenge"
// service.
func NewKnackKodenRegisterClassResponseBody(res *challenge.KnackKodenRegisterClassResult) *KnackKodenRegisterClassResponseBody {
	body := &KnackKodenRegisterClassResponseBody{
		Password: res.Password,
	}
	return body
}

// NewKnackKodenGetClassResponseBody builds the HTTP response body from the
// result of the "KnackKodenGetClass" endpoint of the "challenge" service.
func NewKnackKodenGetClassResponseBody(res *challenge.KnackKodenGetClassResult) *KnackKodenGetClassResponseBody {
	body := &KnackKodenGetClassResponseBody{
		ClassName: res.ClassName,
	}
	if res.Solves != nil {
		body.Solves = make([]string, len(res.Solves))
		for i, val := range res.Solves {
			body.Solves[i] = val
		}
	}
	return body
}

// NewGetCurrentMonthlyChallengeNotFoundResponseBody builds the HTTP response
// body from the result of the "GetCurrentMonthlyChallenge" endpoint of the
// "challenge" service.
func NewGetCurrentMonthlyChallengeNotFoundResponseBody(res *goa.ServiceError) *GetCurrentMonthlyChallengeNotFoundResponseBody {
	body := &GetCurrentMonthlyChallengeNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
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

// NewKnackKodenSubmitFlagAlreadySolvedResponseBody builds the HTTP response
// body from the result of the "KnackKodenSubmitFlag" endpoint of the
// "challenge" service.
func NewKnackKodenSubmitFlagAlreadySolvedResponseBody(res *goa.ServiceError) *KnackKodenSubmitFlagAlreadySolvedResponseBody {
	body := &KnackKodenSubmitFlagAlreadySolvedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewKnackKodenSubmitFlagIncorrectFlagResponseBody builds the HTTP response
// body from the result of the "KnackKodenSubmitFlag" endpoint of the
// "challenge" service.
func NewKnackKodenSubmitFlagIncorrectFlagResponseBody(res *goa.ServiceError) *KnackKodenSubmitFlagIncorrectFlagResponseBody {
	body := &KnackKodenSubmitFlagIncorrectFlagResponseBody{
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
func NewListChallengesPayload(slug *string, ids []string, token *string) *challenge.ListChallengesPayload {
	v := &challenge.ListChallengesPayload{}
	v.Slug = slug
	v.Ids = ids
	v.Token = token

	return v
}

// NewListEventsPayload builds a challenge service ListEvents endpoint payload.
func NewListEventsPayload(token *string) *challenge.ListEventsPayload {
	v := &challenge.ListEventsPayload{}
	v.Token = token

	return v
}

// NewGetCurrentMonthlyChallengePayload builds a challenge service
// GetCurrentMonthlyChallenge endpoint payload.
func NewGetCurrentMonthlyChallengePayload(token *string) *challenge.GetCurrentMonthlyChallengePayload {
	v := &challenge.GetCurrentMonthlyChallengePayload{}
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

// NewSchoolScoreboardPayload builds a challenge service SchoolScoreboard
// endpoint payload.
func NewSchoolScoreboardPayload(token *string) *challenge.SchoolScoreboardPayload {
	v := &challenge.SchoolScoreboardPayload{}
	v.Token = token

	return v
}

// NewUserScoreboardPayload builds a challenge service UserScoreboard endpoint
// payload.
func NewUserScoreboardPayload(token *string) *challenge.UserScoreboardPayload {
	v := &challenge.UserScoreboardPayload{}
	v.Token = token

	return v
}

// NewListAuthorsPayload builds a challenge service ListAuthors endpoint
// payload.
func NewListAuthorsPayload(token *string) *challenge.ListAuthorsPayload {
	v := &challenge.ListAuthorsPayload{}
	v.Token = token

	return v
}

// NewListCoursesPayload builds a challenge service ListCourses endpoint
// payload.
func NewListCoursesPayload(token *string) *challenge.ListCoursesPayload {
	v := &challenge.ListCoursesPayload{}
	v.Token = token

	return v
}

// NewEnrollCoursePayload builds a challenge service EnrollCourse endpoint
// payload.
func NewEnrollCoursePayload(id string, token *string) *challenge.EnrollCoursePayload {
	v := &challenge.EnrollCoursePayload{}
	v.ID = id
	v.Token = token

	return v
}

// NewCompleteCoursePayload builds a challenge service CompleteCourse endpoint
// payload.
func NewCompleteCoursePayload(id string, token *string) *challenge.CompleteCoursePayload {
	v := &challenge.CompleteCoursePayload{}
	v.ID = id
	v.Token = token

	return v
}

// NewKnackKodenSubmitFlagPayload builds a challenge service
// KnackKodenSubmitFlag endpoint payload.
func NewKnackKodenSubmitFlagPayload(body *KnackKodenSubmitFlagRequestBody, challengeID string, token *string) *challenge.KnackKodenSubmitFlagPayload {
	v := &challenge.KnackKodenSubmitFlagPayload{
		Flag:     *body.Flag,
		Password: *body.Password,
	}
	v.ChallengeID = challengeID
	v.Token = token

	return v
}

// NewKnackKodenScoreboardPayload builds a challenge service
// KnackKodenScoreboard endpoint payload.
func NewKnackKodenScoreboardPayload(token *string) *challenge.KnackKodenScoreboardPayload {
	v := &challenge.KnackKodenScoreboardPayload{}
	v.Token = token

	return v
}

// NewKnackKodenRegisterClassPayload builds a challenge service
// KnackKodenRegisterClass endpoint payload.
func NewKnackKodenRegisterClassPayload(body *KnackKodenRegisterClassRequestBody, token *string) *challenge.KnackKodenRegisterClassPayload {
	v := &challenge.KnackKodenRegisterClassPayload{
		TeacherFullName: *body.TeacherFullName,
		TeacherEmail:    *body.TeacherEmail,
		TeacherPhonenr:  *body.TeacherPhonenr,
		SchoolName:      *body.SchoolName,
		ClassName:       *body.ClassName,
		PostalCode:      *body.PostalCode,
	}
	v.Token = token

	return v
}

// NewKnackKodenGetClassPayload builds a challenge service KnackKodenGetClass
// endpoint payload.
func NewKnackKodenGetClassPayload(body *KnackKodenGetClassRequestBody, token *string) *challenge.KnackKodenGetClassPayload {
	v := &challenge.KnackKodenGetClassPayload{
		Password: *body.Password,
	}
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

// ValidateKnackKodenSubmitFlagRequestBody runs the validations defined on
// KnackKodenSubmitFlagRequestBody
func ValidateKnackKodenSubmitFlagRequestBody(body *KnackKodenSubmitFlagRequestBody) (err error) {
	if body.Flag == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("flag", "body"))
	}
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	if body.Flag != nil {
		if utf8.RuneCountInString(*body.Flag) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.flag", *body.Flag, utf8.RuneCountInString(*body.Flag), 200, false))
		}
	}
	return
}

// ValidateKnackKodenRegisterClassRequestBody runs the validations defined on
// KnackKodenRegisterClassRequestBody
func ValidateKnackKodenRegisterClassRequestBody(body *KnackKodenRegisterClassRequestBody) (err error) {
	if body.TeacherFullName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("teacher_full_name", "body"))
	}
	if body.TeacherEmail == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("teacher_email", "body"))
	}
	if body.TeacherPhonenr == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("teacher_phonenr", "body"))
	}
	if body.SchoolName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("school_name", "body"))
	}
	if body.ClassName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("class_name", "body"))
	}
	if body.PostalCode == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("postal_code", "body"))
	}
	return
}

// ValidateKnackKodenGetClassRequestBody runs the validations defined on
// KnackKodenGetClassRequestBody
func ValidateKnackKodenGetClassRequestBody(body *KnackKodenGetClassRequestBody) (err error) {
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	return
}
