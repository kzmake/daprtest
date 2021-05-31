package grpc

import (
	"context"
	"net"

	"google.golang.org/grpc"

	pb "github.com/dapr/go-sdk/dapr/proto/runtime/v1"
	bufconn "google.golang.org/grpc/test/bufconn"
)

type Dapr struct {
	srv *grpc.Server
	api pb.DaprServer
}

func Run() (*Dapr, grpc.ClientConnInterface, net.Listener, error) {
	daprL, conn, err := NewBufListenerAndConnection()
	if err != nil {
		return nil, nil, nil, err
	}

	lis, daprC, err := NewBufListenerAndConnection()
	if err != nil {
		return nil, nil, nil, err
	}

	return RunWithListenerAndConnection(daprL, daprC), conn, lis, nil
}

func RunWithListenerAndConnection(lis net.Listener, conn grpc.ClientConnInterface) *Dapr {
	d := &Dapr{
		srv: NewServer(),
		api: NewAPI(conn),
	}

	pb.RegisterDaprServer(d.srv, d.api)
	go d.srv.Serve(lis)

	return d
}

func (d *Dapr) Close() { d.srv.Stop() }

func NewBufListenerAndConnection() (net.Listener, grpc.ClientConnInterface, error) {
	const size = 1024 * 1024

	lis := bufconn.Listen(size)

	conn, err := grpc.Dial("bufconn",
		grpc.WithContextDialer(func(ctx context.Context, url string) (net.Conn, error) { return lis.Dial() }),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, nil, err
	}

	return lis, conn, nil
}
