// Code generated by goa v3.5.2, DO NOT EDIT.
//
// challenge endpoints
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package challenge

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "challenge" service endpoints.
type Endpoints struct {
	ListChallenges             goa.Endpoint
	ListEvents                 goa.Endpoint
	GetCurrentMonthlyChallenge goa.Endpoint
	ListMonthlyChallenges      goa.Endpoint
	SubmitFlag                 goa.Endpoint
	SchoolScoreboard           goa.Endpoint
}

// NewEndpoints wraps the methods of the "challenge" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		ListChallenges:             NewListChallengesEndpoint(s, a.JWTAuth),
		ListEvents:                 NewListEventsEndpoint(s, a.JWTAuth),
		GetCurrentMonthlyChallenge: NewGetCurrentMonthlyChallengeEndpoint(s, a.JWTAuth),
		ListMonthlyChallenges:      NewListMonthlyChallengesEndpoint(s, a.JWTAuth),
		SubmitFlag:                 NewSubmitFlagEndpoint(s, a.JWTAuth),
		SchoolScoreboard:           NewSchoolScoreboardEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "challenge" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.ListChallenges = m(e.ListChallenges)
	e.ListEvents = m(e.ListEvents)
	e.GetCurrentMonthlyChallenge = m(e.GetCurrentMonthlyChallenge)
	e.ListMonthlyChallenges = m(e.ListMonthlyChallenges)
	e.SubmitFlag = m(e.SubmitFlag)
	e.SchoolScoreboard = m(e.SchoolScoreboard)
}

// NewListChallengesEndpoint returns an endpoint function that calls the method
// "ListChallenges" of service "challenge".
func NewListChallengesEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListChallengesPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.ListChallenges(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedSsmChallengeCollection(res, "default")
		return vres, nil
	}
}

// NewListEventsEndpoint returns an endpoint function that calls the method
// "ListEvents" of service "challenge".
func NewListEventsEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListEventsPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.ListEvents(ctx, p)
	}
}

// NewGetCurrentMonthlyChallengeEndpoint returns an endpoint function that
// calls the method "GetCurrentMonthlyChallenge" of service "challenge".
func NewGetCurrentMonthlyChallengeEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetCurrentMonthlyChallengePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.GetCurrentMonthlyChallenge(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedSsmUserMonthlyChallenge(res, "default")
		return vres, nil
	}
}

// NewListMonthlyChallengesEndpoint returns an endpoint function that calls the
// method "ListMonthlyChallenges" of service "challenge".
func NewListMonthlyChallengesEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListMonthlyChallengesPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.ListMonthlyChallenges(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedSsmUserMonthlyChallengeCollection(res, "default")
		return vres, nil
	}
}

// NewSubmitFlagEndpoint returns an endpoint function that calls the method
// "SubmitFlag" of service "challenge".
func NewSubmitFlagEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SubmitFlagPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.SubmitFlag(ctx, p)
	}
}

// NewSchoolScoreboardEndpoint returns an endpoint function that calls the
// method "SchoolScoreboard" of service "challenge".
func NewSchoolScoreboardEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SchoolScoreboardPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.SchoolScoreboard(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedSsmShoolscoreboard(res, "default")
		return vres, nil
	}
}
