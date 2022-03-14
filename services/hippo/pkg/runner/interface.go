package runner

import (
	"context"
	"time"

	"hippo/pkg/entities"
)

type Registry interface {
	GetModelRunner(ctx context.Context, model entities.Model) ModelRunner
	Release(runner ModelRunner)
}

type ModelRunner interface {
	Run(in entities.TensorList) (entities.TensorList, time.Duration, error)
}
