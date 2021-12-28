package goa

import . "goa.design/goa/v3/dsl"

var _ = Service("challenge", func() {
	HTTP(func() {
		Path("/challenge")
	})
	Security(JWTAuth)
	Method("ListChallenges", func() {
		Result(CollectionOf(ResultChallenge))
		Payload(func() {
			Extend(OptionalTokenPayload)
		})
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})
	Method("SubmitFlag", func() {
		Payload(func() {
			Attribute("flag", String, func() {
				Example("SSM{flag}")
				MaxLength(200)
			})
			Attribute("challengeId", String, func() {
				Format(FormatUUID)
				Example("123")
			})
			Extend(TokenPayload)
			Required("flag", "challengeId")
		})
		Error("already_solved")
		Error("incorrect_flag")
		HTTP(func() {
			POST("/{challengeId}/attempt")
			Response(StatusOK)
			Response("already_solved", StatusConflict, func() {
				Description("If the challenge is already solved")
			})
			Response("incorrect_flag", StatusBadRequest, func() {
				Description("If the challenge is already solved")
			})
		})
	})
})
