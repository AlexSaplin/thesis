package ibis

import (
	"context"
	"github.com/mattn/go-nulltype"

	"github.com/satori/go.uuid"
	"google.golang.org/grpc"

	ibis "rhino/pkg/clients/ibis/pb"
	"rhino/pkg/config"
	"rhino/pkg/entities"
)

type Client interface {
	GetFunction(ctx context.Context, modelID uuid.UUID) (entities.Function, error)
}

type GRPCClient struct {
	client ibis.IbisClient
	config config.IbisClientConfig
}

func NewGRPCIbisClient(ctx context.Context, cfg config.IbisClientConfig) (Client, error) {
	conn, err := grpc.DialContext(ctx, cfg.Target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &GRPCClient{
		client: ibis.NewIbisClient(conn),
		config: cfg,
	}, nil
}

func (g *GRPCClient) GetFunction(ctx context.Context, modelIDReq uuid.UUID) (model entities.Function, err error) {
	var (
		req  *ibis.GetFunctionRequest
		resp *ibis.GetFunctionReply

		functionID uuid.UUID
		ownerID uuid.UUID
		state   entities.FunctionState
	)

	req = &ibis.GetFunctionRequest{
		ID: modelIDReq.String(),
	}

	resp, err = g.client.GetFunction(ctx, req)
	if err != nil {
		return
	}

	functionID, err = uuid.FromString(resp.Function.ID)
	if err != nil {
		return
	}
	ownerID, err = uuid.FromString(resp.Function.OwnerID)
	if err != nil {
		return
	}
	state, err = parseFunctionStateEntity(resp.Function.State)
	if err != nil {
		return
	}
	model = entities.Function{
		ID:       functionID,
		OwnerID:  ownerID,
		Name:     resp.Function.Name,
		State:    state,
		ImageURL: nulltype.NullStringOf(resp.Function.ImageURL),
		CodePath: nulltype.NullStringOf(resp.Function.CodePath),
		Metadata: nulltype.NullStringOf(resp.Function.Metadata),
		Error:    nulltype.NullStringOf(resp.Function.ErrStr),
	}
	return
}
