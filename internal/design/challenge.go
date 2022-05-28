package goa

import . "goa.design/goa/v3/dsl"

var _ = Service("challenge", func() {
	HTTP(func() {
	})
	Security(JWTAuth)
	Method("ListChallenges", func() {
		Result(CollectionOf(ResultChallenge))
		Payload(func() {
			Extend(OptionalTokenPayload)
		})
		HTTP(func() {
			GET("/challenges")
			Response(StatusOK)
		})
	})
	Method("ListMonthlyChallenges", func() {
		Result(ArrayOf(MonthlyChallenge))
		Payload(func() {
			Extend(OptionalTokenPayload)
		})
		HTTP(func() {
			GET("/monthly_challenges")
			Response(StatusOK)
		})
	})
	Method("SubmitFlag", func() {
		Payload(func() {
			Attribute("flag", String, func() {
				Example("SSM{flag}")
				MaxLength(200)
			})
			Extend(ChallengeIDArtifact)
			Extend(TokenPayload)
			Required("flag")
		})
		Error("already_solved")
		Error("incorrect_flag")
		HTTP(func() {
			POST("/challenges/{challengeID}/attempt")
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
