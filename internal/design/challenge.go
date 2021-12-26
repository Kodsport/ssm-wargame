package goa

import . "goa.design/goa/v3/dsl"

var _ = Service("challenge", func() {
	HTTP(func() {
		Path("/challenge")
	})
	Method("ListChallenges", func() {
		Result(CollectionOf(ResultChallenge))
		Payload(func() {
			Attribute("view", func() {
				Default("default")
				Enum("default", "author")
			})
			Required("view")
		})
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})
	Method("CreateChallenge", func() {
		Payload(func() {

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
