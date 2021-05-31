package grpc

import (
	"context"
	"fmt"

	pb "github.com/dapr/go-sdk/dapr/proto/runtime/v1"
)

func (a *api) InvokeBinding(ctx context.Context, req *pb.InvokeBindingRequest) (*pb.InvokeBindingResponse, error) {
	in := &pb.BindingEventRequest{
		Name:     req.GetName(),
		Data:     req.GetData(),
		Metadata: req.GetMetadata(),
	}

	go func() {
		_, err := a.client.OnBindingEvent(context.Background(), in)
		if err != nil {
			fmt.Printf("err: %+v\n", err)
		}
	}()

	return &pb.InvokeBindingResponse{}, nil
}
