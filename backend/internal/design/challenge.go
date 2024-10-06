package goa

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("challenge", func() {
	HTTP(func() {
	})
	Security(JWTAuth)
	Method("ListChallenges", func() {
		Result(CollectionOf(ResultChallenge))
		Payload(func() {
			Extend(OptionalTokenPayload)
			Attribute("slug", String, func() {
				Example("brumm")
				Description("Filter by slug")
			})
			Attribute("ids", ArrayOf(String), func() {
				Example([]string{})
				Description("Selectivly take out certain challs")
			})
		})
		HTTP(func() {
			GET("/challenges")
			Response(StatusOK)
			Param("slug")
			Param("ids")
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
	Method("GetCurrentMonthlyChallenge", func() {
		Result(UserMonthlyChallenge)
		Payload(func() {
			Extend(OptionalTokenPayload)
		})
		Error("not_found")
		HTTP(func() {
			GET("/current_monthly_challenge")
			Response(StatusOK)
			Response("not_found", StatusNotFound, func() {
				Description("If there is no current monthly challenge")
			})
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
			POST("/challenges/{challenge_id}/attempt")
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
	Method("UserScoreboard", func() {
		Result(UserScoreboard)
		Payload(func() {
			Extend(OptionalTokenPayload)
		})
		HTTP(func() {
			GET("/user_scoreboard")
			Response(StatusOK)
		})
	})
	Method("ListAuthors", func() {
		Result(ArrayOf(Author))
		Payload(func() {
			Extend(OptionalTokenPayload)
		})
		HTTP(func() {
			GET("/authors")
			Response(StatusOK)
		})
	})
	Method("ListCourses", func() {
		Result(ArrayOf(Course))
		Payload(func() {
			Extend(OptionalTokenPayload)
		})
		HTTP(func() {
			GET("/courses")
			Response(StatusOK)
		})
	})
	Method("EnrollCourse", func() {
		Payload(func() {
			Extend(OptionalTokenPayload)
			Extend(IDArtifact)
		})
		HTTP(func() {
			POST("/courses/{id}/enroll")
			Response(StatusOK)
		})
	})
	Method("CompleteCourse", func() {
		Payload(func() {
			Extend(OptionalTokenPayload)
			Extend(IDArtifact)
		})
		HTTP(func() {
			POST("/courses/{id}/complete")
			Response(StatusOK)
		})
	})

	// hacky addition of knäck koden routes

	Method("KnackKodenSubmitFlag", func() {
		Payload(func() {
			Attribute("flag", String, func() {
				Example("SSM{flag}")
				MaxLength(200)
			})
			Extend(ChallengeIDArtifact)
			Extend(OptionalTokenPayload)
			Attribute("password", String)
			Required("flag", "password")
		})
		Error("already_solved")
		Error("incorrect_flag")
		HTTP(func() {
			POST("/challenges/{challenge_id}/knack_koden_attempt")
			Response(StatusOK)
			Response("already_solved", StatusConflict, func() {
				Description("If the challenge is already solved")
			})
			Response("incorrect_flag", StatusBadRequest, func() {
				Description("If the challenge is already solved")
			})
		})
	})

	Method("KnackKodenScoreboard", func() {
		Result(SchoolScoreboard)
		Payload(func() {
			Extend(OptionalTokenPayload)
		})
		HTTP(func() {
			GET("/knack_koden_scoreboard")
			Response(StatusOK)
		})
	})

	Method("KnackKodenRegisterClass", func() {
		Payload(func() {
			Extend(OptionalTokenPayload)
			Extend(KnackKodenRegisterClassPayload)
		})
		Result(func() {
			Attribute("password", String)
			Required("password")
		})
		HTTP(func() {
			POST("/knack_koden_register_class")
			Response(StatusOK)

		})
	})
	Method("KnackKodenGetClass", func() {
		Payload(func() {
			Extend(OptionalTokenPayload)
			Attribute("password", String)
			Required("password")
		})
		Result(func() {
			Attribute("class_name", String)
			Attribute("solves", ArrayOf(String))
			Required("class_name", "solves")
		})
		HTTP(func() {
			POST("/knack_koden_get_class")
			Response(StatusOK)

		})
	})

})
