package server

import (
	"context"
	"fmt"
	pb "github.com/walrusyu/gocamp007/demo/api/user/v1"
	cErros "github.com/walrusyu/gocamp007/demo/errors"
)

var _ pb.UserServiceServer = &serviceServer{}

func (s *serviceServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.User, error) {
	c := make(chan *pb.User, 1)
	defer close(c)
	go func() {
		user := s.service.GetUser(req.Id.Value)
		c <- user
	}()
	select {
	case <-ctx.Done():
		return nil, cErros.New(500, "timeout", "can not get user")
	default:
		return <-c, nil
	}
}

func (s *serviceServer) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.User, error) {
	c := make(chan *pb.User, 1)
	defer close(c)
	go func() {
		user := s.service.GetUser(req.User.Id.Value)
		validPaths := req.Mask.GetPaths()
		fmt.Printf("paths:%v", validPaths)
		if isFieldUsed("name", validPaths) {
			user.Name = req.User.Name
		}
		if isFieldUsed("age", validPaths) {
			user.Age = req.User.Age
		}
		user, err := s.service.UpdateUser(user)
		if err != nil {
			return
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
