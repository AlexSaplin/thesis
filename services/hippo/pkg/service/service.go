package service

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"

	"hippo/pkg/clients/ardea"
	"hippo/pkg/entities"
	"hippo/pkg/errors"
	"hippo/pkg/reporter"
	"hippo/pkg/runner"
)

// HippoService describes the service.
type HippoService interface {
	Run(ctx context.Context, modelID uuid.UUID, in entities.TensorList) (out entities.TensorList, err error)
}

type hippoService struct {
	meta     ardea.Client
	registry runner.Registry
	reporter reporter.Reporter
}

func (s *hippoService) Run(
	ctx context.Context, modelID uuid.UUID, in entities.TensorList,
) (out entities.TensorList, err error) {
	if err = in.Valid(); err != nil {
		return
	}

	// Obtain model meta
	var model entities.Model
	if model, err = s.meta.GetModel(ctx, modelID); err != nil {
		return
	}

	// Check model state
	if model.State != entities.ModelStateReady {
		if model.State == entities.ModelStateDeleted {
			err = errors.ErrModelDeleted
			return
		}
		err = errors.ErrModelNotReady
		return
	}

	// Validate input shape
	if !in.ConformsToShape(model.InputShape) {
		err = errors.ErrShapeMismatch
		return
	}

	// Run
	modelRunner := s.registry.GetModelRunner(ctx, model)
	out, runDuration, err := modelRunner.Run(in)
	s.registry.Release(modelRunner)

	s.reporter.Submit(entities.RunReport{
		OwnerID:   model.OwnerID,
		ModelID:   model.ID,
		Duration:  runDuration,
		Timestamp: time.Now(),
	})

	return
}

// NewHippoService returns an implementation of HippoService.
func NewHippoService(meta ardea.Client, registry runner.Registry, rep reporter.Reporter) HippoService {
	return &hippoService{
		meta:     meta,
		registry: registry,
		reporter: rep,
	}
}

// New returns a HippoService with all of the expected middleware wired in.
func New(meta ardea.Client, registry runner.Registry, rep reporter.Reporter, middleware []Middleware) HippoService {
	var svc = NewHippoService(meta, registry, rep)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
