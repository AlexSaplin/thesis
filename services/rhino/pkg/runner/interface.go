package runner

import (
	"context"
	"time"

	"rhino/pkg/entities"
)

type Registry interface {
	GetFunctionRunner(ctx context.Context, model entities.Function) (FunctionRunner, error)
	Release(runner FunctionRunner)
}

type FunctionRunner interface {
	Run(in []byte) (result []byte, duration time.Duration, err error)
}
