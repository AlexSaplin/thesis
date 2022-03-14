package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type RunReport struct {
	OwnerID   uuid.UUID
	ModelID   uuid.UUID
	Duration  time.Duration
	Timestamp time.Time
}
