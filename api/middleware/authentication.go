package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	oidc "github.com/coreos/go-oidc"
	"github.com/mopeneko/mental-clinic/api/consts/ctxkeys"
)

type AuthenticationConfig struct {
	Domain       string
	ClientID     string
	ClientSecret string
	CallbackURL  string
}

type AuthenticationErrorResponse struct {
	Message string `json:"message"`
}

func Authentication(cfg *AuthenticationConfig) func(http.Handler) http.Handler {
	provider, err := oidc.NewProvider(context.Background(), fmt.Sprintf("https://%s/", cfg.Domain))
	if err != nil {
		log.Fatalf("failed to create OpenID provider: %+v", err)
	}
	oidcConfig := &oidc.Config{
		ClientID: cfg.ClientID,
	}

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			header := r.Header.Get("Authorization")
			fields := strings.Fields(header)
			if len(fields) != 2 || strings.ToLower(fields[0]) != "bearer" {
				res := &AuthenticationErrorResponse{
					Message: "Bearer token is missing.",
				}

				b, err := json.Marshal(res)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("Internal Server Error"))
					return
				}

				w.WriteHeader(http.StatusBadRequest)
				w.Write(b)
				return
			}

			token := fields[1]

			idToken, err := provider.Verifier(oidcConfig).Verify(ctx, token)
			if err != nil {
				res := &AuthenticationErrorResponse{
					Message: "Failed to vetify token.",
				}

				b, err := json.Marshal(res)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("Internal Server Error"))
					return
				}

				w.WriteHeader(http.StatusUnauthorized)
				w.Write(b)
			}

			ctx = context.WithValue(ctx, ctxkeys.IDToken, idToken)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
