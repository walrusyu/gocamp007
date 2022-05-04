package server

import (
	context "context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/walrusyu/gocamp007/demo/api/user/v1"
	cErros "github.com/walrusyu/gocamp007/demo/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var _ pb.UserServiceServer = &Server{}

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (*Server) Get(ctx context.Context, req *emptypb.Empty) (*pb.User, error) {
	c := make(chan *pb.User, 1)
	defer close(c)
	go func() {
		user := &pb.User{
			Id:   &wrapperspb.Int32Value{Value: 1},
			Name: "ywf",
			Age:  1,
			Address: &pb.User_Address{
				Province: "sh1",
				City:     "cn1",
				Street:   "xx1",
			}}
		c <- user
	}()
	select {
	case <-ctx.Done():
		return nil, cErros.New(500, "timeout", "can not get user")
	default:
		return <-c, nil
	}
}

func (*Server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.User, error) {
	c := make(chan *pb.User, 1)
	defer close(c)
	go func() {
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
		c <- user
	}()
	select {
	case <-ctx.Done():
		return nil, cErros.New(500, "timeout", "can not update user")
	default:
		return <-c, nil
	}
}

func isFieldUsed(field string, paths []string) bool {
	for i := range paths {
		if paths[i] == field {
			return true
		}
	}
	return false
}
