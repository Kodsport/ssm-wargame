package goa

import (
	. "goa.design/goa/v3/dsl"
)

var CreateChallengePayload = Type("CreateChallengePayload", func() {
	Reference(Challenge)

	Attribute("slug")
	Attribute("title")
	Attribute("description")
	Attribute("score")
	Attribute("publish_at")
	Attribute("ctf_event_id")
})

var TokenPayload = Type("TokenPayload", func() {
	Extend(OptionalTokenPayload)
	Required("token")
})

var OptionalTokenPayload = Type("OptionalTokenPayload", func() {
	Token("token", String, func() {
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6InN1cCAoIDoiLCJpYXQiOjE1MTYyMzkwMjJ9.niAX9xS6jNYQSX6hleuwGmzkUCuR9OXPRb5BksyMlkg")
	})
})
