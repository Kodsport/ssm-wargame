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
	Method("ListEvents", func() {
		Result(ArrayOf(CTFEvent))
		Payload(func() {
			Extend(OptionalTokenPayload)
		})
		HTTP(func() {
			GET("/events")
			Response(StatusOK)
		})
	})
	Method("ListMonthlyChallenges", func() {
		Result(CollectionOf(UserMonthlyChallenge))
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
	Method("SchoolScoreboard", func() {
		Result(SchoolScoreboard)
		Payload(func() {
			Extend(OptionalTokenPayload)
		})
		HTTP(func() {
			GET("/scoreboard")
			Response(StatusOK)
		})
	})
})
