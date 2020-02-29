package main

import (
	"context"
	"fmt"
	helloworld "playground/grpc/pb"
	"time"
)

// greeterServer is used to implement helloworld.GreeterServer.
type greeterServer struct{}

// SayHello implements helloworld.GreeterServer
func (s *greeterServer) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	fmt.Printf("%v: Receive is %s\n", time.Now(), in.Name)
	return &helloworld.HelloReply{Message: "Hello " + in.Name}, nil
}