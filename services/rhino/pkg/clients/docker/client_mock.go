package docker

import (
	"context"
	"math/rand"
	"time"

	"github.com/go-kit/kit/log"

	"rhino/pkg/entities"
)

type mockDockerClient struct {
	logger log.Logger
}

func (m mockDockerClient) StartFunction(ctx context.Context, function entities.Function) (RunningFunction, error) {
	time.Sleep(time.Second * time.Duration(rand.Intn(5)))
	m.logger.Log("msg", "started function")
	return RunningFunction{}, nil
}

func (m mockDockerClient) StopFunction(ctx context.Context, function RunningFunction) error {
	time.Sleep(time.Second * time.Duration(rand.Intn(2)))
	m.logger.Log("msg", "stopped function")
	return nil
}

func (m mockDockerClient) CallFunction(ctx context.Context, function RunningFunction, in []byte) ([]byte, error) {
	time.Sleep(time.Second * time.Duration(rand.Intn(5)))
	m.logger.Log("msg", "stopped function")
	return in, nil
}

func (m mockDockerClient) StopAllFunctions(ctx context.Context) error {
	panic("implement me")
}
