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

	// GetChallengeMeta Doer is the HTTP client used to make requests to the
	// GetChallengeMeta endpoint.
	GetChallengeMetaDoer goahttp.Doer

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

	// ListCategories Doer is the HTTP client used to make requests to the
	// ListCategories endpoint.
	ListCategoriesDoer goahttp.Doer

	// ChalltoolsImport Doer is the HTTP client used to make requests to the
	// ChalltoolsImport endpoint.
	ChalltoolsImportDoer goahttp.Doer

	// ListCTFEvents Doer is the HTTP client used to make requests to the
	// ListCTFEvents endpoint.
	ListCTFEventsDoer goahttp.Doer

	// CreateCTFEvent Doer is the HTTP client used to make requests to the
	// CreateCTFEvent endpoint.
	CreateCTFEventDoer goahttp.Doer

	// DeleteCTFEvent Doer is the HTTP client used to make requests to the
	// DeleteCTFEvent endpoint.
	DeleteCTFEventDoer goahttp.Doer

	// CreateCTFEventImportToken Doer is the HTTP client used to make requests to
	// the CreateCTFEventImportToken endpoint.
	CreateCTFEventImportTokenDoer goahttp.Doer

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
		ListChallengesDoer:            doer,
		GetChallengeMetaDoer:          doer,
		CreateChallengeDoer:           doer,
		UpdateChallengeDoer:           doer,
		PresignChallFileUploadDoer:    doer,
		ListMonthlyChallengesDoer:     doer,
		DeleteMonthlyChallengeDoer:    doer,
		DeleteFileDoer:                doer,
		CreateMonthlyChallengeDoer:    doer,
		ListUsersDoer:                 doer,
		AddFlagDoer:                   doer,
		DeleteFlagDoer:                doer,
		ListCategoriesDoer:            doer,
		ChalltoolsImportDoer:          doer,
		ListCTFEventsDoer:             doer,
		CreateCTFEventDoer:            doer,
		DeleteCTFEventDoer:            doer,
		CreateCTFEventImportTokenDoer: doer,
		RestoreResponseBody:           restoreBody,
		scheme:                        scheme,
		host:                          host,
		decoder:                       dec,
		encoder:                       enc,
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

// GetChallengeMeta returns an endpoint that makes HTTP requests to the admin
// service GetChallengeMeta server.
func (c *Client) GetChallengeMeta() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetChallengeMetaRequest(c.encoder)
		decodeResponse = DecodeGetChallengeMetaResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetChallengeMetaRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetChallengeMetaDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("admin", "GetChallengeMeta", err)
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

// ListCategories returns an endpoint that makes HTTP requests to the admin
// service ListCategories server.
func (c *Client) ListCategories() goa.Endpoint {
	var (
		encodeRequest  = EncodeListCategoriesRequest(c.encoder)
		decodeResponse = DecodeListCategoriesResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListCategoriesRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListCategoriesDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("admin", "ListCategories", err)
		}
		return decodeResponse(resp)
	}
}

// ChalltoolsImport returns an endpoint that makes HTTP requests to the admin
// service ChalltoolsImport server.
func (c *Client) ChalltoolsImport() goa.Endpoint {
	var (
		encodeRequest  = EncodeChalltoolsImportRequest(c.encoder)
		decodeResponse = DecodeChalltoolsImportResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildChalltoolsImportRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ChalltoolsImportDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("admin", "ChalltoolsImport", err)
		}
		return decodeResponse(resp)
	}
}

// ListCTFEvents returns an endpoint that makes HTTP requests to the admin
// service ListCTFEvents server.
func (c *Client) ListCTFEvents() goa.Endpoint {
	var (
		encodeRequest  = EncodeListCTFEventsRequest(c.encoder)
		decodeResponse = DecodeListCTFEventsResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListCTFEventsRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListCTFEventsDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("admin", "ListCTFEvents", err)
		}
		return decodeResponse(resp)
	}
}

// CreateCTFEvent returns an endpoint that makes HTTP requests to the admin
// service CreateCTFEvent server.
func (c *Client) CreateCTFEvent() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateCTFEventRequest(c.encoder)
		decodeResponse = DecodeCreateCTFEventResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateCTFEventRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateCTFEventDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("admin", "CreateCTFEvent", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteCTFEvent returns an endpoint that makes HTTP requests to the admin
// service DeleteCTFEvent server.
func (c *Client) DeleteCTFEvent() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteCTFEventRequest(c.encoder)
		decodeResponse = DecodeDeleteCTFEventResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteCTFEventRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteCTFEventDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("admin", "DeleteCTFEvent", err)
		}
		return decodeResponse(resp)
	}
}

// CreateCTFEventImportToken returns an endpoint that makes HTTP requests to
// the admin service CreateCTFEventImportToken server.
func (c *Client) CreateCTFEventImportToken() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateCTFEventImportTokenRequest(c.encoder)
		decodeResponse = DecodeCreateCTFEventImportTokenResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateCTFEventImportTokenRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateCTFEventImportTokenDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("admin", "CreateCTFEventImportToken", err)
		}
		return decodeResponse(resp)
	}
}
