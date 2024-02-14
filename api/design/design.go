package design

import . "goa.design/goa/v3/dsl"

var _ = API("general", func() {
	Title("mental-clinic.social General Service")
	Description("Service for general users")
	Server("general", func() {
		Host("local", func() {
			URI("http://localhost:8080")
		})
	})
})

var _ = Service("general", func() {
	Description("Service for general users")

	Method("healthCheck", func() {
		Result(String)

		HTTP(func() {
			GET("/health_check")

			Response(StatusOK)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})