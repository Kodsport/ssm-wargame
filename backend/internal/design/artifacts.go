package goa

import (
	. "goa.design/goa/v3/dsl"
)

var ChallengeIDArtifact = Type("ChallengeIDArtifact", func() {
	Attribute("challenge_id", String, "ID of a challenge", func() {
		Example("195229b0-b15f-4ee5-9a99-94bfff492967")
		Format(FormatUUID)
	})
	Required("challenge_id")
})

var FileIDArtifact = Type("FileIDArtifact", func() {
	Attribute("fileID", String, "ID of a file", func() {
		Example("195229b0-b15f-4ee5-9a99-94bfff492967")
		Format(FormatUUID)
	})
	Required("fileID")
})

var IDArtifact = Type("IDArtifact", func() {
	Attribute("id", String, "ID of a file", func() {
		Example("020817da-8b5c-42c4-9e52-0f3a6628c1f8")
		Format(FormatUUID)
	})
	Required("id")
})
