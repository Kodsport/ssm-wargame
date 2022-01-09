package goa

import . "goa.design/goa/v3/dsl"

var _ = Service("admin", func() {
	Error("unauthorized")
	Error("not_found")
	HTTP(func() {
		Response("unauthorized", StatusForbidden, func() {
			Description("When the user does not have the required role")
		})
		Response("not_found", StatusNotFound, func() {
			Description("Resource not found")
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

	Method("ListMonthlyChallenges", func() {
		Payload(func() {
			Extend(TokenPayload)
		})
		Result(ArrayOf(MonthlyChallengeMeta))
		HTTP(func() {
			GET("/monthly_challenges")
		})
	})

	Method("DeleteMonthlyChallenge", func() {
		Payload(func() {
			Extend(TokenPayload)
			Extend(ChallengeIDArtifact)
		})
		HTTP(func() {
			DELETE("/monthly_challenges/{challengeID}")
		})
	})

	Method("CreateMonthlyChallenge", func() {
		Payload(func() {
			Extend(TokenPayload)
			Extend(MonthlyChallengeMeta)
			Extend(ChallengeIDArtifact)
		})
		HTTP(func() {
			POST("/monthly_challenges")
		})
	})

	Method("ListUsers", func() {
		Payload(func() {
			Extend(TokenPayload)
		})
		Result(ArrayOf(User))
		HTTP(func() {
			GET("/users")
		})
	})

})
