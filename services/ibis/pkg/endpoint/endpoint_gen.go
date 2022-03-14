package endpoint

import (
	service "ibis/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	CreateFunctionEndpoint      endpoint.Endpoint
	GetFunctionEndpoint         endpoint.Endpoint
	GetFunctionByNameEndpoint   endpoint.Endpoint
	UpdateFunctionEndpoint      endpoint.Endpoint
	ListFunctionsEndpoint       endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.IbisService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		CreateFunctionEndpoint:      MakeCreateFunctionEndpoint(s),
		GetFunctionEndpoint:         MakeGetFunctionEndpoint(s),
		GetFunctionByNameEndpoint:   MakeGetFunctionByNameEndpoint(s),
		UpdateFunctionEndpoint:      MakeUpdateFunctionEndpoint(s),
		ListFunctionsEndpoint:       MakeListFunctionsEndpoint(s),
	}
	for _, m := range mdw["CreateFunction"] {
		eps.CreateFunctionEndpoint = m(eps.CreateFunctionEndpoint)
	}
	for _, m := range mdw["GetFunction"] {
		eps.GetFunctionEndpoint = m(eps.GetFunctionEndpoint)
	}
	for _, m := range mdw["GetFunctionByName"] {
		eps.GetFunctionEndpoint = m(eps.GetFunctionByNameEndpoint)
	}
	for _, m := range mdw["UpdateFunction"] {
		eps.UpdateFunctionEndpoint = m(eps.UpdateFunctionEndpoint)
	}
	for _, m := range mdw["ListFunctions"] {
		eps.ListFunctionsEndpoint = m(eps.ListFunctionsEndpoint)
	}
	return eps
}
