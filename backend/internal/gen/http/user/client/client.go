// Code generated by goa v3.5.2, DO NOT EDIT.
//
// user client HTTP transport
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the user service endpoint HTTP clients.
type Client struct {
	// GetSelf Doer is the HTTP client used to make requests to the GetSelf
	// endpoint.
	GetSelfDoer goahttp.Doer

	// UpdateSelf Doer is the HTTP client used to make requests to the UpdateSelf
	// endpoint.
	UpdateSelfDoer goahttp.Doer

	// CompleteOnboarding Doer is the HTTP client used to make requests to the
	// CompleteOnboarding endpoint.
	CompleteOnboardingDoer goahttp.Doer

	// JoinSchool Doer is the HTTP client used to make requests to the JoinSchool
	// endpoint.
	JoinSchoolDoer goahttp.Doer

	// LeaveSchool Doer is the HTTP client used to make requests to the LeaveSchool
	// endpoint.
	LeaveSchoolDoer goahttp.Doer

	// SearchSchools Doer is the HTTP client used to make requests to the
	// SearchSchools endpoint.
	SearchSchoolsDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the user service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GetSelfDoer:            doer,
		UpdateSelfDoer:         doer,
		CompleteOnboardingDoer: doer,
		JoinSchoolDoer:         doer,
		LeaveSchoolDoer:        doer,
		SearchSchoolsDoer:      doer,
		RestoreResponseBody:    restoreBody,
		scheme:                 scheme,
		host:                   host,
		decoder:                dec,
		encoder:                enc,
	}
}

// GetSelf returns an endpoint that makes HTTP requests to the user service
// GetSelf server.
func (c *Client) GetSelf() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetSelfRequest(c.encoder)
		decodeResponse = DecodeGetSelfResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetSelfRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetSelfDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("user", "GetSelf", err)
		}
		return decodeResponse(resp)
	}
}

// UpdateSelf returns an endpoint that makes HTTP requests to the user service
// UpdateSelf server.
func (c *Client) UpdateSelf() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateSelfRequest(c.encoder)
		decodeResponse = DecodeUpdateSelfResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUpdateSelfRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateSelfDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("user", "UpdateSelf", err)
		}
		return decodeResponse(resp)
	}
}

// CompleteOnboarding returns an endpoint that makes HTTP requests to the user
// service CompleteOnboarding server.
func (c *Client) CompleteOnboarding() goa.Endpoint {
	var (
		encodeRequest  = EncodeCompleteOnboardingRequest(c.encoder)
		decodeResponse = DecodeCompleteOnboardingResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCompleteOnboardingRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CompleteOnboardingDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("user", "CompleteOnboarding", err)
		}
		return decodeResponse(resp)
	}
}

// JoinSchool returns an endpoint that makes HTTP requests to the user service
// JoinSchool server.
func (c *Client) JoinSchool() goa.Endpoint {
	var (
		encodeRequest  = EncodeJoinSchoolRequest(c.encoder)
		decodeResponse = DecodeJoinSchoolResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildJoinSchoolRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.JoinSchoolDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("user", "JoinSchool", err)
		}
		return decodeResponse(resp)
	}
}

// LeaveSchool returns an endpoint that makes HTTP requests to the user service
// LeaveSchool server.
func (c *Client) LeaveSchool() goa.Endpoint {
	var (
		encodeRequest  = EncodeLeaveSchoolRequest(c.encoder)
		decodeResponse = DecodeLeaveSchoolResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildLeaveSchoolRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.LeaveSchoolDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("user", "LeaveSchool", err)
		}
		return decodeResponse(resp)
	}
}

// SearchSchools returns an endpoint that makes HTTP requests to the user
// service SearchSchools server.
func (c *Client) SearchSchools() goa.Endpoint {
	var (
		encodeRequest  = EncodeSearchSchoolsRequest(c.encoder)
		decodeResponse = DecodeSearchSchoolsResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSearchSchoolsRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SearchSchoolsDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("user", "SearchSchools", err)
		}
		return decodeResponse(resp)
	}
}
