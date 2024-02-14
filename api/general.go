package generalapi

import (
	"context"
	"log"

	general "github.com/mopeneko/mental-clinic/api/gen/general"
)

// general service example implementation.
// The example methods log the requests and return zero values.
type generalsrvc struct {
	logger *log.Logger
}

// NewGeneral returns the general service implementation.
func NewGeneral(logger *log.Logger) general.Service {
	return &generalsrvc{logger}
}

// HealthCheck implements healthCheck.
func (s *generalsrvc) HealthCheck(ctx context.Context) (res string, err error) {
	s.logger.Print("general.healthCheck")
	return
}
