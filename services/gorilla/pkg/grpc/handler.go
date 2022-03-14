package grpc

import (
	"context"
	grpc "github.com/go-kit/kit/transport/grpc"
	uuid "github.com/satori/go.uuid"
	context1 "golang.org/x/net/context"
	endpoint "gorilla/pkg/endpoint"
	"gorilla/pkg/entities"
	pb "gorilla/pkg/grpc/pb"
	"time"
)

// makeAddDeltasHandler creates the handler logic
func makeAddDeltasHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AddDeltasEndpoint, decodeAddDeltasRequest, encodeAddDeltasResponse, options...)
}

// decodeAddDeltasResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain AddDeltas request.
func decodeAddDeltasRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.AddDeltasRequest)
	deltas := make([]entities.Delta, len(req.Deltas))
	for i := 0; i < len(req.Deltas); i++ {
		ownerID, err := uuid.FromString(req.Deltas[i].OwnerID)
		if err != nil {
			return nil, err
		}
		deltas[i] = entities.Delta{
			Date:       time.Unix(req.Deltas[i].Date, 0),
			Category:   req.Deltas[i].Category,
			Balance:    req.Deltas[i].Balance,
			OwnerID:    ownerID,
			ObjectID:   req.Deltas[i].ObjectID,
			ObjectType: req.Deltas[i].ObjectType,
		}
	}
	return endpoint.AddDeltasRequest{
		Deltas: deltas,
	}, nil
}

// encodeAddDeltasResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeAddDeltasResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.AddDeltasResponse)
	return &pb.AddDeltasResponse{}, resp.Err
}

func (g *grpcServer) AddDeltas(ctx context1.Context, req *pb.AddDeltasRequest) (*pb.AddDeltasResponse, error) {
	_, rep, err := g.addDeltas.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AddDeltasResponse), nil
}

// makeGetDeltasHandler creates the handler logic
func makeGetDeltasHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetDeltasEndpoint, decodeGetDeltasRequest, encodeGetDeltasResponse, options...)
}

// decodeGetDeltasResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetDeltas request.
func decodeGetDeltasRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetDeltasRequest)
	ownerID, err := uuid.FromString(req.OwnerID)
	if err != nil {
		return nil, err
	}
	return endpoint.GetDeltasRequest{
		OwnerID:       ownerID,
		ObjectID:      req.ObjectID,
		ObjectType:    req.ObjectType,
		FirstDate:     time.Unix(req.FirstDate, 0),
		LastDate:      time.Unix(req.LastDate, 0),
		UseCategories: req.UseCategories,
	}, nil
}

// encodeGetDeltasResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeGetDeltasResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.GetDeltasResponse)
	deltas := make([]*pb.Delta, len(resp.Deltas))
	for i := 0; i < len(resp.Deltas); i++ {
		delta := makeDeltaPb(resp.Deltas[i])
		deltas[i] = &delta
	}
	return &pb.GetDeltasResponse{
		Deltas: deltas,
	}, resp.Err
}
func (g *grpcServer) GetDeltas(ctx context1.Context, req *pb.GetDeltasRequest) (*pb.GetDeltasResponse, error) {
	_, rep, err := g.getDeltas.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetDeltasResponse), nil
}

// makeGetBalanceHandler creates the handler logic
func makeGetBalanceHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetBalanceEndpoint, decodeGetBalanceRequest, encodeGetBalanceResponse, options...)
}

// decodeGetBalanceResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetBalance request.
func decodeGetBalanceRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetBalanceRequest)
	ownerID, err := uuid.FromString(req.OwnerID)
	if err != nil {
		return nil, err
	}
	return endpoint.GetBalanceRequest{
		OwnerID: ownerID,
	}, nil
}

// encodeGetBalanceResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeGetBalanceResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.GetBalanceResponse)
	return &pb.GetBalanceResponse{
		Balance: resp.Balance,
	}, resp.Err
}
func (g *grpcServer) GetBalance(ctx context1.Context, req *pb.GetBalanceRequest) (*pb.GetBalanceResponse, error) {
	_, rep, err := g.getBalance.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetBalanceResponse), nil
}
