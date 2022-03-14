package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/satori/go.uuid"
)

// Middleware describes a service middleware.
type Middleware func(RhinoService) RhinoService

type loggedRhinoService struct {
	next   RhinoService
	logger log.Logger
}

func LoggedRhinoService(logger log.Logger) Middleware {
	return func(s RhinoService) RhinoService {
		return &loggedRhinoService{
			next: s,
			logger: logger,
		}
	}
}

func (s *loggedRhinoService) Run(
	ctx context.Context, modelID uuid.UUID, in []byte,
) (result []byte, err error) {
	start := time.Now()
	result, err = s.next.Run(ctx, modelID, in)
	duration := time.Since(start)
	_ = s.logger.Log("method", "Run", "functionID", modelID, "latency_human", duration, "error", err)
	return
}
