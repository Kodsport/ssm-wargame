package goa

import (
	. "goa.design/goa/v3/dsl"
)

var User = ResultType("application/vnd.ssm.user", func() {
	Attribute("id", String, func() {
		Example("9f3e4702-d5b8-42c9-a409-a04678f23a33")
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
	Attribute("discord_id", String, func() {
		Example("384389545793845983547349")
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
	Attribute("hide")
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
	Attribute("chall_namespace")

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

var ResultCTFChallenge = ResultType("application/vnd.ssm.ctf.challenge", func() {
	Description("A CTF Wargame challenge")
	Reference(CTFListChallenge)

	Attribute("id")

	Attribute("slug")
	Attribute("title")
	Attribute("description")
	Attribute("score")
	Attribute("services")
	Attribute("files")
	Attribute("solves")

	Attribute("num_team_solves", Int)
	Attribute("num_solves_in_team", Int)

	Attribute("ctf_event_id")
	Attribute("chall_namespace")

	Attribute("solved", Boolean, func() {
		Example(true)
		Description("whether the user has solved the challenge or not")
	})
	Attribute("solved_in_team", Boolean, func() {
		Example(true)
		Description("whether the user has solved the challenge in their team or not")
	})
	Attribute("category", String, func() {
		Example("Misc")
	})

	Attribute("authors", ArrayOf(Author))

	Attribute("solvers", ArrayOf(Solver))
	Attribute("team_solvers", ArrayOf(Solver))
	Attribute("solvers_in_team", ArrayOf(Solver))

	Attribute("display_order", Int)

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
	Attribute("is_university", Boolean, func() {
		Example(true)
	})
	Required("school_name", "score", "is_university")
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

var AdminCourse = ResultType("application/vnd.ssm.admin.course", func() {
	Reference(Course)

	Attribute("id")
	Attribute("title")
	Attribute("slug")
	Attribute("category")
	Attribute("difficulty")
	Attribute("description")
	Attribute("publish")

	Attribute("author_ids", ArrayOf(String), func() {
		Example([]string{"46e0996b-3bee-4836-b5f8-afc2a62fc71b"})
	})
})

// ResultCTFUser now includes team fields for team-based CTFs.
type ResultCTFUser struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	Slug     string  `json:"slug"`
	TeamID   *string `json:"team_id,omitempty"`
	Teamname *string `json:"teamname,omitempty"`
}

// ResultCTFScore now includes team fields for team-based CTFs.
type ResultCTFScore struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	Score    int     `json:"score"`
	TeamID   *string `json:"team_id,omitempty"`
	Teamname *string `json:"teamname,omitempty"`
}

// ResultRegisterUserPayload now supports optional team_code for team-based CTFs.
type ResultRegisterUserPayload struct {
	Slug     string  `json:"slug"`
	Username string  `json:"username"`
	TeamCode *string `json:"team_code,omitempty"`
}

var DiscordUser = ResultType("application/vnd.ssm.discord.user", func() {
	Description("Discord user information")
	Attribute("id", String, func() {
		Example("123456789012345678")
	})
	Attribute("username", String, func() {
		Example("username")
	})
	Attribute("discriminator", String, func() {
		Example("1234")
	})
	Attribute("avatar", String, func() {
		Example("a_1234567890abcdef1234567890abcdef")
	})
	Attribute("global_name", String, func() {
		Example("Display Name")
	})
	Required("id", "username")
})

var SubmissionStats = Type("SubmissionStats", func() {
	Attribute("successful", Int, func() {
		Example(42)
	})
	Attribute("failed", Int, func() {
		Example(18)
	})
	Attribute("total", Int, func() {
		Example(60)
	})
	Attribute("success_rate", Int, func() {
		Example(70)
	})
	Required("successful", "failed", "total", "success_rate")
})

var ChallengeSubmissionsGroup = Type("ChallengeSubmissionsGroup", func() {
	Attribute("challenge_id", String, func() {
		Example("8bf8ae1c-0c49-4ea7-ad0f-0798d7b2728a")
		Format(FormatUUID)
	})
	Attribute("challenge_title", String, func() {
		Example("Binary Exploitation 101")
	})
	Attribute("challenge_slug", String, func() {
		Example("binexp-101")
	})
	Attribute("solved", Boolean, func() {
		Example(true)
	})
	Attribute("submissions", ArrayOf(ChallengeSubmission))
	Required("challenge_id", "challenge_title", "challenge_slug", "solved", "submissions")
})

var UserDetails = ResultType("application/vnd.ssm.admin.userdetails", func() {
	Description("Consolidated user details with all challenge submissions and statistics")

	// user info
	Extend(User)

	// discord user data
	Attribute("discord_user", DiscordUser)

	// solve stats
	Attribute("submission_stats", SubmissionStats)

	// hour grej
	Attribute("hourly_activity", ArrayOf(Int), func() {
		Example([]int{0, 0, 1, 0, 0, 0, 2, 5, 3, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 2, 4, 3, 1, 0})
		MinLength(24)
		MaxLength(24)
	})

	Attribute("challenge_submissions", ArrayOf(ChallengeSubmissionsGroup))

	Required("id", "email", "full_name", "role", "submission_stats", "hourly_activity", "challenge_submissions")
})
