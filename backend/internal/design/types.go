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
