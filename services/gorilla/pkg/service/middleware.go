package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	uuid "github.com/satori/go.uuid"

	"gorilla/pkg/entities"
)

// Middleware describes a service middleware.
type Middleware func(GorillaService) GorillaService

type loggedGorillaService struct {
	next   GorillaService
	logger log.Logger
}

func LoggedGorillaService(logger log.Logger) Middleware {
	return func(s GorillaService) GorillaService {
		return &loggedGorillaService{
			next:   s,
			logger: logger,
		}
	}
}

func (s *loggedGorillaService) AddDeltas(ctx context.Context, deltas []entities.Delta) (err error) {
	start := time.Now()
	err = s.next.AddDeltas(ctx, deltas)
	duration := time.Since(start)
	s.logger.Log("method", "AddDeltas", "deltas_len", len(deltas), "latency_human", duration, "error", err)
	return
}

func (s *loggedGorillaService) GetDeltas(
	ctx context.Context, ownerID uuid.UUID, objectID, objectType string, firstDate, lastDate time.Time, useCategories bool,
) (deltas []entities.Delta, err error) {
	start := time.Now()
	deltas, err = s.next.GetDeltas(ctx, ownerID, objectID, objectType, firstDate, lastDate, useCategories)
	duration := time.Since(start)
	s.logger.Log("method", "GetDeltas", "owner_id", ownerID.String(), "object_id", objectID,
		"object_type", objectType, "first_date", firstDate, "last_date", lastDate, "use_categories", useCategories,
		"deltas_len", len(deltas), "latency_human", duration, "error", err)
	return
}

func (s *loggedGorillaService) GetBalance(ctx context.Context, ownerID uuid.UUID) (balance float64, err error) {
	start := time.Now()
	balance, err = s.next.GetBalance(ctx, ownerID)
	duration := time.Since(start)
	s.logger.Log("method", "GetBalance", "owner_id", ownerID.String(), "balance", balance,
		"latency_human", duration, "error", err)
	return
}
