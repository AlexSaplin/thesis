package entities

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUnknownValueType = status.Error(
		codes.InvalidArgument,
		"unknown tensor value type",
	)

	ErrInvalidShapeNonPositive = status.Error(
		codes.InvalidArgument,
		"invalid shape: all shape values must be positive integers",
	)

	ErrInvalidShapeLength = status.Error(
		codes.InvalidArgument,
		fmt.Sprintf("invalid shape: shape length must be between %d and %d",
			minTensorDimensions, maxTensorDimensions),
	)

	ErrInvalidBatchDimensionValue = status.Error(
		codes.InvalidArgument,
		"invalid shape: first dimension (batch) must always be 1",
	)

	ErrInvalidTensor = status.Error(
		codes.InvalidArgument,
		"tensor data size does not match it's shape and value type",
	)

	ErrMaxSize = status.Error(
		codes.InvalidArgument,
		fmt.Sprintf("tensor data exceeds maximum size of %d bytes", maxTensorBytes),
	)
)
