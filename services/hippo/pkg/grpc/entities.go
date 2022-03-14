package grpc

import (
	"hippo/pkg/endpoint"
	"hippo/pkg/entities"
	"hippo/pkg/errors"
	pb "hippo/pkg/grpc/pb"
)


func serializeValueType(in entities.ValueType) (out pb.ValueType) {
	switch in {
	case entities.ValueTypeUInt8:
		return pb.ValueType_UINT8
	case entities.ValueTypeInt8:
		return pb.ValueType_INT8
	case entities.ValueTypeFloat16:
		return pb.ValueType_FLOAT16
	case entities.ValueTypeUInt16:
		return pb.ValueType_UINT16
	case entities.ValueTypeInt16:
		return pb.ValueType_INT16
	case entities.ValueTypeFloat32:
		return pb.ValueType_FLOAT32
	case entities.ValueTypeUInt32:
		return pb.ValueType_UINT32
	case entities.ValueTypeInt32:
		return pb.ValueType_INT32
	case entities.ValueTypeFloat64:
		return pb.ValueType_FLOAT64
	case entities.ValueTypeUint64:
		return pb.ValueType_UINT64
	case entities.ValueTypeInt64:
		return pb.ValueType_INT64
	case entities.ValueTypeComplex64:
		return pb.ValueType_COMPLEX64
	case entities.ValueTypeComplex128:
		return pb.ValueType_COMPLEX128
	default:
		return pb.ValueType_UNKNOWN
	}
}

func parseValueType(in pb.ValueType) (out entities.ValueType) {
	switch in {
	case pb.ValueType_UINT8:
		return entities.ValueTypeUInt8
	case pb.ValueType_INT8:
		return entities.ValueTypeInt8
	case pb.ValueType_FLOAT16:
		return entities.ValueTypeFloat16
	case pb.ValueType_UINT16:
		return entities.ValueTypeUInt16
	case pb.ValueType_INT16:
		return entities.ValueTypeInt16
	case pb.ValueType_FLOAT32:
		return entities.ValueTypeFloat32
	case pb.ValueType_UINT32:
		return entities.ValueTypeUInt32
	case pb.ValueType_INT32:
		return entities.ValueTypeInt32
	case pb.ValueType_FLOAT64:
		return entities.ValueTypeFloat64
	case pb.ValueType_UINT64:
		return entities.ValueTypeUint64
	case pb.ValueType_INT64:
		return entities.ValueTypeInt64
	case pb.ValueType_COMPLEX64:
		return entities.ValueTypeComplex64
	case pb.ValueType_COMPLEX128:
		return entities.ValueTypeComplex128
	default:
		return entities.ValueTypeUnknown
	}
}

func serializeTensor(tensor endpoint.Tensor) *pb.Tensor {
	return &pb.Tensor{
		Type:  serializeValueType(tensor.Type),
		Shape: &pb.Shape{
			Value: tensor.Shape,
		},
		Data:  tensor.Data,
	}
}

func parseTensor(tensor *pb.Tensor) endpoint.Tensor {
	return endpoint.Tensor{
		Type:  parseValueType(tensor.Type),
		Shape: tensor.Shape.Value,
		Data:  tensor.Data,
	}
}

func serializeTensorList(tensor []endpoint.Tensor) (result []*pb.Tensor) {
	result = make([]*pb.Tensor, 0, len(tensor))
	for _, item := range tensor {
		result = append(result, serializeTensor(item))
	}
	return
}

func parseTensorList(in []*pb.Tensor) (result []endpoint.Tensor, err error) {
	result = make([]endpoint.Tensor, 0, len(in))
	for _, item := range in {
		if item == nil {
			err = errors.ErrNilTensor
		}
		result = append(result, parseTensor(item))
	}
	return
}
