// Code generated by goa v3.12.3, DO NOT EDIT.
//
// challenge HTTP server
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package server

import (
	"context"
	"net/http"

	challenge "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the challenge service endpoint HTTP handlers.
type Server struct {
	Mounts                []*MountPoint
	ListChallenges        http.Handler
	ListMonthlyChallenges http.Handler
	SubmitFlag            http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the challenge service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *challenge.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"ListChallenges", "GET", "/challenges"},
			{"ListMonthlyChallenges", "GET", "/monthly_challenges"},
			{"SubmitFlag", "POST", "/challenges/{challengeID}/attempt"},
		},
		ListChallenges:        NewListChallengesHandler(e.ListChallenges, mux, decoder, encoder, errhandler, formatter),
		ListMonthlyChallenges: NewListMonthlyChallengesHandler(e.ListMonthlyChallenges, mux, decoder, encoder, errhandler, formatter),
		SubmitFlag:            NewSubmitFlagHandler(e.SubmitFlag, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "challenge" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.ListChallenges = m(s.ListChallenges)
	s.ListMonthlyChallenges = m(s.ListMonthlyChallenges)
	s.SubmitFlag = m(s.SubmitFlag)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return challenge.MethodNames[:] }

// Mount configures the mux to serve the challenge endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountListChallengesHandler(mux, h.ListChallenges)
	MountListMonthlyChallengesHandler(mux, h.ListMonthlyChallenges)
	MountSubmitFlagHandler(mux, h.SubmitFlag)
}

// Mount configures the mux to serve the challenge endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountListChallengesHandler configures the mux to serve the "challenge"
// service "ListChallenges" endpoint.
func MountListChallengesHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/challenges", f)
}

// NewListChallengesHandler creates a HTTP handler which loads the HTTP request
// and calls the "challenge" service "ListChallenges" endpoint.
func NewListChallengesHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeListChallengesRequest(mux, decoder)
		encodeResponse = EncodeListChallengesResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "ListChallenges")
		ctx = context.WithValue(ctx, goa.ServiceKey, "challenge")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountListMonthlyChallengesHandler configures the mux to serve the
// "challenge" service "ListMonthlyChallenges" endpoint.
func MountListMonthlyChallengesHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/monthly_challenges", f)
}

// NewListMonthlyChallengesHandler creates a HTTP handler which loads the HTTP
// request and calls the "challenge" service "ListMonthlyChallenges" endpoint.
func NewListMonthlyChallengesHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeListMonthlyChallengesRequest(mux, decoder)
		encodeResponse = EncodeListMonthlyChallengesResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "ListMonthlyChallenges")
		ctx = context.WithValue(ctx, goa.ServiceKey, "challenge")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountSubmitFlagHandler configures the mux to serve the "challenge" service
// "SubmitFlag" endpoint.
func MountSubmitFlagHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/challenges/{challengeID}/attempt", f)
}

// NewSubmitFlagHandler creates a HTTP handler which loads the HTTP request and
// calls the "challenge" service "SubmitFlag" endpoint.
func NewSubmitFlagHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeSubmitFlagRequest(mux, decoder)
		encodeResponse = EncodeSubmitFlagResponse(encoder)
		encodeError    = EncodeSubmitFlagError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "SubmitFlag")
		ctx = context.WithValue(ctx, goa.ServiceKey, "challenge")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}
