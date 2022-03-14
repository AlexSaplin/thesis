package entities

import (
	"github.com/mattn/go-nulltype"
	"hippo/pkg/errors"
)

type ValueType uint8

const (
	ValueTypeUnknown = iota
	ValueTypeFloat16
	ValueTypeFloat32
	ValueTypeFloat64
	ValueTypeUInt8
	ValueTypeUInt16
	ValueTypeUInt32
	ValueTypeUint64
	ValueTypeInt8
	ValueTypeInt16
	ValueTypeInt32
	ValueTypeInt64
	ValueTypeComplex64
	ValueTypeComplex128
)

const (
	typeSize8   = 1
	typeSize16  = 2
	typeSize32  = 4
	typeSize64  = 8
	typeSize128 = 16
)

func (t ValueType) Size() (result int64, err error) {
	switch t {
	case ValueTypeUInt8, ValueTypeInt8:
		result = typeSize8
	case ValueTypeFloat16, ValueTypeUInt16, ValueTypeInt16:
		result = typeSize16
	case ValueTypeFloat32, ValueTypeUInt32, ValueTypeInt32:
		result = typeSize32
	case ValueTypeFloat64, ValueTypeUint64, ValueTypeInt64, ValueTypeComplex64:
		result = typeSize64
	case ValueTypeComplex128:
		result = typeSize128
	default:
		err = ErrUnknownValueType
	}
	return
}

func (t ValueType) String() string {
	switch t {
	case ValueTypeUInt8:
		return "UINT8"
	case ValueTypeInt8:
		return "INT8"
	case ValueTypeFloat16:
		return "FLOAT16"
	case ValueTypeUInt16:
		return "UINT16"
	case ValueTypeInt16:
		return "INT16"
	case ValueTypeFloat32:
		return "FLOAT32"
	case ValueTypeUInt32:
		return "UINT32"
	case ValueTypeInt32:
		return "INT32"
	case ValueTypeFloat64:
		return "FLOAT64"
	case ValueTypeUint64:
		return "UINT64"
	case ValueTypeInt64:
		return "INT64"
	case ValueTypeComplex64:
		return "COMPLEX64"
	case ValueTypeComplex128:
		return "COMPLEX128"
	default:
		return "UNKNOWN"
	}
}

type IOShape []Shape
type Shape []nulltype.NullInt64

type TensorShape []int64

func (s TensorShape) Equals(o Shape) (result bool) {
	if len(s) != len(o) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if o[i].Valid() && s[i] != o[i].Int64Value() {
			return false
		}
	}
	return true
}

type Tensor struct {
	Type  ValueType
	Shape TensorShape
	Data  []byte
}

func (t *Tensor) Valid() (err error) {
	var prod int64 = 1
	for _, dim := range t.Shape {
		prod *= dim
	}
	vs, err := t.Type.Size()
	if err != nil {
		return
	}
	prod *= vs
	if int64(len(t.Data)) != prod {
		err = errors.ErrInvalidTensorShape
	}
	return
}

type TensorList []Tensor

func (t TensorList) Valid() (err error) {
	for _, t := range t {
		err = t.Valid()
		if err != nil {
			return
		}
	}
	return
}

func (t TensorList) ConformsToShape(ioShape IOShape) bool {
	if len(t) != len(ioShape) {
		return false
	}
	for i := range t {
		if !t[i].Shape.Equals(ioShape[i]) {
			return false
		}
	}
	return true
}
