package grpc

import (
	"context"

	grpc "github.com/go-kit/kit/transport/grpc"
	uuid "github.com/satori/go.uuid"
	context1 "golang.org/x/net/context"

	endpoint "rhino/pkg/endpoint"
	pb "rhino/pkg/grpc/pb"
)

// makeRunHandler creates the handler logic
func makeRunHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.RunEndpoint, decodeRunRequest, encodeRunResponse, options...)
}

// decodeRunResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Run request.
func decodeRunRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.RunRequest)
	functionID, err := uuid.FromString(req.FunctionID)
	if err != nil {
		return nil, err
	}
	return endpoint.RunRequest{
		FunctionID: functionID,
		Data:       req.Data,
	}, nil
}

// encodeRunResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeRunResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.RunResponse)
	return &pb.RunReply{
		Data: resp.Result,
	}, resp.Err
}

func (g *grpcServer) Run(ctx context1.Context, req *pb.RunRequest) (*pb.RunReply, error) {
	_, rep, err := g.run.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RunReply), nil
}

