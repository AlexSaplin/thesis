package grpc

import (
	"context"

	grpc "github.com/go-kit/kit/transport/grpc"
	"github.com/mattn/go-nulltype"
	context1 "golang.org/x/net/context"

	endpoint "ardea/pkg/endpoint"
	pb "ardea/pkg/grpc/pb"
)

func makeCreateModelHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateModelEndpoint, decodeCreateModelRequest, encodeCreateModelResponse, options...)
}

func decodeCreateModelRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CreateModelRequest)
	decoded := endpoint.CreateModelRequest{
		OwnerID:     req.OwnerID,
		InputShape:  parseIOShapeEntity(req.InputShape),
		OutputShape: parseIOShapeEntity(req.OutputShape),
		Name:        req.Name,
		ValueType:   parseValueTypeEntity(req.Type),
	}
	return decoded, nil
}

func encodeCreateModelResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.CreateModelResponse)
	model, err := serializeModelEntity(resp.Model)
	if err != nil {
		return nil, err
	}
	return &pb.CreateModelReply{
		Model: &model,
	}, resp.Err
}
func (g *grpcServer) CreateModel(ctx context1.Context, req *pb.CreateModelRequest) (*pb.CreateModelReply, error) {
	_, rep, err := g.createModel.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateModelReply), nil
}

func makeGetModelHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetModelEndpoint, decodeGetModelRequest, encodeGetModelResponse, options...)
}

func decodeGetModelRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetModelRequest)
	return endpoint.GetModelRequest{
		ModelID: req.ID,
	}, nil
}

func encodeGetModelResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.GetModelResponse)
	model, err := serializeModelEntity(resp.Model)
	if err != nil {
		return nil, err
	}
	return &pb.GetModelReply{
		Model: &model,
	}, resp.Err
}
func (g *grpcServer) GetModel(ctx context1.Context, req *pb.GetModelRequest) (*pb.GetModelReply, error) {
	_, rep, err := g.getModel.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetModelReply), nil
}

func makeGetModelByNameHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetModelByNameEndpoint, decodeGetModelByNameRequest, encodeGetModelByNameResponse, options...)
}

func decodeGetModelByNameRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetModelByNameRequest)
	return endpoint.GetModelByNameRequest{
		OwnerID: req.OwnerID,
		Name:    req.Name,
	}, nil
}

func encodeGetModelByNameResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.GetModelByNameResponse)
	model, err := serializeModelEntity(resp.Model)
	if err != nil {
		return nil, err
	}
	return &pb.GetModelReply{
		Model: &model,
	}, resp.Err
}

func (g *grpcServer) GetModelByName(ctx context1.Context, req *pb.GetModelByNameRequest) (*pb.GetModelReply, error) {
	_, rep, err := g.getModelByName.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetModelReply), nil
}

func makeUpdateModelStateHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateModelStateEndpoint, decodeUpdateModelStateRequest, encodeUpdateModelStateResponse, options...)
}

func decodeUpdateModelStateRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.UpdateModelStateRequest)

	state, err := parseModelStateEntity(req.State)
	if err != nil {
		return nil, err
	}
	var errStr nulltype.NullString
	if req.ErrStrSet {
		errStr = nulltype.NullStringOf(req.ErrStr)
	}

	return endpoint.UpdateModelStateRequest{
		ModelID: req.ID,
		State:   state,
		ErrStr:  errStr,
	}, nil
}

func encodeUpdateModelStateResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.UpdateModelStateResponse)
	model, err := serializeModelEntity(resp.Model)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateModelStateReply{
		Model: &model,
	}, resp.Err
}
func (g *grpcServer) UpdateModelState(ctx context1.Context, req *pb.UpdateModelStateRequest) (*pb.UpdateModelStateReply, error) {
	_, rep, err := g.updateModelState.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateModelStateReply), nil
}

func makeUpdateModelPathHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateModelPathEndpoint, decodeUpdateModelPathRequest, encodeUpdateModelPathResponse, options...)
}

func decodeUpdateModelPathRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.UpdateModelPathRequest)
	return endpoint.UpdateModelPathRequest{
		ModelID: req.ID,
		Path:    req.Path,
	}, nil
}

func encodeUpdateModelPathResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.UpdateModelPathResponse)
	model, err := serializeModelEntity(resp.Model)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateModelPathReply{
		Model: &model,
	}, resp.Err
}

func (g *grpcServer) UpdateModelPath(ctx context1.Context, req *pb.UpdateModelPathRequest) (*pb.UpdateModelPathReply, error) {
	_, rep, err := g.updateModelPath.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateModelPathReply), nil
}

func makeListModelsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ListModelsEndpoint, decodeListModelsRequest, encodeListModelsResponse, options...)
}

func decodeListModelsRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.ListModelsRequest)
	return endpoint.ListModelsRequest{
		OwnerID: req.OwnerID,
	}, nil
}

func encodeListModelsResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.ListModelsResponse)
	models := make([]*pb.Model, len(resp.Models))
	for i := 0; i < len(resp.Models); i++ {
		model, err := serializeModelEntity(resp.Models[i])
		if err != nil {
			return nil, err
		}
		models[i] = &model
	}
	return &pb.ListModelsReply{
		Models: models,
	}, resp.Err
}

func (g *grpcServer) ListModels(ctx context1.Context, req *pb.ListModelsRequest) (*pb.ListModelsReply, error) {
	_, rep, err := g.listModels.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ListModelsReply), nil
}
