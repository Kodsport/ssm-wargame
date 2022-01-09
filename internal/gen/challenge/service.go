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
	// ListMonthlyChallenges implements ListMonthlyChallenges.
	ListMonthlyChallenges(context.Context, *ListMonthlyChallengesPayload) (res SsmMonthlyChallengeCollection, err error)
	// SubmitFlag implements SubmitFlag.
	SubmitFlag(context.Context, *SubmitFlagPayload) (err error)
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
var MethodNames = [3]string{"ListChallenges", "ListMonthlyChallenges", "SubmitFlag"}

// ListChallengesPayload is the payload type of the challenge service
// ListChallenges method.
type ListChallengesPayload struct {
	Token *string
}

// SsmChallengeCollection is the result type of the challenge service
// ListChallenges method.
type SsmChallengeCollection []*SsmChallenge

// ListMonthlyChallengesPayload is the payload type of the challenge service
// ListMonthlyChallenges method.
type ListMonthlyChallengesPayload struct {
	Token *string
}

// SsmMonthlyChallengeCollection is the result type of the challenge service
// ListMonthlyChallenges method.
type SsmMonthlyChallengeCollection []*SsmMonthlyChallenge

// SubmitFlagPayload is the payload type of the challenge service SubmitFlag
// method.
type SubmitFlagPayload struct {
	Flag string
	// ID of a challenge
	ChallengeID string
	Token       string
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
	Score     int32
	Services  []*ChallengeService
	Files     []*ChallengeFiles
	Published bool
	// The numer of people who solved the challenge
	Solves int64
}

type ChallengeService struct {
}

type ChallengeFiles struct {
}

// A monthly challenge
type SsmMonthlyChallenge struct {
	// The month(s) that the challenge is assigned for
	DisplayMonth string
	// Starting date of the monthly challenge
	StartDate string
	// Ending date of the monthly challenge
	EndDate string
	ID      string
	// A unique string that can be used in URLs
	Slug string
	// Title displayed to user
	Title string
	// A short text describing the challenge
	Description string
	// The number of points given to the solver
	Score     int32
	Services  []*ChallengeService
	Files     []*ChallengeFiles
	Published bool
	// The numer of people who solved the challenge
	Solves int64
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

// NewSsmMonthlyChallengeCollection initializes result type
// SsmMonthlyChallengeCollection from viewed result type
// SsmMonthlyChallengeCollection.
func NewSsmMonthlyChallengeCollection(vres challengeviews.SsmMonthlyChallengeCollection) SsmMonthlyChallengeCollection {
	return newSsmMonthlyChallengeCollection(vres.Projected)
}

// NewViewedSsmMonthlyChallengeCollection initializes viewed result type
// SsmMonthlyChallengeCollection from result type SsmMonthlyChallengeCollection
// using the given view.
func NewViewedSsmMonthlyChallengeCollection(res SsmMonthlyChallengeCollection, view string) challengeviews.SsmMonthlyChallengeCollection {
	p := newSsmMonthlyChallengeCollectionView(res)
	return challengeviews.SsmMonthlyChallengeCollection{Projected: p, View: "default"}
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
	res := &SsmChallenge{}
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
	if vres.Published != nil {
		res.Published = *vres.Published
	}
	if vres.Solves != nil {
		res.Solves = *vres.Solves
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
		Published:   &res.Published,
		Solves:      &res.Solves,
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
	return vres
}

// newSsmMonthlyChallengeCollection converts projected type
// SsmMonthlyChallengeCollection to service type SsmMonthlyChallengeCollection.
func newSsmMonthlyChallengeCollection(vres challengeviews.SsmMonthlyChallengeCollectionView) SsmMonthlyChallengeCollection {
	res := make(SsmMonthlyChallengeCollection, len(vres))
	for i, n := range vres {
		res[i] = newSsmMonthlyChallenge(n)
	}
	return res
}

// newSsmMonthlyChallengeCollectionView projects result type
// SsmMonthlyChallengeCollection to projected type
// SsmMonthlyChallengeCollectionView using the "default" view.
func newSsmMonthlyChallengeCollectionView(res SsmMonthlyChallengeCollection) challengeviews.SsmMonthlyChallengeCollectionView {
	vres := make(challengeviews.SsmMonthlyChallengeCollectionView, len(res))
	for i, n := range res {
		vres[i] = newSsmMonthlyChallengeView(n)
	}
	return vres
}

// newSsmMonthlyChallenge converts projected type SsmMonthlyChallenge to
// service type SsmMonthlyChallenge.
func newSsmMonthlyChallenge(vres *challengeviews.SsmMonthlyChallengeView) *SsmMonthlyChallenge {
	res := &SsmMonthlyChallenge{}
	if vres.DisplayMonth != nil {
		res.DisplayMonth = *vres.DisplayMonth
	}
	if vres.StartDate != nil {
		res.StartDate = *vres.StartDate
	}
	if vres.EndDate != nil {
		res.EndDate = *vres.EndDate
	}
	return res
}

// newSsmMonthlyChallengeView projects result type SsmMonthlyChallenge to
// projected type SsmMonthlyChallengeView using the "default" view.
func newSsmMonthlyChallengeView(res *SsmMonthlyChallenge) *challengeviews.SsmMonthlyChallengeView {
	vres := &challengeviews.SsmMonthlyChallengeView{
		DisplayMonth: &res.DisplayMonth,
		StartDate:    &res.StartDate,
		EndDate:      &res.EndDate,
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
	res := &ChallengeFiles{}

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
	res := &challengeviews.ChallengeFilesView{}

	return res
}
