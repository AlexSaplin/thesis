package service

import (
	endpoint1 "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"

	"rhino/pkg/config"
	"rhino/pkg/endpoint"
)

func createService(endpoints endpoint.Endpoints, cfg config.GRPCConfig) (g *group.Group) {
	g = &group.Group{}
	initGRPCHandler(endpoints, g, cfg)
	return g
}
func defaultGRPCOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]grpc.ServerOption {
	options := map[string][]grpc.ServerOption{
		"Run": {
			grpc.ServerErrorLogger(logger),
			grpc.ServerBefore(opentracing.GRPCToContext(tracer, "Run", logger)),
		},
	}
	return options
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"Run"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
