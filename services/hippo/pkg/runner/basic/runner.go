package basic

import (
	"context"
	"time"

	"hippo/pkg/entities"
)

type Runner struct {
	ctx      context.Context
	model    entities.Model
	root *Registry
}

func (r *Runner) Run(in entities.TensorList) (result entities.TensorList, duration time.Duration, err error) {
	loadID, err := r.root.client.LoadModel(r.ctx, r.model)
	if err != nil {
		return
	}
	runBegin := time.Now()
	result, err = r.root.client.Run(r.ctx, loadID, in)
	if err != nil {
		return
	}
	duration = time.Since(runBegin)

	_, err = r.root.client.UnloadModel(r.ctx, loadID)
	if err != nil {
		return
	}
	return
}
