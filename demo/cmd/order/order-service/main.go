package main

import (
	"flag"
	"fmt"
	pb "github.com/walrusyu/gocamp007/demo/api/order/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 6001, "The server port")
)

type orderServer struct {
	Server
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, &orderServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
