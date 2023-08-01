// Code generated by goa v3.5.2, DO NOT EDIT.
//
// challenge client
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package challenge

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "challenge" service client.
type Client struct {
	ListChallengesEndpoint        goa.Endpoint
	ListMonthlyChallengesEndpoint goa.Endpoint
	SubmitFlagEndpoint            goa.Endpoint
	SchoolScoreboardEndpoint      goa.Endpoint
}

// NewClient initializes a "challenge" service client given the endpoints.
func NewClient(listChallenges, listMonthlyChallenges, submitFlag, schoolScoreboard goa.Endpoint) *Client {
	return &Client{
		ListChallengesEndpoint:        listChallenges,
		ListMonthlyChallengesEndpoint: listMonthlyChallenges,
		SubmitFlagEndpoint:            submitFlag,
		SchoolScoreboardEndpoint:      schoolScoreboard,
	}
}

// ListChallenges calls the "ListChallenges" endpoint of the "challenge"
// service.
func (c *Client) ListChallenges(ctx context.Context, p *ListChallengesPayload) (res SsmChallengeCollection, err error) {
	var ires interface{}
	ires, err = c.ListChallengesEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(SsmChallengeCollection), nil
}

// ListMonthlyChallenges calls the "ListMonthlyChallenges" endpoint of the
// "challenge" service.
func (c *Client) ListMonthlyChallenges(ctx context.Context, p *ListMonthlyChallengesPayload) (res SsmUsermonthlychallengesCollection, err error) {
	var ires interface{}
	ires, err = c.ListMonthlyChallengesEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(SsmUsermonthlychallengesCollection), nil
}

// SubmitFlag calls the "SubmitFlag" endpoint of the "challenge" service.
// SubmitFlag may return the following errors:
//   - "already_solved" (type *goa.ServiceError)
//   - "incorrect_flag" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) SubmitFlag(ctx context.Context, p *SubmitFlagPayload) (err error) {
	_, err = c.SubmitFlagEndpoint(ctx, p)
	return
}

// SchoolScoreboard calls the "SchoolScoreboard" endpoint of the "challenge"
// service.
func (c *Client) SchoolScoreboard(ctx context.Context, p *SchoolScoreboardPayload) (res *SsmShoolscoreboard, err error) {
	var ires interface{}
	ires, err = c.SchoolScoreboardEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*SsmShoolscoreboard), nil
}
