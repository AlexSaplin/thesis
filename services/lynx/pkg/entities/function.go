package entities

import (
	"github.com/mattn/go-nulltype"
	"github.com/satori/go.uuid"
)


type FunctionState uint8

const (
	FunctionStateUnknown FunctionState = iota
	FunctionStateInit
	FunctionStateProcessing
	FunctionStateReady
	FunctionStateInvalid
	FunctionStateDeleted
)

func (s FunctionState) String() string {
	switch s {
	case FunctionStateInit:
		return "INIT"
	case FunctionStateProcessing:
		return "PROCESSING"
	case FunctionStateReady:
		return "READY"
	case FunctionStateInvalid:
		return "INVALID"
	case FunctionStateDeleted:
		return "DELETED"
	default:
		return "UNKNOWN"
	}
}

type Function struct {
	ID       uuid.UUID
	OwnerID  uuid.UUID
	Name     string
	State    FunctionState
	ImageURL nulltype.NullString
	CodePath nulltype.NullString
	Metadata nulltype.NullString
	Error    nulltype.NullString
}

type FunctionQuery struct {
	ID uuid.UUID
}