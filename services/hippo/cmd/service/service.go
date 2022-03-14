package service

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	grpc1 "google.golang.org/grpc"

	"hippo/pkg/clients/ardea"
	"hippo/pkg/clients/nalogi"
	"hippo/pkg/clients/selachii"
	"hippo/pkg/config"
	endpoint "hippo/pkg/endpoint"
	grpc "hippo/pkg/grpc"
	pb "hippo/pkg/grpc/pb"
	"hippo/pkg/reporter"
	"hippo/pkg/runner/timeout"
	service "hippo/pkg/service"
)

const maxPbMessageSize = 1024 * 1024 * 64

var tracer opentracinggo.Tracer
var logger log.Logger

var fs = flag.NewFlagSet("hippo", flag.ExitOnError)
var configPath = fs.String("config", "", "hippo config path")

func mustLoadConfig(path string)  config.HippoConfig {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var cfg config.HippoConfig
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func Run() {
	fs.Parse(os.Args[1:])

	cfg := mustLoadConfig(*configPath)

	// Create a single logger, which we'll use and give to other components.
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	logger.Log("tracer", "none")
	tracer = opentracinggo.GlobalTracer()

	ardeaClient, err := ardea.NewGRPCArdeaClient(context.Background(), cfg.Ardea)
	if err != nil {
		panic(err)
	}

	selachiiClient, err := selachii.NewGRPCSelachiiClient(context.Background(), cfg.Selachii)
	if err != nil {
		panic(err)
	}

	runnerRegistry := timeout.NewRegistry(selachiiClient, log.With(logger, "module", "registry"))

	client, err := nalogi.NewKafkaClient(cfg.Nalogi)
	if err != nil {
		panic(err)
	}

	rep := reporter.NewReporter(client, log.With(logger, "module", "reporter"))


	svc := service.New(ardeaClient, runnerRegistry, rep, getServiceMiddleware(logger))
	eps := endpoint.New(svc, getEndpointMiddleware(logger))
	g := createService(eps, cfg.GRPC)
	// initMetricsEndpoint(g)
	initCancelInterrupt(g)
	logger.Log("exit", g.Run())

}
func initGRPCHandler(endpoints endpoint.Endpoints, g *group.Group, cfg config.GRPCConfig) {
	options := defaultGRPCOptions(logger, tracer)
	// Add your GRPC options here


	grpcServer := grpc.NewGRPCServer(endpoints, options)
	grpcListener, err := net.Listen("tcp", cfg.Bind)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "gRPC", "addr", cfg.Bind)
		baseServer := grpc1.NewServer(grpc1.MaxRecvMsgSize(maxPbMessageSize), grpc1.MaxSendMsgSize(maxPbMessageSize))
		pb.RegisterHippoServer(baseServer, grpcServer)
		return baseServer.Serve(grpcListener)
	}, func(error) {
		grpcListener.Close()
	})
}

func getServiceMiddleware(logger log.Logger) (mw []service.Middleware) {
	mw = []service.Middleware{}
	// Append your middleware here
	mw = append(mw, service.LoggedHippoService(logger))
	return
}

func getEndpointMiddleware(logger log.Logger) (mw map[string][]endpoint1.Middleware) {
	mw = map[string][]endpoint1.Middleware{}
	// Add you endpoint middleware here

	return
}
/*
func initMetricsEndpoint(g *group.Group) {
	http1.DefaultServeMux.Handle("/metrics", promhttp.Handler())
	debugListener, err := net.Listen("tcp", *debugAddr)
	if err != nil {
		logger.Log("transport", "debug/HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "debug/HTTP", "addr", *debugAddr)
		return http1.Serve(debugListener, http1.DefaultServeMux)
	}, func(error) {
		debugListener.Close()
	})
}
 */

func initCancelInterrupt(g *group.Group) {
	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})
}
