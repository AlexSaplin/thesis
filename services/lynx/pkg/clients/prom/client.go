package prom

import (
	"fmt"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"golang.org/x/net/context"
	"lynx/pkg/config"
	"time"
)

type PromClient struct {
	cfg       config.PromClientConfig
	rawClient api.Client
}

type Client interface {
	GetFunctionMetrics(ctx context.Context, functionID string, days int) (FullMetricsResponse, error)
}

func NewPromClient(cfg config.PromClientConfig) (Client, error) {
	cli, err := api.NewClient(api.Config{
		Address: cfg.Target,
	})
	if err != nil {
		return nil, err
	}
	return &PromClient{
		cfg:       cfg,
		rawClient: cli,
	}, nil

}
func (c *PromClient) GetFunctionMetrics(ctx context.Context, functionID string, days int) (FullMetricsResponse, error) {
	v1api := v1.NewAPI(c.rawClient)
	r := v1.Range{
		Start: time.Now().Add(-time.Duration(days) * 24 * time.Hour),
		End:   time.Now(),
		Step:  time.Hour * 24,
	}
	resultRdur, warnings, err := v1api.QueryRange(ctx, fmt.Sprintf("run_duration_sum{function_id=\"%s\"}", functionID), r)
	if err != nil {
		return FullMetricsResponse{}, err
	}
	if len(warnings) != 0 {
	}

	resultRcnt, warnings, err := v1api.QueryRange(ctx, fmt.Sprintf("run_duration_count{function_id=\"%s\"}", functionID), r)
	if err != nil {
		return FullMetricsResponse{}, err
	}
	if len(warnings) != 0 {
	}

	resultBdur, warnings, err := v1api.QueryRange(ctx, fmt.Sprintf("load_duration_sum{function_id=\"%s\"}", functionID), r)
	if err != nil {
		return FullMetricsResponse{}, err
	}
	if len(warnings) != 0 {
	}

	resultBcnt, warnings, err := v1api.QueryRange(ctx, fmt.Sprintf("load_duration_count{function_id=\"%s\"}", functionID), r)
	if err != nil {
		return FullMetricsResponse{}, err
	}
	if len(warnings) != 0 {
	}

	response := NewFullMetricsResponse(resultRdur, resultRcnt, resultBdur, resultBcnt)
	if len(response.Metrics) < days {
		response.Metrics = append(make([]MetricsResponse, days-len(response.Metrics)), response.Metrics...)
	}

	return response, err
}
