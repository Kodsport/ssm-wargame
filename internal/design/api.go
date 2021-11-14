package goa

import (
	. "goa.design/goa/v3/dsl"
	// _ "goa.design/plugins/v3/goakit"
)

var _ = API("wargame", func() {
	Title("Wargame Service")
	Description("HTTP service for the SSM Wargame")
})

var ChallengeService = Type("ChallengeService", func() {

})

var ChallengeFiles = Type("ChallengeFiles", func() {

})

var Challenge = ResultType("application/vnd.ssm.challenge", func() {
	Description("A Wargame challenge")

	Attributes(func() {
		Attribute("id", String, func() {
			Example("uuid")
		})
		Attribute("title", String, func() {
			Example("pwnme bla bla")
		})
		Attribute("description", String, func() {
			Example("A simple memory exploit bla bla")
		})
		Attribute("score", Int, func() {
			Example(50)
			Minimum(0)
		})
		Attribute("published", Boolean, func() {
			Example(true)
		})
		Attribute("services", ArrayOf(ChallengeService))
		Attribute("files", ArrayOf(ChallengeFiles))
	})

})
