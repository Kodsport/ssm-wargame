package goa

import . "goa.design/goa/v3/dsl"

var TokenPayload = Type("TokenPayload", func() {
	Token("token", String)
})

var _ = Service("challenge", func() {
	HTTP(func() {
		Path("/challenge")
	})
	Security(JWTAuth)
	Method("ListChallenges", func() {
		Result(CollectionOf(ResultChallenge))
		Payload(func() {
			Attribute("view", func() {
				Default("default")
				Enum("default", "author")
			})
			Extend(TokenPayload)
		})
		HTTP(func() {
			GET("/")
			Param("view")
			Response(StatusOK)
		})
	})
	Method("CreateChallenge", func() {
		Payload(func() {
			Extend(CreateChallengePayload)
			Extend(TokenPayload)
			Required("token")
		})
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
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
