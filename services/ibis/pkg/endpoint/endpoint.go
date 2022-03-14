package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mattn/go-nulltype"

	"ibis/pkg/entities"
	"ibis/pkg/service"
)

type Failure interface {
	Failed() error
}

type CreateFunctionRequest struct {
	OwnerID     string                 `json:"owner_id"`
	Name        string                 `json:"name"`
}

type CreateFunctionResponse struct {
	Function entities.Function `json:"Function"`
	Err   error             `json:"err"`
}

func MakeCreateFunctionEndpoint(s service.IbisService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateFunctionRequest)
		Function, err := s.CreateFunction(ctx, req.OwnerID, req.Name)
		return CreateFunctionResponse{
			Function: Function,
			Err:   err,
		}, nil
	}
}

func (r CreateFunctionResponse) Failed() error {
	return r.Err
}

type GetFunctionRequest struct {
	FunctionID string `json:"Function_id"`
}

type GetFunctionResponse struct {
	Function entities.Function `json:"Function"`
	Err   error             `json:"err"`
}

func MakeGetFunctionEndpoint(s service.IbisService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetFunctionRequest)
		Function, err := s.GetFunction(ctx, req.FunctionID)
		return GetFunctionResponse{
			Function: Function,
			Err:   err,
		}, nil
	}
}

func (r GetFunctionResponse) Failed() error {
	return r.Err
}

type GetFunctionByNameRequest struct {
	OwnerID string `json:"owner_id"`
	Name    string `json:"name"`
}

type GetFunctionByNameResponse struct {
	Function entities.Function `json:"Function"`
	Err   error                `json:"err"`
}

func MakeGetFunctionByNameEndpoint(s service.IbisService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetFunctionByNameRequest)
		Function, err := s.GetFunctionByName(ctx, req.OwnerID, req.Name)
		return GetFunctionByNameResponse{
			Function: Function,
			Err:   err,
		}, nil
	}
}

func (r GetFunctionByNameResponse) Failed() error {
	return r.Err
}

type UpdateFunctionRequest struct {
	FunctionID string                   `json:"Function_id"`
	Param   entities.UpdateFunctionParam `json:"param"`
	ErrStr  nulltype.NullString         `json:"err_str"`
}

type UpdateFunctionResponse struct {
	Function entities.Function `json:"Function"`
	Err   error                `json:"err"`
}

func MakeUpdateFunctionEndpoint(s service.IbisService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateFunctionRequest)
		Function, err := s.UpdateFunction(ctx, req.FunctionID, req.Param)
		return UpdateFunctionResponse{
			Function: Function,
			Err:   err,
		}, nil
	}
}

func (r UpdateFunctionResponse) Failed() error {
	return r.Err
}

type ListFunctionsRequest struct {
	OwnerID string `json:"owner_id"`
}

type ListFunctionsResponse struct {
	Functions []entities.Function `json:"Functions"`
	Err    error               `json:"err"`
}

func MakeListFunctionsEndpoint(s service.IbisService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListFunctionsRequest)
		Functions, err := s.ListFunctions(ctx, req.OwnerID)
		return ListFunctionsResponse{
			Functions: Functions,
			Err:    err,
		}, nil
	}
}

func (e Endpoints) CreateFunction(
	ctx context.Context, OwnerID string,
) (Function entities.Function, err error) {
	request := CreateFunctionRequest{
		OwnerID:     OwnerID,
	}
	response, err := e.CreateFunctionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateFunctionResponse).Function, response.(CreateFunctionResponse).Err
}

func (e Endpoints) GetFunction(ctx context.Context, FunctionID string) (Function entities.Function, err error) {
	request := GetFunctionRequest{FunctionID: FunctionID}
	response, err := e.GetFunctionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetFunctionResponse).Function, response.(GetFunctionResponse).Err
}

func (e Endpoints) UpdateFunction(
	ctx context.Context, FunctionID string, param entities.UpdateFunctionParam,
) (Function entities.Function, err error) {
	request := UpdateFunctionRequest{
		FunctionID: FunctionID,
		Param:   param,
	}
	response, err := e.UpdateFunctionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateFunctionResponse).Function, response.(UpdateFunctionResponse).Err
}

func (e Endpoints) ListFunctions(ctx context.Context, ownerID string) (Functions []entities.Function, err error) {
	request := ListFunctionsRequest{
		OwnerID: ownerID,
	}
	response, err := e.ListFunctionsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListFunctionsResponse).Functions, response.(ListFunctionsResponse).Err
}
