// Code generated by goa v3.5.2, DO NOT EDIT.
//
// challenge HTTP client encoders and decoders
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

	challenge "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	challengeviews "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildListChallengesRequest instantiates a HTTP request object with method
// and path set to call the "challenge" service "ListChallenges" endpoint
func (c *Client) BuildListChallengesRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListChallengesChallengePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("challenge", "ListChallenges", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListChallengesRequest returns an encoder for requests sent to the
// challenge ListChallenges server.
func EncodeListChallengesRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*challenge.ListChallengesPayload)
		if !ok {
			return goahttp.ErrInvalidType("challenge", "ListChallenges", "*challenge.ListChallengesPayload", v)
		}
		if p.Token != nil {
			head := *p.Token
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
// challenge ListChallenges endpoint. restoreBody controls whether the response
// body should be restored after having been read.
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
				return nil, goahttp.ErrDecodingError("challenge", "ListChallenges", err)
			}
			p := NewListChallengesSsmChallengeCollectionOK(body)
			view := "default"
			vres := challengeviews.SsmChallengeCollection{Projected: p, View: view}
			if err = challengeviews.ValidateSsmChallengeCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("challenge", "ListChallenges", err)
			}
			res := challenge.NewSsmChallengeCollection(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("challenge", "ListChallenges", resp.StatusCode, string(body))
		}
	}
}

// BuildListMonthlyChallengesRequest instantiates a HTTP request object with
// method and path set to call the "challenge" service "ListMonthlyChallenges"
// endpoint
func (c *Client) BuildListMonthlyChallengesRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListMonthlyChallengesChallengePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("challenge", "ListMonthlyChallenges", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListMonthlyChallengesRequest returns an encoder for requests sent to
// the challenge ListMonthlyChallenges server.
func EncodeListMonthlyChallengesRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*challenge.ListMonthlyChallengesPayload)
		if !ok {
			return goahttp.ErrInvalidType("challenge", "ListMonthlyChallenges", "*challenge.ListMonthlyChallengesPayload", v)
		}
		if p.Token != nil {
			head := *p.Token
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
// by the challenge ListMonthlyChallenges endpoint. restoreBody controls
// whether the response body should be restored after having been read.
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
				return nil, goahttp.ErrDecodingError("challenge", "ListMonthlyChallenges", err)
			}
			p := NewListMonthlyChallengesSsmMonthlyChallengeCollectionOK(body)
			view := "default"
			vres := challengeviews.SsmMonthlyChallengeCollection{Projected: p, View: view}
			if err = challengeviews.ValidateSsmMonthlyChallengeCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("challenge", "ListMonthlyChallenges", err)
			}
			res := challenge.NewSsmMonthlyChallengeCollection(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("challenge", "ListMonthlyChallenges", resp.StatusCode, string(body))
		}
	}
}

// BuildSubmitFlagRequest instantiates a HTTP request object with method and
// path set to call the "challenge" service "SubmitFlag" endpoint
func (c *Client) BuildSubmitFlagRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		challengeID string
	)
	{
		p, ok := v.(*challenge.SubmitFlagPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("challenge", "SubmitFlag", "*challenge.SubmitFlagPayload", v)
		}
		challengeID = p.ChallengeID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SubmitFlagChallengePath(challengeID)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("challenge", "SubmitFlag", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSubmitFlagRequest returns an encoder for requests sent to the
// challenge SubmitFlag server.
func EncodeSubmitFlagRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*challenge.SubmitFlagPayload)
		if !ok {
			return goahttp.ErrInvalidType("challenge", "SubmitFlag", "*challenge.SubmitFlagPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewSubmitFlagRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("challenge", "SubmitFlag", err)
		}
		return nil
	}
}

// DecodeSubmitFlagResponse returns a decoder for responses returned by the
// challenge SubmitFlag endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeSubmitFlagResponse may return the following errors:
//	- "already_solved" (type *goa.ServiceError): http.StatusConflict
//	- "incorrect_flag" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeSubmitFlagResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
			return nil, nil
		case http.StatusConflict:
			var (
				body SubmitFlagAlreadySolvedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("challenge", "SubmitFlag", err)
			}
			err = ValidateSubmitFlagAlreadySolvedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("challenge", "SubmitFlag", err)
			}
			return nil, NewSubmitFlagAlreadySolved(&body)
		case http.StatusBadRequest:
			var (
				body SubmitFlagIncorrectFlagResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("challenge", "SubmitFlag", err)
			}
			err = ValidateSubmitFlagIncorrectFlagResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("challenge", "SubmitFlag", err)
			}
			return nil, NewSubmitFlagIncorrectFlag(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("challenge", "SubmitFlag", resp.StatusCode, string(body))
		}
	}
}

// unmarshalSsmChallengeResponseToChallengeviewsSsmChallengeView builds a value
// of type *challengeviews.SsmChallengeView from a value of type
// *SsmChallengeResponse.
func unmarshalSsmChallengeResponseToChallengeviewsSsmChallengeView(v *SsmChallengeResponse) *challengeviews.SsmChallengeView {
	res := &challengeviews.SsmChallengeView{
		ID:          v.ID,
		Slug:        v.Slug,
		Title:       v.Title,
		Description: v.Description,
		Score:       v.Score,
		Published:   v.Published,
		Solves:      v.Solves,
	}
	if v.Services != nil {
		res.Services = make([]*challengeviews.ChallengeServiceView, len(v.Services))
		for i, val := range v.Services {
			res.Services[i] = unmarshalChallengeServiceResponseToChallengeviewsChallengeServiceView(val)
		}
	}
	if v.Files != nil {
		res.Files = make([]*challengeviews.ChallengeFilesView, len(v.Files))
		for i, val := range v.Files {
			res.Files[i] = unmarshalChallengeFilesResponseToChallengeviewsChallengeFilesView(val)
		}
	}

	return res
}

// unmarshalChallengeServiceResponseToChallengeviewsChallengeServiceView builds
// a value of type *challengeviews.ChallengeServiceView from a value of type
// *ChallengeServiceResponse.
func unmarshalChallengeServiceResponseToChallengeviewsChallengeServiceView(v *ChallengeServiceResponse) *challengeviews.ChallengeServiceView {
	if v == nil {
		return nil
	}
	res := &challengeviews.ChallengeServiceView{}

	return res
}

// unmarshalChallengeFilesResponseToChallengeviewsChallengeFilesView builds a
// value of type *challengeviews.ChallengeFilesView from a value of type
// *ChallengeFilesResponse.
func unmarshalChallengeFilesResponseToChallengeviewsChallengeFilesView(v *ChallengeFilesResponse) *challengeviews.ChallengeFilesView {
	if v == nil {
		return nil
	}
	res := &challengeviews.ChallengeFilesView{
		Filename: v.Filename,
		URL:      v.URL,
	}

	return res
}

// unmarshalSsmMonthlyChallengeResponseToChallengeviewsSsmMonthlyChallengeView
// builds a value of type *challengeviews.SsmMonthlyChallengeView from a value
// of type *SsmMonthlyChallengeResponse.
func unmarshalSsmMonthlyChallengeResponseToChallengeviewsSsmMonthlyChallengeView(v *SsmMonthlyChallengeResponse) *challengeviews.SsmMonthlyChallengeView {
	res := &challengeviews.SsmMonthlyChallengeView{
		DisplayMonth: v.DisplayMonth,
		StartDate:    v.StartDate,
		EndDate:      v.EndDate,
		ID:           v.ID,
		Slug:         v.Slug,
		Title:        v.Title,
		Description:  v.Description,
		Score:        v.Score,
		Published:    v.Published,
		Solves:       v.Solves,
	}
	if v.Services != nil {
		res.Services = make([]*challengeviews.ChallengeServiceView, len(v.Services))
		for i, val := range v.Services {
			res.Services[i] = unmarshalChallengeServiceResponseToChallengeviewsChallengeServiceView(val)
		}
	}
	if v.Files != nil {
		res.Files = make([]*challengeviews.ChallengeFilesView, len(v.Files))
		for i, val := range v.Files {
			res.Files[i] = unmarshalChallengeFilesResponseToChallengeviewsChallengeFilesView(val)
		}
	}

	return res
}
