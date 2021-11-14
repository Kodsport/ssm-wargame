package goa

import . "goa.design/goa/v3/dsl"

var _ = Service("challenge", func() {
	HTTP(func() {
		Path("/challenge")
	})
	Method("ListChallenges", func() {
		Result(CollectionOf(Challenge), func() {
			View("default")
		})
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})
})
