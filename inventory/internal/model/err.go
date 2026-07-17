package model

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrPartNotFound = status.Error(codes.NotFound, "part not found")
