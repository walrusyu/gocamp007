package week3

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type HttpServer interface {
	Reg(string, http.Handler, <-chan struct{})
	Start() error
	Stop() error
}

type MyServer struct {
	sigChan  chan os.Signal
	stopChan <-chan struct{}
	server   http.Server
	ctx      context.Context
}

func (s *MyServer) Reg(addr string, handler http.Handler, stop <-chan struct{}) {
	s.server = http.Server{
		Addr:    addr,
		Handler: handler,
	}
	s.stopChan = stop
}

func (s *MyServer) Start() error {
	if s.sigChan == nil {
		s.sigChan = make(chan os.Signal, 1)
	}
	signal.Notify(s.sigChan, syscall.SIGINT)
	signal.Notify(s.sigChan, syscall.SIGHUP)
	go func() {
		select {
		case <-s.sigChan:
		case <-s.stopChan:
		case <-s.ctx.Done():
		}
		s.Stop()
	}()

	return s.server.ListenAndServe()
}

func (s *MyServer) Stop() error {
	return s.server.Shutdown(s.ctx)
}

func CreateServer(p context.Context) HttpServer {
	ctx, _ := context.WithCancel(p)
	return &MyServer{ctx: ctx}
}
