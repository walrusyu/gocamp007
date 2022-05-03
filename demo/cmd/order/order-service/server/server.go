package server

import (
	context "context"
	pb "github.com/walrusyu/gocamp007/demo/api/order/v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var _ pb.OrderServiceServer = &Server{}

type Server struct {
	pb.UnimplementedOrderServiceServer
}

func (*Server) Get(context.Context, *emptypb.Empty) (*pb.Order, error) {
	return &pb.Order{
		Id: &wrapperspb.Int32Value{Value: 11},
		OrderItems: []*pb.Order_OrderItem{
			&pb.Order_OrderItem{
				Id:       &wrapperspb.Int32Value{Value: 111},
				Offer:    "test",
				Quantity: 10,
			}}}, nil
}
