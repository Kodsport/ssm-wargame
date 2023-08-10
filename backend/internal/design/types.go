package goa

import (
	. "goa.design/goa/v3/dsl"
)

var Category = Type("Category", func() {
	Attribute("id", String, func() {
		Example("e721a338-44de-4de8-a562-43d2db5f4115")
		Format(FormatUUID)
	})
	Attribute("name", String, func() {
		Example("Forensics")
	})

	Required("id", "name")
})

var Challenge = Type("Challenge", func() {
	Attribute("id", String, func() {
		Example("e721a338-44de-4de8-a562-43d2db5f4115")
		Format(FormatUUID)
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

	Attribute("ctf_event_id", String, "The ID of the CTF the challenge was taken from", func() {
		Example("c397efb2-b171-4d77-9166-d105cf4f521a")
		Format(FormatUUID)
	})

	Attribute("category", String, func() {
		Example("Forensics")
	})

	Attribute("services", ArrayOf(ChallengeService))
	Attribute("files", ArrayOf(ChallengeFiles))

	Required("id", "title", "description", "score", "slug", "solves", "category")
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
	Attribute("id", Int, func() {
		Example(78433202)
	})
	Attribute("name", String, func() {
		Example("Engelbrektsskolan")
	})
	Attribute("municipality_name", String, func() {
		Example("Örebro")
	})
	Required("id", "name", "municipality_name")
})

var AdminChallengeFlag = Type("AdminChallengeFlag", func() {
	Attribute("id", String, func() {
		Example("71333e34-4c6b-483e-b3c7-c77d73008cca")
	})
	Attribute("flag", String, func() {
		Example("SSM{yo}")
	})

	Required("id", "flag")
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
	Attribute("id", String, func() {
		Example("71333e34-4c6b-483e-b3c7-c77d73008cca")
	})
	Attribute("filename", String, func() {
		Example("decryptor.exe")
	})
	Attribute("url", String, func() {
		Example("https://s3/bucket/key?signature")
	})

	Required("id", "filename", "url")
})

var CTFEvent = Type("CTFEvent", func() {
	Attribute("id", String, func() {
		Example("71333e34-4c6b-483e-b3c7-c77d73008cca")
	})
	Attribute("name", String, func() {
		Example("Säkerhet-SM 2023")
	})
	Required("id", "name")
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
	Attribute("id", String, func() {
		Example("34062de2-6687-4acb-ab4b-608cc246b133")
		Format(FormatUUID)
	})
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

	Required("id", "user_id", "input", "successful", "submitted_at")
})
