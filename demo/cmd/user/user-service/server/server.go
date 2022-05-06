package server

import (
	"fmt"
	pb "github.com/walrusyu/gocamp007/demo/api/user/v1"
	"github.com/walrusyu/gocamp007/demo/cmd/user/internal/service"
	clog "github.com/walrusyu/gocamp007/demo/internal/log"
	"github.com/walrusyu/gocamp007/demo/internal/middleware"
	"google.golang.org/grpc"
	"log"
	"net"
)

type serviceServer struct {
	pb.UnimplementedUserServiceServer
	service service.Service
}

type server struct {
	*grpc.Server
	address     string
	port        int32
	lis         net.Listener
	middlewares []middleware.Middleware
	dsn         string
	logger      *clog.WrappedLogger
}

type Option func(*server)

func NewServer(opts ...Option) *server {
	svr := &server{
		logger: clog.NewLogger(),
	}
	for _, o := range opts {
		o(svr)
	}
	svr.Server = grpc.NewServer()
	svc, err := service.NewService(svr.dsn)
	if err != nil {
		log.Fatalf("failed to create serviceServer: %v", err)
		return svr
	}
	pb.RegisterUserServiceServer(svr.Server, &serviceServer{service: svc})
	if err != nil {
		log.Fatalf("failed to register serviceServer: %v", err)
		return svr
	}
	return svr
}

func (s *server) listenAndEndpoint() {
	if s.lis == nil {
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.address, s.port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s.lis = lis
	}
}

func (s *server) Start() {
	log.Printf("server listening at %v", s.lis.Addr())
	err := s.Serve(s.lis)
	if err != nil {
		log.Fatalf("failed to start: %v", err)
	}
	s.logger.Info("message", "server started at %v", s.lis.Addr())
}

func (s *server) Stop() {
	s.logger.Info("message", "server stopped")
}

func SetAddress(address string) Option {
	return func(s *server) {
		s.address = address
	}
}

func SetPort(port int32) Option {
	return func(s *server) {
		s.port = port
	}
}

func SetMiddlewares(middlewares ...middleware.Middleware) Option {
	return func(s *server) {
		s.middlewares = middlewares
	}
}

func SetDbConnection(dsn string) Option {
	return func(s *server) {
		s.dsn = dsn
	}
}
