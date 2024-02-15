package generalapi

import (
	"context"
	"fmt"
	"log"

	general "github.com/mopeneko/mental-clinic/api/gen/general"
	"github.com/mopeneko/mental-clinic/api/repository"
)

// general service example implementation.
// The example methods log the requests and return zero values.
type generalsrvc struct {
	logger *log.Logger
	repo   *repository.Auth
}

// NewGeneral returns the general service implementation.
func NewGeneral(logger *log.Logger, repo *repository.Auth) general.Service {
	return &generalsrvc{
		logger: logger,
		repo:   repo,
	}
}

// HealthCheck implements healthCheck.
func (s *generalsrvc) HealthCheck(ctx context.Context) (*general.HealthCheckResult, error) {
	return &general.HealthCheckResult{Status: "OK"}, nil
}

func (s *generalsrvc) Auth(ctx context.Context) (*general.AuthResult, error) {
	authURL, err := s.repo.AuthURL()
	if err != nil {
		return nil, fmt.Errorf("failed to generate auth url: %w", err)
	}
	return &general.AuthResult{
		AuthURL: authURL,
	}, nil
}
