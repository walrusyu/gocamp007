package main

import (
	"context"
	"fmt"
	"github.com/walrusyu/gocamp007/week3"
	"golang.org/x/sync/errgroup"
)

func main() {
	stopChan := make(chan struct{}, 1)
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		server := week3.CreateServer(ctx)
		server.Reg(":8001", &week3.MyHandler{}, stopChan)
		return server.Start()
	})

	g.Go(func() error {
		server := week3.CreateServer(ctx)
		server.Reg(":8002", &week3.MyHandler{}, stopChan)
		return server.Start()
	})
	err := g.Wait()
	fmt.Println(err)
}
