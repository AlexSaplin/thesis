package db

import (
	"database/sql"
	"github.com/go-kit/kit/log"
	"github.com/lib/pq"
	"github.com/mattn/go-nulltype"
	uuid "github.com/satori/go.uuid"

	"ibis/pkg/entities"
	"ibis/pkg/errors"
)

type functionResult struct {
	ID          uuid.UUID  `db:"id"`
	OwnerID     uuid.UUID  `db:"owner_id"`
	Name        string     `db:"name"`
	State       FunctionState `db:"state"`
	Error       *string    `db:"err"`
	ImageURL    *string    `db:"image"`
	CodePath    *string    `db:"code_path"`
	Metadata    *string    `db:"meta"`
}

func makeFunctionEntity(mr functionResult) entities.Function {
	var (
		imageStr nulltype.NullString
		codePathStr nulltype.NullString
		metaStr nulltype.NullString
		errStr nulltype.NullString
	)

	if mr.ImageURL != nil {
		imageStr = nulltype.NullStringOf(*mr.ImageURL)
	}
	if mr.CodePath != nil {
		codePathStr = nulltype.NullStringOf(*mr.CodePath)
	}
	if mr.Metadata != nil {
		metaStr = nulltype.NullStringOf(*mr.Metadata)
	}
	if mr.Error != nil {
		errStr = nulltype.NullStringOf(*mr.Error)
	}

	return entities.Function{
		ID:       mr.ID,
		OwnerID:  mr.OwnerID,
		State:    parseFunctionState(mr.State),
		ImageURL: imageStr,
		CodePath: codePathStr,
		Metadata: metaStr,
		Name:     mr.Name,
		Error:    errStr,
	}
}

type FunctionState string

func serializeFunctionState(s entities.FunctionState) FunctionState {
	switch s {
	case entities.FunctionStateInit:
		return "INIT"
	case entities.FunctionStateProcessing:
		return "PROCESSING"
	case entities.FunctionStateReady:
		return "READY"
	case entities.FunctionStateInvalid:
		return "INVALID"
	case entities.FunctionStateDeleted:
		return "DELETED"
	default:
		return "UNKNOWN"
	}
}

func parseFunctionState(s FunctionState) entities.FunctionState {
	switch s {
	case "INIT":
		return entities.FunctionStateInit
	case "PROCESSING":
		return entities.FunctionStateProcessing
	case "READY":
		return entities.FunctionStateReady
	case "INVALID":
		return entities.FunctionStateInvalid
	case "DELETED":
		return entities.FunctionStateDeleted
	default:
		return entities.FunctionStateUnknown
	}
}

func mapDBError(logger log.Logger, err error) error {
	if err == nil {
		return nil
	}
	if err == sql.ErrNoRows {
		return errors.ErrFunctionNotFound
	}

	pqErr, ok := err.(*pq.Error)
	if !ok {
		logger.Log("msg", "unknown error format", "err", err)
		return errors.ErrInternal
	}

	switch pqErr.Code {
	case "23505": // unique_violation
		return errors.ErrFunctionExists
	default:
		logger.Log("msg", "unknown postgres error code", "code", pqErr.Code, "err", err)
		return errors.ErrInternal
	}
}

