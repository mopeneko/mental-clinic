// Code generated by goa v3.14.6, DO NOT EDIT.
//
// general HTTP client types
//
// Command:
// $ goa gen github.com/mopeneko/mental-clinic/api/design -o api

package client

import (
	general "github.com/mopeneko/mental-clinic/api/gen/general"
	goa "goa.design/goa/v3/pkg"
)

// HealthCheckResponseBody is the type of the "general" service "healthCheck"
// endpoint HTTP response body.
type HealthCheckResponseBody struct {
	// Status of the service
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// AuthResponseBody is the type of the "general" service "auth" endpoint HTTP
// response body.
type AuthResponseBody struct {
	// URL for authorization
	AuthURL *string `form:"authURL,omitempty" json:"authURL,omitempty" xml:"authURL,omitempty"`
}

// NewHealthCheckResultOK builds a "general" service "healthCheck" endpoint
// result from a HTTP "OK" response.
func NewHealthCheckResultOK(body *HealthCheckResponseBody) *general.HealthCheckResult {
	v := &general.HealthCheckResult{
		Status: *body.Status,
	}

	return v
}

// NewAuthResultOK builds a "general" service "auth" endpoint result from a
// HTTP "OK" response.
func NewAuthResultOK(body *AuthResponseBody) *general.AuthResult {
	v := &general.AuthResult{
		AuthURL: *body.AuthURL,
	}

	return v
}

// ValidateHealthCheckResponseBody runs the validations defined on
// HealthCheckResponseBody
func ValidateHealthCheckResponseBody(body *HealthCheckResponseBody) (err error) {
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	return
}

// ValidateAuthResponseBody runs the validations defined on AuthResponseBody
func ValidateAuthResponseBody(body *AuthResponseBody) (err error) {
	if body.AuthURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("authURL", "body"))
	}
	return
}
