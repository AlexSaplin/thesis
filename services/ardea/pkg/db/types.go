package db

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"github.com/go-kit/kit/log"
	"github.com/lib/pq"
	"github.com/mattn/go-nulltype"
	uuid "github.com/satori/go.uuid"

	"ardea/pkg/entities"
	"ardea/pkg/errors"
)

type ModelResult struct {
	ID          uuid.UUID  `db:"id"`
	OwnerID     uuid.UUID  `db:"owner_id"`
	ValueType   valueType  `db:"value_type"`
	Name        string     `db:"name"`
	State       modelState `db:"state"`
	InputShape  ioShape    `db:"input_shape"`
	OutputShape ioShape    `db:"output_shape"`
	Path        *string    `db:"path"`
	Error       *string    `db:"err"`
}

func makeModelEntity(mr ModelResult) entities.Model {
	var pathStr string
	if mr.Path != nil {
		pathStr = *mr.Path
	}
	var errStr nulltype.NullString
	if mr.Error != nil {
		errStr = nulltype.NullStringOf(*mr.Error)
	}
	return entities.Model{
		ID:          mr.ID,
		OwnerID:     mr.OwnerID,
		ValueType:   parseValueType(mr.ValueType),
		State:       parseModelState(mr.State),
		InputShape:  mr.InputShape,
		OutputShape: mr.OutputShape,
		Path:        pathStr,
		Name:        mr.Name,
		Error:       errStr,
	}
}

type modelState string

func serializeModelState(s entities.ModelState) modelState {
	switch s {
	case entities.ModelStateInit:
		return "INIT"
	case entities.ModelStateProcessing:
		return "PROCESSING"
	case entities.ModelStateReady:
		return "READY"
	case entities.ModelStateInvalid:
		return "INVALID"
	case entities.ModelStateDeleted:
		return "DELETED"
	default:
		return "UNKNOWN"
	}
}

func parseModelState(s modelState) entities.ModelState {
	switch s {
	case "INIT":
		return entities.ModelStateInit
	case "PROCESSING":
		return entities.ModelStateProcessing
	case "READY":
		return entities.ModelStateReady
	case "INVALID":
		return entities.ModelStateInvalid
	case "DELETED":
		return entities.ModelStateDeleted
	default:
		return entities.ModelStateUnknown
	}
}

type valueType string

func serializeValueType(t entities.ValueType) valueType {
	switch t {
	case entities.ValueTypeUInt8:
		return "UINT8"
	case entities.ValueTypeInt8:
		return "INT8"
	case entities.ValueTypeFloat16:
		return "FLOAT16"
	case entities.ValueTypeUInt16:
		return "UINT16"
	case entities.ValueTypeInt16:
		return "INT16"
	case entities.ValueTypeFloat32:
		return "FLOAT32"
	case entities.ValueTypeUInt32:
		return "UINT32"
	case entities.ValueTypeInt32:
		return "INT32"
	case entities.ValueTypeFloat64:
		return "FLOAT64"
	case entities.ValueTypeUint64:
		return "UINT64"
	case entities.ValueTypeInt64:
		return "INT64"
	case entities.ValueTypeComplex64:
		return "COMPLEX64"
	case entities.ValueTypeComplex128:
		return "COMPLEX128"
	default:
		return "UNKNOWN"
	}
}

func parseValueType(t valueType) (result entities.ValueType) {
	switch t {
	case "UINT8":
		result = entities.ValueTypeUInt8
	case "INT8":
		result = entities.ValueTypeInt8
	case "FLOAT16":
		result = entities.ValueTypeFloat16
	case "UINT16":
		result = entities.ValueTypeUInt16
	case "INT16":
		result = entities.ValueTypeInt16
	case "FLOAT32":
		result = entities.ValueTypeFloat32
	case "UINT32":
		result = entities.ValueTypeUInt32
	case "INT32":
		result = entities.ValueTypeInt32
	case "FLOAT64":
		result = entities.ValueTypeFloat64
	case "UINT64":
		result = entities.ValueTypeUint64
	case "INT64":
		result = entities.ValueTypeInt64
	case "COMPLEX64":
		result = entities.ValueTypeComplex64
	case "COMPLEX128":
		result = entities.ValueTypeComplex128
	default:
		result = entities.ValueTypeUnknown
	}
	return
}

type ioShape [][]nulltype.NullInt64

func (i ioShape) Value() (driver.Value, error) {
	val, err := json.Marshal(i)
	return string(val), err
}

func (i *ioShape) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), i)
}

func mapDBError(logger log.Logger, err error) error {
	if err == nil {
		return nil
	}
	if err == sql.ErrNoRows {
		return errors.ErrModelNotFound
	}

	pqErr, ok := err.(*pq.Error)
	if !ok {
		logger.Log("msg", "unknown error format", "err", err)
		return errors.ErrInternal
	}

	switch pqErr.Code {
	case "23505": // unique_violation
		return errors.ErrModelExists
	default:
		logger.Log("msg", "unknown postgres error code", "code", pqErr.Code, "err", err)
		return errors.ErrInternal
	}
}
