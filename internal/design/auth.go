package goa

import . "goa.design/goa/v3/dsl"

var _ = Service("auth", func() {
	HTTP(func() {
		Path("/auth")
	})
	Method("GenerateDiscordAuthURL", func() {
		Result(func() {
			Attribute("url", String, func() {
				Example("https://discord.com/api/oauth2/authorize?response_type=code&client_id=157730590492196864&scope=identify%20guilds.join&state=15773059ghq9183habn&redirect_uri=https%3A%2F%2Fnicememe.website&prompt=consent")
			})
			Required("url")
		})
		HTTP(func() {
			GET("/discord/url")
			Response(StatusOK)
		})
	})
	Method("ExchangeDiscord", func() {
		Payload(func() {
			Attribute("code", String, func() {
				Example("123abc")
			})
			Attribute("state", String, func() {
				Example("15773059ghq9183habn")
			})
			Required("code", "state")
		})
		Result(func() {
			Attribute("jwt", String, "A token to authenticate with the SSM API", func() {
				Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
			})
			Required("jwt")
		})
		HTTP(func() {
			POST("/discord/exchange")
			Response(StatusOK)
		})
	})
})
