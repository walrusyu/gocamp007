package server

import (
	context "context"
	pb "github.com/walrusyu/gocamp007/demo/api/order/v1"
	"github.com/walrusyu/gocamp007/demo/cmd/order/internal/service"
	cErros "github.com/walrusyu/gocamp007/demo/errors"
)

var _ pb.OrderServiceServer = &server{}

type server struct {
	pb.UnimplementedOrderServiceServer
	service service.Service
}

func NewServer(dsn string) (pb.OrderServiceServer, error) {
	service, err := service.NewService(dsn)
	if err != nil {
		return nil, err
	}
	return &server{
		service: service,
	}, nil
}

func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.Order, error) {
	c := make(chan *pb.Order, 1)
	defer close(c)

	go func() {
		order := s.service.GetOrder(req.GetId().Value)
		c <- order
	}()

	select {
	case <-ctx.Done():
		return nil, cErros.New(500, "timeout", "can not get order")
	default:
		return <-c, nil
	}
}
