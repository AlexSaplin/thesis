package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"rhino/pkg/reporter"
	"time"

	uuid "github.com/satori/go.uuid"

	"rhino/pkg/clients/ibis"
	"rhino/pkg/entities"
	"rhino/pkg/errors"
	"rhino/pkg/runner"
)

// RhinoService describes the service.
type RhinoService interface {
	Run(ctx context.Context, functionID uuid.UUID, in []byte) (out []byte, err error)
}

type rhinoService struct {
	meta     ibis.Client
	registry runner.Registry
	reporter reporter.Reporter
	logger   log.Logger
}

func (s *rhinoService) Run(ctx context.Context, functionID uuid.UUID, in []byte) (out []byte, err error) {
	begin := time.Now()

	// Obtain model meta
	var function entities.Function
	if function, err = s.meta.GetFunction(ctx, functionID); err != nil {
		return
	}
	_ = s.logger.Log("msg", "got function", "method", "service.Run", "functionID", functionID, "error", err)

	// Check model state
	if function.State != entities.FunctionStateReady {
		if function.State == entities.FunctionStateDeleted {
			err = errors.ErrModelDeleted
			return
		}
		err = errors.ErrModelNotReady
		return
	}

	// Run
	functionRunner, err := s.registry.GetFunctionRunner(ctx, function)
	if err != nil {
		return nil, err
	}
	_ = s.logger.Log("msg", "got runner", "method", "service.Run", "functionID", functionID, "error", err)

	loadDuration := time.Since(begin)
	out, runDuration, err := functionRunner.Run(in)
	s.registry.Release(functionRunner)

	s.reporter.Submit(entities.RunReport{
		OwnerID:      function.OwnerID,
		FunctionID:   function.ID,
		RunDuration:  runDuration,
		LoadDuration: loadDuration,
		Timestamp:    time.Now(),
	})

	return
}

// NewRhinoService returns an implementation of RhinoService.
func NewRhinoService(meta ibis.Client, registry runner.Registry, rep reporter.Reporter, logger log.Logger) RhinoService {
	return &rhinoService{
		meta:     meta,
		registry: registry,
		reporter: rep,
		logger:   logger,
	}
}

// New returns a RhinoService with all of the expected middleware wired in.
func New(meta ibis.Client, registry runner.Registry, rep reporter.Reporter, logger log.Logger, middleware []Middleware) RhinoService {
	var svc = NewRhinoService(meta, registry, rep, logger)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
