package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/satori/go.uuid"

	"hippo/pkg/entities"
)

// Middleware describes a service middleware.
type Middleware func(HippoService) HippoService

type loggedHippoService struct {
	next   HippoService
	logger log.Logger
}

func LoggedHippoService(logger log.Logger) Middleware {
	return func(s HippoService) HippoService{
		return &loggedHippoService{
			next: s,
			logger: logger,
		}
	}
}

func (s *loggedHippoService) Run(
	ctx context.Context, modelID uuid.UUID, tensor entities.TensorList,
) (result entities.TensorList, err error) {
	start := time.Now()
	result, err = s.next.Run(ctx, modelID, tensor)
	duration := time.Since(start)
	_ = s.logger.Log("method", "Run", "modelID", modelID, "latency_human", duration, "error", err)
	return
}
