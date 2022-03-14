package handlers

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrModelAlreadyUploaded = status.Error(
		codes.AlreadyExists,
		"model already uploaded",
	)

	ErrModelDeleted = status.Error(
		codes.NotFound,
		"model deleted",
	)

	ErrModelStateInit = status.Error(
		codes.FailedPrecondition,
		"model not runnable, upload model first",
	)

	ErrModelStateProcessing = status.Error(
		codes.Unavailable,
		"model is still preprocessing, please wait until it is ready",
	)

	ErrModelStateInvalid = status.Error(
		codes.FailedPrecondition,
		"model is invalid, error occurred during model preprocessing",
	)

	ErrModelStateUnknown = status.Error(
		codes.Internal,
		"model is in unknown state",
	)

	ErrShapeMismatch = status.Error(
		codes.InvalidArgument,
		"input count or shape does not match model specification",
	)

	ErrEmptyInput = status.Error(
		codes.InvalidArgument,
		"zero inputs provided",
	)

	ErrTypeMismatch = status.Error(
		codes.InvalidArgument,
		"input tensor data type does not match model input data type",
	)

	ErrInvalidShapeFormat = status.Error(
		codes.InvalidArgument,
		"shape must be a json list of positive integers",
	)

	ErrInvalidTypeName = status.Error(
		codes.InvalidArgument,
		"unknown data type name",
	)

	ErrUnsupportedValueType = status.Error(
		codes.InvalidArgument,
		"unsupported data type",
	)

	ErrInvalidContainerName = status.Error(
		codes.InvalidArgument,
		"container name must consist of lowercase letters, digits or dashes, it must be within 3 and 25 symbols in length",
	)

	ErrInvalidImageName = status.Error(
		codes.InvalidArgument,
		"invalid docker container name",
	)

	ErrInvalidContainerScale = status.Error(
		codes.InvalidArgument,
		"invalid container scale: value should be between 1 and 10",
	)

	ErrUnsupportedInstanceType = status.Error(
		codes.InvalidArgument,
		"unsupported instance type",
	)

	ErrUnauthenticated = status.Error(
		codes.Unauthenticated,
		"not authenticated, token missing",
	)

	ErrPermissionDenied = status.Error(
		codes.PermissionDenied,
		"permission denied",
	)

	ErrInsufficientBalance = status.Error(
		codes.InvalidArgument,
		"account balance must be positive to make this request",
	)

	ErrMultipartBuilding = status.Error(
		codes.Internal,
		"internal server error",
	)

	ErrFunctionStateUnknown = status.Error(
		codes.Internal,
		"function is in unknown state",
	)
)
