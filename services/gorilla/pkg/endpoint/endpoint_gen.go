// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "gorilla/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	AddDeltasEndpoint  endpoint.Endpoint
	GetDeltasEndpoint  endpoint.Endpoint
	GetBalanceEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.GorillaService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		AddDeltasEndpoint:  MakeAddDeltasEndpoint(s),
		GetBalanceEndpoint: MakeGetBalanceEndpoint(s),
		GetDeltasEndpoint:  MakeGetDeltasEndpoint(s),
	}
	for _, m := range mdw["AddDeltas"] {
		eps.AddDeltasEndpoint = m(eps.AddDeltasEndpoint)
	}
	for _, m := range mdw["GetDeltas"] {
		eps.GetDeltasEndpoint = m(eps.GetDeltasEndpoint)
	}
	for _, m := range mdw["GetBalance"] {
		eps.GetBalanceEndpoint = m(eps.GetBalanceEndpoint)
	}
	return eps
}
