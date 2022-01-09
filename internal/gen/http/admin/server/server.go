// Code generated by goa v3.5.2, DO NOT EDIT.
//
// admin HTTP server
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package server

import (
	"context"
	"net/http"

	admin "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the admin service endpoint HTTP handlers.
type Server struct {
	Mounts                 []*MountPoint
	ListChallenges         http.Handler
	CreateChallenge        http.Handler
	ListMonthlyChallenges  http.Handler
	DeleteMonthlyChallenge http.Handler
	CreateMonthlyChallenge http.Handler
	ListUsers              http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
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

// New instantiates HTTP handlers for all the admin service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *admin.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"ListChallenges", "GET", "/admin/challenges"},
			{"CreateChallenge", "POST", "/admin/challenges"},
			{"ListMonthlyChallenges", "GET", "/admin/monthly_challenges"},
			{"DeleteMonthlyChallenge", "DELETE", "/admin/monthly_challenges/{challengeID}"},
			{"CreateMonthlyChallenge", "POST", "/admin/monthly_challenges"},
			{"ListUsers", "GET", "/admin/users"},
		},
		ListChallenges:         NewListChallengesHandler(e.ListChallenges, mux, decoder, encoder, errhandler, formatter),
		CreateChallenge:        NewCreateChallengeHandler(e.CreateChallenge, mux, decoder, encoder, errhandler, formatter),
		ListMonthlyChallenges:  NewListMonthlyChallengesHandler(e.ListMonthlyChallenges, mux, decoder, encoder, errhandler, formatter),
		DeleteMonthlyChallenge: NewDeleteMonthlyChallengeHandler(e.DeleteMonthlyChallenge, mux, decoder, encoder, errhandler, formatter),
		CreateMonthlyChallenge: NewCreateMonthlyChallengeHandler(e.CreateMonthlyChallenge, mux, decoder, encoder, errhandler, formatter),
		ListUsers:              NewListUsersHandler(e.ListUsers, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "admin" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.ListChallenges = m(s.ListChallenges)
	s.CreateChallenge = m(s.CreateChallenge)
	s.ListMonthlyChallenges = m(s.ListMonthlyChallenges)
	s.DeleteMonthlyChallenge = m(s.DeleteMonthlyChallenge)
	s.CreateMonthlyChallenge = m(s.CreateMonthlyChallenge)
	s.ListUsers = m(s.ListUsers)
}

// Mount configures the mux to serve the admin endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountListChallengesHandler(mux, h.ListChallenges)
	MountCreateChallengeHandler(mux, h.CreateChallenge)
	MountListMonthlyChallengesHandler(mux, h.ListMonthlyChallenges)
	MountDeleteMonthlyChallengeHandler(mux, h.DeleteMonthlyChallenge)
	MountCreateMonthlyChallengeHandler(mux, h.CreateMonthlyChallenge)
	MountListUsersHandler(mux, h.ListUsers)
}

// MountListChallengesHandler configures the mux to serve the "admin" service
// "ListChallenges" endpoint.
func MountListChallengesHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/admin/challenges", f)
}

// NewListChallengesHandler creates a HTTP handler which loads the HTTP request
// and calls the "admin" service "ListChallenges" endpoint.
func NewListChallengesHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeListChallengesRequest(mux, decoder)
		encodeResponse = EncodeListChallengesResponse(encoder)
		encodeError    = EncodeListChallengesError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "ListChallenges")
		ctx = context.WithValue(ctx, goa.ServiceKey, "admin")
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

// MountCreateChallengeHandler configures the mux to serve the "admin" service
// "CreateChallenge" endpoint.
func MountCreateChallengeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/admin/challenges", f)
}

// NewCreateChallengeHandler creates a HTTP handler which loads the HTTP
// request and calls the "admin" service "CreateChallenge" endpoint.
func NewCreateChallengeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateChallengeRequest(mux, decoder)
		encodeResponse = EncodeCreateChallengeResponse(encoder)
		encodeError    = EncodeCreateChallengeError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "CreateChallenge")
		ctx = context.WithValue(ctx, goa.ServiceKey, "admin")
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

// MountListMonthlyChallengesHandler configures the mux to serve the "admin"
// service "ListMonthlyChallenges" endpoint.
func MountListMonthlyChallengesHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/admin/monthly_challenges", f)
}

// NewListMonthlyChallengesHandler creates a HTTP handler which loads the HTTP
// request and calls the "admin" service "ListMonthlyChallenges" endpoint.
func NewListMonthlyChallengesHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeListMonthlyChallengesRequest(mux, decoder)
		encodeResponse = EncodeListMonthlyChallengesResponse(encoder)
		encodeError    = EncodeListMonthlyChallengesError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "ListMonthlyChallenges")
		ctx = context.WithValue(ctx, goa.ServiceKey, "admin")
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

// MountDeleteMonthlyChallengeHandler configures the mux to serve the "admin"
// service "DeleteMonthlyChallenge" endpoint.
func MountDeleteMonthlyChallengeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/admin/monthly_challenges/{challengeID}", f)
}

// NewDeleteMonthlyChallengeHandler creates a HTTP handler which loads the HTTP
// request and calls the "admin" service "DeleteMonthlyChallenge" endpoint.
func NewDeleteMonthlyChallengeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteMonthlyChallengeRequest(mux, decoder)
		encodeResponse = EncodeDeleteMonthlyChallengeResponse(encoder)
		encodeError    = EncodeDeleteMonthlyChallengeError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "DeleteMonthlyChallenge")
		ctx = context.WithValue(ctx, goa.ServiceKey, "admin")
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

// MountCreateMonthlyChallengeHandler configures the mux to serve the "admin"
// service "CreateMonthlyChallenge" endpoint.
func MountCreateMonthlyChallengeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/admin/monthly_challenges", f)
}

// NewCreateMonthlyChallengeHandler creates a HTTP handler which loads the HTTP
// request and calls the "admin" service "CreateMonthlyChallenge" endpoint.
func NewCreateMonthlyChallengeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateMonthlyChallengeRequest(mux, decoder)
		encodeResponse = EncodeCreateMonthlyChallengeResponse(encoder)
		encodeError    = EncodeCreateMonthlyChallengeError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "CreateMonthlyChallenge")
		ctx = context.WithValue(ctx, goa.ServiceKey, "admin")
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

// MountListUsersHandler configures the mux to serve the "admin" service
// "ListUsers" endpoint.
func MountListUsersHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/admin/users", f)
}

// NewListUsersHandler creates a HTTP handler which loads the HTTP request and
// calls the "admin" service "ListUsers" endpoint.
func NewListUsersHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeListUsersRequest(mux, decoder)
		encodeResponse = EncodeListUsersResponse(encoder)
		encodeError    = EncodeListUsersError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "ListUsers")
		ctx = context.WithValue(ctx, goa.ServiceKey, "admin")
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
