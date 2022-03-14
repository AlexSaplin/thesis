package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"io/ioutil"
	"log"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	natting "github.com/docker/go-connections/nat"
)

type rawDockerClient struct {
	cli *client.Client
}

func newRawDockerClient(target string) (*rawDockerClient, error) {
	cli, err := client.NewClientWithOpts(client.WithHost(target))
	if err != nil {
		return nil, err
	}
	return &rawDockerClient{cli: cli}, nil
}

func (dc *rawDockerClient) pullImage(ctx context.Context, image string) error {
	ans, err := dc.cli.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	resp, err := ioutil.ReadAll(ans)
	if err != nil {
		return err
	}
	x := string(resp)
	fmt.Println(x)
	_ = ans.Close()
	return nil
}

func (dc *rawDockerClient) runContainer(ctx context.Context, imagename, containername, port, hostPort string, functionID string) error {

	newport, err := natting.NewPort("tcp", port)
	if err != nil {
		fmt.Println("Unable to create docker port")
		return err
	}

	hostConfig := &container.HostConfig{
		PortBindings: natting.PortMap{
			newport: []natting.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: hostPort,
				},
			},
		},
		LogConfig: container.LogConfig{
			Type: "fluentd",
			Config: map[string]string{
				"tag": functionID,
			},
		},
		RestartPolicy: container.RestartPolicy{
			Name: "always",
		},
		Resources: container.Resources{
			Memory:   1024 * 1024 * 1024 * 12, // 12 gb
			NanoCPUs: 4_000_000_000, // 4 cores
			DeviceRequests: []container.DeviceRequest{
				{
					// Driver:       "",
					Count:        1,
					// DeviceIDs:    []string{"0"},
					Capabilities: [][]string{
						{
							"gpu",
						},
					},
					// Options:      nil,
				},
			},
		},
	}

	networkConfig := &network.NetworkingConfig{}

	exposedPorts := map[natting.Port]struct{}{
		newport: {},
	}

	config := &container.Config{
		Image:        imagename,
		ExposedPorts: exposedPorts,
		Labels: map[string]string{
			"deepmux-functions": "",
		},
	}

	cont, err := dc.cli.ContainerCreate(
		ctx,
		config,
		hostConfig,
		networkConfig,
		containername,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return dc.cli.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{})
}

func (dc *rawDockerClient) removeContainer(ctx context.Context, ID string) error {
	timeout := time.Second * 10
	err := dc.cli.ContainerStop(ctx, ID, &timeout)
	if err != nil {
		return err
	}

	err = dc.cli.ContainerRemove(ctx, ID, types.ContainerRemoveOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (dc *rawDockerClient) clearContainers(ctx context.Context) error {
	items, err := dc.cli.ContainerList(ctx, types.ContainerListOptions{
		Filters: filters.NewArgs(filters.Arg("label", "deepmux-functions")),
	})
	if err != nil {
		return err
	}
	for _, v := range items {
		err = dc.removeContainer(ctx, v.ID)
		if err != nil {
			return err
		}
		fmt.Println("removed", v.ID)
	}
	return nil
}
