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
	Status string `form:"status" json:"status" xml:"status"`
}

// AuthResponseBody is the type of the "general" service "auth" endpoint HTTP
// response body.
type AuthResponseBody struct {
	// URL for authorization
	AuthURL string `form:"authURL" json:"authURL" xml:"authURL"`
}

// NewHealthCheckResponseBody builds the HTTP response body from the result of
// the "healthCheck" endpoint of the "general" service.
func NewHealthCheckResponseBody(res *general.HealthCheckResult) *HealthCheckResponseBody {
	body := &HealthCheckResponseBody{
		Status: res.Status,
	}
	return body
}

// NewAuthResponseBody builds the HTTP response body from the result of the
// "auth" endpoint of the "general" service.
func NewAuthResponseBody(res *general.AuthResult) *AuthResponseBody {
	body := &AuthResponseBody{
		AuthURL: res.AuthURL,
	}
	return body
}
