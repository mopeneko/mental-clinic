// Code generated by goa v3.14.6, DO NOT EDIT.
//
// general HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/mopeneko/mental-clinic/api/design -o api

package server

import (
	"context"
	"net/http"

	general "github.com/mopeneko/mental-clinic/api/gen/general"
	goahttp "goa.design/goa/v3/http"
)

// EncodeHealthCheckResponse returns an encoder for responses returned by the
// general healthCheck endpoint.
func EncodeHealthCheckResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*general.HealthCheckResult)
		enc := encoder(ctx, w)
		body := NewHealthCheckResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}
