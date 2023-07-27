// Code generated by goa v3.5.2, DO NOT EDIT.
//
// admin client HTTP transport
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the admin service endpoint HTTP clients.
type Client struct {
	// ListChallenges Doer is the HTTP client used to make requests to the
	// ListChallenges endpoint.
	ListChallengesDoer goahttp.Doer

	// CreateChallenge Doer is the HTTP client used to make requests to the
	// CreateChallenge endpoint.
	CreateChallengeDoer goahttp.Doer

	// UpdateChallenge Doer is the HTTP client used to make requests to the
	// UpdateChallenge endpoint.
	UpdateChallengeDoer goahttp.Doer

	// PresignChallFileUpload Doer is the HTTP client used to make requests to the
	// PresignChallFileUpload endpoint.
	PresignChallFileUploadDoer goahttp.Doer

	// ListMonthlyChallenges Doer is the HTTP client used to make requests to the
	// ListMonthlyChallenges endpoint.
	ListMonthlyChallengesDoer goahttp.Doer

	// DeleteMonthlyChallenge Doer is the HTTP client used to make requests to the
	// DeleteMonthlyChallenge endpoint.
	DeleteMonthlyChallengeDoer goahttp.Doer

	// DeleteFile Doer is the HTTP client used to make requests to the DeleteFile
	// endpoint.
	DeleteFileDoer goahttp.Doer

	// CreateMonthlyChallenge Doer is the HTTP client used to make requests to the
	// CreateMonthlyChallenge endpoint.
	CreateMonthlyChallengeDoer goahttp.Doer

	// ListUsers Doer is the HTTP client used to make requests to the ListUsers
	// endpoint.
	ListUsersDoer goahttp.Doer

	// AddFlag Doer is the HTTP client used to make requests to the AddFlag
	// endpoint.
	AddFlagDoer goahttp.Doer

	// DeleteFlag Doer is the HTTP client used to make requests to the DeleteFlag
	// endpoint.
	DeleteFlagDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the admin service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ListChallengesDoer:         doer,
		CreateChallengeDoer:        doer,
		UpdateChallengeDoer:        doer,
		PresignChallFileUploadDoer: doer,
		ListMonthlyChallengesDoer:  doer,
		DeleteMonthlyChallengeDoer: doer,
		DeleteFileDoer:             doer,
		CreateMonthlyChallengeDoer: doer,
		ListUsersDoer:              doer,
		AddFlagDoer:                doer,
		DeleteFlagDoer:             doer,
		RestoreResponseBody:        restoreBody,
		scheme:                     scheme,
		host:                       host,
		decoder:                    dec,
		encoder:                    enc,
	}
}

// ListChallenges returns an endpoint that makes HTTP requests to the admin
// service ListChallenges server.
func (c *Client) ListChallenges() goa.Endpoint {
	var (
		encodeRequest  = EncodeListChallengesRequest(c.encoder)
		decodeResponse = DecodeListChallengesResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListChallengesRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListChallengesDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("admin", "ListChallenges", err)
		}
		return decodeResponse(resp)
	}
}

// CreateChallenge returns an endpoint that makes HTTP requests to the admin
// service CreateChallenge server.
func (c *Client) CreateChallenge() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateChallengeRequest(c.encoder)
		decodeResponse = DecodeCreateChallengeResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateChallengeRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateChallengeDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("admin", "CreateChallenge", err)
		}
		return decodeResponse(resp)
	}
}

// UpdateChallenge returns an endpoint that makes HTTP requests to the admin
// service UpdateChallenge server.
func (c *Client) UpdateChallenge() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateChallengeRequest(c.encoder)
		decodeResponse = DecodeUpdateChallengeResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUpdateChallengeRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateChallengeDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("admin", "UpdateChallenge", err)
		}
		return decodeResponse(resp)
	}
}

// PresignChallFileUpload returns an endpoint that makes HTTP requests to the
// admin service PresignChallFileUpload server.
func (c *Client) PresignChallFileUpload() goa.Endpoint {
	var (
		encodeRequest  = EncodePresignChallFileUploadRequest(c.encoder)
		decodeResponse = DecodePresignChallFileUploadResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildPresignChallFileUploadRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PresignChallFileUploadDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("admin", "PresignChallFileUpload", err)
		}
		return decodeResponse(resp)
	}
}

// ListMonthlyChallenges returns an endpoint that makes HTTP requests to the
// admin service ListMonthlyChallenges server.
func (c *Client) ListMonthlyChallenges() goa.Endpoint {
	var (
		encodeRequest  = EncodeListMonthlyChallengesRequest(c.encoder)
		decodeResponse = DecodeListMonthlyChallengesResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListMonthlyChallengesRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListMonthlyChallengesDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("admin", "ListMonthlyChallenges", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteMonthlyChallenge returns an endpoint that makes HTTP requests to the
// admin service DeleteMonthlyChallenge server.
func (c *Client) DeleteMonthlyChallenge() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteMonthlyChallengeRequest(c.encoder)
		decodeResponse = DecodeDeleteMonthlyChallengeResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteMonthlyChallengeRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteMonthlyChallengeDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("admin", "DeleteMonthlyChallenge", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteFile returns an endpoint that makes HTTP requests to the admin service
// DeleteFile server.
func (c *Client) DeleteFile() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteFileRequest(c.encoder)
		decodeResponse = DecodeDeleteFileResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteFileRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteFileDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("admin", "DeleteFile", err)
		}
		return decodeResponse(resp)
	}
}

// CreateMonthlyChallenge returns an endpoint that makes HTTP requests to the
// admin service CreateMonthlyChallenge server.
func (c *Client) CreateMonthlyChallenge() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateMonthlyChallengeRequest(c.encoder)
		decodeResponse = DecodeCreateMonthlyChallengeResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateMonthlyChallengeRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateMonthlyChallengeDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("admin", "CreateMonthlyChallenge", err)
		}
		return decodeResponse(resp)
	}
}

// ListUsers returns an endpoint that makes HTTP requests to the admin service
// ListUsers server.
func (c *Client) ListUsers() goa.Endpoint {
	var (
		encodeRequest  = EncodeListUsersRequest(c.encoder)
		decodeResponse = DecodeListUsersResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListUsersRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListUsersDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("admin", "ListUsers", err)
		}
		return decodeResponse(resp)
	}
}

// AddFlag returns an endpoint that makes HTTP requests to the admin service
// AddFlag server.
func (c *Client) AddFlag() goa.Endpoint {
	var (
		encodeRequest  = EncodeAddFlagRequest(c.encoder)
		decodeResponse = DecodeAddFlagResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAddFlagRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AddFlagDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("admin", "AddFlag", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteFlag returns an endpoint that makes HTTP requests to the admin service
// DeleteFlag server.
func (c *Client) DeleteFlag() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteFlagRequest(c.encoder)
		decodeResponse = DecodeDeleteFlagResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteFlagRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteFlagDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("admin", "DeleteFlag", err)
		}
		return decodeResponse(resp)
	}
}
