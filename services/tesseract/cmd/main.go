package main

import (
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"tesseract/pkg/k8s"
	"tesseract/pkg/service"
	pb "tesseract/pkg/service/pb"
)

func main() {
	logger := log.New()

	bind := "0.0.0.0:8092"
	lis, err := net.Listen("tcp", bind)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	k8sClient, err := k8s.NewK8SClient(".kube/config")
	if err != nil {
		logger.Fatalf("failed to init k8s client: %v", err)
	}
	svc := service.NewTesseractGRPCService(k8sClient)
	loggedService := service.NewLoggedTesseractServer(logger, svc)

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterTesseractServer(grpcServer, loggedService)
	logger.WithField("bind", bind).Info("listening")
	logger.Fatal(grpcServer.Serve(lis))
}
