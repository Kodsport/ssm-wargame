package goa

import (
	. "goa.design/goa/v3/dsl"
)

var Category = Type("Category", func() {
	Extend(IDArtifact)

	Attribute("name", String, func() {
		Example("Forensics")
	})

	Required("name")
})

var Challenge = Type("Challenge", func() {
	Extend(IDArtifact)

	Attribute("slug", String, "A unique string that can be used in URLs", func() {
		Example("pwnme")
	})
	Attribute("title", String, "Title displayed to user", func() {
		Example("pwnme")
	})
	Attribute("description", String, "A short text describing the challenge", func() {
		Example("A heap overflow challenge")
	})
	Attribute("score", Int, "The number of points given to the solver", func() {
		Example(50)
	})
	Attribute("publish_at", Int64, func() {
		Example(1638384718)
		Description("unix timestamp")
	})
	Attribute("solves", Int, "The numer of people who solved the challenge", func() {
		Example(3)
	})
	Attribute("hide", Boolean, func() {
		Example(false)
	})
	Attribute("chall_namespace", String, func() {
		Example("knackkoden")
	})

	Attribute("ctf_event_id", String, "The ID of the CTF the challenge was taken from", func() {
		Example("c397efb2-b171-4d77-9166-d105cf4f521a")
		Format(FormatUUID)
	})

	Attribute("category", String, func() {
		Example("Forensics")
	})

	Attribute("services", ArrayOf(ChallengeService))
	Attribute("files", ArrayOf(ChallengeFiles))

	Required("title", "description", "score", "slug", "solves", "category", "hide")
})

var ChallengeService = Type("ChallengeService", func() {
	Attribute("user_display", String, func() {
		Example("nc 0.0.0.0 1234")
	})
	Attribute("hyperlink", Boolean, func() {
		Example(true)
	})
	Required("user_display", "hyperlink")
})

var School = Type("School", func() {
	Extend(IDArtifact)

	Attribute("name", String, func() {
		Example("Engelbrektsskolan")
	})
	Attribute("municipality_name", String, func() {
		Example("Örebro")
	})
	Attribute("is_university", Boolean, func() {
		Example(true)
	})
	Required("name", "municipality_name", "is_university")
})

var AdminChallengeFlag = Type("AdminChallengeFlag", func() {
	Extend(IDArtifact)

	Attribute("flag", String, func() {
		Example("SSM{yo}")
	})

	Required("flag")
})

var MonthlyChallenge = Type("MonthlyChallenge", func() {

	Attribute("challenge_id", String, func() {
		Example("85163218-8735-42ed-a7a6-42a9de2294df")
		Format(FormatUUID)
	})
	Attribute("display_month", String, func() {
		Description("The month(s) that the challenge is assigned for")
		Example("Januari/Februari")
	})
	Attribute("start_date", Int64, func() {
		Description("Starting date of the monthly challenge")
		Example(1690874841)
	})
	Attribute("end_date", Int64, func() {
		Description("Ending date of the monthly challenge")
		Example(1690884841)
	})

	Required("start_date", "end_date", "display_month", "challenge_id")
})

var ChallengeFiles = Type("ChallengeFiles", func() {
	Attribute("filename", String, func() {
		Example("decryptor.exe")
	})
	Attribute("url", String, func() {
		Example("https://s3/bucket/key?signature")
	})
	Required("filename", "url")
})

var AdminChallengeFile = Type("AdminChallengeFiles", func() {
	Extend(IDArtifact)

	Attribute("filename", String, func() {
		Example("decryptor.exe")
	})
	Attribute("url", String, func() {
		Example("https://s3/bucket/key?signature")
	})

	Required("filename", "url")
})

var CTFEvent = Type("CTFEvent", func() {
	Extend(IDArtifact)

	Attribute("name", String, func() {
		Example("Säkerhet-SM 2023")
	})
	Required("name")
})

var ChallengeMeta = Type("ChallengeMeta", func() {
	Attribute("solvers", ArrayOf(ChallengeSolver))
	Attribute("submissions", ArrayOf(ChallengeSubmission))

	Required("solvers", "submissions")
})

var ChallengeSolver = Type("ChallengeSolver", func() {
	Attribute("user_id", String, func() {
		Example("8bf8ae1c-0c49-4ea7-ad0f-0798d7b2728a")
		Format(FormatUUID)
	})
	Attribute("solved_at", Int64)

	Required("user_id", "solved_at")
})

var ChallengeSubmission = Type("ChallengeSubmission", func() {
	Extend(IDArtifact)

	Attribute("input", String, func() {
		Example("SSM{wrong_flag}")
	})
	Attribute("successful", Boolean, func() {
		Example(false)
	})
	Attribute("user_id", String, func() {
		Example("a78c6915-87f4-49e1-8ec9-0bb5c261b56d")
		Format(FormatUUID)
	})

	Attribute("submitted_at", Int64)

	Required("user_id", "input", "successful", "submitted_at")
})

var Author = Type("Author", func() {
	Extend(IDArtifact)

	Attribute("full_name", String, func() {
		Example("Movitz Sunar")
	})
	Attribute("description", String, func() {
		Example("Movitz gör saker")
	})
	Attribute("sponsor", Boolean, func() {
		Example(true)
	})
	Attribute("slug", String, func() {
		Example("movitz")
	})
	Attribute("image_url", String, func() {
		Example("movitz")
	})

	Attribute("publish", Boolean, func() {
		Example(true)
	})

	Required("sponsor", "full_name", "description", "slug", "publish")

})

var Course = Type("Course", func() {
	Extend(IDArtifact)
	Attribute("title", String, func() {
		Example("SQL-Injektioner")
	})
	Attribute("slug", String, func() {
		Example("sqli")
	})
	Attribute("category", String, func() {
		Example("web")
	})
	Attribute("difficulty", String, func() {
		Example("advanced")
	})
	Attribute("description", String, func() {
		Example("markdown text")
	})

	Attribute("enrolled", Boolean, func() {
		Example(false)
	})

	Attribute("publish", Boolean, func() {
		Example(true)
	})
	Attribute("completed", Boolean, func() {
		Example(true)
	})

	Attribute("authors", ArrayOf(Author))
	Attribute("course_items", ArrayOf(CourseItem))

	Required("title", "category", "difficulty", "description", "publish", "slug", "enrolled", "completed")

})

var CourseItem = Type("CourseItem", func() {
	Extend(IDArtifact)
	Extend(ChallengeIDArtifact)
	Attribute("position", Int, func() {
		Example(0)
		Description("to sort after")
	})
	Required("position")
})
