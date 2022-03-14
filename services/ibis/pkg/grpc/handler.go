package grpc

import (
	"context"

	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"

	endpoint "ibis/pkg/endpoint"
	pb "ibis/pkg/grpc/pb"
)

func makeCreateFunctionHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateFunctionEndpoint, decodeCreateFunctionRequest, encodeCreateFunctionResponse, options...)
}

func decodeCreateFunctionRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CreateFunctionRequest)
	decoded := endpoint.CreateFunctionRequest{
		OwnerID:     req.OwnerID,
		Name:        req.Name,
	}
	return decoded, nil
}

func encodeCreateFunctionResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.CreateFunctionResponse)
	Function, err := serializeFunctionEntity(resp.Function)
	if err != nil {
		return nil, err
	}
	return &pb.CreateFunctionReply{
		Function: &Function,
	}, resp.Err
}
func (g *grpcServer) CreateFunction(ctx context1.Context, req *pb.CreateFunctionRequest) (*pb.CreateFunctionReply, error) {
	_, rep, err := g.createFunction.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateFunctionReply), nil
}

func makeGetFunctionHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetFunctionEndpoint, decodeGetFunctionRequest, encodeGetFunctionResponse, options...)
}

func decodeGetFunctionRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetFunctionRequest)
	return endpoint.GetFunctionRequest{
		FunctionID: req.ID,
	}, nil
}

func encodeGetFunctionResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.GetFunctionResponse)
	Function, err := serializeFunctionEntity(resp.Function)
	if err != nil {
		return nil, err
	}
	return &pb.GetFunctionReply{
		Function: &Function,
	}, resp.Err
}
func (g *grpcServer) GetFunction(ctx context1.Context, req *pb.GetFunctionRequest) (*pb.GetFunctionReply, error) {
	_, rep, err := g.getFunction.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetFunctionReply), nil
}

func makeGetFunctionByNameHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetFunctionByNameEndpoint, decodeGetFunctionByNameRequest, encodeGetFunctionByNameResponse, options...)
}

func decodeGetFunctionByNameRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetFunctionByNameRequest)
	return endpoint.GetFunctionByNameRequest{
		OwnerID: req.OwnerID,
		Name:    req.Name,
	}, nil
}

func encodeGetFunctionByNameResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.GetFunctionByNameResponse)
	Function, err := serializeFunctionEntity(resp.Function)
	if err != nil {
		return nil, err
	}
	return &pb.GetFunctionReply{
		Function: &Function,
	}, resp.Err
}

func (g *grpcServer) GetFunctionByName(ctx context1.Context, req *pb.GetFunctionByNameRequest) (*pb.GetFunctionReply, error) {
	_, rep, err := g.getFunctionByName.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetFunctionReply), nil
}

func makeUpdateFunctionHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateFunctionEndpoint, decodeUpdateFunctionRequest, encodeUpdateFunctionResponse, options...)
}

func decodeUpdateFunctionRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.UpdateFunctionRequest)

	param, err := parseUpdateParams(req.Params)
	if err != nil {
		return nil, err
	}

	return endpoint.UpdateFunctionRequest{
		FunctionID: req.ID,
		Param:     param,
	}, nil
}

func encodeUpdateFunctionResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.UpdateFunctionResponse)
	Function, err := serializeFunctionEntity(resp.Function)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateFunctionReply{
		Function: &Function,
	}, resp.Err
}
func (g *grpcServer) UpdateFunction(ctx context1.Context, req *pb.UpdateFunctionRequest) (*pb.UpdateFunctionReply, error) {
	_, rep, err := g.updateFunction.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateFunctionReply), nil
}

func makeListFunctionsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ListFunctionsEndpoint, decodeListFunctionsRequest, encodeListFunctionsResponse, options...)
}

func decodeListFunctionsRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.ListFunctionsRequest)
	return endpoint.ListFunctionsRequest{
		OwnerID: req.OwnerID,
	}, nil
}

func encodeListFunctionsResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.ListFunctionsResponse)
	Functions := make([]*pb.Function, len(resp.Functions))
	for i := 0; i < len(resp.Functions); i++ {
		Function, err := serializeFunctionEntity(resp.Functions[i])
		if err != nil {
			return nil, err
		}
		Functions[i] = &Function
	}
	return &pb.ListFunctionsReply{
		Functions: Functions,
	}, resp.Err
}

func (g *grpcServer) ListFunctions(ctx context1.Context, req *pb.ListFunctionsRequest) (*pb.ListFunctionsReply, error) {
	_, rep, err := g.listFunctions.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ListFunctionsReply), nil
}
