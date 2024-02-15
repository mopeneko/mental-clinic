// Code generated by goa v3.14.6, DO NOT EDIT.
//
// HTTP request path constructors for the general service.
//
// Command:
// $ goa gen github.com/mopeneko/mental-clinic/api/design -o api

package server

// HealthCheckGeneralPath returns the URL path to the general service healthCheck HTTP endpoint.
func HealthCheckGeneralPath() string {
	return "/health_check"
}

// AuthGeneralPath returns the URL path to the general service auth HTTP endpoint.
func AuthGeneralPath() string {
	return "/auth"
}
