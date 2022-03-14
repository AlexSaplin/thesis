package docker

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
	"rhino/pkg/config"
	"rhino/pkg/entities"
	"strconv"
	"time"
)

type RunningFunction struct {
	ID       string
	endpoint string
	parent   *dockerClient
}

var portId = 0

type Client interface {
	StartFunction(ctx context.Context, function entities.Function) (RunningFunction, error)
	StopFunction(ctx context.Context, function RunningFunction) error
	CallFunction(ctx context.Context, function RunningFunction, in []byte) ([]byte, error)
	StopAllFunctions(ctx context.Context) error
}

type dockerClient struct {
	logger     log.Logger
	cfg        config.DockerClientConfig
	cli        *rawDockerClient
	httpClient *http.Client
	created    int
}

func NewDockerClient(cfg config.DockerClientConfig, logger log.Logger) (Client, error) {
	cli, err := newRawDockerClient(cfg.Target)
	if err != nil {
		return nil, err
	}
	httpCli := &http.Client{
		Transport: &http.Transport{
			MaxConnsPerHost: 1,
		},
		Timeout: time.Duration(300) * time.Second,
	}

	err = cli.clearContainers(context.Background())
	if err != nil {
		return nil, err
	}

	return &dockerClient{
		cli:        cli,
		cfg:        cfg,
		httpClient: httpCli,
		logger:     logger,
	}, nil
}

func (d *dockerClient) StartFunction(ctx context.Context, function entities.Function) (res RunningFunction, err error) {
	runID := uuid.FromBytesOrNil(append(function.ID.Bytes()[:8], uuid.NewV4().Bytes()[:8]...))
	logger := log.With(d.logger, "name", function.Name, "ID", function.ID.String(), "loadID", runID.String())
	logger.Log("msg", "pulling container")
	// pull image
	err = d.cli.pullImage(ctx, function.ImageURL.String())
	if err != nil {
		return
	}
	logger.Log("msg", "starting container")
	// start container
	hostPort := strconv.Itoa(29000 + (portId % 5000))
	portId++

	err = d.cli.runContainer(ctx, function.ImageURL.String(), runID.String(), "80", hostPort, function.ID.String())
	if err != nil {
		return
	}
	logger.Log("msg", "verifying container")

	host := fmt.Sprintf("%s:%s", d.cfg.ContainerHost, hostPort)
	validateEndpont := fmt.Sprintf("%s%s", host, "/load")
	runEndpont := fmt.Sprintf("%s%s", host, "/run")

	// verify function is started correctly
	ok := false
	for i := 0; i <= 10; i++ {
		var resp *http.Response
		resp, err = http.Get(validateEndpont)
		if err != nil {
			logger.Log("msg", "backing off", "endpoint", validateEndpont, "err", err)
			time.Sleep(time.Duration(i+1) * 100 * time.Millisecond)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			var body []byte
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				_ = resp.Body.Close()
				return
			}
			errBody := fmt.Sprintf("function failed to start: %s", string(body))
			err = status.Error(codes.InvalidArgument, errBody)
			removeErr := d.cli.removeContainer(ctx, runID.String())
			if removeErr != nil {
				logger.Log("msg", "failed to remove container that didn't start properly", "err", removeErr)
			}
			_ = resp.Body.Close()
			return
		}
		ok = true
		break
	}
	if !ok {
		removeErr := d.cli.removeContainer(ctx, runID.String())
		if removeErr != nil {
			logger.Log("msg", "failed to remove container that didn't start properly", "err", removeErr)
		}
		errStr := "failed to initialize container with all retries"
		logger.Log("msg", errStr)
		err = status.Error(codes.Internal, errStr)
		return
	}

	res = RunningFunction{
		ID:       runID.String(),
		endpoint: runEndpont,
		parent:   d,
	}
	logger.Log("msg", "created container")
	return
}

func (d *dockerClient) StopFunction(ctx context.Context, function RunningFunction) error {
	logger := log.With(d.logger, "loadID", function.ID)
	logger.Log("msg", "stopping container")
	defer logger.Log("msg", "stopped container")
	return d.cli.removeContainer(ctx, function.ID)
}

func (d *dockerClient) CallFunction(ctx context.Context, function RunningFunction, in []byte) ([]byte, error) {
	logger := log.With(d.logger, "loadID", function.ID)
	logger.Log("msg", "calling function")
	req, err := http.NewRequestWithContext(ctx, "POST", function.endpoint, bytes.NewReader(in))
	if err != nil {
		return nil, err
	}
	resp, err := d.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		logger.Log("msg", "called function successfully")
		return respBody, nil
	}
	if resp.StatusCode == http.StatusBadRequest {
		logger.Log("msg", "called function with function failure")
		errStr := fmt.Sprintf("function call failed: %s", string(respBody))
		return nil, status.Error(codes.InvalidArgument, errStr)
	}
	errStr := fmt.Sprintf("function call failed with internal error: %s", string(respBody))
	logger.Log("msg", "called function with internal failure")
	return nil, status.Error(codes.Internal, errStr)
}

func (d *dockerClient) StopAllFunctions(ctx context.Context) error {
	d.logger.Log("msg", "stopping all container")
	defer d.logger.Log("msg", "stopped all containers")
	return d.cli.clearContainers(ctx)
}
