package goa

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("wargame", func() {
	Title("Wargame Service")
	Description("HTTP service for the SSM Wargame")
})

var JWTAuth = JWTSecurity("jwt")
