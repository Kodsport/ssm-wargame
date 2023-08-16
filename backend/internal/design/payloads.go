package goa

import (
	. "goa.design/goa/v3/dsl"
)

var CreateChallengePayload = Type("CreateChallengePayload", func() {
	Reference(Challenge)

	Attribute("slug")
	Attribute("title")
	Attribute("description")
	Attribute("publish_at")
	Attribute("ctf_event_id")
	Attribute("hide")

	Attribute("static_score", Int, func() {
		Example(50)
	})

	Attribute("category_id", String, func() {
		Example("1b678292-6737-4cc7-8eae-aa821100293f")
		Format(FormatUUID)
	})

	Attribute("authors", ArrayOf(String), func() {
		Example([]string{"1b678293-6737-4cc7-8eae-aa821100293f"})
	})

	Required("category_id")
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

var ImportChallService = Type("ImportChallService", func() {
	Reference(ChallengeService)
	Attribute("user_display", String)
	Attribute("hyperlink", Boolean)
})

var ImportChallFlag = Type("ImportChallFlag", func() {
	Attribute("type", String, func() {
		Example("regex")
	})
	Attribute("flag", String, func() {
		Example("fl4g_l0l")
	})
	Required("flag", "type")
})

var ChallImport = Type("ChallImport", func() {

	Attribute("title", String, func() {
		Example("DNS 101")
	})
	Attribute("description", String, func() {
		Example("how to dns")
	})
	Attribute("authors", ArrayOf(String), func() {
		Example([]string{"Movitz Sunar"})
	})
	Attribute("categories", ArrayOf(String), func() {
		Example([]string{"web"})
	})
	Attribute("score", Int, func() {
		Example(100)
	})
	Attribute("challenge_id", String, func() {
		Example("225ada44-3fde-460d-84a4-2f16ff579618")
		Format(FormatUUID)
	})

	Attribute("flag_format_prefix", String, func() {
		Example("SSM{")
	})
	Attribute("flag_format_suffix", String, func() {
		Example("}")
	})
	Attribute("file_urls", ArrayOf(String), func() {
		Example([]string{"https://bucket/key"})
	})
	Attribute("flags", ArrayOf(ImportChallFlag))
	Attribute("order", Int, func() {
		Example(5)
	})
	Attribute("services", ArrayOf(ImportChallService))

	Required("title", "description", "challenge_id")
})

var CreateAuthorPayload = Type("CreateAuthorPayload", func() {
	Reference(Author)
	Attribute("full_name")
	Attribute("description")
	Attribute("sponsor")
	Attribute("slug")
	Attribute("image_url")
	Attribute("publish")

})

var CreateCoursePayload = Type("CreateCoursePayload", func() {
	Reference(Course)

	Attribute("title")
	Attribute("slug")
	Attribute("category")
	Attribute("difficulty")
	Attribute("description")
	Attribute("publish")

	Attribute("author_ids", ArrayOf(String), func() {
		Example([]string{"b29869ec-6eef-451c-a389-249a6d6ac47b"})
	})

})
