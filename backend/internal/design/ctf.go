package goa

import . "goa.design/goa/v3/dsl"

var _ = Service("ctf", func() {
	Description("CTF mini-competition endpoints for school visits.")
	HTTP(func() {
		Path("/ctfs")
	})

	Method("Get", func() {
		Description("Get info about a ctf.")
		Payload(func() {
			Attribute("slug", String)
			Required("slug")
		})
		Result(CTFInfo)
		HTTP(func() {
			GET("/{slug}")
			Response(StatusOK)
		})
	})

	Method("RegisterUser", func() {
		Description("Register a user for a ctf with a unique username.")
		Payload(func() {
			Attribute("slug", String)
			Attribute("username", String)
			Required("slug", "username")
		})
		Result(CTFUser)
		Error("UsernameTaken", func() {
			Description("Username already exists for this ctf.")
		})
		HTTP(func() {
			POST("/{slug}/users")
			Response(StatusCreated)
			Response("UsernameTaken", StatusConflict)
		})
	})

	Method("GetUser", func() {
		Description("Get a user for a ctf.")
		Payload(func() {
			Attribute("password", String)
			Attribute("slug", String)
			Required("slug", "password")
		})
		Result(CTFUser)
		Error("not_found")
		Error("invalid_password")
		HTTP(func() {
			GET("/{slug}/user")
			Param("password")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response("invalid_password", StatusUnauthorized)
		})
	})
	Method("GetUserSolves", func() {
		Description("Get a user's solves for a ctf.")
		Payload(func() {
			Attribute("slug", String)
			Attribute("id", String)
			Required("slug", "id")
		})
		Result(CTFUserSolves)
		HTTP(func() {
			GET("/{slug}/user/{id}/solves")
			Param("id")
			Response(StatusOK)
		})
	})
	Method("ListChallenges", func() {
		Description("List challenges for a ctf.")
		Payload(func() {
			Attribute("slug", String)
			Attribute("password", String)
			Required("slug")
		})
		Result(CollectionOf(ResultChallenge))
		HTTP(func() {
			GET("/{slug}/challenges")
			Param("password")
			Response(StatusOK)
		})
	})

	Method("Scoreboard", func() {
		Description("Get scoreboard for a ctf.")
		Payload(func() {
			Attribute("slug", String)
			Required("slug")
		})
		Result(ArrayOf(CTFScore))
		HTTP(func() {
			GET("/{slug}/scoreboard")
			Response(StatusOK)
		})
	})

	Method("SubmitFlag", func() {
		Description("Submit a flag for a user in a ctf.")
		Payload(func() {
			Extend(FlagSubmission)
		})
		Result(CTFSolve)
		Error("already_solved")
		Error("incorrect_flag")
		Error("ctf_not_active")
		HTTP(func() {
			POST("/{slug}/attempt")
			Response(StatusOK)
			Response("already_solved", StatusConflict, func() {
				Description("If the challenge is already solved")
			})
			Response("incorrect_flag", StatusBadRequest, func() {
				Description("If the challenge is already solved")
			})
			Response("ctf_not_active", StatusForbidden, func() {
				Description("If the ctf is not active")
			})
		})
	})
})
