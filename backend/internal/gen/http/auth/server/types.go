// Code generated by goa v3.5.2, DO NOT EDIT.
//
// auth HTTP server types
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package server

import (
	auth "github.com/sakerhetsm/ssm-wargame/internal/gen/auth"
	goa "goa.design/goa/v3/pkg"
)

// ExchangeDiscordRequestBody is the type of the "auth" service
// "ExchangeDiscord" endpoint HTTP request body.
type ExchangeDiscordRequestBody struct {
	Code  *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	State *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
}

// GenerateDiscordAuthURLResponseBody is the type of the "auth" service
// "GenerateDiscordAuthURL" endpoint HTTP response body.
type GenerateDiscordAuthURLResponseBody struct {
	URL string `form:"url" json:"url" xml:"url"`
}

// ExchangeDiscordResponseBody is the type of the "auth" service
// "ExchangeDiscord" endpoint HTTP response body.
type ExchangeDiscordResponseBody struct {
	// A token to authenticate with the SSM API
	JWT string `form:"jwt" json:"jwt" xml:"jwt"`
}

// NewGenerateDiscordAuthURLResponseBody builds the HTTP response body from the
// result of the "GenerateDiscordAuthURL" endpoint of the "auth" service.
func NewGenerateDiscordAuthURLResponseBody(res *auth.GenerateDiscordAuthURLResult) *GenerateDiscordAuthURLResponseBody {
	body := &GenerateDiscordAuthURLResponseBody{
		URL: res.URL,
	}
	return body
}

// NewExchangeDiscordResponseBody builds the HTTP response body from the result
// of the "ExchangeDiscord" endpoint of the "auth" service.
func NewExchangeDiscordResponseBody(res *auth.ExchangeDiscordResult) *ExchangeDiscordResponseBody {
	body := &ExchangeDiscordResponseBody{
		JWT: res.JWT,
	}
	return body
}

// NewExchangeDiscordPayload builds a auth service ExchangeDiscord endpoint
// payload.
func NewExchangeDiscordPayload(body *ExchangeDiscordRequestBody) *auth.ExchangeDiscordPayload {
	v := &auth.ExchangeDiscordPayload{
		Code:  *body.Code,
		State: *body.State,
	}

	return v
}

// ValidateExchangeDiscordRequestBody runs the validations defined on
// ExchangeDiscordRequestBody
func ValidateExchangeDiscordRequestBody(body *ExchangeDiscordRequestBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.State == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("state", "body"))
	}
	return
}
