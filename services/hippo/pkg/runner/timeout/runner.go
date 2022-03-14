package timeout

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"

	"hippo/pkg/entities"
)

type Runner struct {
	runnerCtx  context.Context
	requestCtx context.Context

	loadID uuid.NullUUID
	model  entities.Model
	root   *Registry

	deadline time.Time
	active   bool
}

func (r *Runner) Run(in entities.TensorList) (result entities.TensorList, duration time.Duration, err error) {
	loadID, err := r.ensureLoaded()
	if err != nil {
		return
	}
	runBegin := time.Now()
	result, err = r.root.client.Run(r.requestCtx, loadID, in)
	if err != nil {
		return
	}
	duration = time.Since(runBegin)

	r.deadline = time.Now().Add(unloadTimeout)
	r.active = true
	return
}

func (r *Runner) ensureLoaded() (loadID uuid.UUID, err error) {
	if r.loadID.Valid {
		loadID = r.loadID.UUID
		return
	}
	r.root.logger.Log("msg", "loading model", "modelID", r.model.ID)

	loadID, err = r.root.client.LoadModel(r.requestCtx, r.model)
	if err != nil {
		r.root.logger.Log("msg", "loading model failed", "modelID", r.model.ID, "err", err)
		return
	}

	r.deadline = time.Now().Add(unloadTimeout)

	go r.dropModelAfterTimeout(unloadTimeout)
	r.loadID = uuid.NullUUID{Valid:true, UUID: loadID}
	r.root.logger.Log("msg", "loading model successful", "modelID", r.model.ID)
	return
}


func (r *Runner) dropModelAfterTimeout(timeout time.Duration) {
	for {
		select {
		case <- r.runnerCtx.Done():
			return
		case <- time.After(timeout):
			if r.deadline.After(time.Now()) {
				timeout = time.Until(r.deadline)
				continue
			}
			r.active = false
			timeout = time.Second
		}
	}
}


func (r *Runner) unload() (err error) {
	if r.loadID.Valid {
		r.root.logger.Log("msg", "actually unloading model", "loadID", r.loadID.UUID)
		_, err = r.root.client.UnloadModel(r.runnerCtx, r.loadID.UUID)
	}
	return
}
