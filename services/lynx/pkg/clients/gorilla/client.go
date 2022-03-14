package gorilla

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"

	gorilla "lynx/pkg/clients/gorilla/pb"
	"lynx/pkg/config"
)

type Client interface {
	GetBalance(ctx context.Context, ownerID uuid.UUID) (balance float64, err error)
}

func NewGRPCClient(ctx context.Context, cfg config.GorillaClientConfig) (Client, error) {
	conn, err := grpc.DialContext(ctx,
		cfg.Target,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}
	return &gRPCClient{
		config: cfg,
		client: gorilla.NewGorillaClient(conn),
	}, nil
}

type gRPCClient struct {
	config config.GorillaClientConfig
	client gorilla.GorillaClient
}

func (g *gRPCClient) GetBalance(ctx context.Context, ownerID uuid.UUID) (balance float64, err error) {
	resp, err := g.client.GetBalance(ctx,
		&gorilla.GetBalanceRequest{
			OwnerID: ownerID.String(),
		})
	if err != nil {
		return
	}
	balance = resp.Balance
	return
}
