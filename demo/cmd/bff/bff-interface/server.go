package main

import (
	context "context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/walrusyu/gocamp007/demo/api/bff/v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Server struct {
	pb.UnimplementedBffServiceServer
}

func (*Server) GetOrder(context.Context, *emptypb.Empty) (*pb.Order, error) {
	return &pb.Order{
		Id: &wrapperspb.Int32Value{Value: 1},
		OrderItems: []*pb.Order_OrderItem{
			&pb.Order_OrderItem{
				Id:       &wrapperspb.Int32Value{Value: 1},
				Offer:    "test",
				Quantity: 10,
			}}}, nil
}

func (*Server) GetUser(context.Context, *emptypb.Empty) (*pb.User, error) {
	return &pb.User{
		Id:   &wrapperspb.Int32Value{Value: 1},
		Name: "ywf",
		Age:  1,
		Address: &pb.User_Address{
			Province: "sh1",
			City:     "cn1",
			Street:   "xx1",
		}}, nil
}

func (*Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	user := &pb.User{
		Id:   &wrappers.Int32Value{Value: 1},
		Name: "ywf2",
		Age:  18,
		Address: &pb.User_Address{
			Province: "sh2",
			City:     "hk2",
			Street:   "lc2",
		},
	}
	validPaths := req.Mask.GetPaths()
	fmt.Printf("paths:%v", validPaths)
	if isFieldUsed("name", validPaths) {
		user.Name = req.User.Name
	}
	if isFieldUsed("age", validPaths) {
		user.Age = req.User.Age
	}
	return user, nil
}

func isFieldUsed(field string, paths []string) bool {
	for i := range paths {
		if paths[i] == field {
			return true
		}
	}
	return false
}
