package binding

import (
	"context"
	"fmt"

	dapr "github.com/dapr/go-sdk/client"
	grpc "google.golang.org/grpc"
)

type Alice struct {
	client dapr.Client
}

func NewAlice(conn *grpc.ClientConn) *Alice {
	return &Alice{client: dapr.NewClientWithConnection(conn)}
}

func (a *Alice) InvokeBinding(ctx context.Context, in *dapr.InvokeBindingRequest) error {
	fmt.Println("alice: invoking")
	defer fmt.Println("alice: completed")

	return a.client.InvokeOutputBinding(ctx, in)
}

func (a *Alice) Close() { a.client.Close() }
