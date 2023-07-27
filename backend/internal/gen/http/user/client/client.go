// Code generated by goa v3.5.2, DO NOT EDIT.
//
// user client HTTP transport
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

// Client lists the user service endpoint HTTP clients.
type Client struct {
	// GetSelf Doer is the HTTP client used to make requests to the GetSelf
	// endpoint.
	GetSelfDoer goahttp.Doer

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
		GetSelfDoer:         doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
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
