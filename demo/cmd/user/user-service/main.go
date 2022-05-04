package main

import (
	"flag"
	"fmt"
	pb "github.com/walrusyu/gocamp007/demo/api/user/v1"
	"github.com/walrusyu/gocamp007/demo/cmd/user/user-service/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 6002, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server, err := server.NewServer("dsn")
	pb.RegisterUserServiceServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
