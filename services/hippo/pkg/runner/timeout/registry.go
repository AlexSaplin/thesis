package timeout

import (
	"context"
	"sync"
	"time"

	"github.com/go-kit/kit/log"

	"hippo/pkg/clients/selachii"
	"hippo/pkg/entities"
	"hippo/pkg/runner"
)

const (
	unloadTimeout = time.Second * 30
	modelChanSize = 1024
	checkInactiveInterval = time.Second * 30
)

type Registry struct {

	client  selachii.SelachiiClient

	runners         sync.Map
	inactiveRunners chan *Runner

	logger log.Logger

	ctx context.Context
}

func NewRegistry(client selachii.SelachiiClient, logger log.Logger) *Registry {

	r := &Registry{
		client: client,
		ctx: context.Background(),
		logger: logger,
		inactiveRunners: make(chan *Runner, modelChanSize),
	}
	go r.checkInactive()
	go r.dropInactive()

	return r
}

func (r *Registry) GetModelRunner(ctx context.Context, model entities.Model) runner.ModelRunner {
	if rc, ok := r.runners.Load(model.ID); ok {
		chanLoop:
		for {
			select {
			case modelRunner := <- rc.(chan *Runner):
				if !modelRunner.active {
					_ = r.logger.Log("msg", "model inactive via get", "modelID", model.ID)
					r.inactiveRunners <- modelRunner
					continue
				}
				modelRunner.requestCtx = ctx
				return modelRunner
			default:
				break chanLoop
			}
		}
	}

	modelRunner := &Runner{
		runnerCtx:  r.ctx,
		requestCtx: ctx,
		model:      model,
		root:       r,
	}
	return modelRunner
}

func (r *Registry) Release(modelRunnerInt runner.ModelRunner) {
	modelRunner, ok := modelRunnerInt.(*Runner)
	if !ok {
		panic("trying to release different implementation of ModelRunner")
	}

	modelRunner.requestCtx = nil

	rc, ok := r.runners.Load(modelRunner.model.ID)
	if !ok {
		rc, _ = r.runners.LoadOrStore(modelRunner.model.ID, make(chan *Runner, modelChanSize))
	}
	select {
	case rc.(chan *Runner) <- modelRunner:
	default:
		r.inactiveRunners <- modelRunner // drop extra runner if the channel is overflowing
	}
}


func (r *Registry) checkInactive() {
	for {
		select {
		case <- time.After(checkInactiveInterval):
			r.runners.Range(func(key, value interface{}) bool {
				rc := value.(chan *Runner)
				rcLen := len(rc)
				rcLoop:
				for i := 0; i < rcLen; i++ {
					select {
					case item := <- rc:
						if item.active {
							rc <- item
						} else {
							_ = r.logger.Log("msg", "model inactive via check", "modelID", item.model.ID)
							r.inactiveRunners <- item
						}
					default:
						break rcLoop
					}
				}
				return true
			})
		}
	}
}

func (r *Registry) dropInactive() {
	for {
		select {
		case <- r.ctx.Done():
			r.logger.Log("msg", "dropInactive terminating")
			return
		case item := <- r.inactiveRunners:
			r.logger.Log("msg", "unloading model", "modelID", item.model.ID)
			err := item.unload()
			if err != nil {
				r.logger.Log("msg", "failed to unload model", "err", err, "modelID", item.model.ID)
			}
			r.logger.Log("msg", "unloaded inactive model", "modelID", item.model.ID)
		}
	}
}