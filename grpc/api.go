package grpc

import (
	pb "github.com/dapr/go-sdk/dapr/proto/runtime/v1"
	"google.golang.org/grpc"
)

type api struct {
	pb.UnimplementedDaprServer

	client pb.AppCallbackClient
}

func NewAPI(conn grpc.ClientConnInterface) pb.DaprServer {
	return &api{
		client: pb.NewAppCallbackClient(conn),
	}
}
