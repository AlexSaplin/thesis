package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/mattn/go-nulltype"

	"ardea/pkg/entities"
)

// Middleware describes a service middleware.
type Middleware func(ArdeaService) ArdeaService

type loggedArdeaService struct {
	next   ArdeaService
	logger log.Logger
}

func LoggedArdeaService(logger log.Logger) Middleware {
	return func(s ArdeaService) ArdeaService {
		return &loggedArdeaService{
			next:   s,
			logger: logger,
		}
	}
}

func (s *loggedArdeaService) CreateModel(
	ctx context.Context, ownerID, name string, inputShape, outputShape [][]nulltype.NullInt64, valueType entities.ValueType,
) (model entities.Model, err error) {
	model, err = s.next.CreateModel(ctx, ownerID, name, inputShape, outputShape, valueType)
	s.logger.Log("method", "CreateModel", "ownerID", ownerID, "modelName", name,
		"valueType", valueType.String(), "error", err)
	return
}

func (s *loggedArdeaService) GetModel(ctx context.Context, modelID string) (model entities.Model, err error) {
	model, err = s.next.GetModel(ctx, modelID)
	s.logger.Log("method", "GetModel", "modelID", modelID, "error", err)
	return
}

func (s *loggedArdeaService) GetModelByName(
	ctx context.Context, ownerID, modelName string,
) (model entities.Model, err error) {
	model, err = s.next.GetModelByName(ctx, ownerID, modelName)
	s.logger.Log("method", "GetModelByName", "ownerID", ownerID, "modelName", modelName, "error", err)
	return
}

func (s *loggedArdeaService) UpdateModelState(
	ctx context.Context, modelID string, state entities.ModelState, errStr nulltype.NullString,
) (model entities.Model, err error) {
	model, err = s.next.UpdateModelState(ctx, modelID, state, errStr)
	s.logger.Log("method", "UpdateModelState", "modelID", modelID,
		"state", state.String(), "errStr", errStr.String(), "error", err)
	return
}

func (s *loggedArdeaService) UpdateModelPath(
	ctx context.Context, modelID, path string,
) (model entities.Model, err error) {
	model, err = s.next.UpdateModelPath(ctx, modelID, path)
	s.logger.Log("method", "UpdateModelPath", "modelID", modelID, "path", path, "error", err)
	return
}

func (s *loggedArdeaService) ListModels(ctx context.Context, ownerID string) (models []entities.Model, err error) {
	models, err = s.next.ListModels(ctx, ownerID)
	s.logger.Log("method", "ListModels", "ownerID", ownerID, "error", err)
	return
}
