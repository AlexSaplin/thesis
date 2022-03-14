package hippo

import (
	"bytes"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"lynx/pkg/clients/hippo/pb"
	"lynx/pkg/config"
	"lynx/pkg/entities"
)

const maxMsgSizeBytes = 1024 * 1024 * 64

type Client interface {
	Run(ctx context.Context, model entities.Model, tensor entities.TensorList) ([]byte, [][]int64, error)
}

type GRPCClient struct {
	config config.HippoClientConfig
	client hippo.HippoClient
}

func NewGRPCClient(ctx context.Context, config config.HippoClientConfig) (Client, error) {
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
		client: hippo.NewHippoClient(conn),
		config: config,
	}, nil
}

func (c *GRPCClient) Run(
	ctx context.Context, model entities.Model, tensors entities.TensorList,
) (result []byte, shape [][]int64, err error) {

	req := &hippo.RunRequest{
		ModelID: model.ID.String(),
		Tensors: serializeTensorListEntity(tensors),
	}
	resp, err := c.client.Run(ctx, req,
		grpc.MaxCallRecvMsgSize(maxMsgSizeBytes),
		grpc.MaxCallSendMsgSize(maxMsgSizeBytes),
	)
	if err != nil {
		return
	}

	var bufResult bytes.Buffer

	for _, item := range resp.Tensors {
		_, err = bufResult.Write(item.Data)
		if err != nil {
			return
		}
		shape = append(shape, item.Shape.GetValue())
	}
	result = bufResult.Bytes()
	return
}
