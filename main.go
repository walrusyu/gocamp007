package main

import (
	"context"
	"fmt"
	"github.com/walrusyu/gocamp007/week3"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		server := week3.CreateServer(ctx, "server1")
		server.Reg(":8001", &week3.MyHandler{})
		return server.Start()
	})

	g.Go(func() error {
		server := week3.CreateServer(ctx, "server2")
		server.Reg(":8002", &week3.MyHandler{})

		time.AfterFunc(time.Second*5, func() {
			server.Stop()
		})
		return server.Start()
	})
	err := g.Wait()
	fmt.Println(err)
}
