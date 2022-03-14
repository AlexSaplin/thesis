package grpc

import (
	"ardea/pkg/entities"
	pb "ardea/pkg/grpc/pb"
	"github.com/mattn/go-nulltype"
)

func parseModelStateEntity(in pb.ModelState) (result entities.ModelState, err error) {
	switch in {
	case pb.ModelState_STATE_INIT:
		result = entities.ModelStateInit
	case pb.ModelState_STATE_PROCESSING:
		result = entities.ModelStateProcessing
	case pb.ModelState_STATE_READY:
		result = entities.ModelStateReady
	case pb.ModelState_STATE_INVALID:
		result = entities.ModelStateInvalid
	case pb.ModelState_STATE_DELETED:
		result = entities.ModelStateDeleted
	default:
		result = entities.ModelStateUnknown
	}
	return
}

func serializeModelStateEntity(in entities.ModelState) (result pb.ModelState, err error) {
	switch in {
	case entities.ModelStateInit:
		result = pb.ModelState_STATE_INIT
	case entities.ModelStateProcessing:
		result = pb.ModelState_STATE_PROCESSING
	case entities.ModelStateReady:
		result = pb.ModelState_STATE_READY
	case entities.ModelStateInvalid:
		result = pb.ModelState_STATE_INVALID
	case entities.ModelStateDeleted:
		result = pb.ModelState_STATE_DELETED
	default:
		result = pb.ModelState_STATE_UNKNOWN
	}
	return
}

func serializeValueTypeEntity(in entities.ValueType) (out pb.ValueType) {
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

func parseValueTypeEntity(in pb.ValueType) (out entities.ValueType) {
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

func serializeModelEntity(in entities.Model) (result pb.Model, err error) {
	var state pb.ModelState
	state, err = serializeModelStateEntity(in.State)
	if err != nil {
		return
	}
	return pb.Model{
		ID:          in.ID.String(),
		OwnerID:     in.OwnerID.String(),
		Type:        serializeValueTypeEntity(in.ValueType),
		State:       state,
		InputShape:  serializeIOShapeEntity(in.InputShape),
		OutputShape: serializeIOShapeEntity(in.OutputShape),
		Path:        in.Path,
		Name:        in.Name,
		ErrStrSet:   in.Error.Valid(),
		ErrStr:      in.Error.String(),
	}, nil
}

func parseIOShapeEntity(shape []*pb.Shape) (result [][]nulltype.NullInt64) {
	result = make([][]nulltype.NullInt64, 0, len(shape))
	for _, item := range shape {
		if item != nil {
			currentShape := make([]nulltype.NullInt64, 0, len(item.Value))
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

func serializeIOShapeEntity(shape [][]nulltype.NullInt64) (result []*pb.Shape) {
	result = make([]*pb.Shape, 0, len(shape))
	for _, item := range shape {
		currentShape := make([]*pb.NullInt64, 0, len(item))
		for _, elem := range item {
			currentShape = append(currentShape, &pb.NullInt64{
				Value:  elem.Int64Value(),
				IsValid: elem.Valid(),
			})
		}
		result = append(result, &pb.Shape{
			Value: currentShape,
		})
	}
	return
}
