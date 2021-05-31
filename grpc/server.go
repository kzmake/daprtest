package grpc

import (
	"google.golang.org/grpc"
)

func NewServer(opts ...grpc.ServerOption) *grpc.Server {
	return grpc.NewServer(opts...)
}
