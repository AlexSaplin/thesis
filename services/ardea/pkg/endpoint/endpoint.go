package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mattn/go-nulltype"

	"ardea/pkg/entities"
	"ardea/pkg/service"
)

type Failure interface {
	Failed() error
}

type CreateModelRequest struct {
	OwnerID     string                 `json:"owner_id"`
	InputShape  [][]nulltype.NullInt64 `json:"input_shape"`
	OutputShape [][]nulltype.NullInt64 `json:"output_shape"`
	Name        string                 `json:"name"`
	ValueType   entities.ValueType     `json:"value_type"`
}

type CreateModelResponse struct {
	Model entities.Model `json:"model"`
	Err   error          `json:"err"`
}

func MakeCreateModelEndpoint(s service.ArdeaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateModelRequest)
		model, err := s.CreateModel(ctx, req.OwnerID, req.Name, req.InputShape, req.OutputShape, req.ValueType)
		return CreateModelResponse{
			Model: model,
			Err:   err,
		}, nil
	}
}

func (r CreateModelResponse) Failed() error {
	return r.Err
}

type GetModelRequest struct {
	ModelID string `json:"model_id"`
}

type GetModelResponse struct {
	Model entities.Model `json:"model"`
	Err   error          `json:"err"`
}

func MakeGetModelEndpoint(s service.ArdeaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetModelRequest)
		model, err := s.GetModel(ctx, req.ModelID)
		return GetModelResponse{
			Model: model,
			Err:   err,
		}, nil
	}
}

func (r GetModelResponse) Failed() error {
	return r.Err
}

type GetModelByNameRequest struct {
	OwnerID string `json:"owner_id"`
	Name    string `json:"name"`
}

type GetModelByNameResponse struct {
	Model entities.Model `json:"model"`
	Err   error          `json:"err"`
}

func MakeGetModelByNameEndpoint(s service.ArdeaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetModelByNameRequest)
		model, err := s.GetModelByName(ctx, req.OwnerID, req.Name)
		return GetModelByNameResponse{
			Model: model,
			Err:   err,
		}, nil
	}
}

func (r GetModelByNameResponse) Failed() error {
	return r.Err
}

type UpdateModelStateRequest struct {
	ModelID string              `json:"model_id"`
	State   entities.ModelState `json:"state"`
	ErrStr  nulltype.NullString `json:"err_str"`
}

type UpdateModelStateResponse struct {
	Model entities.Model `json:"model"`
	Err   error          `json:"err"`
}

func MakeUpdateModelStateEndpoint(s service.ArdeaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateModelStateRequest)
		model, err := s.UpdateModelState(ctx, req.ModelID, req.State, req.ErrStr)
		return UpdateModelStateResponse{
			Model: model,
			Err:   err,
		}, nil
	}
}

func (r UpdateModelStateResponse) Failed() error {
	return r.Err
}

type UpdateModelPathRequest struct {
	ModelID string `json:"model_id"`
	Path    string `json:"path"`
}

type UpdateModelPathResponse struct {
	Model entities.Model `json:"model"`
	Err   error          `json:"err"`
}

func MakeUpdateModelPathEndpoint(s service.ArdeaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateModelPathRequest)
		model, err := s.UpdateModelPath(ctx, req.ModelID, req.Path)
		return UpdateModelPathResponse{
			Model: model,
			Err:   err,
		}, nil
	}
}

type ListModelsRequest struct {
	OwnerID string `json:"owner_id"`
}

type ListModelsResponse struct {
	Models []entities.Model `json:"models"`
	Err    error            `json:"err"`
}

func MakeListModelsEndpoint(s service.ArdeaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListModelsRequest)
		models, err := s.ListModels(ctx, req.OwnerID)
		return ListModelsResponse{
			Models: models,
			Err:    err,
		}, nil
	}
}

func (r UpdateModelPathResponse) Failed() error {
	return r.Err
}

func (e Endpoints) CreateModel(
	ctx context.Context, OwnerID string, inputShape, outputShape [][]nulltype.NullInt64, valueType entities.ValueType,
) (model entities.Model, err error) {
	request := CreateModelRequest{
		InputShape:  inputShape,
		OutputShape: outputShape,
		OwnerID:     OwnerID,
		ValueType:   valueType,
	}
	response, err := e.CreateModelEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateModelResponse).Model, response.(CreateModelResponse).Err
}

func (e Endpoints) GetModel(ctx context.Context, modelID string) (model entities.Model, err error) {
	request := GetModelRequest{ModelID: modelID}
	response, err := e.GetModelEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetModelResponse).Model, response.(GetModelResponse).Err
}

func (e Endpoints) UpdateModelState(
	ctx context.Context, modelID string, state entities.ModelState, errStr nulltype.NullString,
) (model entities.Model, err error) {
	request := UpdateModelStateRequest{
		ModelID: modelID,
		State:   state,
		ErrStr:  errStr,
	}
	response, err := e.UpdateModelStateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateModelStateResponse).Model, response.(UpdateModelStateResponse).Err
}

func (e Endpoints) UpdateModelPath(ctx context.Context, modelID string, path string) (model entities.Model, err error) {
	request := UpdateModelPathRequest{
		ModelID: modelID,
		Path:    path,
	}
	response, err := e.UpdateModelPathEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateModelPathResponse).Model, response.(UpdateModelPathResponse).Err
}

func (e Endpoints) ListModels(ctx context.Context, ownerID string) (models []entities.Model, err error) {
	request := ListModelsRequest{
		OwnerID: ownerID,
	}
	response, err := e.ListModelsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListModelsResponse).Models, response.(ListModelsResponse).Err
}
