package entities

type ValueType uint8

const (
	// Warning: these values are mapped to DB
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
