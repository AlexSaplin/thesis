package ardea

import (
	"github.com/mattn/go-nulltype"
	ardea "hippo/pkg/clients/ardea/pb"
	"hippo/pkg/entities"
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
