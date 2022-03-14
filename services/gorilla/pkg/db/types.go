package db

import (
	uuid "github.com/satori/go.uuid"
	"gorilla/pkg/entities"
	"time"
)

type DeltaResult struct {
	OwnerID    uuid.UUID `db:"owner_id"`
	ObjectID   uuid.UUID `db:"object_id"`
	ObjectType string    `db:"object_type"`
	Category   string    `db:"category"`
	Date       time.Time `db:"date"`
	Balance    float64   `db:"balance"`
}

func makeDeltaEntity(dr DeltaResult) entities.Delta {
	return entities.Delta{
		OwnerID:    dr.OwnerID,
		ObjectID:   dr.ObjectID.String(),
		ObjectType: dr.ObjectType,
		Category:   dr.Category,
		Date:       dr.Date,
		Balance:    dr.Balance,
	}
}

type BalanceResult struct {
	Balance float64 `db:"balance"`
}
