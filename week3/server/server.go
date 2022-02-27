package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type HttpServer interface {
	Reg(string, http.Handler)
	Start() error
	Stop() error
}

type MyServer struct {
	name    string
	sigChan chan os.Signal
	server  http.Server
	ctx     context.Context
}

func (s *MyServer) Reg(addr string, handler http.Handler) {
	s.server = http.Server{
		Addr:    addr,
		Handler: handler,
	}
}

func (s *MyServer) Start() error {
	fmt.Printf("server:%s is starting\n", s.name)
	if s.sigChan == nil {
		s.sigChan = make(chan os.Signal, 1)
	}
	signal.Notify(s.sigChan, syscall.SIGINT, syscall.SIGHUP)
	go func() {
		select {
		case <-s.sigChan:
			fmt.Printf("server:%s received signal from sigChan\n", s.name)
		case <-s.ctx.Done():
			fmt.Printf("server:%s is canceled\n", s.name)
		}
		s.Stop()
	}()

	fmt.Printf("server:%s started\n", s.name)
	return s.server.ListenAndServe()
}

func (s *MyServer) Stop() error {
	fmt.Printf("server:%s is stopping\n", s.name)
	err := s.server.Shutdown(s.ctx)
	fmt.Printf("server:%s stopped\n", s.name)
	return err
}

func CreateServer(p context.Context, name string) HttpServer {
	ctx, _ := context.WithCancel(p)
	return &MyServer{ctx: ctx, name: name}
}
