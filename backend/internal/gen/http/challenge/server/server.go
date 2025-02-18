// Code generated by goa v3.5.2, DO NOT EDIT.
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
	Mounts                     []*MountPoint
	ListChallenges             http.Handler
	ListEvents                 http.Handler
	GetCurrentMonthlyChallenge http.Handler
	ListMonthlyChallenges      http.Handler
	SubmitFlag                 http.Handler
	SchoolScoreboard           http.Handler
	UserScoreboard             http.Handler
	ListAuthors                http.Handler
	ListCourses                http.Handler
	EnrollCourse               http.Handler
	CompleteCourse             http.Handler
	KnackKodenSubmitFlag       http.Handler
	KnackKodenScoreboard       http.Handler
	KnackKodenRegisterClass    http.Handler
	KnackKodenGetClass         http.Handler
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
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"ListChallenges", "GET", "/challenges"},
			{"ListEvents", "GET", "/events"},
			{"GetCurrentMonthlyChallenge", "GET", "/current_monthly_challenge"},
			{"ListMonthlyChallenges", "GET", "/monthly_challenges"},
			{"SubmitFlag", "POST", "/challenges/{challenge_id}/attempt"},
			{"SchoolScoreboard", "GET", "/scoreboard"},
			{"UserScoreboard", "GET", "/user_scoreboard"},
			{"ListAuthors", "GET", "/authors"},
			{"ListCourses", "GET", "/courses"},
			{"EnrollCourse", "POST", "/courses/{id}/enroll"},
			{"CompleteCourse", "POST", "/courses/{id}/complete"},
			{"KnackKodenSubmitFlag", "POST", "/challenges/{challenge_id}/knack_koden_attempt"},
			{"KnackKodenScoreboard", "GET", "/knack_koden_scoreboard"},
			{"KnackKodenRegisterClass", "POST", "/knack_koden_register_class"},
			{"KnackKodenGetClass", "POST", "/knack_koden_get_class"},
		},
		ListChallenges:             NewListChallengesHandler(e.ListChallenges, mux, decoder, encoder, errhandler, formatter),
		ListEvents:                 NewListEventsHandler(e.ListEvents, mux, decoder, encoder, errhandler, formatter),
		GetCurrentMonthlyChallenge: NewGetCurrentMonthlyChallengeHandler(e.GetCurrentMonthlyChallenge, mux, decoder, encoder, errhandler, formatter),
		ListMonthlyChallenges:      NewListMonthlyChallengesHandler(e.ListMonthlyChallenges, mux, decoder, encoder, errhandler, formatter),
		SubmitFlag:                 NewSubmitFlagHandler(e.SubmitFlag, mux, decoder, encoder, errhandler, formatter),
		SchoolScoreboard:           NewSchoolScoreboardHandler(e.SchoolScoreboard, mux, decoder, encoder, errhandler, formatter),
		UserScoreboard:             NewUserScoreboardHandler(e.UserScoreboard, mux, decoder, encoder, errhandler, formatter),
		ListAuthors:                NewListAuthorsHandler(e.ListAuthors, mux, decoder, encoder, errhandler, formatter),
		ListCourses:                NewListCoursesHandler(e.ListCourses, mux, decoder, encoder, errhandler, formatter),
		EnrollCourse:               NewEnrollCourseHandler(e.EnrollCourse, mux, decoder, encoder, errhandler, formatter),
		CompleteCourse:             NewCompleteCourseHandler(e.CompleteCourse, mux, decoder, encoder, errhandler, formatter),
		KnackKodenSubmitFlag:       NewKnackKodenSubmitFlagHandler(e.KnackKodenSubmitFlag, mux, decoder, encoder, errhandler, formatter),
		KnackKodenScoreboard:       NewKnackKodenScoreboardHandler(e.KnackKodenScoreboard, mux, decoder, encoder, errhandler, formatter),
		KnackKodenRegisterClass:    NewKnackKodenRegisterClassHandler(e.KnackKodenRegisterClass, mux, decoder, encoder, errhandler, formatter),
		KnackKodenGetClass:         NewKnackKodenGetClassHandler(e.KnackKodenGetClass, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "challenge" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.ListChallenges = m(s.ListChallenges)
	s.ListEvents = m(s.ListEvents)
	s.GetCurrentMonthlyChallenge = m(s.GetCurrentMonthlyChallenge)
	s.ListMonthlyChallenges = m(s.ListMonthlyChallenges)
	s.SubmitFlag = m(s.SubmitFlag)
	s.SchoolScoreboard = m(s.SchoolScoreboard)
	s.UserScoreboard = m(s.UserScoreboard)
	s.ListAuthors = m(s.ListAuthors)
	s.ListCourses = m(s.ListCourses)
	s.EnrollCourse = m(s.EnrollCourse)
	s.CompleteCourse = m(s.CompleteCourse)
	s.KnackKodenSubmitFlag = m(s.KnackKodenSubmitFlag)
	s.KnackKodenScoreboard = m(s.KnackKodenScoreboard)
	s.KnackKodenRegisterClass = m(s.KnackKodenRegisterClass)
	s.KnackKodenGetClass = m(s.KnackKodenGetClass)
}

// Mount configures the mux to serve the challenge endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountListChallengesHandler(mux, h.ListChallenges)
	MountListEventsHandler(mux, h.ListEvents)
	MountGetCurrentMonthlyChallengeHandler(mux, h.GetCurrentMonthlyChallenge)
	MountListMonthlyChallengesHandler(mux, h.ListMonthlyChallenges)
	MountSubmitFlagHandler(mux, h.SubmitFlag)
	MountSchoolScoreboardHandler(mux, h.SchoolScoreboard)
	MountUserScoreboardHandler(mux, h.UserScoreboard)
	MountListAuthorsHandler(mux, h.ListAuthors)
	MountListCoursesHandler(mux, h.ListCourses)
	MountEnrollCourseHandler(mux, h.EnrollCourse)
	MountCompleteCourseHandler(mux, h.CompleteCourse)
	MountKnackKodenSubmitFlagHandler(mux, h.KnackKodenSubmitFlag)
	MountKnackKodenScoreboardHandler(mux, h.KnackKodenScoreboard)
	MountKnackKodenRegisterClassHandler(mux, h.KnackKodenRegisterClass)
	MountKnackKodenGetClassHandler(mux, h.KnackKodenGetClass)
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
	formatter func(err error) goahttp.Statuser,
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

// MountListEventsHandler configures the mux to serve the "challenge" service
// "ListEvents" endpoint.
func MountListEventsHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/events", f)
}

// NewListEventsHandler creates a HTTP handler which loads the HTTP request and
// calls the "challenge" service "ListEvents" endpoint.
func NewListEventsHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeListEventsRequest(mux, decoder)
		encodeResponse = EncodeListEventsResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "ListEvents")
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

// MountGetCurrentMonthlyChallengeHandler configures the mux to serve the
// "challenge" service "GetCurrentMonthlyChallenge" endpoint.
func MountGetCurrentMonthlyChallengeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/current_monthly_challenge", f)
}

// NewGetCurrentMonthlyChallengeHandler creates a HTTP handler which loads the
// HTTP request and calls the "challenge" service "GetCurrentMonthlyChallenge"
// endpoint.
func NewGetCurrentMonthlyChallengeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetCurrentMonthlyChallengeRequest(mux, decoder)
		encodeResponse = EncodeGetCurrentMonthlyChallengeResponse(encoder)
		encodeError    = EncodeGetCurrentMonthlyChallengeError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "GetCurrentMonthlyChallenge")
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
	formatter func(err error) goahttp.Statuser,
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
	mux.Handle("POST", "/challenges/{challenge_id}/attempt", f)
}

// NewSubmitFlagHandler creates a HTTP handler which loads the HTTP request and
// calls the "challenge" service "SubmitFlag" endpoint.
func NewSubmitFlagHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
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

// MountSchoolScoreboardHandler configures the mux to serve the "challenge"
// service "SchoolScoreboard" endpoint.
func MountSchoolScoreboardHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/scoreboard", f)
}

// NewSchoolScoreboardHandler creates a HTTP handler which loads the HTTP
// request and calls the "challenge" service "SchoolScoreboard" endpoint.
func NewSchoolScoreboardHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeSchoolScoreboardRequest(mux, decoder)
		encodeResponse = EncodeSchoolScoreboardResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "SchoolScoreboard")
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

// MountUserScoreboardHandler configures the mux to serve the "challenge"
// service "UserScoreboard" endpoint.
func MountUserScoreboardHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/user_scoreboard", f)
}

// NewUserScoreboardHandler creates a HTTP handler which loads the HTTP request
// and calls the "challenge" service "UserScoreboard" endpoint.
func NewUserScoreboardHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUserScoreboardRequest(mux, decoder)
		encodeResponse = EncodeUserScoreboardResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "UserScoreboard")
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

// MountListAuthorsHandler configures the mux to serve the "challenge" service
// "ListAuthors" endpoint.
func MountListAuthorsHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/authors", f)
}

// NewListAuthorsHandler creates a HTTP handler which loads the HTTP request
// and calls the "challenge" service "ListAuthors" endpoint.
func NewListAuthorsHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeListAuthorsRequest(mux, decoder)
		encodeResponse = EncodeListAuthorsResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "ListAuthors")
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

// MountListCoursesHandler configures the mux to serve the "challenge" service
// "ListCourses" endpoint.
func MountListCoursesHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/courses", f)
}

// NewListCoursesHandler creates a HTTP handler which loads the HTTP request
// and calls the "challenge" service "ListCourses" endpoint.
func NewListCoursesHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeListCoursesRequest(mux, decoder)
		encodeResponse = EncodeListCoursesResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "ListCourses")
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

// MountEnrollCourseHandler configures the mux to serve the "challenge" service
// "EnrollCourse" endpoint.
func MountEnrollCourseHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/courses/{id}/enroll", f)
}

// NewEnrollCourseHandler creates a HTTP handler which loads the HTTP request
// and calls the "challenge" service "EnrollCourse" endpoint.
func NewEnrollCourseHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeEnrollCourseRequest(mux, decoder)
		encodeResponse = EncodeEnrollCourseResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "EnrollCourse")
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

// MountCompleteCourseHandler configures the mux to serve the "challenge"
// service "CompleteCourse" endpoint.
func MountCompleteCourseHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/courses/{id}/complete", f)
}

// NewCompleteCourseHandler creates a HTTP handler which loads the HTTP request
// and calls the "challenge" service "CompleteCourse" endpoint.
func NewCompleteCourseHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCompleteCourseRequest(mux, decoder)
		encodeResponse = EncodeCompleteCourseResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "CompleteCourse")
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

// MountKnackKodenSubmitFlagHandler configures the mux to serve the "challenge"
// service "KnackKodenSubmitFlag" endpoint.
func MountKnackKodenSubmitFlagHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/challenges/{challenge_id}/knack_koden_attempt", f)
}

// NewKnackKodenSubmitFlagHandler creates a HTTP handler which loads the HTTP
// request and calls the "challenge" service "KnackKodenSubmitFlag" endpoint.
func NewKnackKodenSubmitFlagHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeKnackKodenSubmitFlagRequest(mux, decoder)
		encodeResponse = EncodeKnackKodenSubmitFlagResponse(encoder)
		encodeError    = EncodeKnackKodenSubmitFlagError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "KnackKodenSubmitFlag")
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

// MountKnackKodenScoreboardHandler configures the mux to serve the "challenge"
// service "KnackKodenScoreboard" endpoint.
func MountKnackKodenScoreboardHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/knack_koden_scoreboard", f)
}

// NewKnackKodenScoreboardHandler creates a HTTP handler which loads the HTTP
// request and calls the "challenge" service "KnackKodenScoreboard" endpoint.
func NewKnackKodenScoreboardHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeKnackKodenScoreboardRequest(mux, decoder)
		encodeResponse = EncodeKnackKodenScoreboardResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "KnackKodenScoreboard")
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

// MountKnackKodenRegisterClassHandler configures the mux to serve the
// "challenge" service "KnackKodenRegisterClass" endpoint.
func MountKnackKodenRegisterClassHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/knack_koden_register_class", f)
}

// NewKnackKodenRegisterClassHandler creates a HTTP handler which loads the
// HTTP request and calls the "challenge" service "KnackKodenRegisterClass"
// endpoint.
func NewKnackKodenRegisterClassHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeKnackKodenRegisterClassRequest(mux, decoder)
		encodeResponse = EncodeKnackKodenRegisterClassResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "KnackKodenRegisterClass")
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

// MountKnackKodenGetClassHandler configures the mux to serve the "challenge"
// service "KnackKodenGetClass" endpoint.
func MountKnackKodenGetClassHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/knack_koden_get_class", f)
}

// NewKnackKodenGetClassHandler creates a HTTP handler which loads the HTTP
// request and calls the "challenge" service "KnackKodenGetClass" endpoint.
func NewKnackKodenGetClassHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeKnackKodenGetClassRequest(mux, decoder)
		encodeResponse = EncodeKnackKodenGetClassResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "KnackKodenGetClass")
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
