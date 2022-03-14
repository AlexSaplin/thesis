// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "gorilla/pkg/endpoint"
	pb "gorilla/pkg/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	addDeltas  grpc.Handler
	getDeltas  grpc.Handler
	getBalance grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.GorillaServer {
	return &grpcServer{
		addDeltas:  makeAddDeltasHandler(endpoints, options["AddDeltas"]),
		getBalance: makeGetBalanceHandler(endpoints, options["GetBalance"]),
		getDeltas:  makeGetDeltasHandler(endpoints, options["GetDeltas"]),
	}
}
