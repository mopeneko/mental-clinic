// Code generated by goa v3.14.6, DO NOT EDIT.
//
// general endpoints
//
// Command:
// $ goa gen github.com/mopeneko/mental-clinic/api/design -o api

package general

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "general" service endpoints.
type Endpoints struct {
	HealthCheck goa.Endpoint
}

// NewEndpoints wraps the methods of the "general" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		HealthCheck: NewHealthCheckEndpoint(s),
	}
}

// Use applies the given middleware to all the "general" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.HealthCheck = m(e.HealthCheck)
}

// NewHealthCheckEndpoint returns an endpoint function that calls the method
// "healthCheck" of service "general".
func NewHealthCheckEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return s.HealthCheck(ctx)
	}
}