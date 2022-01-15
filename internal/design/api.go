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
	Required("id", "email", "first_name", "last_name", "role")
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

var AdminChallengeFiles = Type("AdminChallengeFiles", func() {
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

var MonthlyChallengeMeta = Type("MonthlyChallengeMeta", func() {
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

	Required("start_date", "end_date", "display_month")
})

var MonthlyChallenge = ResultType("application/vnd.ssm.monthly.challenge", func() {
	Description("A monthly challenge")

	Extend(ResultChallenge)

	Reference(MonthlyChallengeMeta)
	Attribute("display_month")
	Attribute("start_date")
	Attribute("end_date")

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
	Attribute("published")
	Attribute("solves")

	Attribute("files", ArrayOf(AdminChallengeFiles))

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
	Attribute("published")
	Attribute("solves")

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
	Attribute("solves", Int64, "The numer of people who solved the challenge", func() {
		Example(3)
	})

	Attribute("services", ArrayOf(ChallengeService))
	Attribute("files", ArrayOf(ChallengeFiles))

	Required("id", "title", "description", "score", "published", "slug", "solves")
})

var AdminChallenge = Type("AdminChallenge", func() {
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
	Attribute("solves", Int64, "The numer of people who solved the challenge", func() {
		Example(3)
	})

	Required("id", "title", "description", "score", "published", "slug", "solves")
})
