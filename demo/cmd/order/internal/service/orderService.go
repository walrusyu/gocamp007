package service

import (
	pb "github.com/walrusyu/gocamp007/demo/api/order/v1"
	cErros "github.com/walrusyu/gocamp007/demo/errors"
	"google.golang.org/protobuf/types/known/wrapperspb"
)
import "github.com/walrusyu/gocamp007/demo/cmd/order/internal/biz"

type Service interface {
	GetOrder(int32) *pb.Order
}

var _ Service = &orderService{}

type orderService struct {
	biz biz.Biz
}

func NewService(dsn string) (Service, error) {
	biz, err := biz.NewBiz(dsn)
	if err != nil {
		return nil, cErros.New(500, "incorrect config", "the mysql connection string is invalid")
	}
	return &orderService{
		biz: biz,
	}, nil
}

func (s *orderService) GetOrder(id int32) *pb.Order {
	order := s.biz.GetOrder(id)
	ret := &pb.Order{
		Id:          &wrapperspb.Int32Value{Value: order.Id},
		Description: order.Description,
		OrderItems:  []*pb.Order_OrderItem{},
	}
	for _, item := range order.OrderItems {
		ret.OrderItems = append(ret.OrderItems, &pb.Order_OrderItem{
			Id:       &wrapperspb.Int32Value{Value: item.Id},
			Offer:    item.Offer,
			Quantity: item.Quantity,
		})
	}
	return ret
}
