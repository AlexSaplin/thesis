package endpoint

import (
	service "ardea/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	CreateModelEndpoint      endpoint.Endpoint
	GetModelEndpoint         endpoint.Endpoint
	GetModelByNameEndpoint   endpoint.Endpoint
	UpdateModelStateEndpoint endpoint.Endpoint
	UpdateModelPathEndpoint  endpoint.Endpoint
	ListModelsEndpoint       endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.ArdeaService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		CreateModelEndpoint:      MakeCreateModelEndpoint(s),
		GetModelEndpoint:         MakeGetModelEndpoint(s),
		GetModelByNameEndpoint:   MakeGetModelByNameEndpoint(s),
		UpdateModelPathEndpoint:  MakeUpdateModelPathEndpoint(s),
		UpdateModelStateEndpoint: MakeUpdateModelStateEndpoint(s),
		ListModelsEndpoint:       MakeListModelsEndpoint(s),
	}
	for _, m := range mdw["CreateModel"] {
		eps.CreateModelEndpoint = m(eps.CreateModelEndpoint)
	}
	for _, m := range mdw["GetModel"] {
		eps.GetModelEndpoint = m(eps.GetModelEndpoint)
	}
	for _, m := range mdw["GetModelByName"] {
		eps.GetModelEndpoint = m(eps.GetModelByNameEndpoint)
	}
	for _, m := range mdw["UpdateModelState"] {
		eps.UpdateModelStateEndpoint = m(eps.UpdateModelStateEndpoint)
	}
	for _, m := range mdw["UpdateModelPath"] {
		eps.UpdateModelPathEndpoint = m(eps.UpdateModelPathEndpoint)
	}
	for _, m := range mdw["ListModels"] {
		eps.ListModelsEndpoint = m(eps.ListModelsEndpoint)
	}
	return eps
}
