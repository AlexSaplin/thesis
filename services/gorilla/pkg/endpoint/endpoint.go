package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	gouuid "github.com/satori/go.uuid"
	"gorilla/pkg/entities"
	service "gorilla/pkg/service"
	"time"
)

// AddDeltasRequest collects the request parameters for the AddDeltas method.
type AddDeltasRequest struct {
	Deltas []entities.Delta `json:"deltas"`
}

// AddDeltasResponse collects the response parameters for the AddDeltas method.
type AddDeltasResponse struct {
	Err error `json:"err"`
}

// MakeAddDeltasEndpoint returns an endpoint that invokes AddDeltas on the service.
func MakeAddDeltasEndpoint(s service.GorillaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddDeltasRequest)
		err := s.AddDeltas(ctx, req.Deltas)
		return AddDeltasResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r AddDeltasResponse) Failed() error {
	return r.Err
}

// GetDeltasRequest collects the request parameters for the GetDeltas method.
type GetDeltasRequest struct {
	OwnerID       gouuid.UUID `json:"owner_id"`
	ObjectID      string      `json:"object_id"`
	ObjectType    string      `json:"object_type"`
	FirstDate     time.Time   `json:"first_date"`
	LastDate      time.Time   `json:"last_date"`
	UseCategories bool        `json:"use_categories"`
}

// GetDeltasResponse collects the response parameters for the GetDeltas method.
type GetDeltasResponse struct {
	Deltas []entities.Delta `json:"deltas"`
	Err    error            `json:"err"`
}

// MakeGetDeltasEndpoint returns an endpoint that invokes GetDeltas on the service.
func MakeGetDeltasEndpoint(s service.GorillaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDeltasRequest)
		deltas, err := s.GetDeltas(ctx, req.OwnerID, req.ObjectID, req.ObjectType, req.FirstDate, req.LastDate, req.UseCategories)
		return GetDeltasResponse{
			Deltas: deltas,
			Err:    err,
		}, nil
	}
}

// Failed implements Failer.
func (r GetDeltasResponse) Failed() error {
	return r.Err
}

// GetBalanceRequest collects the request parameters for the GetBalance method.
type GetBalanceRequest struct {
	OwnerID gouuid.UUID `json:"owner_id"`
}

// GetBalanceResponse collects the response parameters for the GetBalance method.
type GetBalanceResponse struct {
	Balance float64 `json:"balance"`
	Err     error   `json:"err"`
}

// MakeGetBalanceEndpoint returns an endpoint that invokes GetBalance on the service.
func MakeGetBalanceEndpoint(s service.GorillaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetBalanceRequest)
		balance, err := s.GetBalance(ctx, req.OwnerID)
		return GetBalanceResponse{
			Balance: balance,
			Err:     err,
		}, nil
	}
}

// Failed implements Failer.
func (r GetBalanceResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// AddDeltas implements Service. Primarily useful in a client.
func (e Endpoints) AddDeltas(ctx context.Context, deltas []entities.Delta) (err error) {
	request := AddDeltasRequest{Deltas: deltas}
	response, err := e.AddDeltasEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddDeltasResponse).Err
}

// GetDeltas implements Service. Primarily useful in a client.
func (e Endpoints) GetDeltas(ctx context.Context, ownerID gouuid.UUID, objectID, objectType string, firstDate time.Time, lastDate time.Time, useCategories bool) (deltas []entities.Delta, err error) {
	request := GetDeltasRequest{
		FirstDate:     firstDate,
		LastDate:      lastDate,
		ObjectID:      objectID,
		ObjectType:    objectType,
		OwnerID:       ownerID,
		UseCategories: useCategories,
	}
	response, err := e.GetDeltasEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetDeltasResponse).Deltas, response.(GetDeltasResponse).Err
}

// GetBalance implements Service. Primarily useful in a client.
func (e Endpoints) GetBalance(ctx context.Context, ownerID gouuid.UUID) (balance float64, err error) {
	request := GetBalanceRequest{OwnerID: ownerID}
	response, err := e.GetBalanceEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetBalanceResponse).Balance, response.(GetBalanceResponse).Err
}
