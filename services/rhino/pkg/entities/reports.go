package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type RunReport struct {
	OwnerID      uuid.UUID
	FunctionID   uuid.UUID
	RunDuration  time.Duration
	LoadDuration time.Duration
	Timestamp    time.Time
}
