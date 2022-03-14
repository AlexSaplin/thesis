package rhino

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	rhino "lynx/pkg/clients/rhino/pb"
	"lynx/pkg/config"
	"lynx/pkg/entities"
)

const maxMsgSizeBytes = 1024 * 1024 * 64

type Client interface {
	Run(ctx context.Context, function entities.Function, in []byte) ([]byte, error)
}

type GRPCClient struct {
	config config.RhinoClientConfig
	client rhino.RhinoClient
}

func NewGRPCClient(ctx context.Context, config config.RhinoClientConfig) (Client, error) {
	conn, err := grpc.DialContext(ctx,
		config.Target,
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(maxMsgSizeBytes),
			grpc.MaxCallSendMsgSize(maxMsgSizeBytes),
		),
	)
	if err != nil {
		return nil, err
	}
	return &GRPCClient{
		client: rhino.NewRhinoClient(conn),
		config: config,
	}, nil
}

func (c *GRPCClient) Run(ctx context.Context, function entities.Function, in []byte) (out []byte, err error) {

	req := &rhino.RunRequest{
		FunctionID: function.ID.String(),
		Data: in,
	}
	var resp *rhino.RunReply
	resp, err = c.client.Run(ctx, req,
		grpc.MaxCallRecvMsgSize(maxMsgSizeBytes),
		grpc.MaxCallSendMsgSize(maxMsgSizeBytes),
	)
	if err != nil {
		return
	}
	out = resp.Data
	return
}
