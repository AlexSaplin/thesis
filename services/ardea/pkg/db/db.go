package db

import (
	"context"
	"github.com/mattn/go-nulltype"

	uuid "github.com/satori/go.uuid"

	"ardea/pkg/entities"
)

type modelUpdateType uint8

const (
	modelUpdateTypePath modelUpdateType = iota
	modelUpdateTypeState
	modelUpdateTypeErrStr
)

type ModelUpdateParam struct {
	field modelUpdateType
	value interface{}
}

func PathUpdateParam(path string) ModelUpdateParam {
	return ModelUpdateParam{
		field: modelUpdateTypePath,
		value: path,
	}
}

func StateUpdateParam(state entities.ModelState) ModelUpdateParam {
	return ModelUpdateParam{
		field: modelUpdateTypeState,
		value: serializeModelState(state),
	}
}

func ErrStrUpdateParam(errStr string) ModelUpdateParam {
	return ModelUpdateParam{
		field: modelUpdateTypeErrStr,
		value: errStr,
	}
}

type ModelDB interface {
	CreateModel(
		ctx context.Context, modelID, ownerID uuid.UUID, name string,
		inputShape, outputShape [][]nulltype.NullInt64, valueType entities.ValueType,
	) (model entities.Model, err error)
	GetModel(ctx context.Context, modelID uuid.UUID) (entities.Model, error)
	GetModelByName(ctx context.Context, ownerID uuid.UUID, name string) (entities.Model, error)
	UpdateModel(ctx context.Context, modelID uuid.UUID, params ...ModelUpdateParam) (entities.Model, error)
	ListModels(ctx context.Context, ownerID uuid.UUID) ([]entities.Model, error)
}
