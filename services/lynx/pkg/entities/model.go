package entities

import (
	"github.com/mattn/go-nulltype"
	"github.com/satori/go.uuid"
)

type ModelState uint8

const (
	ModelStateUnknown ModelState = iota
	ModelStateInit
	ModelStateProcessing
	ModelStateReady
	ModelStateInvalid
	ModelStateDeleted
)

func (s ModelState) String() string {
	switch s {
	case ModelStateInit:
		return "CREATED"
	case ModelStateProcessing:
		return "PROCESSING"
	case ModelStateReady:
		return "READY"
	case ModelStateInvalid:
		return "ERROR"
	case ModelStateDeleted:
		return "DELETED"
	default:
		return "UNKNOWN"
	}
}

type Model struct {
	ID          uuid.UUID
	Name        string
	OwnerID     uuid.UUID
	ValueType   ValueType
	State       ModelState
	InputShape  IOShape
	OutputShape IOShape
	Path        string
	Error       nulltype.NullString
}
