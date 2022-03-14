package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrFunctionNotFound = status.Error(codes.NotFound, "function not found")
	ErrFunctionExists   = status.Error(codes.AlreadyExists, "function with given name exists")
	ErrInternal         = status.Error(codes.Internal, "internal error")
	ErrFunctionDeleted  = status.Error(codes.NotFound, "function with given name deleted")
)
