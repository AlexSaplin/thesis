package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrModelNotFound = status.Error(codes.NotFound, "model not found")
	ErrModelExists   = status.Error(codes.AlreadyExists, "model with given name exists")
	ErrInternal      = status.Error(codes.Internal, "internal error")
	ErrModelDeleted  = status.Error(codes.NotFound, "model with given name deleted")
)
