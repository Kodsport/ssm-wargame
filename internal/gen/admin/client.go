// Code generated by goa v3.5.2, DO NOT EDIT.
//
// admin client
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package admin

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "admin" service client.
type Client struct {
	ListChallengesEndpoint         goa.Endpoint
	CreateChallengeEndpoint        goa.Endpoint
	ListMonthlyChallengesEndpoint  goa.Endpoint
	DeleteMonthlyChallengeEndpoint goa.Endpoint
	CreateMonthlyChallengeEndpoint goa.Endpoint
	ListUsersEndpoint              goa.Endpoint
}

// NewClient initializes a "admin" service client given the endpoints.
func NewClient(listChallenges, createChallenge, listMonthlyChallenges, deleteMonthlyChallenge, createMonthlyChallenge, listUsers goa.Endpoint) *Client {
	return &Client{
		ListChallengesEndpoint:         listChallenges,
		CreateChallengeEndpoint:        createChallenge,
		ListMonthlyChallengesEndpoint:  listMonthlyChallenges,
		DeleteMonthlyChallengeEndpoint: deleteMonthlyChallenge,
		CreateMonthlyChallengeEndpoint: createMonthlyChallenge,
		ListUsersEndpoint:              listUsers,
	}
}

// ListChallenges calls the "ListChallenges" endpoint of the "admin" service.
func (c *Client) ListChallenges(ctx context.Context, p *ListChallengesPayload) (res SsmChallengeCollection, err error) {
	var ires interface{}
	ires, err = c.ListChallengesEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(SsmChallengeCollection), nil
}

// CreateChallenge calls the "CreateChallenge" endpoint of the "admin" service.
func (c *Client) CreateChallenge(ctx context.Context, p *CreateChallengePayload) (err error) {
	_, err = c.CreateChallengeEndpoint(ctx, p)
	return
}

// ListMonthlyChallenges calls the "ListMonthlyChallenges" endpoint of the
// "admin" service.
func (c *Client) ListMonthlyChallenges(ctx context.Context, p *ListMonthlyChallengesPayload) (res []*MonthlyChallengeMeta, err error) {
	var ires interface{}
	ires, err = c.ListMonthlyChallengesEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]*MonthlyChallengeMeta), nil
}

// DeleteMonthlyChallenge calls the "DeleteMonthlyChallenge" endpoint of the
// "admin" service.
func (c *Client) DeleteMonthlyChallenge(ctx context.Context, p *DeleteMonthlyChallengePayload) (err error) {
	_, err = c.DeleteMonthlyChallengeEndpoint(ctx, p)
	return
}

// CreateMonthlyChallenge calls the "CreateMonthlyChallenge" endpoint of the
// "admin" service.
func (c *Client) CreateMonthlyChallenge(ctx context.Context, p *CreateMonthlyChallengePayload) (err error) {
	_, err = c.CreateMonthlyChallengeEndpoint(ctx, p)
	return
}

// ListUsers calls the "ListUsers" endpoint of the "admin" service.
func (c *Client) ListUsers(ctx context.Context, p *ListUsersPayload) (res []*SsmUser, err error) {
	var ires interface{}
	ires, err = c.ListUsersEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]*SsmUser), nil
}
