package server

import (
	context "context"
	pb "github.com/walrusyu/gocamp007/demo/api/order/v1"
	cErros "github.com/walrusyu/gocamp007/demo/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var _ pb.OrderServiceServer = &Server{}

type Server struct {
	pb.UnimplementedOrderServiceServer
}

func (*Server) Get(ctx context.Context, req *emptypb.Empty) (*pb.Order, error) {
	c := make(chan *pb.Order, 1)
	defer close(c)

	go func() {
		order := &pb.Order{
			Id: &wrapperspb.Int32Value{Value: 11},
			OrderItems: []*pb.Order_OrderItem{
				&pb.Order_OrderItem{
					Id:       &wrapperspb.Int32Value{Value: 111},
					Offer:    "test",
					Quantity: 10,
				}}}
		c <- order
	}()

	select {
	case <-ctx.Done():
		return nil, cErros.New(500, "timeout", "can not get order")
	default:
		return <-c, nil
	}
}
