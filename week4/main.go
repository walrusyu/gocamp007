package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/walrusyu/gocamp007/week4/api/user"
	"github.com/walrusyu/gocamp007/week4/configs"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() error {
		config := configs.GetConfig()
		flag.Parse()
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", config.Server.Grpc.Port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		server := InitUserServer()
		pb.RegisterUserServer(s, server)
		log.Printf("server listening at %v", lis.Addr())

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGHUP)
		go func() {
			select {
			case <-sigChan:
				fmt.Printf("received shutdown signal\n")
				s.Stop()
			}
		}()
		return s.Serve(lis)
	})
	err := g.Wait()
	fmt.Println(err)
}
