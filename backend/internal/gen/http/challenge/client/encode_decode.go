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
	goa "goa.design/goa/v3/pkg"
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
		values := req.URL.Query()
		if p.Slug != nil {
			values.Add("slug", *p.Slug)
		}
		req.URL.RawQuery = values.Encode()
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

// BuildListEventsRequest instantiates a HTTP request object with method and
// path set to call the "challenge" service "ListEvents" endpoint
func (c *Client) BuildListEventsRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListEventsChallengePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("challenge", "ListEvents", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListEventsRequest returns an encoder for requests sent to the
// challenge ListEvents server.
func EncodeListEventsRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*challenge.ListEventsPayload)
		if !ok {
			return goahttp.ErrInvalidType("challenge", "ListEvents", "*challenge.ListEventsPayload", v)
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

// DecodeListEventsResponse returns a decoder for responses returned by the
// challenge ListEvents endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeListEventsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body ListEventsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("challenge", "ListEvents", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateCTFEventResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("challenge", "ListEvents", err)
			}
			res := NewListEventsCTFEventOK(body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("challenge", "ListEvents", resp.StatusCode, string(body))
		}
	}
}

// BuildGetCurrentMonthlyChallengeRequest instantiates a HTTP request object
// with method and path set to call the "challenge" service
// "GetCurrentMonthlyChallenge" endpoint
func (c *Client) BuildGetCurrentMonthlyChallengeRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetCurrentMonthlyChallengeChallengePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("challenge", "GetCurrentMonthlyChallenge", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetCurrentMonthlyChallengeRequest returns an encoder for requests sent
// to the challenge GetCurrentMonthlyChallenge server.
func EncodeGetCurrentMonthlyChallengeRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*challenge.GetCurrentMonthlyChallengePayload)
		if !ok {
			return goahttp.ErrInvalidType("challenge", "GetCurrentMonthlyChallenge", "*challenge.GetCurrentMonthlyChallengePayload", v)
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

// DecodeGetCurrentMonthlyChallengeResponse returns a decoder for responses
// returned by the challenge GetCurrentMonthlyChallenge endpoint. restoreBody
// controls whether the response body should be restored after having been read.
// DecodeGetCurrentMonthlyChallengeResponse may return the following errors:
//   - "not_found" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeGetCurrentMonthlyChallengeResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body GetCurrentMonthlyChallengeResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("challenge", "GetCurrentMonthlyChallenge", err)
			}
			p := NewGetCurrentMonthlyChallengeSsmUserMonthlyChallengeOK(&body)
			view := "default"
			vres := &challengeviews.SsmUserMonthlyChallenge{Projected: p, View: view}
			if err = challengeviews.ValidateSsmUserMonthlyChallenge(vres); err != nil {
				return nil, goahttp.ErrValidationError("challenge", "GetCurrentMonthlyChallenge", err)
			}
			res := challenge.NewSsmUserMonthlyChallenge(vres)
			return res, nil
		case http.StatusNotFound:
			var (
				body GetCurrentMonthlyChallengeNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("challenge", "GetCurrentMonthlyChallenge", err)
			}
			err = ValidateGetCurrentMonthlyChallengeNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("challenge", "GetCurrentMonthlyChallenge", err)
			}
			return nil, NewGetCurrentMonthlyChallengeNotFound(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("challenge", "GetCurrentMonthlyChallenge", resp.StatusCode, string(body))
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
			p := NewListMonthlyChallengesSsmUserMonthlyChallengeCollectionOK(body)
			view := "default"
			vres := challengeviews.SsmUserMonthlyChallengeCollection{Projected: p, View: view}
			if err = challengeviews.ValidateSsmUserMonthlyChallengeCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("challenge", "ListMonthlyChallenges", err)
			}
			res := challenge.NewSsmUserMonthlyChallengeCollection(vres)
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
//   - "already_solved" (type *goa.ServiceError): http.StatusConflict
//   - "incorrect_flag" (type *goa.ServiceError): http.StatusBadRequest
//   - error: internal error
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

// BuildSchoolScoreboardRequest instantiates a HTTP request object with method
// and path set to call the "challenge" service "SchoolScoreboard" endpoint
func (c *Client) BuildSchoolScoreboardRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SchoolScoreboardChallengePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("challenge", "SchoolScoreboard", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSchoolScoreboardRequest returns an encoder for requests sent to the
// challenge SchoolScoreboard server.
func EncodeSchoolScoreboardRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*challenge.SchoolScoreboardPayload)
		if !ok {
			return goahttp.ErrInvalidType("challenge", "SchoolScoreboard", "*challenge.SchoolScoreboardPayload", v)
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

// DecodeSchoolScoreboardResponse returns a decoder for responses returned by
// the challenge SchoolScoreboard endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeSchoolScoreboardResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body SchoolScoreboardResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("challenge", "SchoolScoreboard", err)
			}
			p := NewSchoolScoreboardSsmSchoolScoreboardOK(&body)
			view := "default"
			vres := &challengeviews.SsmSchoolScoreboard{Projected: p, View: view}
			if err = challengeviews.ValidateSsmSchoolScoreboard(vres); err != nil {
				return nil, goahttp.ErrValidationError("challenge", "SchoolScoreboard", err)
			}
			res := challenge.NewSsmSchoolScoreboard(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("challenge", "SchoolScoreboard", resp.StatusCode, string(body))
		}
	}
}

// BuildUserScoreboardRequest instantiates a HTTP request object with method
// and path set to call the "challenge" service "UserScoreboard" endpoint
func (c *Client) BuildUserScoreboardRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UserScoreboardChallengePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("challenge", "UserScoreboard", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUserScoreboardRequest returns an encoder for requests sent to the
// challenge UserScoreboard server.
func EncodeUserScoreboardRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*challenge.UserScoreboardPayload)
		if !ok {
			return goahttp.ErrInvalidType("challenge", "UserScoreboard", "*challenge.UserScoreboardPayload", v)
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

// DecodeUserScoreboardResponse returns a decoder for responses returned by the
// challenge UserScoreboard endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeUserScoreboardResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body UserScoreboardResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("challenge", "UserScoreboard", err)
			}
			p := NewUserScoreboardSsmUserScoreboardOK(&body)
			view := "default"
			vres := &challengeviews.SsmUserScoreboard{Projected: p, View: view}
			if err = challengeviews.ValidateSsmUserScoreboard(vres); err != nil {
				return nil, goahttp.ErrValidationError("challenge", "UserScoreboard", err)
			}
			res := challenge.NewSsmUserScoreboard(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("challenge", "UserScoreboard", resp.StatusCode, string(body))
		}
	}
}

// BuildListAuthorsRequest instantiates a HTTP request object with method and
// path set to call the "challenge" service "ListAuthors" endpoint
func (c *Client) BuildListAuthorsRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListAuthorsChallengePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("challenge", "ListAuthors", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListAuthorsRequest returns an encoder for requests sent to the
// challenge ListAuthors server.
func EncodeListAuthorsRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*challenge.ListAuthorsPayload)
		if !ok {
			return goahttp.ErrInvalidType("challenge", "ListAuthors", "*challenge.ListAuthorsPayload", v)
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

// DecodeListAuthorsResponse returns a decoder for responses returned by the
// challenge ListAuthors endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeListAuthorsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body ListAuthorsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("challenge", "ListAuthors", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateAuthorResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("challenge", "ListAuthors", err)
			}
			res := NewListAuthorsAuthorOK(body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("challenge", "ListAuthors", resp.StatusCode, string(body))
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
		Solves:      v.Solves,
		CtfEventID:  v.CtfEventID,
		Solved:      v.Solved,
		Category:    v.Category,
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
	if v.Authors != nil {
		res.Authors = make([]*challengeviews.AuthorView, len(v.Authors))
		for i, val := range v.Authors {
			res.Authors[i] = unmarshalAuthorResponseToChallengeviewsAuthorView(val)
		}
	}
	if v.Solvers != nil {
		res.Solvers = make([]*challengeviews.SsmSolverView, len(v.Solvers))
		for i, val := range v.Solvers {
			res.Solvers[i] = unmarshalSsmSolverResponseToChallengeviewsSsmSolverView(val)
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
	res := &challengeviews.ChallengeServiceView{
		UserDisplay: v.UserDisplay,
		Hyperlink:   v.Hyperlink,
	}

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

// unmarshalAuthorResponseToChallengeviewsAuthorView builds a value of type
// *challengeviews.AuthorView from a value of type *AuthorResponse.
func unmarshalAuthorResponseToChallengeviewsAuthorView(v *AuthorResponse) *challengeviews.AuthorView {
	if v == nil {
		return nil
	}
	res := &challengeviews.AuthorView{
		ID:          v.ID,
		FullName:    v.FullName,
		Description: v.Description,
		Sponsor:     v.Sponsor,
		Slug:        v.Slug,
		ImageURL:    v.ImageURL,
		Publish:     v.Publish,
	}

	return res
}

// unmarshalSsmSolverResponseToChallengeviewsSsmSolverView builds a value of
// type *challengeviews.SsmSolverView from a value of type *SsmSolverResponse.
func unmarshalSsmSolverResponseToChallengeviewsSsmSolverView(v *SsmSolverResponse) *challengeviews.SsmSolverView {
	if v == nil {
		return nil
	}
	res := &challengeviews.SsmSolverView{
		ID:       v.ID,
		FullName: v.FullName,
		SolvedAt: v.SolvedAt,
	}

	return res
}

// unmarshalCTFEventResponseToChallengeCTFEvent builds a value of type
// *challenge.CTFEvent from a value of type *CTFEventResponse.
func unmarshalCTFEventResponseToChallengeCTFEvent(v *CTFEventResponse) *challenge.CTFEvent {
	res := &challenge.CTFEvent{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}

// unmarshalSsmChallengeResponseBodyToChallengeviewsSsmChallengeView builds a
// value of type *challengeviews.SsmChallengeView from a value of type
// *SsmChallengeResponseBody.
func unmarshalSsmChallengeResponseBodyToChallengeviewsSsmChallengeView(v *SsmChallengeResponseBody) *challengeviews.SsmChallengeView {
	res := &challengeviews.SsmChallengeView{
		ID:          v.ID,
		Slug:        v.Slug,
		Title:       v.Title,
		Description: v.Description,
		Score:       v.Score,
		Solves:      v.Solves,
		CtfEventID:  v.CtfEventID,
		Solved:      v.Solved,
		Category:    v.Category,
	}
	if v.Services != nil {
		res.Services = make([]*challengeviews.ChallengeServiceView, len(v.Services))
		for i, val := range v.Services {
			res.Services[i] = unmarshalChallengeServiceResponseBodyToChallengeviewsChallengeServiceView(val)
		}
	}
	if v.Files != nil {
		res.Files = make([]*challengeviews.ChallengeFilesView, len(v.Files))
		for i, val := range v.Files {
			res.Files[i] = unmarshalChallengeFilesResponseBodyToChallengeviewsChallengeFilesView(val)
		}
	}
	if v.Authors != nil {
		res.Authors = make([]*challengeviews.AuthorView, len(v.Authors))
		for i, val := range v.Authors {
			res.Authors[i] = unmarshalAuthorResponseBodyToChallengeviewsAuthorView(val)
		}
	}
	if v.Solvers != nil {
		res.Solvers = make([]*challengeviews.SsmSolverView, len(v.Solvers))
		for i, val := range v.Solvers {
			res.Solvers[i] = unmarshalSsmSolverResponseBodyToChallengeviewsSsmSolverView(val)
		}
	}

	return res
}

// unmarshalChallengeServiceResponseBodyToChallengeviewsChallengeServiceView
// builds a value of type *challengeviews.ChallengeServiceView from a value of
// type *ChallengeServiceResponseBody.
func unmarshalChallengeServiceResponseBodyToChallengeviewsChallengeServiceView(v *ChallengeServiceResponseBody) *challengeviews.ChallengeServiceView {
	if v == nil {
		return nil
	}
	res := &challengeviews.ChallengeServiceView{
		UserDisplay: v.UserDisplay,
		Hyperlink:   v.Hyperlink,
	}

	return res
}

// unmarshalChallengeFilesResponseBodyToChallengeviewsChallengeFilesView builds
// a value of type *challengeviews.ChallengeFilesView from a value of type
// *ChallengeFilesResponseBody.
func unmarshalChallengeFilesResponseBodyToChallengeviewsChallengeFilesView(v *ChallengeFilesResponseBody) *challengeviews.ChallengeFilesView {
	if v == nil {
		return nil
	}
	res := &challengeviews.ChallengeFilesView{
		Filename: v.Filename,
		URL:      v.URL,
	}

	return res
}

// unmarshalAuthorResponseBodyToChallengeviewsAuthorView builds a value of type
// *challengeviews.AuthorView from a value of type *AuthorResponseBody.
func unmarshalAuthorResponseBodyToChallengeviewsAuthorView(v *AuthorResponseBody) *challengeviews.AuthorView {
	if v == nil {
		return nil
	}
	res := &challengeviews.AuthorView{
		ID:          v.ID,
		FullName:    v.FullName,
		Description: v.Description,
		Sponsor:     v.Sponsor,
		Slug:        v.Slug,
		ImageURL:    v.ImageURL,
		Publish:     v.Publish,
	}

	return res
}

// unmarshalSsmSolverResponseBodyToChallengeviewsSsmSolverView builds a value
// of type *challengeviews.SsmSolverView from a value of type
// *SsmSolverResponseBody.
func unmarshalSsmSolverResponseBodyToChallengeviewsSsmSolverView(v *SsmSolverResponseBody) *challengeviews.SsmSolverView {
	if v == nil {
		return nil
	}
	res := &challengeviews.SsmSolverView{
		ID:       v.ID,
		FullName: v.FullName,
		SolvedAt: v.SolvedAt,
	}

	return res
}

// unmarshalSsmUserMonthlyChallengeResponseToChallengeviewsSsmUserMonthlyChallengeView
// builds a value of type *challengeviews.SsmUserMonthlyChallengeView from a
// value of type *SsmUserMonthlyChallengeResponse.
func unmarshalSsmUserMonthlyChallengeResponseToChallengeviewsSsmUserMonthlyChallengeView(v *SsmUserMonthlyChallengeResponse) *challengeviews.SsmUserMonthlyChallengeView {
	res := &challengeviews.SsmUserMonthlyChallengeView{
		ChallengeID:  v.ChallengeID,
		DisplayMonth: v.DisplayMonth,
		StartDate:    v.StartDate,
		EndDate:      v.EndDate,
	}
	res.Challenge = unmarshalSsmChallengeResponseToChallengeviewsSsmChallengeView(v.Challenge)

	return res
}

// unmarshalSchoolScoreboardScoreResponseBodyToChallengeviewsSchoolScoreboardScoreView
// builds a value of type *challengeviews.SchoolScoreboardScoreView from a
// value of type *SchoolScoreboardScoreResponseBody.
func unmarshalSchoolScoreboardScoreResponseBodyToChallengeviewsSchoolScoreboardScoreView(v *SchoolScoreboardScoreResponseBody) *challengeviews.SchoolScoreboardScoreView {
	res := &challengeviews.SchoolScoreboardScoreView{
		Score:      v.Score,
		SchoolName: v.SchoolName,
	}

	return res
}

// unmarshalUserScoreboardScoreResponseBodyToChallengeviewsUserScoreboardScoreView
// builds a value of type *challengeviews.UserScoreboardScoreView from a value
// of type *UserScoreboardScoreResponseBody.
func unmarshalUserScoreboardScoreResponseBodyToChallengeviewsUserScoreboardScoreView(v *UserScoreboardScoreResponseBody) *challengeviews.UserScoreboardScoreView {
	res := &challengeviews.UserScoreboardScoreView{
		UserID:     v.UserID,
		Name:       v.Name,
		SchoolName: v.SchoolName,
		Score:      v.Score,
	}

	return res
}

// unmarshalAuthorResponseToChallengeAuthor builds a value of type
// *challenge.Author from a value of type *AuthorResponse.
func unmarshalAuthorResponseToChallengeAuthor(v *AuthorResponse) *challenge.Author {
	res := &challenge.Author{
		ID:          *v.ID,
		FullName:    *v.FullName,
		Description: *v.Description,
		Sponsor:     *v.Sponsor,
		Slug:        *v.Slug,
		ImageURL:    v.ImageURL,
		Publish:     *v.Publish,
	}

	return res
}
