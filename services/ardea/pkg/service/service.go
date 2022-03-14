package service

import (
	"context"

	"github.com/mattn/go-nulltype"
	uuid "github.com/satori/go.uuid"

	"ardea/pkg/db"
	"ardea/pkg/entities"
)

// ArdeaService describes the service.
type ArdeaService interface {
	CreateModel(
		ctx context.Context, ownerID, name string, inputShape, outputShape [][]nulltype.NullInt64, valueType entities.ValueType,
	) (model entities.Model, err error)
	GetModel(ctx context.Context, modelID string) (entities.Model, error)
	GetModelByName(ctx context.Context, ownerID, modelName string) (entities.Model, error)
	UpdateModelState(ctx context.Context, modelID string, state entities.ModelState, errStr nulltype.NullString) (entities.Model, error)
	UpdateModelPath(ctx context.Context, modelID, path string) (entities.Model, error)
	ListModels(ctx context.Context, ownerID string) ([]entities.Model, error)
}

type ArdeaServiceImpl struct {
	db db.ModelDB
}

// NewArdeaService returns an implementation of ArdeaService.
func NewArdeaService(modelDB db.ModelDB) ArdeaService {
	return &ArdeaServiceImpl{
		db: modelDB,
	}
}

// New returns a ArdeaService with all of the expected middleware wired in.
func New(modelDB db.ModelDB, middleware []Middleware) ArdeaService {
	var svc = NewArdeaService(modelDB)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (s *ArdeaServiceImpl) CreateModel(
	ctx context.Context, ownerStr, name string, inputShape, outputShape [][]nulltype.NullInt64, valueType entities.ValueType,
) (model entities.Model, err error) {
	var ownerID uuid.UUID
	ownerID, err = uuid.FromString(ownerStr)
	if err != nil {
		return
	}
	modelID := uuid.NewV4()

	return s.db.CreateModel(ctx, modelID, ownerID, name, inputShape, outputShape, valueType)
}
func (s *ArdeaServiceImpl) GetModel(ctx context.Context, modelStr string) (model entities.Model, err error) {
	modelID, err := uuid.FromString(modelStr)
	if err != nil {
		return
	}
	return s.db.GetModel(ctx, modelID)
}

func (s *ArdeaServiceImpl) GetModelByName(
	ctx context.Context, ownerIDStr, modelName string,
) (model entities.Model, err error) {
	ownerID, err := uuid.FromString(ownerIDStr)
	if err != nil {
		return
	}
	return s.db.GetModelByName(ctx, ownerID, modelName)
}

func (s *ArdeaServiceImpl) UpdateModelState(
	ctx context.Context, modelStr string, state entities.ModelState, errStr nulltype.NullString,
) (model entities.Model, err error) {
	modelID, err := uuid.FromString(modelStr)
	if err != nil {
		return
	}
	params := []db.ModelUpdateParam{db.StateUpdateParam(state)}
	if errStr.Valid() {
		params = append(params, db.ErrStrUpdateParam(errStr.String()))
	}

	return s.db.UpdateModel(ctx, modelID, params...)
}
func (s *ArdeaServiceImpl) UpdateModelPath(
	ctx context.Context, modelStr string, path string,
) (model entities.Model, err error) {
	modelID, err := uuid.FromString(modelStr)
	if err != nil {
		return
	}
	return s.db.UpdateModel(ctx, modelID, db.PathUpdateParam(path))
}

func (s *ArdeaServiceImpl) ListModels(ctx context.Context, ownerIDStr string) (models []entities.Model, err error) {
	ownerID, err := uuid.FromString(ownerIDStr)
	if err != nil {
		return
	}
	return s.db.ListModels(ctx, ownerID)
}
