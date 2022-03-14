package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	uuid "github.com/satori/go.uuid"

	service "rhino/pkg/service"
)

// RunRequest collects the request parameters for the Run method.
type RunRequest struct {
	FunctionID uuid.UUID `json:"model_id"`
	Data       []byte    `json:"data"`
}

// RunResponse collects the response parameters for the Run method.
type RunResponse struct {
	Result []byte `json:"result"`
	Err    error  `json:"err"`
}

// MakeRunEndpoint returns an endpoint that invokes Run on the service.
func MakeRunEndpoint(s service.RhinoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RunRequest)
		result, err := s.Run(ctx, req.FunctionID, req.Data)
		return RunResponse{
			Err:    err,
			Result: result,
		}, nil
	}
}

// Failed implements Failer.
func (r RunResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Run implements Service. Primarily useful in a client.
func (e Endpoints) Run(ctx context.Context, modelID uuid.UUID, in []byte) (result []byte, err error) {
	request := RunRequest{
		FunctionID: modelID,
		Data:       in,
	}
	response, err := e.RunEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RunResponse).Result, response.(RunResponse).Err
}
