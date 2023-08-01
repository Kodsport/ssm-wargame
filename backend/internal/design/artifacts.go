package goa

import (
	. "goa.design/goa/v3/dsl"
)

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
