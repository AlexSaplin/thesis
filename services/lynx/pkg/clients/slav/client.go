package slav

import (
	"context"
	"github.com/mattn/go-nulltype"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	slav "lynx/pkg/clients/slav/pb"
	"lynx/pkg/config"
	"lynx/pkg/entities"
)

type Client interface {
	CreateContainer(
		ctx context.Context, ownerID uuid.UUID, name string, scale uint32,
		instance entities.ContainerInstanceType, image string, port uint32,
		env map[string]string, auth []entities.ContainerAuth,
	) (container entities.Container, err error)
	UpdateContainer(
		ctx context.Context, ownerID uuid.UUID, name string, scale entities.NullUInt32,
		instance entities.NullInstance, image entities.NullString,
	) (container entities.Container, err error)
	GetContainer(ctx context.Context, ownerID uuid.UUID, name string) (container entities.ContainerFull, err error)
	DeleteContainer(ctx context.Context, ownerID uuid.UUID, name string) (err error)
	ListContainers(ctx context.Context, ownerID uuid.UUID) (containers []entities.Container, err error)
}

type GRPCClient struct {
	config config.SlavClientConfig
	client slav.SlavClient
}

func NewGRPCClient(ctx context.Context, config config.SlavClientConfig) (Client, error) {
	conn, err := grpc.DialContext(ctx, config.Target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &GRPCClient{
		client: slav.NewSlavClient(conn),
		config: config,
	}, nil
}

func (g *GRPCClient) CreateContainer(
	ctx context.Context, ownerID uuid.UUID, name string, scale uint32,
	instance entities.ContainerInstanceType, image string, port uint32,
	env map[string]string, auth []entities.ContainerAuth,
) (container entities.Container, err error) {
	var slavEnv []*slav.KV
	for k, v := range env {
		slavEnv = append(slavEnv, &slav.KV{
			Key:   k,
			Value: v,
		})
	}

	var slavAuth []*slav.AuthItem
	for _, item := range auth {
		slavAuth = append(slavAuth, &slav.AuthItem{
			Username: item.Username,
			Password: item.Password,
			Registry: item.Registry,
		})
	}

	req := &slav.CreateContainerRequest{
		OwnerID:  ownerID.String(),
		Name:     name,
		Scale:    scale,
		Instance: serializeInstanceTypeEntity(instance),
		Image:    image,
		Port:     port,
		Env:      slavEnv,
		Auth:     slavAuth,
	}
	resp, err := g.client.CreateContainer(ctx, req)
	if err != nil {
		return
	}
	return makeContainerEntity(resp.Container), nil
}

func (g *GRPCClient) UpdateContainer(
	ctx context.Context, ownerID uuid.UUID, name string, scale entities.NullUInt32,
	instance entities.NullInstance, image entities.NullString,
) (container entities.Container, err error) {
	req := &slav.UpdateContainerRequest{
		Name:    name,
		OwnerID: ownerID.String(),
		Scale: &slav.NullUInt32{
			Value:   scale.Value,
			IsValid: scale.IsValid,
		},
		Instance: &slav.NullInstanceType{
			Value:   serializeInstanceTypeEntity(instance.Value),
			IsValid: instance.IsValid,
		},
		Image: &slav.NullString{
			Value:   image.Value,
			IsValid: image.IsValid,
		},
	}
	resp, err := g.client.UpdateContainer(ctx, req)
	if err != nil {
		return
	}
	return makeContainerEntity(resp.Container), nil
}

func (g *GRPCClient) GetContainer(
	ctx context.Context, ownerID uuid.UUID, name string,
) (container entities.ContainerFull, err error) {
	req := &slav.GetContainerRequest{
		Name:    name,
		OwnerID: ownerID.String(),
	}
	resp, err := g.client.GetContainer(ctx, req)
	if err != nil {
		return
	}
	return makeContainerFullEntity(resp.Container, resp.State, resp.Error), nil
}

func (g *GRPCClient) DeleteContainer(
	ctx context.Context, ownerID uuid.UUID, name string,
) (err error) {
	req := &slav.DeleteContainerRequest{
		Name:    name,
		OwnerID: ownerID.String(),
	}
	_, err = g.client.DeleteContainer(ctx, req)
	if err != nil {
		return
	}
	return nil
}

func (g *GRPCClient) ListContainers(
	ctx context.Context, ownerID uuid.UUID,
) (containers []entities.Container, err error) {
	req := &slav.ListContainersRequest{
		OwnerID: ownerID.String(),
	}
	resp, err := g.client.ListContainers(ctx, req)
	if err != nil {
		return
	}
	for _, container := range resp.Containers {
		containers = append(containers, makeContainerEntity(container))
	}
	return containers, nil
}

func makeContainerEntity(in *slav.Container) (result entities.Container) {
	env := make(map[string]string)
	for _, item := range in.Env {
		env[item.Key] = item.Value
	}
	return entities.Container{
		Name:         in.Name,
		Scale:        in.Scale,
		InstanceType: parseInstanceTypeEntity(in.Instance),
		Image:        in.Image,
		Port:         in.Port,
		URL:          in.URL,
		Env:          env,
		Auth:         in.Auth,
	}
}

func makeContainerFullEntity(in *slav.Container, state slav.StateType, error string) (result entities.ContainerFull) {
	var errStr nulltype.NullString
	if error != "" {
		errStr = nulltype.NullStringOf(error)
	}
	env := make(map[string]string)
	for _, item := range in.Env {
		env[item.Key] = item.Value
	}

	return entities.ContainerFull{
		Name:         in.Name,
		State:        parseStateTypeEntity(state),
		Scale:        in.Scale,
		InstanceType: parseInstanceTypeEntity(in.Instance),
		Image:        in.Image,
		Port:         in.Port,
		URL:          in.URL,
		Error:        errStr,
		Env:          env,
		Auth:         in.Auth,
	}
}
