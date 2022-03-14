package hippo

import (
	hippo "lynx/pkg/clients/hippo/pb"
	"lynx/pkg/entities"
)

func makeValueTypePb(in entities.ValueType) (out hippo.ValueType) {
	switch in {
	case entities.ValueTypeUInt8:
		return hippo.ValueType_UINT8
	case entities.ValueTypeInt8:
		return hippo.ValueType_INT8
	case entities.ValueTypeFloat16:
		return hippo.ValueType_FLOAT16
	case entities.ValueTypeUInt16:
		return hippo.ValueType_UINT16
	case entities.ValueTypeInt16:
		return hippo.ValueType_INT16
	case entities.ValueTypeFloat32:
		return hippo.ValueType_FLOAT32
	case entities.ValueTypeUInt32:
		return hippo.ValueType_UINT32
	case entities.ValueTypeInt32:
		return hippo.ValueType_INT32
	case entities.ValueTypeFloat64:
		return hippo.ValueType_FLOAT64
	case entities.ValueTypeUint64:
		return hippo.ValueType_UINT64
	case entities.ValueTypeInt64:
		return hippo.ValueType_INT64
	case entities.ValueTypeComplex64:
		return hippo.ValueType_COMPLEX64
	case entities.ValueTypeComplex128:
		return hippo.ValueType_COMPLEX128
	default:
		return hippo.ValueType_UNKNOWN
	}
}

func makeValueTypeEntity(in hippo.ValueType) (out entities.ValueType) {
	switch in {
	case hippo.ValueType_UINT8:
		return entities.ValueTypeUInt8
	case hippo.ValueType_INT8:
		return entities.ValueTypeInt8
	case hippo.ValueType_FLOAT16:
		return entities.ValueTypeFloat16
	case hippo.ValueType_UINT16:
		return entities.ValueTypeUInt16
	case hippo.ValueType_INT16:
		return entities.ValueTypeInt16
	case hippo.ValueType_FLOAT32:
		return entities.ValueTypeFloat32
	case hippo.ValueType_UINT32:
		return entities.ValueTypeUInt32
	case hippo.ValueType_INT32:
		return entities.ValueTypeInt32
	case hippo.ValueType_FLOAT64:
		return entities.ValueTypeFloat64
	case hippo.ValueType_UINT64:
		return entities.ValueTypeUint64
	case hippo.ValueType_INT64:
		return entities.ValueTypeInt64
	case hippo.ValueType_COMPLEX64:
		return entities.ValueTypeComplex64
	case hippo.ValueType_COMPLEX128:
		return entities.ValueTypeComplex128
	default:
		return entities.ValueTypeUnknown
	}
}

func serializeTensorListEntity(in entities.TensorList) (result []*hippo.Tensor) {
	result = make([]*hippo.Tensor, 0, len(in))
	for _, item := range in {
		parsed := &hippo.Tensor{
			Type: makeValueTypePb(item.Type),
			Shape: &hippo.Shape{
				Value: item.Shape,
			},
			Data: item.Data,
		}
		result = append(result, parsed)
	}
	return
}
