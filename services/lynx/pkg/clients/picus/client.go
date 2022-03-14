package picus

import (
	"context"
	"google.golang.org/grpc"
	"io"
	picus "lynx/pkg/clients/picus/pb"
	"lynx/pkg/config"
	"lynx/pkg/entities"
)

type Client interface {
	GetOnlineLogs(
		ctx context.Context, function entities.FunctionQuery, ch chan LogResponse,
	) error
	GetFunctionLogs(ctx context.Context, function entities.FunctionQuery) (FullLogResponse, error)
}

type GRPCClient struct {
	config config.PicusClientConfig
	client picus.PicusClient
}

func NewGRPCClient(ctx context.Context, config config.PicusClientConfig) (Client, error) {
	conn, err := grpc.DialContext(ctx, config.Target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &GRPCClient{
		client: picus.NewPicusClient(conn),
		config: config,
	}, nil
}

func (c *GRPCClient) GetOnlineLogs(ctx context.Context, function entities.FunctionQuery,
	ch chan LogResponse) error {
	req := &picus.StreamFunctionLogsRequest{FunctionId: function.ID.String()}
	stream, err := c.client.StreamFunctionLogs(ctx, req)
	if err != nil {
		return err
	}
	for {
		logEnt, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		ch <- LogResponse{
			Timestamp: *logEnt.Time,
			Message:   logEnt.Message,
		}
	}
	close(ch)
	return nil
}

func (c *GRPCClient) GetFunctionLogs(ctx context.Context, function entities.FunctionQuery) (FullLogResponse, error) {
	resp, err := c.client.GetFunctionLogs(ctx, &picus.GetFunctionLogsRequest{
		FunctionId: function.ID.String(),
	})
	if err != nil {
		return FullLogResponse{}, err
	}
	return newFullLogResponse(resp.Entries), nil
}
