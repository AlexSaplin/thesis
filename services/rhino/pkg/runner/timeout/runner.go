package timeout

import (
	"context"
	"rhino/pkg/clients/docker"
	"time"

	"rhino/pkg/entities"
)

type Runner struct {
	client docker.Client

	runnerCtx    context.Context
	runnerCancel context.CancelFunc

	requestCtx context.Context

	runningFunction *docker.RunningFunction
	function        entities.Function

	root *Registry

	deadline time.Time
	active   bool
}

func (r *Runner) Run(in []byte) (result []byte, duration time.Duration, err error) {
	runningFunction, err := r.ensureLoaded()
	if err != nil {
		return
	}
	runBegin := time.Now()
	result, err = r.client.CallFunction(r.requestCtx, runningFunction, in)
	if err != nil {
		return
	}
	duration = time.Since(runBegin)

	r.deadline = time.Now().Add(unloadTimeout)
	r.active = true
	return
}

func (r *Runner) ensureLoaded() (res docker.RunningFunction, err error) {
	if r.runningFunction != nil {
		return *r.runningFunction, nil
	}
	r.root.logger.Log("msg", "loading function", "functionID", r.function.ID)

	var rf docker.RunningFunction
	loadCtx, cancel := context.WithTimeout(context.Background(), loadTimeout)
	defer cancel()

	rf, err = r.client.StartFunction(loadCtx, r.function)
	if err != nil {
		r.root.logger.Log("msg", "loading function failed", "functionID", r.function.ID, "err", err)
		return
	}

	r.deadline = time.Now().Add(unloadTimeout)
	r.active = true

	go r.dropfunctionAfterTimeout(unloadTimeout)
	r.runningFunction = &rf
	r.root.logger.Log("msg", "loading function successful", "functionID", r.function.ID)
	return rf, nil
}

func (r *Runner) dropfunctionAfterTimeout(timeout time.Duration) {
	for {
		select {
		case <-r.runnerCtx.Done():
			return
		case <-time.After(timeout):
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
	if r.runningFunction != nil {
		r.root.logger.Log("msg", "actually unloading function", "loadID", r.runningFunction.ID)
		err = r.client.StopFunction(r.runnerCtx, *r.runningFunction)
		r.runnerCancel()
	}
	return
}
