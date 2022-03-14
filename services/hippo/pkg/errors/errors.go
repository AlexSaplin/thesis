package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUnknownValueType = status.Error(
		codes.Internal,
		"unknown tensor value type",
	)

	ErrNilTensor = status.Error(
		codes.Internal,
		"tried parse nil tensor from pb",
	)

	ErrInvalidTensorShape = status.Error(
		codes.InvalidArgument,
		"tensor size does not match required shape and value type",
	)

	ErrModelNotReady = status.Error(
		codes.InvalidArgument,
		"trying to run a model that is not ready",
	)

	ErrShapeMismatch = status.Error(
		codes.InvalidArgument,
		"input shape does not match model input shape",
	)

	ErrModelDeleted = status.Error(
		codes.InvalidArgument,
		"trying to run a deleted model",
	)
)
