// Code generated by goa v3.5.2, DO NOT EDIT.
//
// admin HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	admin "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
	adminviews "github.com/sakerhetsm/ssm-wargame/internal/gen/admin/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildListChallengesRequest instantiates a HTTP request object with method
// and path set to call the "admin" service "ListChallenges" endpoint
func (c *Client) BuildListChallengesRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListChallengesAdminPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("admin", "ListChallenges", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListChallengesRequest returns an encoder for requests sent to the
// admin ListChallenges server.
func EncodeListChallengesRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*admin.ListChallengesPayload)
		if !ok {
			return goahttp.ErrInvalidType("admin", "ListChallenges", "*admin.ListChallengesPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeListChallengesResponse returns a decoder for responses returned by the
// admin ListChallenges endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeListChallengesResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusForbidden
//	- "not_found" (type *goa.ServiceError): http.StatusNotFound
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeListChallengesResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListChallengesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "ListChallenges", err)
			}
			p := NewListChallengesSsmAdminChallengeCollectionOK(body)
			view := "default"
			vres := adminviews.SsmAdminChallengeCollection{Projected: p, View: view}
			if err = adminviews.ValidateSsmAdminChallengeCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("admin", "ListChallenges", err)
			}
			res := admin.NewSsmAdminChallengeCollection(vres)
			return res, nil
		case http.StatusForbidden:
			var (
				body ListChallengesUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "ListChallenges", err)
			}
			err = ValidateListChallengesUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "ListChallenges", err)
			}
			return nil, NewListChallengesUnauthorized(&body)
		case http.StatusNotFound:
			var (
				body ListChallengesNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "ListChallenges", err)
			}
			err = ValidateListChallengesNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "ListChallenges", err)
			}
			return nil, NewListChallengesNotFound(&body)
		case http.StatusBadRequest:
			var (
				body ListChallengesBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "ListChallenges", err)
			}
			err = ValidateListChallengesBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "ListChallenges", err)
			}
			return nil, NewListChallengesBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("admin", "ListChallenges", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateChallengeRequest instantiates a HTTP request object with method
// and path set to call the "admin" service "CreateChallenge" endpoint
func (c *Client) BuildCreateChallengeRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateChallengeAdminPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("admin", "CreateChallenge", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateChallengeRequest returns an encoder for requests sent to the
// admin CreateChallenge server.
func EncodeCreateChallengeRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*admin.CreateChallengePayload)
		if !ok {
			return goahttp.ErrInvalidType("admin", "CreateChallenge", "*admin.CreateChallengePayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewCreateChallengeRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("admin", "CreateChallenge", err)
		}
		return nil
	}
}

// DecodeCreateChallengeResponse returns a decoder for responses returned by
// the admin CreateChallenge endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeCreateChallengeResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusForbidden
//	- "not_found" (type *goa.ServiceError): http.StatusNotFound
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeCreateChallengeResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			return nil, nil
		case http.StatusForbidden:
			var (
				body CreateChallengeUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "CreateChallenge", err)
			}
			err = ValidateCreateChallengeUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "CreateChallenge", err)
			}
			return nil, NewCreateChallengeUnauthorized(&body)
		case http.StatusNotFound:
			var (
				body CreateChallengeNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "CreateChallenge", err)
			}
			err = ValidateCreateChallengeNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "CreateChallenge", err)
			}
			return nil, NewCreateChallengeNotFound(&body)
		case http.StatusBadRequest:
			var (
				body CreateChallengeBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "CreateChallenge", err)
			}
			err = ValidateCreateChallengeBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "CreateChallenge", err)
			}
			return nil, NewCreateChallengeBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("admin", "CreateChallenge", resp.StatusCode, string(body))
		}
	}
}

// BuildPresignChallFileUploadRequest instantiates a HTTP request object with
// method and path set to call the "admin" service "PresignChallFileUpload"
// endpoint
func (c *Client) BuildPresignChallFileUploadRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		challengeID string
	)
	{
		p, ok := v.(*admin.PresignChallFileUploadPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("admin", "PresignChallFileUpload", "*admin.PresignChallFileUploadPayload", v)
		}
		challengeID = p.ChallengeID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: PresignChallFileUploadAdminPath(challengeID)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("admin", "PresignChallFileUpload", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodePresignChallFileUploadRequest returns an encoder for requests sent to
// the admin PresignChallFileUpload server.
func EncodePresignChallFileUploadRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*admin.PresignChallFileUploadPayload)
		if !ok {
			return goahttp.ErrInvalidType("admin", "PresignChallFileUpload", "*admin.PresignChallFileUploadPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewPresignChallFileUploadRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("admin", "PresignChallFileUpload", err)
		}
		return nil
	}
}

// DecodePresignChallFileUploadResponse returns a decoder for responses
// returned by the admin PresignChallFileUpload endpoint. restoreBody controls
// whether the response body should be restored after having been read.
// DecodePresignChallFileUploadResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusForbidden
//	- "not_found" (type *goa.ServiceError): http.StatusNotFound
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodePresignChallFileUploadResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body PresignChallFileUploadResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "PresignChallFileUpload", err)
			}
			err = ValidatePresignChallFileUploadResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "PresignChallFileUpload", err)
			}
			res := NewPresignChallFileUploadResultOK(&body)
			return res, nil
		case http.StatusForbidden:
			var (
				body PresignChallFileUploadUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "PresignChallFileUpload", err)
			}
			err = ValidatePresignChallFileUploadUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "PresignChallFileUpload", err)
			}
			return nil, NewPresignChallFileUploadUnauthorized(&body)
		case http.StatusNotFound:
			var (
				body PresignChallFileUploadNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "PresignChallFileUpload", err)
			}
			err = ValidatePresignChallFileUploadNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "PresignChallFileUpload", err)
			}
			return nil, NewPresignChallFileUploadNotFound(&body)
		case http.StatusBadRequest:
			var (
				body PresignChallFileUploadBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "PresignChallFileUpload", err)
			}
			err = ValidatePresignChallFileUploadBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "PresignChallFileUpload", err)
			}
			return nil, NewPresignChallFileUploadBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("admin", "PresignChallFileUpload", resp.StatusCode, string(body))
		}
	}
}

// BuildListMonthlyChallengesRequest instantiates a HTTP request object with
// method and path set to call the "admin" service "ListMonthlyChallenges"
// endpoint
func (c *Client) BuildListMonthlyChallengesRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListMonthlyChallengesAdminPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("admin", "ListMonthlyChallenges", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListMonthlyChallengesRequest returns an encoder for requests sent to
// the admin ListMonthlyChallenges server.
func EncodeListMonthlyChallengesRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*admin.ListMonthlyChallengesPayload)
		if !ok {
			return goahttp.ErrInvalidType("admin", "ListMonthlyChallenges", "*admin.ListMonthlyChallengesPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeListMonthlyChallengesResponse returns a decoder for responses returned
// by the admin ListMonthlyChallenges endpoint. restoreBody controls whether
// the response body should be restored after having been read.
// DecodeListMonthlyChallengesResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusForbidden
//	- "not_found" (type *goa.ServiceError): http.StatusNotFound
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeListMonthlyChallengesResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListMonthlyChallengesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "ListMonthlyChallenges", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateMonthlyChallengeMetaResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "ListMonthlyChallenges", err)
			}
			res := NewListMonthlyChallengesMonthlyChallengeMetaOK(body)
			return res, nil
		case http.StatusForbidden:
			var (
				body ListMonthlyChallengesUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "ListMonthlyChallenges", err)
			}
			err = ValidateListMonthlyChallengesUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "ListMonthlyChallenges", err)
			}
			return nil, NewListMonthlyChallengesUnauthorized(&body)
		case http.StatusNotFound:
			var (
				body ListMonthlyChallengesNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "ListMonthlyChallenges", err)
			}
			err = ValidateListMonthlyChallengesNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "ListMonthlyChallenges", err)
			}
			return nil, NewListMonthlyChallengesNotFound(&body)
		case http.StatusBadRequest:
			var (
				body ListMonthlyChallengesBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "ListMonthlyChallenges", err)
			}
			err = ValidateListMonthlyChallengesBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "ListMonthlyChallenges", err)
			}
			return nil, NewListMonthlyChallengesBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("admin", "ListMonthlyChallenges", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteMonthlyChallengeRequest instantiates a HTTP request object with
// method and path set to call the "admin" service "DeleteMonthlyChallenge"
// endpoint
func (c *Client) BuildDeleteMonthlyChallengeRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		challengeID string
	)
	{
		p, ok := v.(*admin.DeleteMonthlyChallengePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("admin", "DeleteMonthlyChallenge", "*admin.DeleteMonthlyChallengePayload", v)
		}
		challengeID = p.ChallengeID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteMonthlyChallengeAdminPath(challengeID)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("admin", "DeleteMonthlyChallenge", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteMonthlyChallengeRequest returns an encoder for requests sent to
// the admin DeleteMonthlyChallenge server.
func EncodeDeleteMonthlyChallengeRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*admin.DeleteMonthlyChallengePayload)
		if !ok {
			return goahttp.ErrInvalidType("admin", "DeleteMonthlyChallenge", "*admin.DeleteMonthlyChallengePayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeDeleteMonthlyChallengeResponse returns a decoder for responses
// returned by the admin DeleteMonthlyChallenge endpoint. restoreBody controls
// whether the response body should be restored after having been read.
// DecodeDeleteMonthlyChallengeResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusForbidden
//	- "not_found" (type *goa.ServiceError): http.StatusNotFound
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeDeleteMonthlyChallengeResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusForbidden:
			var (
				body DeleteMonthlyChallengeUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "DeleteMonthlyChallenge", err)
			}
			err = ValidateDeleteMonthlyChallengeUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "DeleteMonthlyChallenge", err)
			}
			return nil, NewDeleteMonthlyChallengeUnauthorized(&body)
		case http.StatusNotFound:
			var (
				body DeleteMonthlyChallengeNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "DeleteMonthlyChallenge", err)
			}
			err = ValidateDeleteMonthlyChallengeNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "DeleteMonthlyChallenge", err)
			}
			return nil, NewDeleteMonthlyChallengeNotFound(&body)
		case http.StatusBadRequest:
			var (
				body DeleteMonthlyChallengeBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "DeleteMonthlyChallenge", err)
			}
			err = ValidateDeleteMonthlyChallengeBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "DeleteMonthlyChallenge", err)
			}
			return nil, NewDeleteMonthlyChallengeBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("admin", "DeleteMonthlyChallenge", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteFileRequest instantiates a HTTP request object with method and
// path set to call the "admin" service "DeleteFile" endpoint
func (c *Client) BuildDeleteFileRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		fileID string
	)
	{
		p, ok := v.(*admin.DeleteFilePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("admin", "DeleteFile", "*admin.DeleteFilePayload", v)
		}
		fileID = p.FileID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteFileAdminPath(fileID)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("admin", "DeleteFile", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteFileRequest returns an encoder for requests sent to the admin
// DeleteFile server.
func EncodeDeleteFileRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*admin.DeleteFilePayload)
		if !ok {
			return goahttp.ErrInvalidType("admin", "DeleteFile", "*admin.DeleteFilePayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeDeleteFileResponse returns a decoder for responses returned by the
// admin DeleteFile endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeDeleteFileResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusForbidden
//	- "not_found" (type *goa.ServiceError): http.StatusNotFound
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeDeleteFileResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusForbidden:
			var (
				body DeleteFileUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "DeleteFile", err)
			}
			err = ValidateDeleteFileUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "DeleteFile", err)
			}
			return nil, NewDeleteFileUnauthorized(&body)
		case http.StatusNotFound:
			var (
				body DeleteFileNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "DeleteFile", err)
			}
			err = ValidateDeleteFileNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "DeleteFile", err)
			}
			return nil, NewDeleteFileNotFound(&body)
		case http.StatusBadRequest:
			var (
				body DeleteFileBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "DeleteFile", err)
			}
			err = ValidateDeleteFileBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "DeleteFile", err)
			}
			return nil, NewDeleteFileBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("admin", "DeleteFile", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateMonthlyChallengeRequest instantiates a HTTP request object with
// method and path set to call the "admin" service "CreateMonthlyChallenge"
// endpoint
func (c *Client) BuildCreateMonthlyChallengeRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateMonthlyChallengeAdminPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("admin", "CreateMonthlyChallenge", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateMonthlyChallengeRequest returns an encoder for requests sent to
// the admin CreateMonthlyChallenge server.
func EncodeCreateMonthlyChallengeRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*admin.CreateMonthlyChallengePayload)
		if !ok {
			return goahttp.ErrInvalidType("admin", "CreateMonthlyChallenge", "*admin.CreateMonthlyChallengePayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewCreateMonthlyChallengeRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("admin", "CreateMonthlyChallenge", err)
		}
		return nil
	}
}

// DecodeCreateMonthlyChallengeResponse returns a decoder for responses
// returned by the admin CreateMonthlyChallenge endpoint. restoreBody controls
// whether the response body should be restored after having been read.
// DecodeCreateMonthlyChallengeResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusForbidden
//	- "not_found" (type *goa.ServiceError): http.StatusNotFound
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeCreateMonthlyChallengeResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusForbidden:
			var (
				body CreateMonthlyChallengeUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "CreateMonthlyChallenge", err)
			}
			err = ValidateCreateMonthlyChallengeUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "CreateMonthlyChallenge", err)
			}
			return nil, NewCreateMonthlyChallengeUnauthorized(&body)
		case http.StatusNotFound:
			var (
				body CreateMonthlyChallengeNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "CreateMonthlyChallenge", err)
			}
			err = ValidateCreateMonthlyChallengeNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "CreateMonthlyChallenge", err)
			}
			return nil, NewCreateMonthlyChallengeNotFound(&body)
		case http.StatusBadRequest:
			var (
				body CreateMonthlyChallengeBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "CreateMonthlyChallenge", err)
			}
			err = ValidateCreateMonthlyChallengeBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "CreateMonthlyChallenge", err)
			}
			return nil, NewCreateMonthlyChallengeBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("admin", "CreateMonthlyChallenge", resp.StatusCode, string(body))
		}
	}
}

// BuildListUsersRequest instantiates a HTTP request object with method and
// path set to call the "admin" service "ListUsers" endpoint
func (c *Client) BuildListUsersRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListUsersAdminPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("admin", "ListUsers", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListUsersRequest returns an encoder for requests sent to the admin
// ListUsers server.
func EncodeListUsersRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*admin.ListUsersPayload)
		if !ok {
			return goahttp.ErrInvalidType("admin", "ListUsers", "*admin.ListUsersPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeListUsersResponse returns a decoder for responses returned by the
// admin ListUsers endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeListUsersResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusForbidden
//	- "not_found" (type *goa.ServiceError): http.StatusNotFound
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeListUsersResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListUsersResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "ListUsers", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateSsmUserResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "ListUsers", err)
			}
			res := NewListUsersSsmUserOK(body)
			return res, nil
		case http.StatusForbidden:
			var (
				body ListUsersUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "ListUsers", err)
			}
			err = ValidateListUsersUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "ListUsers", err)
			}
			return nil, NewListUsersUnauthorized(&body)
		case http.StatusNotFound:
			var (
				body ListUsersNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "ListUsers", err)
			}
			err = ValidateListUsersNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "ListUsers", err)
			}
			return nil, NewListUsersNotFound(&body)
		case http.StatusBadRequest:
			var (
				body ListUsersBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "ListUsers", err)
			}
			err = ValidateListUsersBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "ListUsers", err)
			}
			return nil, NewListUsersBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("admin", "ListUsers", resp.StatusCode, string(body))
		}
	}
}

// unmarshalSsmAdminChallengeResponseToAdminviewsSsmAdminChallengeView builds a
// value of type *adminviews.SsmAdminChallengeView from a value of type
// *SsmAdminChallengeResponse.
func unmarshalSsmAdminChallengeResponseToAdminviewsSsmAdminChallengeView(v *SsmAdminChallengeResponse) *adminviews.SsmAdminChallengeView {
	res := &adminviews.SsmAdminChallengeView{
		ID:          v.ID,
		Slug:        v.Slug,
		Title:       v.Title,
		Description: v.Description,
		Score:       v.Score,
		Published:   v.Published,
		Solves:      v.Solves,
	}
	if v.Services != nil {
		res.Services = make([]*adminviews.ChallengeServiceView, len(v.Services))
		for i, val := range v.Services {
			res.Services[i] = unmarshalChallengeServiceResponseToAdminviewsChallengeServiceView(val)
		}
	}
	res.Files = make([]*adminviews.AdminChallengeFilesView, len(v.Files))
	for i, val := range v.Files {
		res.Files[i] = unmarshalAdminChallengeFilesResponseToAdminviewsAdminChallengeFilesView(val)
	}

	return res
}

// unmarshalChallengeServiceResponseToAdminviewsChallengeServiceView builds a
// value of type *adminviews.ChallengeServiceView from a value of type
// *ChallengeServiceResponse.
func unmarshalChallengeServiceResponseToAdminviewsChallengeServiceView(v *ChallengeServiceResponse) *adminviews.ChallengeServiceView {
	if v == nil {
		return nil
	}
	res := &adminviews.ChallengeServiceView{}

	return res
}

// unmarshalAdminChallengeFilesResponseToAdminviewsAdminChallengeFilesView
// builds a value of type *adminviews.AdminChallengeFilesView from a value of
// type *AdminChallengeFilesResponse.
func unmarshalAdminChallengeFilesResponseToAdminviewsAdminChallengeFilesView(v *AdminChallengeFilesResponse) *adminviews.AdminChallengeFilesView {
	res := &adminviews.AdminChallengeFilesView{
		ID:       v.ID,
		Filename: v.Filename,
		URL:      v.URL,
		Bucket:   v.Bucket,
		Key:      v.Key,
		Size:     v.Size,
		Md5:      v.Md5,
	}

	return res
}

// unmarshalMonthlyChallengeMetaResponseToAdminMonthlyChallengeMeta builds a
// value of type *admin.MonthlyChallengeMeta from a value of type
// *MonthlyChallengeMetaResponse.
func unmarshalMonthlyChallengeMetaResponseToAdminMonthlyChallengeMeta(v *MonthlyChallengeMetaResponse) *admin.MonthlyChallengeMeta {
	res := &admin.MonthlyChallengeMeta{
		DisplayMonth: *v.DisplayMonth,
		StartDate:    *v.StartDate,
		EndDate:      *v.EndDate,
	}

	return res
}

// unmarshalSsmUserResponseToAdminSsmUser builds a value of type *admin.SsmUser
// from a value of type *SsmUserResponse.
func unmarshalSsmUserResponseToAdminSsmUser(v *SsmUserResponse) *admin.SsmUser {
	res := &admin.SsmUser{
		ID:        *v.ID,
		Email:     *v.Email,
		FirstName: *v.FirstName,
		LastName:  *v.LastName,
		Role:      *v.Role,
	}

	return res
}
