package service

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"ardea/pkg/config"
	"ardea/pkg/db"
	"ardea/pkg/endpoint"
	"ardea/pkg/grpc"
	pb "ardea/pkg/grpc/pb"
	"ardea/pkg/service"

	endpoint1 "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/oklog/oklog/pkg/group"
	"github.com/opentracing/opentracing-go"
	grpc1 "google.golang.org/grpc"
)

var logger log.Logger

var fs = flag.NewFlagSet("ardea", flag.ExitOnError)
var configPath = fs.String("config", "", "Config file path")

func Run() {
	fs.Parse(os.Args[1:])

	// Create a single logger, which we'll use and give to other components.
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	logger.Log("tracer", "none")

	cfg, err := loadConfig(*configPath)
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}

	modelDB, err := db.NewPostgresModelDB(cfg.DB, log.With(logger, "module", "db"))
	fmt.Println(cfg.DB)
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}
	svc := service.New(modelDB, getServiceMiddleware(logger))
	eps := endpoint.New(svc, getEndpointMiddleware(logger))
	g := createService(eps, cfg.Server)
	initCancelInterrupt(g)
	logger.Log("exit", g.Run())

}
func initGRPCHandler(endpoints endpoint.Endpoints, g *group.Group, cfg config.ServerConfig) {
	options := defaultGRPCOptions(logger, opentracing.GlobalTracer())
	// Add your GRPC options here

	grpcServer := grpc.NewGRPCServer(endpoints, options)
	grpcListener, err := net.Listen("tcp", cfg.Bind)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "gRPC", "addr", cfg.Bind)
		baseServer := grpc1.NewServer()
		pb.RegisterArdeaServer(baseServer, grpcServer)
		return baseServer.Serve(grpcListener)
	}, func(error) {
		grpcListener.Close()
	})

}
func getServiceMiddleware(logger log.Logger) (mw []service.Middleware) {
	mw = []service.Middleware{}
	mw = append(mw, service.LoggedArdeaService(log.With(logger, "module", "service")))
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

func loadConfig(path string) (out config.ArdeaConfig, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	err = json.NewDecoder(file).Decode(&out)
	return
}
