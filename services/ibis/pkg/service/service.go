package service

import (
	"context"

	uuid "github.com/satori/go.uuid"

	"ibis/pkg/db"
	"ibis/pkg/entities"
)

// IbisService describes the service.
type IbisService interface {
	CreateFunction(ctx context.Context, ownerID, name string) (Function entities.Function, err error)
	GetFunction(ctx context.Context, FunctionID string) (entities.Function, error)
	GetFunctionByName(ctx context.Context, ownerID, FunctionName string) (entities.Function, error)
	UpdateFunction(ctx context.Context, FunctionID string, param entities.UpdateFunctionParam) (entities.Function, error)
	ListFunctions(ctx context.Context, ownerID string) ([]entities.Function, error)
}

type IbisServiceImpl struct {
	db db.FunctionDB
}


// NewIbisService returns an implementation of IbisService.
func NewIbisService(FunctionDB db.FunctionDB) IbisService {
	return &IbisServiceImpl{
		db: FunctionDB,
	}
}

// New returns a ArdeaService with all of the expected middleware wired in.
func New(FunctionDB db.FunctionDB, middleware []Middleware) IbisService {
	var svc = NewIbisService(FunctionDB)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (s *IbisServiceImpl) CreateFunction(
	ctx context.Context, ownerStr, name string,
) (Function entities.Function, err error) {
	var ownerID uuid.UUID
	ownerID, err = uuid.FromString(ownerStr)
	if err != nil {
		return
	}
	functionID := uuid.NewV4()

	return s.db.CreateFunction(ctx, functionID, ownerID, name)
}
func (s *IbisServiceImpl) GetFunction(ctx context.Context, FunctionStr string) (function entities.Function, err error) {
	functionID, err := uuid.FromString(FunctionStr)
	if err != nil {
		return
	}
	return s.db.GetFunction(ctx, functionID)
}

func (s *IbisServiceImpl) GetFunctionByName(
	ctx context.Context, ownerIDStr, FunctionName string,
) (Function entities.Function, err error) {
	ownerID, err := uuid.FromString(ownerIDStr)
	if err != nil {
		return
	}
	return s.db.GetFunctionByName(ctx, ownerID, FunctionName)
}

func (s *IbisServiceImpl) UpdateFunction(
	ctx context.Context, FunctionStr string, param entities.UpdateFunctionParam,
) (function entities.Function, err error) {
	functionID, err := uuid.FromString(FunctionStr)

	return s.db.UpdateFunction(ctx, functionID, param)
}

func (s *IbisServiceImpl) ListFunctions(ctx context.Context, ownerIDStr string) (result []entities.Function, err error) {
	ownerID, err := uuid.FromString(ownerIDStr)
	if err != nil {
		return
	}
	return s.db.ListFunctions(ctx, ownerID)
}