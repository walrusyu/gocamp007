package main

import (
	"flag"
	"fmt"
	pb "github.com/walrusyu/gocamp007/demo/api/bff/v1"
	"github.com/walrusyu/gocamp007/demo/cmd/bff/bff-interface/server"
	"github.com/walrusyu/gocamp007/demo/config"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 5001, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// read config from file
	bffConfig := config.BffServerConfig{
		OrderServiceIP:   "localhost",
		OrderServicePort: 6001,
		UserServiceIP:    "localhost",
		UserServicePort:  6002,
	}
	pb.RegisterBffServiceServer(s, server.NewServer(bffConfig))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
