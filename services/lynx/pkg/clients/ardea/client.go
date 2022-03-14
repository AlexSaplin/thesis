package ardea

import (
	"context"

	"github.com/mattn/go-nulltype"
	"github.com/satori/go.uuid"
	"google.golang.org/grpc"

	ardea "lynx/pkg/clients/ardea/pb"
	"lynx/pkg/config"
	"lynx/pkg/entities"
)

type Client interface {
	CreateModel(
		ctx context.Context, ownerID uuid.UUID, name string,
		inputShape, outputShape entities.IOShape, valueType entities.ValueType,
	) (model entities.Model, err error)
	GetModelByName(ctx context.Context, ownerID uuid.UUID, modelName string) (entities.Model, error)
	UpdateModelState(ctx context.Context, modelID uuid.UUID, state entities.ModelState) (entities.Model, error)
	UpdateModelPath(ctx context.Context, modelID uuid.UUID, path string) (entities.Model, error)
	ListModels(ctx context.Context, ownerID uuid.UUID) ([]entities.Model, error)
}

type GRPCClient struct {
	config config.ArdeaClientConfig
	client ardea.ArdeaClient
}

func NewGRPCClient(ctx context.Context, config config.ArdeaClientConfig) (Client, error) {
	conn, err := grpc.DialContext(ctx, config.Target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &GRPCClient{
		client: ardea.NewArdeaClient(conn),
		config: config,
	}, nil
}

func (g *GRPCClient) CreateModel(
	ctx context.Context, ownerID uuid.UUID, name string,
	inputShape, outputShape entities.IOShape, valueType entities.ValueType,
) (model entities.Model, err error) {
	req := &ardea.CreateModelRequest{
		OwnerID:     ownerID.String(),
		Name:        name,
		InputShape:  serializeIOShapeEntity(inputShape),
		OutputShape: serializeIOShapeEntity(outputShape),
		Type:        makeValueTypePb(valueType),
	}
	resp, err := g.client.CreateModel(ctx, req)
	if err != nil {
		return
	}
	return makeModelEntity(resp.Model)
}

func (g *GRPCClient) GetModelByName(
	ctx context.Context, ownerID uuid.UUID, modelName string,
) (model entities.Model, err error) {
	req := &ardea.GetModelByNameRequest{
		OwnerID: ownerID.String(),
		Name:    modelName,
	}

	resp, err := g.client.GetModelByName(ctx, req)
	if err != nil {
		return
	}

	return makeModelEntity(resp.Model)
}

func (g *GRPCClient) UpdateModelState(
	ctx context.Context, modelID uuid.UUID, state entities.ModelState,
) (model entities.Model, err error) {
	statePb, err := makeModelStatePb(state)
	if err != nil {
		return
	}
	req := &ardea.UpdateModelStateRequest{
		ID:    modelID.String(),
		State: statePb,
	}
	resp, err := g.client.UpdateModelState(ctx, req)
	if err != nil {
		return
	}
	return makeModelEntity(resp.Model)
}

func (g *GRPCClient) UpdateModelPath(
	ctx context.Context, modelID uuid.UUID, path string,
) (model entities.Model, err error) {
	req := &ardea.UpdateModelPathRequest{
		ID:   modelID.String(),
		Path: path,
	}
	resp, err := g.client.UpdateModelPath(ctx, req)
	if err != nil {
		return
	}
	return makeModelEntity(resp.Model)
}

func (g *GRPCClient) ListModels(ctx context.Context, ownerID uuid.UUID) (models []entities.Model, err error) {
	req := &ardea.ListModelsRequest{
		OwnerID: ownerID.String(),
	}

	resp, err := g.client.ListModels(ctx, req)
	if err != nil {
		return
	}
	result := make([]entities.Model, 0, len(resp.Models))
	for i := 0; i < len(resp.Models); i++ {
		model, err := makeModelEntity(resp.Models[i])
		if err != nil {
			return nil, err
		}
		result = append(result, model)
	}
	return result, nil
}

func makeModelEntity(in *ardea.Model) (result entities.Model, err error) {
	modelID, err := uuid.FromString(in.ID)
	if err != nil {
		return
	}

	ownerID, err := uuid.FromString(in.OwnerID)
	if err != nil {
		return
	}

	state, err := parseModelStateEntity(in.State)
	if err != nil {
		return
	}
	var errStr nulltype.NullString
	if in.ErrStrSet {
		errStr = nulltype.NullStringOf(in.ErrStr)
	}

	result = entities.Model{
		ID:          modelID,
		Name:        in.Name,
		OwnerID:     ownerID,
		ValueType:   parseValueTypeEntity(in.Type),
		State:       state,
		InputShape:  parseIOShapeEntity(in.InputShape),
		OutputShape: parseIOShapeEntity(in.OutputShape),
		Path:        in.Path,
		Error:       errStr,
	}
	return
}
