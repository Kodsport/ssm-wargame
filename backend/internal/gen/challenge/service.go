// Code generated by goa v3.5.2, DO NOT EDIT.
//
// challenge service
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package challenge

import (
	"context"

	challengeviews "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge/views"
	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Service is the challenge service interface.
type Service interface {
	// ListChallenges implements ListChallenges.
	ListChallenges(context.Context, *ListChallengesPayload) (res SsmChallengeCollection, err error)
	// ListEvents implements ListEvents.
	ListEvents(context.Context, *ListEventsPayload) (res []*CTFEvent, err error)
	// ListMonthlyChallenges implements ListMonthlyChallenges.
	ListMonthlyChallenges(context.Context, *ListMonthlyChallengesPayload) (res SsmUsermonthlychallengesCollection, err error)
	// SubmitFlag implements SubmitFlag.
	SubmitFlag(context.Context, *SubmitFlagPayload) (err error)
	// SchoolScoreboard implements SchoolScoreboard.
	SchoolScoreboard(context.Context, *SchoolScoreboardPayload) (res *SsmShoolscoreboard, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "challenge"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"ListChallenges", "ListEvents", "ListMonthlyChallenges", "SubmitFlag", "SchoolScoreboard"}

// ListChallengesPayload is the payload type of the challenge service
// ListChallenges method.
type ListChallengesPayload struct {
	Token *string
}

// SsmChallengeCollection is the result type of the challenge service
// ListChallenges method.
type SsmChallengeCollection []*SsmChallenge

// ListEventsPayload is the payload type of the challenge service ListEvents
// method.
type ListEventsPayload struct {
	Token *string
}

// ListMonthlyChallengesPayload is the payload type of the challenge service
// ListMonthlyChallenges method.
type ListMonthlyChallengesPayload struct {
	Token *string
}

// SsmUsermonthlychallengesCollection is the result type of the challenge
// service ListMonthlyChallenges method.
type SsmUsermonthlychallengesCollection []*SsmUsermonthlychallenges

// SubmitFlagPayload is the payload type of the challenge service SubmitFlag
// method.
type SubmitFlagPayload struct {
	Flag string
	// ID of a challenge
	ChallengeID string
	Token       string
}

// SchoolScoreboardPayload is the payload type of the challenge service
// SchoolScoreboard method.
type SchoolScoreboardPayload struct {
	Token *string
}

// SsmShoolscoreboard is the result type of the challenge service
// SchoolScoreboard method.
type SsmShoolscoreboard struct {
	Scores []*SchoolScoreboardScore
}

// A Wargame challenge
type SsmChallenge struct {
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
	Files    []*ChallengeFiles
	// The numer of people who solved the challenge
	Solves int
	// The ID of the CTF the challenge was taken from
	CtfEventID *string
	// whether the user has solved the challenge or not
	Solved   bool
	Category string
	Authors  []*SsmUser
	Solvers  []*SsmSolver
}

type ChallengeService struct {
}

type ChallengeFiles struct {
	Filename string
	URL      string
}

type SsmUser struct {
	ID       string
	Email    string
	FullName string
	Role     string
	SchoolID *int
}

type SsmSolver struct {
	ID       string
	FullName string
	SolvedAt int64
}

type CTFEvent struct {
	ID   string
	Name string
}

type SsmUsermonthlychallenges struct {
	ChallengeID string
	// The month(s) that the challenge is assigned for
	DisplayMonth string
	// Starting date of the monthly challenge
	StartDate int64
	// Ending date of the monthly challenge
	EndDate   int64
	Challenge *SsmChallenge
}

type SchoolScoreboardScore struct {
	Score      int
	SchoolName string
}

// MakeAlreadySolved builds a goa.ServiceError from an error.
func MakeAlreadySolved(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "already_solved",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeIncorrectFlag builds a goa.ServiceError from an error.
func MakeIncorrectFlag(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "incorrect_flag",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewSsmChallengeCollection initializes result type SsmChallengeCollection
// from viewed result type SsmChallengeCollection.
func NewSsmChallengeCollection(vres challengeviews.SsmChallengeCollection) SsmChallengeCollection {
	return newSsmChallengeCollection(vres.Projected)
}

// NewViewedSsmChallengeCollection initializes viewed result type
// SsmChallengeCollection from result type SsmChallengeCollection using the
// given view.
func NewViewedSsmChallengeCollection(res SsmChallengeCollection, view string) challengeviews.SsmChallengeCollection {
	p := newSsmChallengeCollectionView(res)
	return challengeviews.SsmChallengeCollection{Projected: p, View: "default"}
}

// NewSsmUsermonthlychallengesCollection initializes result type
// SsmUsermonthlychallengesCollection from viewed result type
// SsmUsermonthlychallengesCollection.
func NewSsmUsermonthlychallengesCollection(vres challengeviews.SsmUsermonthlychallengesCollection) SsmUsermonthlychallengesCollection {
	return newSsmUsermonthlychallengesCollection(vres.Projected)
}

// NewViewedSsmUsermonthlychallengesCollection initializes viewed result type
// SsmUsermonthlychallengesCollection from result type
// SsmUsermonthlychallengesCollection using the given view.
func NewViewedSsmUsermonthlychallengesCollection(res SsmUsermonthlychallengesCollection, view string) challengeviews.SsmUsermonthlychallengesCollection {
	p := newSsmUsermonthlychallengesCollectionView(res)
	return challengeviews.SsmUsermonthlychallengesCollection{Projected: p, View: "default"}
}

// NewSsmShoolscoreboard initializes result type SsmShoolscoreboard from viewed
// result type SsmShoolscoreboard.
func NewSsmShoolscoreboard(vres *challengeviews.SsmShoolscoreboard) *SsmShoolscoreboard {
	return newSsmShoolscoreboard(vres.Projected)
}

// NewViewedSsmShoolscoreboard initializes viewed result type
// SsmShoolscoreboard from result type SsmShoolscoreboard using the given view.
func NewViewedSsmShoolscoreboard(res *SsmShoolscoreboard, view string) *challengeviews.SsmShoolscoreboard {
	p := newSsmShoolscoreboardView(res)
	return &challengeviews.SsmShoolscoreboard{Projected: p, View: "default"}
}

// newSsmChallengeCollection converts projected type SsmChallengeCollection to
// service type SsmChallengeCollection.
func newSsmChallengeCollection(vres challengeviews.SsmChallengeCollectionView) SsmChallengeCollection {
	res := make(SsmChallengeCollection, len(vres))
	for i, n := range vres {
		res[i] = newSsmChallenge(n)
	}
	return res
}

// newSsmChallengeCollectionView projects result type SsmChallengeCollection to
// projected type SsmChallengeCollectionView using the "default" view.
func newSsmChallengeCollectionView(res SsmChallengeCollection) challengeviews.SsmChallengeCollectionView {
	vres := make(challengeviews.SsmChallengeCollectionView, len(res))
	for i, n := range res {
		vres[i] = newSsmChallengeView(n)
	}
	return vres
}

// newSsmChallenge converts projected type SsmChallenge to service type
// SsmChallenge.
func newSsmChallenge(vres *challengeviews.SsmChallengeView) *SsmChallenge {
	res := &SsmChallenge{
		CtfEventID: vres.CtfEventID,
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
	if vres.Solved != nil {
		res.Solved = *vres.Solved
	}
	if vres.Category != nil {
		res.Category = *vres.Category
	}
	if vres.Services != nil {
		res.Services = make([]*ChallengeService, len(vres.Services))
		for i, val := range vres.Services {
			res.Services[i] = transformChallengeviewsChallengeServiceViewToChallengeService(val)
		}
	}
	if vres.Files != nil {
		res.Files = make([]*ChallengeFiles, len(vres.Files))
		for i, val := range vres.Files {
			res.Files[i] = transformChallengeviewsChallengeFilesViewToChallengeFiles(val)
		}
	}
	if vres.Authors != nil {
		res.Authors = make([]*SsmUser, len(vres.Authors))
		for i, val := range vres.Authors {
			res.Authors[i] = transformChallengeviewsSsmUserViewToSsmUser(val)
		}
	}
	if vres.Solvers != nil {
		res.Solvers = make([]*SsmSolver, len(vres.Solvers))
		for i, val := range vres.Solvers {
			res.Solvers[i] = transformChallengeviewsSsmSolverViewToSsmSolver(val)
		}
	}
	return res
}

// newSsmChallengeView projects result type SsmChallenge to projected type
// SsmChallengeView using the "default" view.
func newSsmChallengeView(res *SsmChallenge) *challengeviews.SsmChallengeView {
	vres := &challengeviews.SsmChallengeView{
		ID:          &res.ID,
		Slug:        &res.Slug,
		Title:       &res.Title,
		Description: &res.Description,
		Score:       &res.Score,
		Solves:      &res.Solves,
		CtfEventID:  res.CtfEventID,
		Solved:      &res.Solved,
		Category:    &res.Category,
	}
	if res.Services != nil {
		vres.Services = make([]*challengeviews.ChallengeServiceView, len(res.Services))
		for i, val := range res.Services {
			vres.Services[i] = transformChallengeServiceToChallengeviewsChallengeServiceView(val)
		}
	}
	if res.Files != nil {
		vres.Files = make([]*challengeviews.ChallengeFilesView, len(res.Files))
		for i, val := range res.Files {
			vres.Files[i] = transformChallengeFilesToChallengeviewsChallengeFilesView(val)
		}
	}
	if res.Authors != nil {
		vres.Authors = make([]*challengeviews.SsmUserView, len(res.Authors))
		for i, val := range res.Authors {
			vres.Authors[i] = transformSsmUserToChallengeviewsSsmUserView(val)
		}
	}
	if res.Solvers != nil {
		vres.Solvers = make([]*challengeviews.SsmSolverView, len(res.Solvers))
		for i, val := range res.Solvers {
			vres.Solvers[i] = transformSsmSolverToChallengeviewsSsmSolverView(val)
		}
	}
	return vres
}

// newSsmUser converts projected type SsmUser to service type SsmUser.
func newSsmUser(vres *challengeviews.SsmUserView) *SsmUser {
	res := &SsmUser{
		SchoolID: vres.SchoolID,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Email != nil {
		res.Email = *vres.Email
	}
	if vres.FullName != nil {
		res.FullName = *vres.FullName
	}
	if vres.Role != nil {
		res.Role = *vres.Role
	}
	return res
}

// newSsmUserView projects result type SsmUser to projected type SsmUserView
// using the "default" view.
func newSsmUserView(res *SsmUser) *challengeviews.SsmUserView {
	vres := &challengeviews.SsmUserView{
		ID:       &res.ID,
		Email:    &res.Email,
		FullName: &res.FullName,
		Role:     &res.Role,
		SchoolID: res.SchoolID,
	}
	return vres
}

// newSsmSolver converts projected type SsmSolver to service type SsmSolver.
func newSsmSolver(vres *challengeviews.SsmSolverView) *SsmSolver {
	res := &SsmSolver{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.FullName != nil {
		res.FullName = *vres.FullName
	}
	if vres.SolvedAt != nil {
		res.SolvedAt = *vres.SolvedAt
	}
	return res
}

// newSsmSolverView projects result type SsmSolver to projected type
// SsmSolverView using the "default" view.
func newSsmSolverView(res *SsmSolver) *challengeviews.SsmSolverView {
	vres := &challengeviews.SsmSolverView{
		ID:       &res.ID,
		FullName: &res.FullName,
		SolvedAt: &res.SolvedAt,
	}
	return vres
}

// newSsmUsermonthlychallengesCollection converts projected type
// SsmUsermonthlychallengesCollection to service type
// SsmUsermonthlychallengesCollection.
func newSsmUsermonthlychallengesCollection(vres challengeviews.SsmUsermonthlychallengesCollectionView) SsmUsermonthlychallengesCollection {
	res := make(SsmUsermonthlychallengesCollection, len(vres))
	for i, n := range vres {
		res[i] = newSsmUsermonthlychallenges(n)
	}
	return res
}

// newSsmUsermonthlychallengesCollectionView projects result type
// SsmUsermonthlychallengesCollection to projected type
// SsmUsermonthlychallengesCollectionView using the "default" view.
func newSsmUsermonthlychallengesCollectionView(res SsmUsermonthlychallengesCollection) challengeviews.SsmUsermonthlychallengesCollectionView {
	vres := make(challengeviews.SsmUsermonthlychallengesCollectionView, len(res))
	for i, n := range res {
		vres[i] = newSsmUsermonthlychallengesView(n)
	}
	return vres
}

// newSsmUsermonthlychallenges converts projected type SsmUsermonthlychallenges
// to service type SsmUsermonthlychallenges.
func newSsmUsermonthlychallenges(vres *challengeviews.SsmUsermonthlychallengesView) *SsmUsermonthlychallenges {
	res := &SsmUsermonthlychallenges{}
	if vres.ChallengeID != nil {
		res.ChallengeID = *vres.ChallengeID
	}
	if vres.DisplayMonth != nil {
		res.DisplayMonth = *vres.DisplayMonth
	}
	if vres.StartDate != nil {
		res.StartDate = *vres.StartDate
	}
	if vres.EndDate != nil {
		res.EndDate = *vres.EndDate
	}
	if vres.Challenge != nil {
		res.Challenge = newSsmChallenge(vres.Challenge)
	}
	return res
}

// newSsmUsermonthlychallengesView projects result type
// SsmUsermonthlychallenges to projected type SsmUsermonthlychallengesView
// using the "default" view.
func newSsmUsermonthlychallengesView(res *SsmUsermonthlychallenges) *challengeviews.SsmUsermonthlychallengesView {
	vres := &challengeviews.SsmUsermonthlychallengesView{
		ChallengeID:  &res.ChallengeID,
		DisplayMonth: &res.DisplayMonth,
		StartDate:    &res.StartDate,
		EndDate:      &res.EndDate,
	}
	if res.Challenge != nil {
		vres.Challenge = newSsmChallengeView(res.Challenge)
	}
	return vres
}

// newSsmShoolscoreboard converts projected type SsmShoolscoreboard to service
// type SsmShoolscoreboard.
func newSsmShoolscoreboard(vres *challengeviews.SsmShoolscoreboardView) *SsmShoolscoreboard {
	res := &SsmShoolscoreboard{}
	if vres.Scores != nil {
		res.Scores = make([]*SchoolScoreboardScore, len(vres.Scores))
		for i, val := range vres.Scores {
			res.Scores[i] = transformChallengeviewsSchoolScoreboardScoreViewToSchoolScoreboardScore(val)
		}
	}
	return res
}

// newSsmShoolscoreboardView projects result type SsmShoolscoreboard to
// projected type SsmShoolscoreboardView using the "default" view.
func newSsmShoolscoreboardView(res *SsmShoolscoreboard) *challengeviews.SsmShoolscoreboardView {
	vres := &challengeviews.SsmShoolscoreboardView{}
	if res.Scores != nil {
		vres.Scores = make([]*challengeviews.SchoolScoreboardScoreView, len(res.Scores))
		for i, val := range res.Scores {
			vres.Scores[i] = transformSchoolScoreboardScoreToChallengeviewsSchoolScoreboardScoreView(val)
		}
	}
	return vres
}

// transformChallengeviewsChallengeServiceViewToChallengeService builds a value
// of type *ChallengeService from a value of type
// *challengeviews.ChallengeServiceView.
func transformChallengeviewsChallengeServiceViewToChallengeService(v *challengeviews.ChallengeServiceView) *ChallengeService {
	if v == nil {
		return nil
	}
	res := &ChallengeService{}

	return res
}

// transformChallengeviewsChallengeFilesViewToChallengeFiles builds a value of
// type *ChallengeFiles from a value of type *challengeviews.ChallengeFilesView.
func transformChallengeviewsChallengeFilesViewToChallengeFiles(v *challengeviews.ChallengeFilesView) *ChallengeFiles {
	if v == nil {
		return nil
	}
	res := &ChallengeFiles{
		Filename: *v.Filename,
		URL:      *v.URL,
	}

	return res
}

// transformChallengeviewsSsmUserViewToSsmUser builds a value of type *SsmUser
// from a value of type *challengeviews.SsmUserView.
func transformChallengeviewsSsmUserViewToSsmUser(v *challengeviews.SsmUserView) *SsmUser {
	if v == nil {
		return nil
	}
	res := &SsmUser{
		ID:       *v.ID,
		Email:    *v.Email,
		FullName: *v.FullName,
		Role:     *v.Role,
		SchoolID: v.SchoolID,
	}

	return res
}

// transformChallengeviewsSsmSolverViewToSsmSolver builds a value of type
// *SsmSolver from a value of type *challengeviews.SsmSolverView.
func transformChallengeviewsSsmSolverViewToSsmSolver(v *challengeviews.SsmSolverView) *SsmSolver {
	if v == nil {
		return nil
	}
	res := &SsmSolver{
		ID:       *v.ID,
		FullName: *v.FullName,
		SolvedAt: *v.SolvedAt,
	}

	return res
}

// transformChallengeServiceToChallengeviewsChallengeServiceView builds a value
// of type *challengeviews.ChallengeServiceView from a value of type
// *ChallengeService.
func transformChallengeServiceToChallengeviewsChallengeServiceView(v *ChallengeService) *challengeviews.ChallengeServiceView {
	if v == nil {
		return nil
	}
	res := &challengeviews.ChallengeServiceView{}

	return res
}

// transformChallengeFilesToChallengeviewsChallengeFilesView builds a value of
// type *challengeviews.ChallengeFilesView from a value of type *ChallengeFiles.
func transformChallengeFilesToChallengeviewsChallengeFilesView(v *ChallengeFiles) *challengeviews.ChallengeFilesView {
	if v == nil {
		return nil
	}
	res := &challengeviews.ChallengeFilesView{
		Filename: &v.Filename,
		URL:      &v.URL,
	}

	return res
}

// transformSsmUserToChallengeviewsSsmUserView builds a value of type
// *challengeviews.SsmUserView from a value of type *SsmUser.
func transformSsmUserToChallengeviewsSsmUserView(v *SsmUser) *challengeviews.SsmUserView {
	if v == nil {
		return nil
	}
	res := &challengeviews.SsmUserView{
		ID:       &v.ID,
		Email:    &v.Email,
		FullName: &v.FullName,
		Role:     &v.Role,
		SchoolID: v.SchoolID,
	}

	return res
}

// transformSsmSolverToChallengeviewsSsmSolverView builds a value of type
// *challengeviews.SsmSolverView from a value of type *SsmSolver.
func transformSsmSolverToChallengeviewsSsmSolverView(v *SsmSolver) *challengeviews.SsmSolverView {
	if v == nil {
		return nil
	}
	res := &challengeviews.SsmSolverView{
		ID:       &v.ID,
		FullName: &v.FullName,
		SolvedAt: &v.SolvedAt,
	}

	return res
}

// transformChallengeviewsSchoolScoreboardScoreViewToSchoolScoreboardScore
// builds a value of type *SchoolScoreboardScore from a value of type
// *challengeviews.SchoolScoreboardScoreView.
func transformChallengeviewsSchoolScoreboardScoreViewToSchoolScoreboardScore(v *challengeviews.SchoolScoreboardScoreView) *SchoolScoreboardScore {
	if v == nil {
		return nil
	}
	res := &SchoolScoreboardScore{
		Score:      *v.Score,
		SchoolName: *v.SchoolName,
	}

	return res
}

// transformSchoolScoreboardScoreToChallengeviewsSchoolScoreboardScoreView
// builds a value of type *challengeviews.SchoolScoreboardScoreView from a
// value of type *SchoolScoreboardScore.
func transformSchoolScoreboardScoreToChallengeviewsSchoolScoreboardScoreView(v *SchoolScoreboardScore) *challengeviews.SchoolScoreboardScoreView {
	res := &challengeviews.SchoolScoreboardScoreView{
		Score:      &v.Score,
		SchoolName: &v.SchoolName,
	}

	return res
}
