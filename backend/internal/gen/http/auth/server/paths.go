// Code generated by goa v3.5.2, DO NOT EDIT.
//
// HTTP request path constructors for the auth service.
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design

package server

// GenerateDiscordAuthURLAuthPath returns the URL path to the auth service GenerateDiscordAuthURL HTTP endpoint.
func GenerateDiscordAuthURLAuthPath() string {
	return "/auth/discord/url"
}

// ExchangeDiscordAuthPath returns the URL path to the auth service ExchangeDiscord HTTP endpoint.
func ExchangeDiscordAuthPath() string {
	return "/auth/discord/exchange"
}
