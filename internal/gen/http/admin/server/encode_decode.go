// Code generated by goa v3.5.2, DO NOT EDIT.
//
// admin HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package server

import (
	"context"
	"io"
	"net/http"
	"strings"

	admin "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
	adminviews "github.com/sakerhetsm/ssm-wargame/internal/gen/admin/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeListChallengesResponse returns an encoder for responses returned by
// the admin ListChallenges endpoint.
func EncodeListChallengesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(adminviews.SsmChallengeCollection)
		enc := encoder(ctx, w)
		body := NewSsmChallengeResponseCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListChallengesRequest returns a decoder for requests sent to the admin
// ListChallenges endpoint.
func DecodeListChallengesRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			token string
			err   error
		)
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewListChallengesPayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeListChallengesError returns an encoder for errors returned by the
// ListChallenges admin endpoint.
func EncodeListChallengesError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListChallengesUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListChallengesNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateChallengeResponse returns an encoder for responses returned by
// the admin CreateChallenge endpoint.
func EncodeCreateChallengeResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusCreated)
		return nil
	}
}

// DecodeCreateChallengeRequest returns a decoder for requests sent to the
// admin CreateChallenge endpoint.
func DecodeCreateChallengeRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateChallengeRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateChallengeRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			token string
		)
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewCreateChallengePayload(&body, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeCreateChallengeError returns an encoder for errors returned by the
// CreateChallenge admin endpoint.
func EncodeCreateChallengeError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewCreateChallengeUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewCreateChallengeNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeListMonthlyChallengesResponse returns an encoder for responses
// returned by the admin ListMonthlyChallenges endpoint.
func EncodeListMonthlyChallengesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.([]*admin.MonthlyChallengeMeta)
		enc := encoder(ctx, w)
		body := NewListMonthlyChallengesResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListMonthlyChallengesRequest returns a decoder for requests sent to
// the admin ListMonthlyChallenges endpoint.
func DecodeListMonthlyChallengesRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			token string
			err   error
		)
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewListMonthlyChallengesPayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeListMonthlyChallengesError returns an encoder for errors returned by
// the ListMonthlyChallenges admin endpoint.
func EncodeListMonthlyChallengesError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListMonthlyChallengesUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListMonthlyChallengesNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteMonthlyChallengeResponse returns an encoder for responses
// returned by the admin DeleteMonthlyChallenge endpoint.
func EncodeDeleteMonthlyChallengeResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeDeleteMonthlyChallengeRequest returns a decoder for requests sent to
// the admin DeleteMonthlyChallenge endpoint.
func DecodeDeleteMonthlyChallengeRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			monthlyChallengeID string
			token              string
			err                error

			params = mux.Vars(r)
		)
		monthlyChallengeID = params["monthlyChallengeID"]
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeleteMonthlyChallengePayload(monthlyChallengeID, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeDeleteMonthlyChallengeError returns an encoder for errors returned by
// the DeleteMonthlyChallenge admin endpoint.
func EncodeDeleteMonthlyChallengeError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeleteMonthlyChallengeUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeleteMonthlyChallengeNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateMonthlyChallengeResponse returns an encoder for responses
// returned by the admin CreateMonthlyChallenge endpoint.
func EncodeCreateMonthlyChallengeResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeCreateMonthlyChallengeRequest returns a decoder for requests sent to
// the admin CreateMonthlyChallenge endpoint.
func DecodeCreateMonthlyChallengeRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateMonthlyChallengeRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateMonthlyChallengeRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			token string
		)
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewCreateMonthlyChallengePayload(&body, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeCreateMonthlyChallengeError returns an encoder for errors returned by
// the CreateMonthlyChallenge admin endpoint.
func EncodeCreateMonthlyChallengeError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewCreateMonthlyChallengeUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewCreateMonthlyChallengeNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalAdminviewsSsmChallengeViewToSsmChallengeResponse builds a value of
// type *SsmChallengeResponse from a value of type *adminviews.SsmChallengeView.
func marshalAdminviewsSsmChallengeViewToSsmChallengeResponse(v *adminviews.SsmChallengeView) *SsmChallengeResponse {
	res := &SsmChallengeResponse{
		ID:          *v.ID,
		Slug:        *v.Slug,
		Title:       *v.Title,
		Description: *v.Description,
		Score:       *v.Score,
		Published:   *v.Published,
		Solves:      *v.Solves,
	}
	if v.Services != nil {
		res.Services = make([]*ChallengeServiceResponse, len(v.Services))
		for i, val := range v.Services {
			res.Services[i] = marshalAdminviewsChallengeServiceViewToChallengeServiceResponse(val)
		}
	}
	if v.Files != nil {
		res.Files = make([]*ChallengeFilesResponse, len(v.Files))
		for i, val := range v.Files {
			res.Files[i] = marshalAdminviewsChallengeFilesViewToChallengeFilesResponse(val)
		}
	}

	return res
}

// marshalAdminviewsChallengeServiceViewToChallengeServiceResponse builds a
// value of type *ChallengeServiceResponse from a value of type
// *adminviews.ChallengeServiceView.
func marshalAdminviewsChallengeServiceViewToChallengeServiceResponse(v *adminviews.ChallengeServiceView) *ChallengeServiceResponse {
	if v == nil {
		return nil
	}
	res := &ChallengeServiceResponse{}

	return res
}

// marshalAdminviewsChallengeFilesViewToChallengeFilesResponse builds a value
// of type *ChallengeFilesResponse from a value of type
// *adminviews.ChallengeFilesView.
func marshalAdminviewsChallengeFilesViewToChallengeFilesResponse(v *adminviews.ChallengeFilesView) *ChallengeFilesResponse {
	if v == nil {
		return nil
	}
	res := &ChallengeFilesResponse{}

	return res
}

// marshalAdminMonthlyChallengeMetaToMonthlyChallengeMetaResponse builds a
// value of type *MonthlyChallengeMetaResponse from a value of type
// *admin.MonthlyChallengeMeta.
func marshalAdminMonthlyChallengeMetaToMonthlyChallengeMetaResponse(v *admin.MonthlyChallengeMeta) *MonthlyChallengeMetaResponse {
	res := &MonthlyChallengeMetaResponse{
		DisplayMonth: v.DisplayMonth,
		StartDate:    v.StartDate,
		EndDate:      v.EndDate,
	}

	return res
}
