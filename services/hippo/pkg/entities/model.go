package entities

import "github.com/satori/go.uuid"

type ModelState uint8

const (
	ModelStateUnknown ModelState = iota
	ModelStateInit
	ModelStateProcessing
	ModelStateReady
	ModelStateInvalid
	ModelStateDeleted
)

type Model struct {
	ID          uuid.UUID
	OwnerID     uuid.UUID
	ValueType   ValueType
	State       ModelState
	InputShape  IOShape
	OutputShape IOShape
	Path        string
}
