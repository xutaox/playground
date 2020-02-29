package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"playground/grpc/pb"
	"syscall"
	"time"
)

var (
	port = flag.Int("port", 50001, "listening port")
)

func main() {

	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		panic(err)
	}
	serviceKey, err := registerToEtcd(time.Second, 10, *port)
	if err != nil {
		panic(err)
	}

	go handleSysSignal(func() error {
		return UnRegisterToEtcd(serviceKey)
	})

	log.Printf("starting hello service at %d", *port)

	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &greeterServer{})
	panic(s.Serve(lis))
}

func handleSysSignal(f func() error) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		err := f()
		if err != nil {
			panic(err)
		}
		log.Printf("receive signal '%v'", s)
		os.Exit(1)
	}()
}
