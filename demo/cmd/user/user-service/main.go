package main

import (
	"github.com/walrusyu/gocamp007/demo/cmd/user/user-service/server"
	"github.com/walrusyu/gocamp007/demo/internal/middleware/tracing"
)

func main() {
	svr := server.NewServer(
		server.SetAddress("localhost"),
		server.SetPort(6002),
		server.SetDbConnection(""),
		server.SetMiddlewares(tracing.TracingMiddlewareOnServer()))
	svr.Start()
	defer svr.Stop()
}
