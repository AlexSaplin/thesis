package selachii

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"

	"hippo/pkg/clients/selachii/pb"
	"hippo/pkg/config"
	"hippo/pkg/entities"
)

const maxMsgSizeBytes = 1024 * 1024 * 64

type SelachiiClient interface {
	LoadModel(ctx context.Context, model entities.Model) (uuid.UUID, error)
	UnloadModel(ctx context.Context, loadID uuid.UUID) (bool, error)
	Run(ctx context.Context, loadID uuid.UUID, tensor entities.TensorList) (entities.TensorList, error)
}

type GRPCClient struct {
	client selachii.SelachiiClient
	config config.SelachiiClientConfig
}

func NewGRPCSelachiiClient(ctx context.Context, cfg config.SelachiiClientConfig) (SelachiiClient, error) {
	conn, err := grpc.DialContext(ctx,
		cfg.Target,
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(maxMsgSizeBytes),
			grpc.MaxCallSendMsgSize(maxMsgSizeBytes),
		),
	)
	if err != nil {
		return nil, err
	}
	return &GRPCClient{
		client: selachii.NewSelachiiClient(conn),
		config: cfg,
	}, nil
}

func (g *GRPCClient) LoadModel(ctx context.Context, model entities.Model) (loadID uuid.UUID, err error) {
	var (
		resp *selachii.LoadResponse
	)
	req := &selachii.LoadRequest{
		Model: &selachii.ModelMeta{
			ID:          model.ID.String(),
			InputShape:  serializeIOShapeEntity(model.InputShape),
			OutputShape: serializeIOShapeEntity(model.OutputShape),
			Path:        model.Path,
			Type:        serializeValueTypeEntity(model.ValueType),
		},
	}

	resp, err = g.client.LoadModel(ctx, req)
	if err != nil {
		return
	}

	return uuid.FromString(resp.LoadID)
}

func (g *GRPCClient) UnloadModel(ctx context.Context, loadID uuid.UUID) (changed bool, err error) {
	var (
		resp *selachii.UnloadResponse
	)
	req := &selachii.UnloadRequest{
		LoadID: loadID.String(),
	}

	resp, err = g.client.UnloadModel(ctx, req)
	if err != nil {
		return
	}
	return resp.DidChange, nil
}

func (g *GRPCClient) Run(
	ctx context.Context, loadID uuid.UUID, tensor entities.TensorList,
) (result entities.TensorList, err error) {
	var (
		resp *selachii.RunResponse
	)
	req := &selachii.RunRequest{
		LoadID: loadID.String(),
		Tensor: serializeTensorListEntity(tensor),
	}

	resp, err = g.client.Run(ctx, req,
		grpc.MaxCallRecvMsgSize(maxMsgSizeBytes),
		grpc.MaxCallSendMsgSize(maxMsgSizeBytes),
	)
	if err != nil {
		return
	}
	return parseTensorListEntity(resp.Tensor)
}
