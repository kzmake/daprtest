package binding

import (
	"context"
	"fmt"
	"net"

	pb "github.com/dapr/go-sdk/dapr/proto/runtime/v1"
	common "github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/grpc"
	"google.golang.org/grpc"
)

type Bob struct {
	daprd.Server
	invokeHandlers  map[string]func(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error)
	bindingHandlers map[string]func(ctx context.Context, in *common.BindingEvent) (out []byte, err error)
	listener        net.Listener
	grpcServer      *grpc.Server
}

func NewBob(lis net.Listener) *Bob {
	srv := daprd.NewServiceWithListener(lis)
	srv.AddBindingInvocationHandler("bob", handle)
	bob := &Bob{
		Server:          *srv.(*daprd.Server),
		invokeHandlers:  make(map[string]func(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error)),
		bindingHandlers: make(map[string]func(ctx context.Context, in *common.BindingEvent) (out []byte, err error)),
		grpcServer:      grpc.NewServer(),
		listener:        lis,
	}

	pb.RegisterAppCallbackServer(bob.grpcServer, bob)

	return bob
}

func (b *Bob) Start() error  { return b.grpcServer.Serve(b.listener) }
func (b *Bob) Stop()         { b.grpcServer.Stop() }
func (b *Bob) GracefulStop() { b.grpcServer.GracefulStop() }

func handle(ctx context.Context, in *common.BindingEvent) ([]byte, error) {
	fmt.Println("bob: recieved")
	defer fmt.Println("bob: completed")

	return nil, nil
}
