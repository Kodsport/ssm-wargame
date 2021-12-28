package goa

import . "goa.design/goa/v3/dsl"

var _ = Service("admin", func() {
	HTTP(func() {
		Path("/admin")
	})
	Security(JWTAuth)
	Method("ListChallenges", func() {
		Payload(func() {
			Extend(TokenPayload)
		})
		Result(CollectionOf(ResultChallenge))
		HTTP(func() {
			GET("/challenges")
			Response(StatusOK)
		})
	})
	Method("CreateChallenge", func() {
		Payload(func() {
			Extend(CreateChallengePayload)
			Extend(TokenPayload)
		})
		HTTP(func() {
			POST("/challenges")
			Response(StatusCreated)
		})
	})
})
