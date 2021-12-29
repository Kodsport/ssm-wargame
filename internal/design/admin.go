package goa

import . "goa.design/goa/v3/dsl"

var _ = Service("admin", func() {
	Error("unauthorized")
	HTTP(func() {
		Response("unauthorized", StatusForbidden, func() {
			Description("When the user does not have the required role")
		})
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
