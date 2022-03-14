package ardea

import (
	"github.com/mattn/go-nulltype"
	ardea "lynx/pkg/clients/ardea/pb"
	"lynx/pkg/entities"
)

func parseModelStateEntity(in ardea.ModelState) (result entities.ModelState, err error) {
	switch in {
	case ardea.ModelState_STATE_INIT:
		result = entities.ModelStateInit
	case ardea.ModelState_STATE_PROCESSING:
		result = entities.ModelStateProcessing
	case ardea.ModelState_STATE_READY:
		result = entities.ModelStateReady
	case ardea.ModelState_STATE_INVALID:
		result = entities.ModelStateInvalid
	case ardea.ModelState_STATE_DELETED:
		result = entities.ModelStateDeleted
	default:
		result = entities.ModelStateUnknown
	}
	return
}

func makeModelStatePb(in entities.ModelState) (result ardea.ModelState, err error) {
	switch in {
	case entities.ModelStateInit:
		result = ardea.ModelState_STATE_INIT
	case entities.ModelStateProcessing:
		result = ardea.ModelState_STATE_PROCESSING
	case entities.ModelStateReady:
		result = ardea.ModelState_STATE_READY
	case entities.ModelStateInvalid:
		result = ardea.ModelState_STATE_INVALID
	case entities.ModelStateDeleted:
		result = ardea.ModelState_STATE_DELETED
	default:
		result = ardea.ModelState_STATE_UNKNOWN
	}
	return
}

func makeValueTypePb(in entities.ValueType) (out ardea.ValueType) {
	switch in {
	case entities.ValueTypeUInt8:
		return ardea.ValueType_UINT8
	case entities.ValueTypeInt8:
		return ardea.ValueType_INT8
	case entities.ValueTypeFloat16:
		return ardea.ValueType_FLOAT16
	case entities.ValueTypeUInt16:
		return ardea.ValueType_UINT16
	case entities.ValueTypeInt16:
		return ardea.ValueType_INT16
	case entities.ValueTypeFloat32:
		return ardea.ValueType_FLOAT32
	case entities.ValueTypeUInt32:
		return ardea.ValueType_UINT32
	case entities.ValueTypeInt32:
		return ardea.ValueType_INT32
	case entities.ValueTypeFloat64:
		return ardea.ValueType_FLOAT64
	case entities.ValueTypeUint64:
		return ardea.ValueType_UINT64
	case entities.ValueTypeInt64:
		return ardea.ValueType_INT64
	case entities.ValueTypeComplex64:
		return ardea.ValueType_COMPLEX64
	case entities.ValueTypeComplex128:
		return ardea.ValueType_COMPLEX128
	default:
		return ardea.ValueType_UNKNOWN
	}
}

func parseValueTypeEntity(in ardea.ValueType) (out entities.ValueType) {
	switch in {
	case ardea.ValueType_UINT8:
		return entities.ValueTypeUInt8
	case ardea.ValueType_INT8:
		return entities.ValueTypeInt8
	case ardea.ValueType_FLOAT16:
		return entities.ValueTypeFloat16
	case ardea.ValueType_UINT16:
		return entities.ValueTypeUInt16
	case ardea.ValueType_INT16:
		return entities.ValueTypeInt16
	case ardea.ValueType_FLOAT32:
		return entities.ValueTypeFloat32
	case ardea.ValueType_UINT32:
		return entities.ValueTypeUInt32
	case ardea.ValueType_INT32:
		return entities.ValueTypeInt32
	case ardea.ValueType_FLOAT64:
		return entities.ValueTypeFloat64
	case ardea.ValueType_UINT64:
		return entities.ValueTypeUint64
	case ardea.ValueType_INT64:
		return entities.ValueTypeInt64
	case ardea.ValueType_COMPLEX64:
		return entities.ValueTypeComplex64
	case ardea.ValueType_COMPLEX128:
		return entities.ValueTypeComplex128
	default:
		return entities.ValueTypeUnknown
	}
}

func parseIOShapeEntity(shape []*ardea.Shape) (result entities.IOShape) {
	result = make([]entities.Shape, 0, len(shape))
	for _, item := range shape {
		if item != nil {
			currentShape := make(entities.Shape, 0, len(item.Value))
			for _, elem := range item.Value {
				var currentValue nulltype.NullInt64
				if elem.GetIsValid() {
					currentValue.Set(elem.GetValue())
				} else {
					currentValue.Reset()
				}
				currentShape = append(currentShape, currentValue)
			}
			result = append(result, currentShape)
		}
	}
	return
}

func serializeIOShapeEntity(shape entities.IOShape) (result []*ardea.Shape) {
	result = make([]*ardea.Shape, 0, len(shape))
	for _, item := range shape {
		currentShape := make([]*ardea.NullInt64, 0, len(item))
		for _, elem := range item {
			currentShape = append(currentShape, &ardea.NullInt64{
				Value:  elem.Int64Value(),
				IsValid: elem.Valid(),
			})
		}
		result = append(result, &ardea.Shape{
			Value: currentShape,
		})
	}
	return
}
