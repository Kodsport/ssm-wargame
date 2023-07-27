// Code generated by goa v3.5.2, DO NOT EDIT.
//
// auth endpoints
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package auth

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "auth" service endpoints.
type Endpoints struct {
	GenerateDiscordAuthURL goa.Endpoint
	ExchangeDiscord        goa.Endpoint
}

// NewEndpoints wraps the methods of the "auth" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		GenerateDiscordAuthURL: NewGenerateDiscordAuthURLEndpoint(s),
		ExchangeDiscord:        NewExchangeDiscordEndpoint(s),
	}
}

// Use applies the given middleware to all the "auth" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GenerateDiscordAuthURL = m(e.GenerateDiscordAuthURL)
	e.ExchangeDiscord = m(e.ExchangeDiscord)
}

// NewGenerateDiscordAuthURLEndpoint returns an endpoint function that calls
// the method "GenerateDiscordAuthURL" of service "auth".
func NewGenerateDiscordAuthURLEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GenerateDiscordAuthURL(ctx)
	}
}

// NewExchangeDiscordEndpoint returns an endpoint function that calls the
// method "ExchangeDiscord" of service "auth".
func NewExchangeDiscordEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ExchangeDiscordPayload)
		return s.ExchangeDiscord(ctx, p)
	}
}
