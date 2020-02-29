package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/naming"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	helloworld "playground/grpc/pb"
	"strconv"
	"time"
)

func main() {

	const (
		serviceName = "hello_world"
		ectdAddr    = "http://localhost:2379"
	)
	var (
		grpcConn *grpc.ClientConn
		ctx      = context.Background()
		err      error
	)

	etcdCli, err := clientv3.NewFromURL(ectdAddr)
	if err != nil {
		panic(err)
	}
	grpcConn, err = grpc.DialContext(ctx, serviceName,
		grpc.WithInsecure(),
		grpc.WithBalancer(grpc.RoundRobin(&naming.GRPCResolver{Client: etcdCli})))
	if err != nil {
		panic(err)
	}
	grpcCli := helloworld.NewGreeterClient(grpcConn)

	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		resp, err := grpcCli.SayHello(ctx, &helloworld.HelloRequest{Name: "world " + strconv.Itoa(t.Second())})
		if err != nil {
			fmt.Printf("err: %v\n", err)
			continue
		}
		fmt.Printf("%v: Reply is %s\n", t, resp.Message)
	}
}
