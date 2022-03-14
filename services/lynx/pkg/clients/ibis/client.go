package ibis

import (
	"context"
	"github.com/mattn/go-nulltype"

	"github.com/satori/go.uuid"
	"google.golang.org/grpc"

	ibis "lynx/pkg/clients/ibis/pb"
	"lynx/pkg/config"
	"lynx/pkg/entities"
)

type Client interface {
	CreateFunction(
		ctx context.Context, ownerID uuid.UUID, name string,
	) (function entities.Function, err error)
	GetFunctionByName(ctx context.Context, ownerID uuid.UUID, functionName string) (entities.Function, error)
	UpdateFunction(ctx context.Context, modelID uuid.UUID, param UpdateFunctionParam) (fn entities.Function, err error)
	ListFunctions(ctx context.Context, ownerID uuid.UUID) ([]entities.Function, error)
}

type UpdateFunctionParam struct {
	State    *entities.FunctionState
	ErrStr   nulltype.NullString
	ImageURL nulltype.NullString
	CodePath nulltype.NullString
	Metadata nulltype.NullString
}

type GRPCClient struct {
	config config.IbisClientConfig
	client ibis.IbisClient
}

func NewGRPCClient(ctx context.Context, config config.IbisClientConfig) (Client, error) {
	conn, err := grpc.DialContext(ctx, config.Target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &GRPCClient{
		client: ibis.NewIbisClient(conn),
		config: config,
	}, nil
}

func (g *GRPCClient) CreateFunction(
	ctx context.Context, ownerID uuid.UUID, name string,
) (function entities.Function, err error) {
	req := &ibis.CreateFunctionRequest{
		OwnerID:     ownerID.String(),
		Name:        name,
	}
	resp, err := g.client.CreateFunction(ctx, req)
	if err != nil {
		return
	}
	return makeFunctionEntity(resp.Function)
}

func (g *GRPCClient) GetFunctionByName(
	ctx context.Context, ownerID uuid.UUID, name string,
) (fn entities.Function, err error) {
	req := &ibis.GetFunctionByNameRequest{
		OwnerID: ownerID.String(),
		Name:    name,
	}

	resp, err := g.client.GetFunctionByName(ctx, req)
	if err != nil {
		return
	}

	return makeFunctionEntity(resp.Function)
}



func (g *GRPCClient) UpdateFunction(
	ctx context.Context, modelID uuid.UUID, param UpdateFunctionParam,
) (fn entities.Function, err error) {
	var updateParams []*ibis.UpdateFunctionParam

	if param.State != nil {
		var statePb ibis.FunctionState
		statePb, err = makeFunctionStatePb(*param.State)
		if err != nil {
			return
		}
		updateParams = append(updateParams, &ibis.UpdateFunctionParam{
			Param: &ibis.UpdateFunctionParam_State{
				State: statePb,
			},
		})
	}

	if param.ErrStr.Valid() {
		updateParams = append(updateParams, &ibis.UpdateFunctionParam{
			Param: &ibis.UpdateFunctionParam_ErrStr{
				ErrStr: param.ErrStr.String(),
			},
		})
	}

	if param.ImageURL.Valid() {
		updateParams = append(updateParams, &ibis.UpdateFunctionParam{
			Param: &ibis.UpdateFunctionParam_ImageURL{
				ImageURL: param.ImageURL.String(),
			},
		})
	}

	if param.CodePath.Valid() {
		updateParams = append(updateParams, &ibis.UpdateFunctionParam{
			Param: &ibis.UpdateFunctionParam_CodePath{
				CodePath: param.CodePath.String(),
			},
		})
	}

	if param.Metadata.Valid() {
		updateParams = append(updateParams, &ibis.UpdateFunctionParam{
			Param: &ibis.UpdateFunctionParam_Metadata{
				Metadata: param.Metadata.String(),
			},
		})
	}

	req := &ibis.UpdateFunctionRequest{
		ID:   modelID.String(),
		Params: updateParams,
	}


	resp, err := g.client.UpdateFunction(ctx, req)
	if err != nil {
		return
	}

	return makeFunctionEntity(resp.Function)
}

func (g *GRPCClient) ListFunctions(ctx context.Context, ownerID uuid.UUID) (functions []entities.Function, err error) {
	req := &ibis.ListFunctionsRequest{
		OwnerID: ownerID.String(),
	}

	resp, err := g.client.ListFunctions(ctx, req)
	if err != nil {
		return
	}
	result := make([]entities.Function, 0, len(resp.Functions))
	for i := 0; i < len(resp.Functions); i++ {
		f, err := makeFunctionEntity(resp.Functions[i])
		if err != nil {
			return nil, err
		}
		result = append(result, f)
	}
	return result, nil
}

