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
)

// Endpoints wraps the "challenge" service endpoints.
type Endpoints struct {
	ListChallenges  goa.Endpoint
	CreateChallenge goa.Endpoint
	SubmitFlag      goa.Endpoint
}

// NewEndpoints wraps the methods of the "challenge" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		ListChallenges:  NewListChallengesEndpoint(s),
		CreateChallenge: NewCreateChallengeEndpoint(s),
		SubmitFlag:      NewSubmitFlagEndpoint(s),
	}
}

// Use applies the given middleware to all the "challenge" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.ListChallenges = m(e.ListChallenges)
	e.CreateChallenge = m(e.CreateChallenge)
	e.SubmitFlag = m(e.SubmitFlag)
}

// NewListChallengesEndpoint returns an endpoint function that calls the method
// "ListChallenges" of service "challenge".
func NewListChallengesEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListChallengesPayload)
		res, view, err := s.ListChallenges(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedSsmChallengeCollection(res, view)
		return vres, nil
	}
}

// NewCreateChallengeEndpoint returns an endpoint function that calls the
// method "CreateChallenge" of service "challenge".
func NewCreateChallengeEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, s.CreateChallenge(ctx)
	}
}

// NewSubmitFlagEndpoint returns an endpoint function that calls the method
// "SubmitFlag" of service "challenge".
func NewSubmitFlagEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SubmitFlagPayload)
		return nil, s.SubmitFlag(ctx, p)
	}
}
