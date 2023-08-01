// Code generated by goa v3.5.2, DO NOT EDIT.
//
// challenge HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package server

import (
	"context"
	"io"
	"net/http"
	"strings"

	challenge "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	challengeviews "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeListChallengesResponse returns an encoder for responses returned by
// the challenge ListChallenges endpoint.
func EncodeListChallengesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(challengeviews.SsmChallengeCollection)
		enc := encoder(ctx, w)
		body := NewSsmChallengeResponseCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListChallengesRequest returns a decoder for requests sent to the
// challenge ListChallenges endpoint.
func DecodeListChallengesRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			token *string
		)
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		payload := NewListChallengesPayload(token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeListMonthlyChallengesResponse returns an encoder for responses
// returned by the challenge ListMonthlyChallenges endpoint.
func EncodeListMonthlyChallengesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.([]*challenge.MonthlyChallenge)
		enc := encoder(ctx, w)
		body := NewListMonthlyChallengesResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListMonthlyChallengesRequest returns a decoder for requests sent to
// the challenge ListMonthlyChallenges endpoint.
func DecodeListMonthlyChallengesRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			token *string
		)
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		payload := NewListMonthlyChallengesPayload(token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeSubmitFlagResponse returns an encoder for responses returned by the
// challenge SubmitFlag endpoint.
func EncodeSubmitFlagResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeSubmitFlagRequest returns a decoder for requests sent to the challenge
// SubmitFlag endpoint.
func DecodeSubmitFlagRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body SubmitFlagRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateSubmitFlagRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			challengeID string
			token       string

			params = mux.Vars(r)
		)
		challengeID = params["challengeID"]
		err = goa.MergeErrors(err, goa.ValidateFormat("challengeID", challengeID, goa.FormatUUID))

		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewSubmitFlagPayload(&body, challengeID, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeSubmitFlagError returns an encoder for errors returned by the
// SubmitFlag challenge endpoint.
func EncodeSubmitFlagError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "already_solved":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewSubmitFlagAlreadySolvedResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusConflict)
			return enc.Encode(body)
		case "incorrect_flag":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewSubmitFlagIncorrectFlagResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalChallengeviewsSsmChallengeViewToSsmChallengeResponse builds a value
// of type *SsmChallengeResponse from a value of type
// *challengeviews.SsmChallengeView.
func marshalChallengeviewsSsmChallengeViewToSsmChallengeResponse(v *challengeviews.SsmChallengeView) *SsmChallengeResponse {
	res := &SsmChallengeResponse{
		ID:          *v.ID,
		Slug:        *v.Slug,
		Title:       *v.Title,
		Description: *v.Description,
		Score:       *v.Score,
		Solves:      *v.Solves,
		Solved:      *v.Solved,
		Category:    *v.Category,
	}
	if v.Services != nil {
		res.Services = make([]*ChallengeServiceResponse, len(v.Services))
		for i, val := range v.Services {
			res.Services[i] = marshalChallengeviewsChallengeServiceViewToChallengeServiceResponse(val)
		}
	}
	if v.Files != nil {
		res.Files = make([]*ChallengeFilesResponse, len(v.Files))
		for i, val := range v.Files {
			res.Files[i] = marshalChallengeviewsChallengeFilesViewToChallengeFilesResponse(val)
		}
	}

	return res
}

// marshalChallengeviewsChallengeServiceViewToChallengeServiceResponse builds a
// value of type *ChallengeServiceResponse from a value of type
// *challengeviews.ChallengeServiceView.
func marshalChallengeviewsChallengeServiceViewToChallengeServiceResponse(v *challengeviews.ChallengeServiceView) *ChallengeServiceResponse {
	if v == nil {
		return nil
	}
	res := &ChallengeServiceResponse{}

	return res
}

// marshalChallengeviewsChallengeFilesViewToChallengeFilesResponse builds a
// value of type *ChallengeFilesResponse from a value of type
// *challengeviews.ChallengeFilesView.
func marshalChallengeviewsChallengeFilesViewToChallengeFilesResponse(v *challengeviews.ChallengeFilesView) *ChallengeFilesResponse {
	if v == nil {
		return nil
	}
	res := &ChallengeFilesResponse{
		Filename: *v.Filename,
		URL:      *v.URL,
	}

	return res
}

// marshalChallengeMonthlyChallengeToMonthlyChallengeResponse builds a value of
// type *MonthlyChallengeResponse from a value of type
// *challenge.MonthlyChallenge.
func marshalChallengeMonthlyChallengeToMonthlyChallengeResponse(v *challenge.MonthlyChallenge) *MonthlyChallengeResponse {
	res := &MonthlyChallengeResponse{
		Slug:         v.Slug,
		Title:        v.Title,
		Description:  v.Description,
		ChallengeID:  v.ChallengeID,
		DisplayMonth: v.DisplayMonth,
		StartDate:    v.StartDate,
		EndDate:      v.EndDate,
	}

	return res
}
