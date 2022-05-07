package server

import (
	"context"
	"fmt"
	pb "github.com/walrusyu/gocamp007/demo/api/user/v1"
	"github.com/walrusyu/gocamp007/demo/cmd/user/internal/service"
	clog "github.com/walrusyu/gocamp007/demo/internal/log"
	"github.com/walrusyu/gocamp007/demo/internal/middleware"
	"github.com/walrusyu/gocamp007/demo/internal/transport"
	tg "github.com/walrusyu/gocamp007/demo/internal/transport/grpc"
	"google.golang.org/grpc"
	gmetadata "google.golang.org/grpc/metadata"
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
	endpoint    string
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

	interceptors := []grpc.UnaryServerInterceptor{
		svr.unaryServerInterceptor(),
	}

	grpcOpts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(interceptors...),
	}

	svr.Server = grpc.NewServer(grpcOpts...)
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

func (s *server) unaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, _ := gmetadata.FromIncomingContext(ctx)
		replyHeader := gmetadata.MD{}
		ctx = transport.NewServerContext(ctx, tg.NewTransport(s.endpoint, info.FullMethod, tg.HeaderCarrier(md), tg.HeaderCarrier(replyHeader)))
		//if s.timeout > 0 {
		//	ctx, cancel = context.WithTimeout(ctx, s.timeout)
		//	defer cancel()
		//}
		h := func(ctx context.Context, req interface{}) (interface{}, error) {
			return handler(ctx, req)
		}
		if len(s.middlewares) > 0 {
			h = middleware.Chain(s.middlewares...)(h)
		}
		reply, err := h(ctx, req)
		if len(replyHeader) > 0 {
			_ = grpc.SetHeader(ctx, replyHeader)
		}
		return reply, err
	}
}
