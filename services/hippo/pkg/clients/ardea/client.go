package ardea

import (
	"context"

	"github.com/satori/go.uuid"
	"google.golang.org/grpc"

	"hippo/pkg/clients/ardea/pb"
	"hippo/pkg/config"
	"hippo/pkg/entities"
)

type Client interface {
	GetModel(ctx context.Context, modelID uuid.UUID) (entities.Model, error)
}

type GRPCClient struct {
	client ardea.ArdeaClient
	config config.ArdeaClientConfig
}

func NewGRPCArdeaClient(ctx context.Context, cfg config.ArdeaClientConfig) (Client, error) {
	conn, err := grpc.DialContext(ctx, cfg.Target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &GRPCClient{
		client: ardea.NewArdeaClient(conn),
		config: cfg,
	}, nil
}

func (g *GRPCClient) GetModel(ctx context.Context, modelIDReq uuid.UUID) (model entities.Model, err error) {
	var (
		req  *ardea.GetModelRequest
		resp *ardea.GetModelReply

		modelID uuid.UUID
		ownerID uuid.UUID
		state   entities.ModelState
	)

	req = &ardea.GetModelRequest{
		ID: modelIDReq.String(),
	}

	resp, err = g.client.GetModel(ctx, req)
	if err != nil {
		return
	}

	modelID, err = uuid.FromString(resp.Model.ID)
	if err != nil {
		return
	}
	ownerID, err = uuid.FromString(resp.Model.OwnerID)
	if err != nil {
		return
	}
	state, err = parseModelStateEntity(resp.Model.State)
	if err != nil {
		return
	}
	model = entities.Model{
		ID:          modelID,
		OwnerID:     ownerID,
		ValueType:   parseValueTypeEntity(resp.Model.Type),
		State:       state,
		InputShape:  parseIOShapeEntity(resp.Model.InputShape),
		OutputShape: parseIOShapeEntity(resp.Model.OutputShape),
		Path:        resp.Model.Path,
	}
	return
}
