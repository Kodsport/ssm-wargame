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
		})
		HTTP(func() {
			GET("/self")
			Response(StatusOK)
		})
	})
})
