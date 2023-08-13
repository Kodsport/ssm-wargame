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
	Attribute("full_name", String, func() {
		Example("Movitz Sunar")
	})
	Attribute("role", String, func() {
		Example("admin")
	})
	Attribute("school_id", String, func() {
		Example("a0b05541-9211-4faf-82cc-d3bd64370bf4")
	})
	Required("id", "email", "full_name", "role")
})

var Solver = ResultType("application/vnd.ssm.solver", func() {
	Reference(User)

	Attribute("id")
	Attribute("full_name")
	Attribute("solved_at", Int64, func() {
		Example(16123128)
	})

	Required("solved_at")
})

var ResultAdminChallenge = ResultType("application/vnd.ssm.admin.challenge", func() {
	Description("A Wargame challenge")
	Reference(Challenge)

	Attribute("id")

	Attribute("slug")
	Attribute("title")
	Attribute("description")
	Attribute("services")
	Attribute("files")
	Attribute("publish_at")
	Attribute("solves")
	Attribute("ctf_event_id")

	Attribute("files", ArrayOf(AdminChallengeFile))
	Attribute("flags", ArrayOf(AdminChallengeFlag))

	Attribute("static_score", Int, func() {
		Example(50)
	})

	Attribute("category_id", String, func() {
		Example("12b8dc3a-10ae-49ed-9d69-5208ccd92ed1")
		Format(FormatUUID)
	})

	Attribute("authors", ArrayOf(String), func() {
		Example([]string{"dfaa1a6a-9051-4c52-9441-5d664536cf24"})
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
	Attribute("ctf_event_id")

	Attribute("solved", Boolean, func() {
		Example(true)
		Description("whether the user has solved the challenge or not")
	})
	Attribute("category", String, func() {
		Example("Misc")
	})

	Attribute("authors", ArrayOf(Author))

	Attribute("solvers", ArrayOf(Solver))

	Required("solved", "category")
})

var SchoolScoreboard = ResultType("application/vnd.ssm.school.scoreboard", func() {
	Description("A scoreboard of schools")
	Attribute("scores", ArrayOf(SchoolScoreboardScore))
	Required("scores")
})

var UserScoreboard = ResultType("application/vnd.ssm.user.scoreboard", func() {
	Description("A scoreboard of user")
	Attribute("scores", ArrayOf(UserScoreboardScore))
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

var UserScoreboardScore = Type("UserScoreboardScore", func() {
	Attribute("user_id", String, func() {
		Example("97ee0ee4-b65f-445b-97cc-b61d96db8d8e")
	})
	Attribute("name", String, func() {
		Example("Movitz Sunar")
	})
	Attribute("school_name", String, func() {
		Example("Stockholm Science and Innovation School")
	})
	Attribute("score", Int, func() {
		Example(1337)
	})
	Required("user_id", "school_name", "score", "name")
})

var UserMonthlyChallenge = ResultType("application/vnd.ssm.user.monthly.challenge", func() {
	Extend(MonthlyChallenge)
	Attribute("challenge_id")
	Attribute("display_month")
	Attribute("start_date")
	Attribute("end_date")

	Reference(Challenge)

	Attribute("challenge", ResultChallenge)
	Required("challenge")
})
