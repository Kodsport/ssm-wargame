package goa

import (
	. "goa.design/goa/v3/dsl"
)

var User = ResultType("application/vnd.ssm.user", func() {
	Attribute("id", String, func() {
		Example("uuid todo")
	})
	Attribute("email", String, func() {
		Example("movitz.sunar@ssm.example")
	})
	Attribute("first_name", String, func() {
		Example("Movitz")
	})
	Attribute("last_name", String, func() {
		Example("Sunar")
	})
	Attribute("role", String, func() {
		Example("admin")
	})
	Attribute("school_id", Int, func() {
		Example(78433202)
	})
	Required("id", "email", "first_name", "last_name", "role")
})

var ResultAdminChallenge = ResultType("application/vnd.ssm.admin.challenge", func() {
	Description("A Wargame challenge")
	Reference(Challenge)

	Attribute("id")

	Attribute("slug")
	Attribute("title")
	Attribute("description")
	Attribute("score")
	Attribute("services")
	Attribute("files")
	Attribute("publish_at")
	Attribute("solves")

	Attribute("files", ArrayOf(AdminChallengeFile))
	Attribute("flags", ArrayOf(AdminChallengeFlag))

	Attribute("category_id", String, func() {
		Example("12b8dc3a-10ae-49ed-9d69-5208ccd92ed1")
		Format(FormatUUID)
	})

	Required("files", "category_id")
})

var ResultChallenge = ResultType("application/vnd.ssm.challenge", func() {
	Description("A Wargame challenge")
	Reference(Challenge)

	Attribute("id")

	Attribute("slug")
	Attribute("title")
	Attribute("description")
	Attribute("score")
	Attribute("services")
	Attribute("files")
	Attribute("solves")

	Attribute("solved", Boolean, func() {
		Example(true)
		Description("whether the user has solved the challenge or not")
	})
	Attribute("category", String, func() {
		Example("Misc")
	})
	Required("solved", "category")
})

var SchoolScoreboard = ResultType("application/vnd.ssm.shoolscoreboard", func() {
	Description("A scoreboard of schools")
	Attribute("scores", ArrayOf(SchoolScoreboardScore))
	Required("scores")
})

var SchoolScoreboardScore = Type("SchoolScoreboardScore", func() {
	Attribute("score", Int, func() {
		Example(1337)
	})
	Attribute("school_name", String, func() {
		Example("Stockholm Science and Innovation School")
	})
	Required("school_name", "score")
})

var UserMonthlyChallenge = ResultType("application/vnd.ssm.usermonthlychallenges", func() {
	Extend(MonthlyChallenge)
	Attribute("challenge_id")
	Attribute("display_month")
	Attribute("start_date")
	Attribute("end_date")

	Reference(Challenge)

	Attribute("challenge", ResultChallenge)
	Required("challenge")
})
