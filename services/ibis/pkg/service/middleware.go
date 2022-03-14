package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"ibis/pkg/entities"
)

// Middleware describes a service middleware.
type Middleware func(IbisService) IbisService

type loggedIbisService struct {
	next   IbisService
	logger log.Logger
}

func LoggedIbisService(logger log.Logger) Middleware {
	return func(s IbisService) IbisService {
		return &loggedIbisService{
			next:   s,
			logger: logger,
		}
	}
}

func (s *loggedIbisService) CreateFunction(
	ctx context.Context, ownerID, name string,
) (Function entities.Function, err error) {
	Function, err = s.next.CreateFunction(ctx, ownerID, name)
	s.logger.Log("method", "CreateFunction", "ownerID", ownerID, "FunctionName", name, "error", err)
	return
}

func (s *loggedIbisService) GetFunction(ctx context.Context, FunctionID string) (Function entities.Function, err error) {
	Function, err = s.next.GetFunction(ctx, FunctionID)
	s.logger.Log("method", "GetFunction", "FunctionID", FunctionID, "error", err)
	return
}

func (s *loggedIbisService) GetFunctionByName(
	ctx context.Context, ownerID, FunctionName string,
) (Function entities.Function, err error) {
	Function, err = s.next.GetFunctionByName(ctx, ownerID, FunctionName)
	s.logger.Log("method", "GetFunctionByName", "ownerID", ownerID, "FunctionName", FunctionName, "error", err)
	return
}

func (s *loggedIbisService) UpdateFunction(
	ctx context.Context, FunctionID string, param entities.UpdateFunctionParam,
) (Function entities.Function, err error) {
	Function, err = s.next.UpdateFunction(ctx, FunctionID, param)
	s.logger.Log("method", "UpdateFunction", "FunctionID", FunctionID, "param", param, "error", err)
	return
}

func (s *loggedIbisService) ListFunctions(ctx context.Context, ownerID string) (Functions []entities.Function, err error) {
	Functions, err = s.next.ListFunctions(ctx, ownerID)
	s.logger.Log("method", "ListFunctions", "ownerID", ownerID, "error", err)
	return
}
