package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	uuid "github.com/satori/go.uuid"

	"hippo/pkg/entities"
	service "hippo/pkg/service"
)



// Tensor represents data tensor
type Tensor struct {
	Type  entities.ValueType `json:"type"`
	Shape []int64            `json:"shape"`
	Data  []byte             `json:"data"`
}


func serializeTensor(tensor entities.Tensor) Tensor {
	return Tensor{
		Type:  tensor.Type,
		Shape: tensor.Shape,
		Data:  tensor.Data,
	}
}

func parseTensor(tensor Tensor) (result entities.Tensor) {
	result = entities.Tensor{
		Type:  tensor.Type,
		Shape: tensor.Shape,
		Data:  tensor.Data,
	}
	return
}

func serializeTensorList(tensor entities.TensorList) (result []Tensor) {
	result = make([]Tensor, 0, len(tensor))
	for _, item := range tensor {
		result = append(result, serializeTensor(item))
	}
	return
}

func parseTensorList(tensor []Tensor) (result entities.TensorList) {
	result = make(entities.TensorList, 0, len(tensor))
	for _, item := range tensor {
		result = append(result, parseTensor(item))
	}
	return
}

// RunRequest collects the request parameters for the Run method.
type RunRequest struct {
	ModelID  uuid.UUID `json:"model_id"`
	Tensors  []Tensor  `json:"tensor"`
}

// RunResponse collects the response parameters for the Run method.
type RunResponse struct {
	Result []Tensor `json:"result"`
	Err    error    `json:"err"`
}

// MakeRunEndpoint returns an endpoint that invokes Run on the service.
func MakeRunEndpoint(s service.HippoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RunRequest)
		pt := parseTensorList(req.Tensors)
		result, err := s.Run(ctx, req.ModelID, pt)
		return RunResponse{
			Err:    err,
			Result: serializeTensorList(result),
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
func (e Endpoints) Run(ctx context.Context, modelID uuid.UUID, tensor []Tensor) (result []Tensor, err error) {
	request := RunRequest{
		ModelID: modelID,
		Tensors:  tensor,
	}
	response, err := e.RunEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RunResponse).Result, response.(RunResponse).Err
}
