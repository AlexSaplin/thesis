package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "rhino/pkg/endpoint"
	pb "rhino/pkg/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	run grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.RhinoServer {
	return &grpcServer{run: makeRunHandler(endpoints, options["Run"])}
}
