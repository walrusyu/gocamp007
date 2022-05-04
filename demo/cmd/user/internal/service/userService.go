package service

import (
	pb "github.com/walrusyu/gocamp007/demo/api/user/v1"
	"github.com/walrusyu/gocamp007/demo/cmd/user/internal/biz"
	cErros "github.com/walrusyu/gocamp007/demo/errors"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Service interface {
	GetUser(int32) *pb.User
	UpdateUser(*pb.User) (*pb.User, error)
}

var _ Service = &userService{}

type userService struct {
	biz biz.Biz
}

func NewService(dsn string) (Service, error) {
	biz, err := biz.NewBiz(dsn)
	if err != nil {
		return nil, cErros.New(500, "incorrect config", "the mysql connection string is invalid")
	}
	return &userService{
		biz: biz,
	}, nil
}

func (s *userService) GetUser(id int32) *pb.User {
	user := s.biz.GetUser(id)
	ret := &pb.User{
		Id:   &wrapperspb.Int32Value{Value: user.Id},
		Name: user.Name,
		Age:  user.Age,
		Address: &pb.User_Address{
			Id:       &wrapperspb.Int32Value{Value: user.Address.Id},
			Province: user.Address.Province,
			City:     user.Address.City,
			Street:   user.Address.Street,
		},
	}
	return ret
}

func (s *userService) UpdateUser(updateUser *pb.User) (*pb.User, error) {
	user, err := s.biz.UpdateUser(&biz.User{
		Id:   updateUser.Id.Value,
		Name: updateUser.Name,
		Age:  updateUser.Age,
		Address: &biz.Address{
			Id:       updateUser.Address.Id.Value,
			Province: updateUser.Address.Province,
			City:     updateUser.Address.City,
			Street:   updateUser.Address.Street,
		},
	})
	if err != nil {
		return nil, err
	} else {
		return &pb.User{
			Id:   &wrapperspb.Int32Value{Value: user.Id},
			Name: user.Name,
			Age:  user.Age,
			Address: &pb.User_Address{
				Id:       &wrapperspb.Int32Value{Value: user.Address.Id},
				Province: user.Address.Province,
				City:     user.Address.City,
				Street:   user.Address.Street,
			},
		}, nil
	}
}
