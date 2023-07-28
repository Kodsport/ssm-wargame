package goa

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("wargame", func() {
	Title("Wargame Service")
	Description("HTTP service for the SSM Wargame")
})

var JWTAuth = JWTSecurity("jwt")

var ChallengeIDArtifact = Type("ChallengeIDArtifact", func() {
	Attribute("challengeID", String, "ID of a challenge", func() {
		Example("195229b0-b15f-4ee5-9a99-94bfff492967")
		Format(FormatUUID)
	})
	Required("challengeID")
})

var FileIDArtifact = Type("FileIDArtifact", func() {
	Attribute("fileID", String, "ID of a file", func() {
		Example("195229b0-b15f-4ee5-9a99-94bfff492967")
		Format(FormatUUID)
	})
	Required("fileID")
})

var OptionalTokenPayload = Type("OptionalTokenPayload", func() {
	Token("token", String, func() {
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6InN1cCAoIDoiLCJpYXQiOjE1MTYyMzkwMjJ9.niAX9xS6jNYQSX6hleuwGmzkUCuR9OXPRb5BksyMlkg")
	})
})

var TokenPayload = Type("TokenPayload", func() {
	Extend(OptionalTokenPayload)
	Required("token")
})

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

var ChallengeService = Type("ChallengeService", func() {
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
	Attribute("bucket", String, func() {
		Example("ssm-prod-bucket")
	})
	Attribute("key", String, func() {
		Example("file.exe")
	})
	Attribute("size", Int64, func() {
		Example(87123)
	})
	Attribute("md5", String, "MD5 hash of the file content in base64", func() {
		Example("cq02dBbcuugBHM1oKyvMlQ==")
	})

	Required("id", "filename", "url", "key", "bucket", "size", "md5")
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
	Reference(Challenge)

	Attribute("slug")
	Attribute("title")
	Attribute("description")

	Attribute("challenge_id", String, func() {
		Example("85163218-8735-42ed-a7a6-42a9de2294df")
		Format(FormatUUID)
	})
	Attribute("display_month", String, func() {
		Description("The month(s) that the challenge is assigned for")
		Example("Januari/Februari")
	})
	Attribute("start_date", String, func() {
		Description("Starting date of the monthly challenge")
		Example("2006-02-01")
	})
	Attribute("end_date", String, func() {
		Description("Ending date of the monthly challenge")
		Example("2006-02-01")
	})

	Required("start_date", "end_date", "display_month", "challenge_id")
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

	Required("files")
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
})

var CreateChallengePayload = Type("CreateChallengePayload", func() {
	Reference(Challenge)

	Attribute("slug")
	Attribute("title")
	Attribute("description")
	Attribute("score")
	Attribute("publish_at")
	Attribute("ctf_event_id")
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

	Attribute("services", ArrayOf(ChallengeService))
	Attribute("files", ArrayOf(ChallengeFiles))

	Required("id", "title", "description", "score", "slug", "solves")
})
