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

// EncodeListEventsResponse returns an encoder for responses returned by the
// challenge ListEvents endpoint.
func EncodeListEventsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.([]*challenge.CTFEvent)
		enc := encoder(ctx, w)
		body := NewListEventsResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListEventsRequest returns a decoder for requests sent to the challenge
// ListEvents endpoint.
func DecodeListEventsRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			token *string
		)
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		payload := NewListEventsPayload(token)
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

// EncodeGetCurrentMonthlyChallengeResponse returns an encoder for responses
// returned by the challenge GetCurrentMonthlyChallenge endpoint.
func EncodeGetCurrentMonthlyChallengeResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*challengeviews.SsmUserMonthlyChallenge)
		enc := encoder(ctx, w)
		body := NewGetCurrentMonthlyChallengeResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetCurrentMonthlyChallengeRequest returns a decoder for requests sent
// to the challenge GetCurrentMonthlyChallenge endpoint.
func DecodeGetCurrentMonthlyChallengeRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			token *string
		)
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		payload := NewGetCurrentMonthlyChallengePayload(token)
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
		res := v.(challengeviews.SsmUserMonthlyChallengeCollection)
		enc := encoder(ctx, w)
		body := NewSsmUserMonthlyChallengeResponseCollection(res.Projected)
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

// EncodeSchoolScoreboardResponse returns an encoder for responses returned by
// the challenge SchoolScoreboard endpoint.
func EncodeSchoolScoreboardResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*challengeviews.SsmShoolscoreboard)
		enc := encoder(ctx, w)
		body := NewSchoolScoreboardResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeSchoolScoreboardRequest returns a decoder for requests sent to the
// challenge SchoolScoreboard endpoint.
func DecodeSchoolScoreboardRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			token *string
		)
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		payload := NewSchoolScoreboardPayload(token)
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
		CtfEventID:  v.CtfEventID,
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
	if v.OtherAuthors != nil {
		res.OtherAuthors = make([]string, len(v.OtherAuthors))
		for i, val := range v.OtherAuthors {
			res.OtherAuthors[i] = val
		}
	}
	if v.Authors != nil {
		res.Authors = make([]*SsmUserResponse, len(v.Authors))
		for i, val := range v.Authors {
			res.Authors[i] = marshalChallengeviewsSsmUserViewToSsmUserResponse(val)
		}
	}
	if v.Solvers != nil {
		res.Solvers = make([]*SsmSolverResponse, len(v.Solvers))
		for i, val := range v.Solvers {
			res.Solvers[i] = marshalChallengeviewsSsmSolverViewToSsmSolverResponse(val)
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
	res := &ChallengeServiceResponse{
		UserDisplay: *v.UserDisplay,
		Hyperlink:   *v.Hyperlink,
	}

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

// marshalChallengeviewsSsmUserViewToSsmUserResponse builds a value of type
// *SsmUserResponse from a value of type *challengeviews.SsmUserView.
func marshalChallengeviewsSsmUserViewToSsmUserResponse(v *challengeviews.SsmUserView) *SsmUserResponse {
	if v == nil {
		return nil
	}
	res := &SsmUserResponse{
		ID:       *v.ID,
		Email:    *v.Email,
		FullName: *v.FullName,
		Role:     *v.Role,
		SchoolID: v.SchoolID,
	}

	return res
}

// marshalChallengeviewsSsmSolverViewToSsmSolverResponse builds a value of type
// *SsmSolverResponse from a value of type *challengeviews.SsmSolverView.
func marshalChallengeviewsSsmSolverViewToSsmSolverResponse(v *challengeviews.SsmSolverView) *SsmSolverResponse {
	if v == nil {
		return nil
	}
	res := &SsmSolverResponse{
		ID:       *v.ID,
		FullName: *v.FullName,
		SolvedAt: *v.SolvedAt,
	}

	return res
}

// marshalChallengeCTFEventToCTFEventResponse builds a value of type
// *CTFEventResponse from a value of type *challenge.CTFEvent.
func marshalChallengeCTFEventToCTFEventResponse(v *challenge.CTFEvent) *CTFEventResponse {
	res := &CTFEventResponse{
		ID:   v.ID,
		Name: v.Name,
	}

	return res
}

// marshalChallengeviewsSsmChallengeViewToSsmChallengeResponseBody builds a
// value of type *SsmChallengeResponseBody from a value of type
// *challengeviews.SsmChallengeView.
func marshalChallengeviewsSsmChallengeViewToSsmChallengeResponseBody(v *challengeviews.SsmChallengeView) *SsmChallengeResponseBody {
	res := &SsmChallengeResponseBody{
		ID:          *v.ID,
		Slug:        *v.Slug,
		Title:       *v.Title,
		Description: *v.Description,
		Score:       *v.Score,
		Solves:      *v.Solves,
		CtfEventID:  v.CtfEventID,
		Solved:      *v.Solved,
		Category:    *v.Category,
	}
	if v.Services != nil {
		res.Services = make([]*ChallengeServiceResponseBody, len(v.Services))
		for i, val := range v.Services {
			res.Services[i] = marshalChallengeviewsChallengeServiceViewToChallengeServiceResponseBody(val)
		}
	}
	if v.Files != nil {
		res.Files = make([]*ChallengeFilesResponseBody, len(v.Files))
		for i, val := range v.Files {
			res.Files[i] = marshalChallengeviewsChallengeFilesViewToChallengeFilesResponseBody(val)
		}
	}
	if v.OtherAuthors != nil {
		res.OtherAuthors = make([]string, len(v.OtherAuthors))
		for i, val := range v.OtherAuthors {
			res.OtherAuthors[i] = val
		}
	}
	if v.Authors != nil {
		res.Authors = make([]*SsmUserResponseBody, len(v.Authors))
		for i, val := range v.Authors {
			res.Authors[i] = marshalChallengeviewsSsmUserViewToSsmUserResponseBody(val)
		}
	}
	if v.Solvers != nil {
		res.Solvers = make([]*SsmSolverResponseBody, len(v.Solvers))
		for i, val := range v.Solvers {
			res.Solvers[i] = marshalChallengeviewsSsmSolverViewToSsmSolverResponseBody(val)
		}
	}

	return res
}

// marshalChallengeviewsChallengeServiceViewToChallengeServiceResponseBody
// builds a value of type *ChallengeServiceResponseBody from a value of type
// *challengeviews.ChallengeServiceView.
func marshalChallengeviewsChallengeServiceViewToChallengeServiceResponseBody(v *challengeviews.ChallengeServiceView) *ChallengeServiceResponseBody {
	if v == nil {
		return nil
	}
	res := &ChallengeServiceResponseBody{
		UserDisplay: *v.UserDisplay,
		Hyperlink:   *v.Hyperlink,
	}

	return res
}

// marshalChallengeviewsChallengeFilesViewToChallengeFilesResponseBody builds a
// value of type *ChallengeFilesResponseBody from a value of type
// *challengeviews.ChallengeFilesView.
func marshalChallengeviewsChallengeFilesViewToChallengeFilesResponseBody(v *challengeviews.ChallengeFilesView) *ChallengeFilesResponseBody {
	if v == nil {
		return nil
	}
	res := &ChallengeFilesResponseBody{
		Filename: *v.Filename,
		URL:      *v.URL,
	}

	return res
}

// marshalChallengeviewsSsmUserViewToSsmUserResponseBody builds a value of type
// *SsmUserResponseBody from a value of type *challengeviews.SsmUserView.
func marshalChallengeviewsSsmUserViewToSsmUserResponseBody(v *challengeviews.SsmUserView) *SsmUserResponseBody {
	if v == nil {
		return nil
	}
	res := &SsmUserResponseBody{
		ID:       *v.ID,
		Email:    *v.Email,
		FullName: *v.FullName,
		Role:     *v.Role,
		SchoolID: v.SchoolID,
	}

	return res
}

// marshalChallengeviewsSsmSolverViewToSsmSolverResponseBody builds a value of
// type *SsmSolverResponseBody from a value of type
// *challengeviews.SsmSolverView.
func marshalChallengeviewsSsmSolverViewToSsmSolverResponseBody(v *challengeviews.SsmSolverView) *SsmSolverResponseBody {
	if v == nil {
		return nil
	}
	res := &SsmSolverResponseBody{
		ID:       *v.ID,
		FullName: *v.FullName,
		SolvedAt: *v.SolvedAt,
	}

	return res
}

// marshalChallengeviewsSsmUserMonthlyChallengeViewToSsmUserMonthlyChallengeResponse
// builds a value of type *SsmUserMonthlyChallengeResponse from a value of type
// *challengeviews.SsmUserMonthlyChallengeView.
func marshalChallengeviewsSsmUserMonthlyChallengeViewToSsmUserMonthlyChallengeResponse(v *challengeviews.SsmUserMonthlyChallengeView) *SsmUserMonthlyChallengeResponse {
	res := &SsmUserMonthlyChallengeResponse{
		ChallengeID:  *v.ChallengeID,
		DisplayMonth: *v.DisplayMonth,
		StartDate:    *v.StartDate,
		EndDate:      *v.EndDate,
	}
	if v.Challenge != nil {
		res.Challenge = marshalChallengeviewsSsmChallengeViewToSsmChallengeResponse(v.Challenge)
	}

	return res
}

// marshalChallengeviewsSchoolScoreboardScoreViewToSchoolScoreboardScoreResponseBody
// builds a value of type *SchoolScoreboardScoreResponseBody from a value of
// type *challengeviews.SchoolScoreboardScoreView.
func marshalChallengeviewsSchoolScoreboardScoreViewToSchoolScoreboardScoreResponseBody(v *challengeviews.SchoolScoreboardScoreView) *SchoolScoreboardScoreResponseBody {
	res := &SchoolScoreboardScoreResponseBody{
		Score:      *v.Score,
		SchoolName: *v.SchoolName,
	}

	return res
}
