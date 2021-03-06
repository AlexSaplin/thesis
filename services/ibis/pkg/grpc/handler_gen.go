// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	"github.com/go-kit/kit/transport/grpc"

	"ibis/pkg/endpoint"
	pb "ibis/pkg/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	createFunction    grpc.Handler
	getFunction       grpc.Handler
	getFunctionByName grpc.Handler
	updateFunction    grpc.Handler
	listFunctions     grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.IbisServer {
	return &grpcServer{
		createFunction:    makeCreateFunctionHandler(endpoints, options["CreateFunction"]),
		getFunction:       makeGetFunctionHandler(endpoints, options["GetFunction"]),
		getFunctionByName: makeGetFunctionByNameHandler(endpoints, options["GetFunctionByName"]),
		updateFunction:    makeUpdateFunctionHandler(endpoints, options["UpdateFunction"]),
		listFunctions:     makeListFunctionsHandler(endpoints, options["ListFunctions"]),
	}
}
