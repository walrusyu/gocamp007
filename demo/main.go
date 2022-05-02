package main

import (
	"flag"
	"fmt"
	pb2 "github.com/walrusyu/gocamp007/demo/api/order/v1"
	pb "github.com/walrusyu/gocamp007/demo/api/user/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 5001, "The server port")
)

type server struct {
	pb.MyUserServiceServer
}

type server2 struct {
	pb2.MyOrderServiceServer
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	pb2.RegisterOrderServiceServer(s, &server2{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
