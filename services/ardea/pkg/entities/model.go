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
		return "INIT"
	case ModelStateProcessing:
		return "PROCESSING"
	case ModelStateReady:
		return "READY"
	case ModelStateInvalid:
		return "INVALID"
	case ModelStateDeleted:
		return "DELETED"
	default:
		return "UNKNOWN"
	}
}

type Model struct {
	ID          uuid.UUID
	OwnerID     uuid.UUID
	ValueType   ValueType
	State       ModelState
	InputShape  [][]nulltype.NullInt64
	OutputShape [][]nulltype.NullInt64
	Path        string
	Name        string
	Error       nulltype.NullString
}
