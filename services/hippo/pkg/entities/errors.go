package entities

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUnknownValueType = status.Error(
		codes.Internal,
		"unknown tensor value type",
	)
)
