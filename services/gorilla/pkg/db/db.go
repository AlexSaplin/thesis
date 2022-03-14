package db

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"gorilla/pkg/entities"
	"time"
)

type DeltaDB interface {
	AddDeltas(ctx context.Context, deltas []entities.Delta) (err error)
	GetDeltas(
		ctx context.Context, ownerID uuid.UUID, objectID, objectType string, firstDate, lastDate time.Time, useCategories bool,
	) (deltas []entities.Delta, err error)
	GetBalance(ctx context.Context, ownerID uuid.UUID) (balance float64, err error)
}
