// Code generated by goa v3.5.2, DO NOT EDIT.
//
// auth client
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package auth

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "auth" service client.
type Client struct {
	GenerateDiscordAuthURLEndpoint goa.Endpoint
	ExchangeDiscordEndpoint        goa.Endpoint
}

// NewClient initializes a "auth" service client given the endpoints.
func NewClient(generateDiscordAuthURL, exchangeDiscord goa.Endpoint) *Client {
	return &Client{
		GenerateDiscordAuthURLEndpoint: generateDiscordAuthURL,
		ExchangeDiscordEndpoint:        exchangeDiscord,
	}
}

// GenerateDiscordAuthURL calls the "GenerateDiscordAuthURL" endpoint of the
// "auth" service.
func (c *Client) GenerateDiscordAuthURL(ctx context.Context) (res *GenerateDiscordAuthURLResult, err error) {
	var ires interface{}
	ires, err = c.GenerateDiscordAuthURLEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*GenerateDiscordAuthURLResult), nil
}

// ExchangeDiscord calls the "ExchangeDiscord" endpoint of the "auth" service.
func (c *Client) ExchangeDiscord(ctx context.Context, p *ExchangeDiscordPayload) (res *ExchangeDiscordResult, err error) {
	var ires interface{}
	ires, err = c.ExchangeDiscordEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ExchangeDiscordResult), nil
}
