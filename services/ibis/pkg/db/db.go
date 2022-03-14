package db

import (
	"context"
	uuid "github.com/satori/go.uuid"

	"ibis/pkg/entities"
)

type FunctionUpdateType uint8

type FunctionDB interface {
	CreateFunction(
		ctx context.Context, FunctionID, ownerID uuid.UUID, name string,
	) (Function entities.Function, err error)
	GetFunction(ctx context.Context, FunctionID uuid.UUID) (entities.Function, error)
	GetFunctionByName(ctx context.Context, ownerID uuid.UUID, name string) (entities.Function, error)
	UpdateFunction(ctx context.Context, FunctionID uuid.UUID, param entities.UpdateFunctionParam) (entities.Function, error)
	ListFunctions(ctx context.Context, ownerID uuid.UUID) ([]entities.Function, error)
}
