// Code generated by goa v3.5.2, DO NOT EDIT.
//
// user HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package server

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"

	user "github.com/sakerhetsm/ssm-wargame/internal/gen/user"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeGetSelfResponse returns an encoder for responses returned by the user
// GetSelf endpoint.
func EncodeGetSelfResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*user.GetSelfResult)
		enc := encoder(ctx, w)
		body := NewGetSelfResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetSelfRequest returns a decoder for requests sent to the user GetSelf
// endpoint.
func DecodeGetSelfRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
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
		payload := NewGetSelfPayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeUpdateSelfResponse returns an encoder for responses returned by the
// user UpdateSelf endpoint.
func EncodeUpdateSelfResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeUpdateSelfRequest returns a decoder for requests sent to the user
// UpdateSelf endpoint.
func DecodeUpdateSelfRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body UpdateSelfRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateSelfRequestBody(&body)
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
		payload := NewUpdateSelfPayload(&body, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeCompleteOnboardingResponse returns an encoder for responses returned
// by the user CompleteOnboarding endpoint.
func EncodeCompleteOnboardingResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeCompleteOnboardingRequest returns a decoder for requests sent to the
// user CompleteOnboarding endpoint.
func DecodeCompleteOnboardingRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
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
		payload := NewCompleteOnboardingPayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeJoinSchoolResponse returns an encoder for responses returned by the
// user JoinSchool endpoint.
func EncodeJoinSchoolResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeJoinSchoolRequest returns a decoder for requests sent to the user
// JoinSchool endpoint.
func DecodeJoinSchoolRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body JoinSchoolRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateJoinSchoolRequestBody(&body)
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
		payload := NewJoinSchoolPayload(&body, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeLeaveSchoolResponse returns an encoder for responses returned by the
// user LeaveSchool endpoint.
func EncodeLeaveSchoolResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeLeaveSchoolRequest returns a decoder for requests sent to the user
// LeaveSchool endpoint.
func DecodeLeaveSchoolRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
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
		payload := NewLeaveSchoolPayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeSearchSchoolsResponse returns an encoder for responses returned by the
// user SearchSchools endpoint.
func EncodeSearchSchoolsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.([]*user.School)
		enc := encoder(ctx, w)
		body := NewSearchSchoolsResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeSearchSchoolsRequest returns a decoder for requests sent to the user
// SearchSchools endpoint.
func DecodeSearchSchoolsRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q          string
			university *bool
			token      string
			err        error
		)
		q = r.URL.Query().Get("q")
		if q == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
		}
		{
			universityRaw := r.URL.Query().Get("university")
			if universityRaw != "" {
				v, err2 := strconv.ParseBool(universityRaw)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("university", universityRaw, "boolean"))
				}
				university = &v
			}
		}
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewSearchSchoolsPayload(q, university, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// marshalUserSchoolToSchoolResponse builds a value of type *SchoolResponse
// from a value of type *user.School.
func marshalUserSchoolToSchoolResponse(v *user.School) *SchoolResponse {
	res := &SchoolResponse{
		Name:             v.Name,
		MunicipalityName: v.MunicipalityName,
		IsUniversity:     v.IsUniversity,
		ID:               v.ID,
	}

	return res
}
