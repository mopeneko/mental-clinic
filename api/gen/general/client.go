// Code generated by goa v3.14.6, DO NOT EDIT.
//
// general client
//
// Command:
// $ goa gen github.com/mopeneko/mental-clinic/api/design -o api

package general

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "general" service client.
type Client struct {
	HealthCheckEndpoint goa.Endpoint
}

// NewClient initializes a "general" service client given the endpoints.
func NewClient(healthCheck goa.Endpoint) *Client {
	return &Client{
		HealthCheckEndpoint: healthCheck,
	}
}

// HealthCheck calls the "healthCheck" endpoint of the "general" service.
func (c *Client) HealthCheck(ctx context.Context) (res *HealthCheckResult, err error) {
	var ires any
	ires, err = c.HealthCheckEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*HealthCheckResult), nil
}
