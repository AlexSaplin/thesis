package service

import (
	"encoding/json"
	"flag"
	"fmt"
	"gorilla/pkg/config"
	"gorilla/pkg/db"
	"net"
	"os"
	"os/signal"
	"syscall"

	endpoint1 "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	grpc1 "google.golang.org/grpc"

	"gorilla/pkg/endpoint"
	"gorilla/pkg/grpc"
	pb "gorilla/pkg/grpc/pb"
	"gorilla/pkg/service"
)

var tracer opentracinggo.Tracer
var logger log.Logger

// Define our flags. Your service probably won't need to bind listeners for
// all* supported transports, but we do it here for demonstration purposes.
var fs = flag.NewFlagSet("gorilla", flag.ExitOnError)
var configPath = fs.String("config", "", "Config file path")

func Run() {
	fs.Parse(os.Args[1:])

	// Create a single logger, which we'll use and give to other components.
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	logger.Log("tracer", "none")
	tracer = opentracinggo.GlobalTracer()

	cfg, err := loadConfig(*configPath)
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}

	deltaDB, err := db.NewPostgresDeltaDB(cfg.DB, log.With(logger, "module", "db"))
	fmt.Println(cfg.DB)
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}

	svc := service.New(deltaDB, getServiceMiddleware(logger))
	eps := endpoint.New(svc, getEndpointMiddleware(logger))
	g := createService(eps, cfg.Server)
	initCancelInterrupt(g)
	logger.Log("exit", g.Run())

}
func initGRPCHandler(endpoints endpoint.Endpoints, g *group.Group, cfg config.ServerConfig) {
	options := defaultGRPCOptions(logger, tracer)
	// Add your GRPC options here

	grpcServer := grpc.NewGRPCServer(endpoints, options)
	grpcListener, err := net.Listen("tcp", cfg.Bind)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "gRPC", "addr", cfg.Bind)
		baseServer := grpc1.NewServer()
		pb.RegisterGorillaServer(baseServer, grpcServer)
		return baseServer.Serve(grpcListener)
	}, func(error) {
		grpcListener.Close()
	})
}

func getServiceMiddleware(logger log.Logger) (mw []service.Middleware) {
	mw = []service.Middleware{}
	mw = append(mw, service.LoggedGorillaService(log.With(logger, "module", "service")))
	// Append your middleware here
	return
}
func getEndpointMiddleware(logger log.Logger) (mw map[string][]endpoint1.Middleware) {
	mw = map[string][]endpoint1.Middleware{}
	return
}

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

func loadConfig(path string) (out config.GorillaConfig, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	err = json.NewDecoder(file).Decode(&out)
	return
}
