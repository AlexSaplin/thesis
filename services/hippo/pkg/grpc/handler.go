package grpc

import (
	"context"

	grpc "github.com/go-kit/kit/transport/grpc"
	uuid "github.com/satori/go.uuid"
	context1 "golang.org/x/net/context"

	endpoint "hippo/pkg/endpoint"
	pb "hippo/pkg/grpc/pb"
)

// makeRunHandler creates the handler logic
func makeRunHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.RunEndpoint, decodeRunRequest, encodeRunResponse, options...)
}

// decodeRunResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Run request.
func decodeRunRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.RunRequest)
	modelID, err := uuid.FromString(req.ModelID)
	if err != nil {
		return nil, err
	}
	parsed, err := parseTensorList(req.Tensors)
	if err != nil {
		return nil, err
	}
	return endpoint.RunRequest{
		ModelID: modelID,
		Tensors: parsed,
	}, nil
}

// encodeRunResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeRunResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.RunResponse)
	return &pb.RunReply{
		Tensors: serializeTensorList(resp.Result),
	}, resp.Err
}

func (g *grpcServer) Run(ctx context1.Context, req *pb.RunRequest) (*pb.RunReply, error) {
	_, rep, err := g.run.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RunReply), nil
}

