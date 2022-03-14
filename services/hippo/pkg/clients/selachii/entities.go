package selachii

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	selachii "hippo/pkg/clients/selachii/pb"
	"hippo/pkg/entities"
	"hippo/pkg/errors"
)

func serializeValueTypeEntity(in entities.ValueType) (out selachii.ValueType) {
	switch in {
	case entities.ValueTypeUInt8:
		return selachii.ValueType_UINT8
	case entities.ValueTypeInt8:
		return selachii.ValueType_INT8
	case entities.ValueTypeFloat16:
		return selachii.ValueType_FLOAT16
	case entities.ValueTypeUInt16:
		return selachii.ValueType_UINT16
	case entities.ValueTypeInt16:
		return selachii.ValueType_INT16
	case entities.ValueTypeFloat32:
		return selachii.ValueType_FLOAT32
	case entities.ValueTypeUInt32:
		return selachii.ValueType_UINT32
	case entities.ValueTypeInt32:
		return selachii.ValueType_INT32
	case entities.ValueTypeFloat64:
		return selachii.ValueType_FLOAT64
	case entities.ValueTypeUint64:
		return selachii.ValueType_UINT64
	case entities.ValueTypeInt64:
		return selachii.ValueType_INT64
	case entities.ValueTypeComplex64:
		return selachii.ValueType_COMPLEX64
	case entities.ValueTypeComplex128:
		return selachii.ValueType_COMPLEX128
	default:
		return selachii.ValueType_UNKNOWN
	}
}

func parseValueTypeEntity(in selachii.ValueType) (out entities.ValueType) {
	switch in {
	case selachii.ValueType_UINT8:
		return entities.ValueTypeUInt8
	case selachii.ValueType_INT8:
		return entities.ValueTypeInt8
	case selachii.ValueType_FLOAT16:
		return entities.ValueTypeFloat16
	case selachii.ValueType_UINT16:
		return entities.ValueTypeUInt16
	case selachii.ValueType_INT16:
		return entities.ValueTypeInt16
	case selachii.ValueType_FLOAT32:
		return entities.ValueTypeFloat32
	case selachii.ValueType_UINT32:
		return entities.ValueTypeUInt32
	case selachii.ValueType_INT32:
		return entities.ValueTypeInt32
	case selachii.ValueType_FLOAT64:
		return entities.ValueTypeFloat64
	case selachii.ValueType_UINT64:
		return entities.ValueTypeUint64
	case selachii.ValueType_INT64:
		return entities.ValueTypeInt64
	case selachii.ValueType_COMPLEX64:
		return entities.ValueTypeComplex64
	case selachii.ValueType_COMPLEX128:
		return entities.ValueTypeComplex128
	default:
		return entities.ValueTypeUnknown
	}
}

func serializeTensorEntity(tensor entities.Tensor) *selachii.Tensor {
	return &selachii.Tensor{
		Type: serializeValueTypeEntity(tensor.Type),
		Shape: &selachii.Shape{
			Value: tensor.Shape,
		},
		Data: tensor.Data,
	}
}

func parseTensorEntity(tensor *selachii.Tensor) (result entities.Tensor, err error) {
	if tensor == nil {
		err = status.Error(codes.Internal, "parsing nil tensor")
		return
	}
	result = entities.Tensor{
		Type: parseValueTypeEntity(tensor.Type),
		Data: tensor.Data,
	}
	if tensor.Shape != nil {
		result.Shape = tensor.Shape.Value
	}
	return
}

func serializeTensorListEntity(in entities.TensorList) (result []*selachii.Tensor) {
	result = make([]*selachii.Tensor, 0, len(in))
	for _, item := range in {
		result = append(result, serializeTensorEntity(item))
	}
	return
}

func parseTensorListEntity(in []*selachii.Tensor) (result entities.TensorList, err error) {
	result = make([]entities.Tensor, 0, len(in))
	for _, item := range in {
		if item == nil {
			err = errors.ErrNilTensor
		}
		var parsed entities.Tensor
		parsed, err = parseTensorEntity(item)
		if err != nil {
			return
		}
		result = append(result, parsed)
	}
	return
}

func serializeIOShapeEntity(shape entities.IOShape) (result []*selachii.Shape) {
	result = make([]*selachii.Shape, 0, len(shape))
	for _, item := range shape {
		currentShape := make([]int64, 0, len(item))
		for _, elem := range item {
			var currentValue int64
			if elem.Valid() {
				currentValue = elem.Int64Value()
			} else {
				currentValue = 0
			}
			currentShape = append(currentShape, currentValue)
		}
		result = append(result, &selachii.Shape{
			Value: currentShape,
		})
	}
	return
}
