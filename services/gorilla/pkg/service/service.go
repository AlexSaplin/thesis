package service

import (
	"context"
	"gorilla/pkg/db"
	"time"

	uuid "github.com/satori/go.uuid"

	"gorilla/pkg/entities"
)

// GorillaService describes the service.
type GorillaService interface {
	// Add your methods here
	AddDeltas(ctx context.Context, deltas []entities.Delta) (err error)
	GetDeltas(
		ctx context.Context, ownerID uuid.UUID, objectID, objectType string, firstDate, lastDate time.Time, useCategories bool,
	) (deltas []entities.Delta, err error)
	GetBalance(ctx context.Context, ownerID uuid.UUID) (balance float64, err error)
}

type gorillaService struct {
	db db.DeltaDB
}

func (b *gorillaService) AddDeltas(ctx context.Context, deltas []entities.Delta) (err error) {
	return b.db.AddDeltas(ctx, deltas)
}
func (b *gorillaService) GetDeltas(ctx context.Context, ownerID uuid.UUID, objectID, objectType string, firstDate time.Time, lastDate time.Time, useCategories bool) (deltas []entities.Delta, err error) {
	return b.db.GetDeltas(ctx, ownerID, objectID, objectType, firstDate, lastDate, useCategories)
}
func (b *gorillaService) GetBalance(ctx context.Context, ownerID uuid.UUID) (balance float64, err error) {
	return b.db.GetBalance(ctx, ownerID)
}

// NewBasicGorillaService returns a naive, stateless implementation of GorillaService.
func NewBasicGorillaService(deltaDB db.DeltaDB) GorillaService {
	return &gorillaService{
		db: deltaDB,
	}
}

// New returns a GorillaService with all of the expected middleware wired in.
func New(deltaDB db.DeltaDB, middleware []Middleware) GorillaService {
	var svc GorillaService = NewBasicGorillaService(deltaDB)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
