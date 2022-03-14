package basic

import (
	"context"

	"hippo/pkg/clients/selachii"
	"hippo/pkg/entities"
	"hippo/pkg/runner"
)

type Registry struct {
	client selachii.SelachiiClient
}

func NewRegistry(client selachii.SelachiiClient) *Registry {
	return &Registry{
		client: client,
	}
}

func (b *Registry) GetModelRunner(ctx context.Context, model entities.Model) runner.ModelRunner {
	return &Runner{
		ctx:   ctx,
		model: model,
		root:  b,
	}
}

func (r *Registry) Release(modelRunnerInt runner.ModelRunner) {}
