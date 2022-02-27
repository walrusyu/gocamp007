package main

import (
	"context"
	"fmt"
	myServer "github.com/walrusyu/gocamp007/week3/server"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		server := myServer.CreateServer(ctx, "server1")
		server.Reg(":8001", &myServer.MyHandler{})
		return server.Start()
	})

	g.Go(func() error {
		server := myServer.CreateServer(ctx, "server2")
		server.Reg(":8002", &myServer.MyHandler{})

		time.AfterFunc(time.Second*5, func() {
			server.Stop()
		})
		return server.Start()
	})
	err := g.Wait()
	fmt.Println(err)
}
