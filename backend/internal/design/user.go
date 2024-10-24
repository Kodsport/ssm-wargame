package goa

import . "goa.design/goa/v3/dsl"

var _ = Service("user", func() {
	HTTP(func() {
		Path("/user")
	})
	Security(JWTAuth)
	Method("GetSelf", func() {
		Payload(func() {
			Extend(TokenPayload)
		})
		Result(func() {
			Extend(User)
			Attribute("school_name", String)
			Attribute("onboarding_done", Boolean, func() {
				Example(true)
			})
			Required("onboarding_done")
		})
		HTTP(func() {
			GET("/self")
			Response(StatusOK)
		})
	})
	Method("UpdateSelf", func() {
		Error("InvalidUsername")
		HTTP(func() {
			Response("InvalidUsername", StatusBadRequest)

		})
		Payload(func() {
			Extend(TokenPayload)
			Reference(User)
			Attribute("full_name")
		})
		HTTP(func() {
			POST("/self")
			Response(StatusOK)
		})
	})
	Method("CompleteOnboarding", func() {
		Payload(func() {
			Extend(TokenPayload)
		})
		HTTP(func() {
			POST("/complete_onboarding")
			Response(StatusOK)
		})
	})
	Method("JoinSchool", func() {
		Payload(func() {
			Extend(TokenPayload)
			Attribute("school_id", String, func() {
				Example("b6d5e5eb-3272-4795-ba1c-73a8efc1cf19")
			})
			Required("school_id")
		})
		HTTP(func() {
			POST("/join_school")
			Response(StatusOK)
		})
	})
	Method("LeaveSchool", func() {
		Payload(func() {
			Extend(TokenPayload)
		})
		HTTP(func() {
			POST("/leave_school")
			Response(StatusOK)
		})
	})
	Method("SearchSchools", func() {
		Payload(func() {
			Extend(TokenPayload)
			Attribute("q", String, func() {
				Example("engelbrektssko")
			})
			Attribute("university", Boolean, func() {
				Example(true)
			})
			Required("q")
		})
		Result(ArrayOf(School))
		HTTP(func() {
			GET("/schools")
			Response(StatusOK)
			Param("q")
			Param("university")
		})
	})

})
