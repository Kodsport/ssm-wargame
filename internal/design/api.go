package goa

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("wargame", func() {
	Title("Wargame Service")
	Description("HTTP service for the SSM Wargame")
})

var JWTAuth = JWTSecurity("jwt")

var OptionalTokenPayload = Type("OptionalTokenPayload", func() {
	Token("token", String)
})

var TokenPayload = Type("TokenPayload", func() {
	Token("token", String)
	Required("token")
})

var ChallengeService = Type("ChallengeService", func() {

})

var ChallengeFiles = Type("ChallengeFiles", func() {

})

var MonthlyChallenge = ResultType("application/vnd.ssm.monthly.challenge", func() {
	Description("A monthly challenge")
	Reference(Challenge)

	Attributes(func() {
		Attribute("id")
		Attribute("slug")
		Attribute("title")
		Attribute("description")
		Attribute("score")
		Attribute("services")
		Attribute("files")

		Attribute("solves", Int64, "The numer of people who solved the challenge", func() {
			Example(3)
		})

		Required("solves")

	})

})

var ResultChallenge = ResultType("application/vnd.ssm.challenge", func() {
	Description("A Wargame challenge")
	Reference(Challenge)

	Attributes(func() {
		Attribute("id")

		Attribute("slug")
		Attribute("title")
		Attribute("description")
		Attribute("score")
		Attribute("services")
		Attribute("files")
		Attribute("published")

		Attribute("solves", Int64, "The numer of people who solved the challenge", func() {
			Example(3)
		})

		Required("solves")

	})

})

var CreateChallengePayload = Type("CreateChallengePayload", func() {
	Reference(Challenge)

	Attribute("slug")
	Attribute("title")
	Attribute("description")
	Attribute("score")

})

var Challenge = Type("Challenge", func() {
	Attribute("id", String, func() {
		Example("e721a338-44de-4de8-a562-43d2db5f4115")
	})
	Attribute("slug", String, "A unique string that can be used in URLs", func() {
		Example("pwnme")
	})
	Attribute("title", String, "Title displayed to user", func() {
		Example("pwnme")
	})
	Attribute("description", String, "A short text describing the challenge", func() {
		Example("A heap overflow challenge")
	})
	Attribute("score", Int32, "The number of points given to the solver", func() {
		Example(50)
	})
	Attribute("published", Boolean, func() {
		Example(true)
	})

	Attribute("services", ArrayOf(ChallengeService))
	Attribute("files", ArrayOf(ChallengeFiles))

	Required("id", "title", "description", "score", "published", "slug")
})
