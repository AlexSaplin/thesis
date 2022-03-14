package entities

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Delta struct {
	Date     time.Time // Unix timestamp of date in UTC
	Category string
	Balance  float64
	OwnerID  uuid.UUID
	ObjectID  string
	ObjectType string
}
