package repository

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type Auth struct {
	provider *oidc.Provider
	cfg      *oauth2.Config
}

func NewAuth(ctx context.Context, domain string) (*Auth, error) {
	provider, err := oidc.NewProvider(context.Background(), fmt.Sprintf("https://%s/", domain))
	if err != nil {
		log.Fatalf("failed to create OpenID provider: %+v", err)
	}

	cfg := &oauth2.Config{
		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("AUTH0_CALLBACK_URL"),
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Auth{
		provider: provider,
		cfg:      cfg,
	}, nil
}

func (r *Auth) AuthURL() (string, error) {
	b := make([]byte, 128)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to generate state: %w", err)
	}
	state := base64.StdEncoding.EncodeToString(b)

	return r.cfg.AuthCodeURL(state), nil
}
