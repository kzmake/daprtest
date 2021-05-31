package binding

import (
	"context"
	"time"

	dapr "github.com/dapr/go-sdk/client"
	localdapr "github.com/kzmake/daprtest/grpc"
	grpc "google.golang.org/grpc"
)

const size = 1024 * 1024

func Example_InvokeBinding() {
	// setup
	d, conn, lis, err := localdapr.Run()
	if err != nil {
		panic(err)
	}
	defer d.Close()

	alice := NewAlice(conn.(*grpc.ClientConn))
	defer alice.Close()

	bob := NewBob(lis)
	go bob.Start()
	defer bob.GracefulStop()

	// test
	in := &dapr.InvokeBindingRequest{
		Name:      "bob",
		Operation: "create",
		Data:      []byte("hi, bob"),
		Metadata:  map[string]string{"hoge": "fuga"},
	}
	alice.InvokeBinding(context.TODO(), in)

	// wait
	time.Sleep(1 * time.Second)

	// Output:
	// alice: invoking
	// alice: completed
	// bob: recieved
	// bob: completed
}
