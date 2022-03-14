package service

import (
	"context"
	"fmt"

	"tesseract/pkg/k8s"
	pb "tesseract/pkg/service/pb"
)

type TesseractGRPCService struct {
	client *k8s.K8SClient
}

func NewTesseractGRPCService(client *k8s.K8SClient) *TesseractGRPCService {
	return &TesseractGRPCService{client: client}
}

func (t *TesseractGRPCService) Apply(ctx context.Context, req *pb.ApplyRequest) (*pb.ApplyResponse, error) {
	args := k8s.TemplateArgs{
		Namespace: req.ID,
		Name:      req.Name,
		DNS:       req.DNS,
		Image:     req.Image,
		Port:      uint(req.Port),
		Scale:     uint(req.Scale),
		CPU:       uint(req.CPU * 100),
		MemoryMB:  uint(req.RAM),
		GPU:       req.GPU,
		Env:       req.Env,
		Auth:      req.Auth,
	}

	err := t.client.Apply(ctx, args)
	return &pb.ApplyResponse{}, err
}

func (t *TesseractGRPCService) GetStatus(ctx context.Context, req *pb.GetStatusRequest) (*pb.GetStatusResponse, error) {
	res, err := t.client.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return mapContainerStatus(res), nil
}

func (t *TesseractGRPCService) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return &pb.DeleteResponse{}, t.client.Down(ctx, req.ID)
}

func mapContainerStatus(in string) *pb.GetStatusResponse {
	var (
		status pb.Status
		err    string
	)
	switch in {
	case "ContainerCreating", "MODIFYING":
		status = pb.Status_UPDATING
	case "OK":
		status = pb.Status_RUNNING
	case "ErrImagePull":
		status = pb.Status_ERROR
		err = fmt.Sprintf("Failed to pull image")
	case "ContainersNotReady":
		status = pb.Status_ERROR
		err = fmt.Sprintf("Port is not accepting TCP connections")
	case "Unschedulable":
		status = pb.Status_ERROR
		err = fmt.Sprintf("Internal error: can't schedule container")
	default:
		status = pb.Status_ERROR
		err = fmt.Sprintf("Invalid container status: %s", in)
	}
	return &pb.GetStatusResponse{
		Status: status,
		Error:  err,
	}
}
