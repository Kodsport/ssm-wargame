// Code generated by goa v3.12.3, DO NOT EDIT.
//
// admin HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package server

import (
	"context"
	"errors"
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
func EncodeListChallengesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res := v.(adminviews.SsmAdminChallengeCollection)
		enc := encoder(ctx, w)
		body := NewSsmAdminChallengeResponseCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListChallengesRequest returns a decoder for requests sent to the admin
// ListChallenges endpoint.
func DecodeListChallengesRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			token string
			err   error
		)
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("token", "header"))
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
func EncodeListChallengesError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "unauthorized":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewListChallengesUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not_found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewListChallengesNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "bad_request":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewListChallengesBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateChallengeResponse returns an encoder for responses returned by
// the admin CreateChallenge endpoint.
func EncodeCreateChallengeResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusCreated)
		return nil
	}
}

// DecodeCreateChallengeRequest returns a decoder for requests sent to the
// admin CreateChallenge endpoint.
func DecodeCreateChallengeRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
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
			err = goa.MergeErrors(err, goa.MissingFieldError("token", "header"))
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
func EncodeCreateChallengeError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "unauthorized":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewCreateChallengeUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not_found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewCreateChallengeNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "bad_request":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewCreateChallengeBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUpdateChallengeResponse returns an encoder for responses returned by
// the admin UpdateChallenge endpoint.
func EncodeUpdateChallengeResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusCreated)
		return nil
	}
}

// DecodeUpdateChallengeRequest returns a decoder for requests sent to the
// admin UpdateChallenge endpoint.
func DecodeUpdateChallengeRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body UpdateChallengeRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateChallengeRequestBody(&body)
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
			err = goa.MergeErrors(err, goa.MissingFieldError("token", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewUpdateChallengePayload(&body, challengeID, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeUpdateChallengeError returns an encoder for errors returned by the
// UpdateChallenge admin endpoint.
func EncodeUpdateChallengeError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "unauthorized":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUpdateChallengeUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not_found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUpdateChallengeNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "bad_request":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUpdateChallengeBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodePresignChallFileUploadResponse returns an encoder for responses
// returned by the admin PresignChallFileUpload endpoint.
func EncodePresignChallFileUploadResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*admin.PresignChallFileUploadResult)
		enc := encoder(ctx, w)
		body := NewPresignChallFileUploadResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodePresignChallFileUploadRequest returns a decoder for requests sent to
// the admin PresignChallFileUpload endpoint.
func DecodePresignChallFileUploadRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body PresignChallFileUploadRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidatePresignChallFileUploadRequestBody(&body)
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
			err = goa.MergeErrors(err, goa.MissingFieldError("token", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewPresignChallFileUploadPayload(&body, challengeID, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodePresignChallFileUploadError returns an encoder for errors returned by
// the PresignChallFileUpload admin endpoint.
func EncodePresignChallFileUploadError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "unauthorized":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewPresignChallFileUploadUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not_found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewPresignChallFileUploadNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "bad_request":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewPresignChallFileUploadBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeListMonthlyChallengesResponse returns an encoder for responses
// returned by the admin ListMonthlyChallenges endpoint.
func EncodeListMonthlyChallengesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.([]*admin.MonthlyChallenge)
		enc := encoder(ctx, w)
		body := NewListMonthlyChallengesResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListMonthlyChallengesRequest returns a decoder for requests sent to
// the admin ListMonthlyChallenges endpoint.
func DecodeListMonthlyChallengesRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			token string
			err   error
		)
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("token", "header"))
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
func EncodeListMonthlyChallengesError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "unauthorized":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewListMonthlyChallengesUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not_found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewListMonthlyChallengesNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "bad_request":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewListMonthlyChallengesBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteMonthlyChallengeResponse returns an encoder for responses
// returned by the admin DeleteMonthlyChallenge endpoint.
func EncodeDeleteMonthlyChallengeResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeDeleteMonthlyChallengeRequest returns a decoder for requests sent to
// the admin DeleteMonthlyChallenge endpoint.
func DecodeDeleteMonthlyChallengeRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			challengeID string
			token       string
			err         error

			params = mux.Vars(r)
		)
		challengeID = params["challengeID"]
		err = goa.MergeErrors(err, goa.ValidateFormat("challengeID", challengeID, goa.FormatUUID))
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("token", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeleteMonthlyChallengePayload(challengeID, token)
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
func EncodeDeleteMonthlyChallengeError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "unauthorized":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteMonthlyChallengeUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not_found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteMonthlyChallengeNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "bad_request":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteMonthlyChallengeBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteFileResponse returns an encoder for responses returned by the
// admin DeleteFile endpoint.
func EncodeDeleteFileResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeDeleteFileRequest returns a decoder for requests sent to the admin
// DeleteFile endpoint.
func DecodeDeleteFileRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			fileID string
			token  string
			err    error

			params = mux.Vars(r)
		)
		fileID = params["fileID"]
		err = goa.MergeErrors(err, goa.ValidateFormat("fileID", fileID, goa.FormatUUID))
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("token", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeleteFilePayload(fileID, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeDeleteFileError returns an encoder for errors returned by the
// DeleteFile admin endpoint.
func EncodeDeleteFileError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "unauthorized":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteFileUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not_found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteFileNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "bad_request":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteFileBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateMonthlyChallengeResponse returns an encoder for responses
// returned by the admin CreateMonthlyChallenge endpoint.
func EncodeCreateMonthlyChallengeResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeCreateMonthlyChallengeRequest returns a decoder for requests sent to
// the admin CreateMonthlyChallenge endpoint.
func DecodeCreateMonthlyChallengeRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
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
			err = goa.MergeErrors(err, goa.MissingFieldError("token", "header"))
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
func EncodeCreateMonthlyChallengeError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "unauthorized":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewCreateMonthlyChallengeUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not_found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewCreateMonthlyChallengeNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "bad_request":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewCreateMonthlyChallengeBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeListUsersResponse returns an encoder for responses returned by the
// admin ListUsers endpoint.
func EncodeListUsersResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.([]*admin.SsmUser)
		enc := encoder(ctx, w)
		body := NewListUsersResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListUsersRequest returns a decoder for requests sent to the admin
// ListUsers endpoint.
func DecodeListUsersRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			token string
			err   error
		)
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("token", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewListUsersPayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeListUsersError returns an encoder for errors returned by the ListUsers
// admin endpoint.
func EncodeListUsersError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "unauthorized":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewListUsersUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not_found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewListUsersNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "bad_request":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewListUsersBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeAddFlagResponse returns an encoder for responses returned by the admin
// AddFlag endpoint.
func EncodeAddFlagResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeAddFlagRequest returns a decoder for requests sent to the admin
// AddFlag endpoint.
func DecodeAddFlagRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body AddFlagRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateAddFlagRequestBody(&body)
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
			err = goa.MergeErrors(err, goa.MissingFieldError("token", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewAddFlagPayload(&body, challengeID, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeAddFlagError returns an encoder for errors returned by the AddFlag
// admin endpoint.
func EncodeAddFlagError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "unauthorized":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewAddFlagUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not_found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewAddFlagNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "bad_request":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewAddFlagBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteFlagResponse returns an encoder for responses returned by the
// admin DeleteFlag endpoint.
func EncodeDeleteFlagResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeDeleteFlagRequest returns a decoder for requests sent to the admin
// DeleteFlag endpoint.
func DecodeDeleteFlagRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			challengeID string
			flagID      string
			token       string
			err         error

			params = mux.Vars(r)
		)
		challengeID = params["challengeID"]
		err = goa.MergeErrors(err, goa.ValidateFormat("challengeID", challengeID, goa.FormatUUID))
		flagID = params["flagID"]
		err = goa.MergeErrors(err, goa.ValidateFormat("flagID", flagID, goa.FormatUUID))
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("token", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeleteFlagPayload(challengeID, flagID, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeDeleteFlagError returns an encoder for errors returned by the
// DeleteFlag admin endpoint.
func EncodeDeleteFlagError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "unauthorized":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteFlagUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not_found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteFlagNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "bad_request":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteFlagBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalAdminviewsSsmAdminChallengeViewToSsmAdminChallengeResponse builds a
// value of type *SsmAdminChallengeResponse from a value of type
// *adminviews.SsmAdminChallengeView.
func marshalAdminviewsSsmAdminChallengeViewToSsmAdminChallengeResponse(v *adminviews.SsmAdminChallengeView) *SsmAdminChallengeResponse {
	res := &SsmAdminChallengeResponse{
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
		res.Files = make([]*AdminChallengeFilesResponse, len(v.Files))
		for i, val := range v.Files {
			res.Files[i] = marshalAdminviewsAdminChallengeFilesViewToAdminChallengeFilesResponse(val)
		}
	} else {
		res.Files = []*AdminChallengeFilesResponse{}
	}
	if v.Flags != nil {
		res.Flags = make([]*AdminChallengeFlagResponse, len(v.Flags))
		for i, val := range v.Flags {
			res.Flags[i] = marshalAdminviewsAdminChallengeFlagViewToAdminChallengeFlagResponse(val)
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

// marshalAdminviewsAdminChallengeFilesViewToAdminChallengeFilesResponse builds
// a value of type *AdminChallengeFilesResponse from a value of type
// *adminviews.AdminChallengeFilesView.
func marshalAdminviewsAdminChallengeFilesViewToAdminChallengeFilesResponse(v *adminviews.AdminChallengeFilesView) *AdminChallengeFilesResponse {
	res := &AdminChallengeFilesResponse{
		ID:       *v.ID,
		Filename: *v.Filename,
		URL:      *v.URL,
		Bucket:   *v.Bucket,
		Key:      *v.Key,
		Size:     *v.Size,
		Md5:      *v.Md5,
	}

	return res
}

// marshalAdminviewsAdminChallengeFlagViewToAdminChallengeFlagResponse builds a
// value of type *AdminChallengeFlagResponse from a value of type
// *adminviews.AdminChallengeFlagView.
func marshalAdminviewsAdminChallengeFlagViewToAdminChallengeFlagResponse(v *adminviews.AdminChallengeFlagView) *AdminChallengeFlagResponse {
	if v == nil {
		return nil
	}
	res := &AdminChallengeFlagResponse{
		ID:   *v.ID,
		Flag: *v.Flag,
	}

	return res
}

// marshalAdminMonthlyChallengeToMonthlyChallengeResponse builds a value of
// type *MonthlyChallengeResponse from a value of type *admin.MonthlyChallenge.
func marshalAdminMonthlyChallengeToMonthlyChallengeResponse(v *admin.MonthlyChallenge) *MonthlyChallengeResponse {
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

// marshalAdminSsmUserToSsmUserResponse builds a value of type *SsmUserResponse
// from a value of type *admin.SsmUser.
func marshalAdminSsmUserToSsmUserResponse(v *admin.SsmUser) *SsmUserResponse {
	res := &SsmUserResponse{
		ID:        v.ID,
		Email:     v.Email,
		FirstName: v.FirstName,
		LastName:  v.LastName,
		Role:      v.Role,
	}

	return res
}
