// Code generated by goa v3.14.6, DO NOT EDIT.
//
// general HTTP server types
//
// Command:
// $ goa gen github.com/mopeneko/mental-clinic/api/design -o api

package server

import (
	general "github.com/mopeneko/mental-clinic/api/gen/general"
)

// HealthCheckResponseBody is the type of the "general" service "healthCheck"
// endpoint HTTP response body.
type HealthCheckResponseBody struct {
	// Status of the service
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// NewHealthCheckResponseBody builds the HTTP response body from the result of
// the "healthCheck" endpoint of the "general" service.
func NewHealthCheckResponseBody(res *general.HealthCheckResult) *HealthCheckResponseBody {
	body := &HealthCheckResponseBody{
		Status: res.Status,
	}
	return body
}
