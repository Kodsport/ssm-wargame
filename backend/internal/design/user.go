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
		})
		HTTP(func() {
			GET("/self")
			Response(StatusOK)
		})
	})
	Method("JoinSchool", func() {
		Payload(func() {
			Extend(TokenPayload)
			Attribute("school_id", Int, func() {
				Example(78433202)
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
			Required("q")
		})
		Result(ArrayOf(School))
		HTTP(func() {
			GET("/schools")
			Response(StatusOK)
			Param("q")
		})
	})
})
