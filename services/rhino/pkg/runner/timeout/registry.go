package timeout

import (
	"context"
	"sync"
	"time"

	"github.com/go-kit/kit/log"
	"golang.org/x/sync/semaphore"

	"rhino/pkg/clients/docker"
	"rhino/pkg/config"
	"rhino/pkg/entities"
	"rhino/pkg/runner"
)

const (
	unloadTimeout         = time.Second * 120
	inactiveChanSize      = 1024
	checkInactiveInterval = time.Second * 30
	loadTimeout           = time.Second * 30
)

type registryItem struct {
	sema    *semaphore.Weighted
	runners chan *Runner
}

type Registry struct {
	clients []docker.Client
	config  config.RunnerConfig

	totalClients    chan docker.Client
	runners         sync.Map
	inactiveRunners chan *Runner

	logger log.Logger

	cleanupLock sync.RWMutex

	ctx context.Context
}

func NewRegistry(cfg config.RunnerConfig, logger log.Logger) (*Registry, error) {

	var clients []docker.Client
	for _, dockerCfg := range cfg.Clients {
		dockerClient, err := docker.NewDockerClient(dockerCfg, logger)
		if err != nil {
			return nil, err
		}
		clients = append(clients, dockerClient)
	}

	r := &Registry{
		config:          cfg,
		clients:         clients,
		ctx:             context.Background(),
		logger:          logger,
		inactiveRunners: make(chan *Runner, inactiveChanSize),
		totalClients:    make(chan docker.Client, len(clients)),
	}
	for i := range clients {
		r.totalClients <- clients[i] // could be replaced with structs bound to specific runner nodes
	}

	go r.checkInactive()
	go r.dropInactive()

	return r, nil
}

func (r *Registry) GetFunctionRunner(ctx context.Context, fn entities.Function) (runner.FunctionRunner, error) {

	// Load queue for given function
	fnQeueueInt, ok := r.runners.Load(fn.ID)
	if !ok {
		fnQeueueInt, _ = r.runners.LoadOrStore(fn.ID, r.newRegistryItem())
	}
	fnQeueue := fnQeueueInt.(*registryItem)

	// Maximum functions of same kind running simultaneously
	err := fnQeueue.sema.Acquire(ctx, 1)
	if err != nil {
		return nil, err
	}

	// Cleanup lock
	r.cleanupLock.RLock()
	defer r.cleanupLock.RUnlock()

	for {
		select {
		case fnRunner := <-fnQeueue.runners:
			if !fnRunner.active {
				_ = r.logger.Log("msg", "function inactive via get", "functionID", fn.ID)
				r.inactiveRunners <- fnRunner
				continue
			}
			fnRunner.requestCtx = ctx
			return fnRunner, nil
		default:
			select {
			case client := <-r.totalClients:
				rCtx, cancel := context.WithCancel(r.ctx)
				return &Runner{
					client:       client,
					runnerCtx:    rCtx,
					runnerCancel: cancel,
					requestCtx:   ctx,
					function:     fn,
					root:         r,
				}, nil
			case <-ctx.Done():
				fnQeueue.sema.Release(1)
				return nil, ctx.Err()
			default:
				time.Sleep(time.Millisecond * 50) // KILL ME
			}
		}
	}
}

func (r *Registry) Release(functionRunnerInt runner.FunctionRunner) {
	// Cleanup lock
	r.logger.Log("msg", "releasing runner")
	r.cleanupLock.RLock()
	defer r.cleanupLock.RUnlock()
	r.logger.Log("msg", "releasing runner got cleanup lock")
	functionRunner, ok := functionRunnerInt.(*Runner)
	if !ok {
		panic("trying to release different implementation of FunctionRunner")
	}

	functionRunner.requestCtx = nil

	fnQeueueInt, ok := r.runners.Load(functionRunner.function.ID)
	if !ok {
		panic("releasing to functions without creating")
	}
	fnQeueue := fnQeueueInt.(*registryItem)
	loadID := "not_loaded"
	if functionRunner.runningFunction != nil {
		loadID = functionRunner.runningFunction.ID
	}

	r.logger.Log("msg", "releasing runner got fnQeueueInt", "loadID", loadID)
	select {
	case fnQeueue.runners <- functionRunner:
	default:
		r.logger.Log("msg", "dropping extra runner, should not be happening", "loadID", loadID)
		r.inactiveRunners <- functionRunner // drop extra runner if the channel is overflowing
	}

	fnQeueue.sema.Release(1)
	r.logger.Log("msg", "released sema runner got fnQeueueInt", "loadID", loadID)
}

func (r *Registry) newRegistryItem() *registryItem {
	return &registryItem{
		sema:    semaphore.NewWeighted(int64(r.config.MaxClientsPerFunction)),
		runners: make(chan *Runner, r.config.MaxClientsPerFunction),
	}
}

func (r *Registry) checkInactive() {
	for {
		select {
		case <-time.After(checkInactiveInterval):
			r.logger.Log("msg", "initiating cleanup")
			r.cleanupLock.Lock()
			r.logger.Log("msg", "cleanup lock aquired")
			r.runners.Range(func(key, value interface{}) bool {
				rc := value.(*registryItem)
				rcLen := len(rc.runners)
			rcLoop:
				for i := 0; i < rcLen; i++ {
					select {
					case item := <-rc.runners:
						if item.active {
							rc.runners <- item
						} else {
							_ = r.logger.Log("msg", "function inactive via check", "functionID", item.function.ID)
							r.inactiveRunners <- item
						}
					default:
						break rcLoop
					}
				}
				return true
			})
			r.cleanupLock.Unlock()
			r.logger.Log("msg", "cleanup complete")
		}
	}
}

func (r *Registry) dropInactive() {
	for {
		select {
		case <-r.ctx.Done():
			r.logger.Log("msg", "dropInactive terminating")
			return
		case item := <-r.inactiveRunners:
			r.logger.Log("msg", "unloading function", "functionID", item.function.ID)
			err := item.unload()
			if err != nil {
				r.logger.Log("msg", "failed to unload function", "err", err, "functionID", item.function.ID)
			}
			r.totalClients <- item.client
			r.logger.Log("msg", "unloaded inactive model", "functionID", item.function.ID)
		}
	}
}
