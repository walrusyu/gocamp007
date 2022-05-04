package biz

import (
	"github.com/walrusyu/gocamp007/demo/cmd/user/internal/data"
	"golang.org/x/sync/errgroup"
)

type Biz interface {
	GetUser(int32) *User
	UpdateUser(*User) (*User, error)
}

type User struct {
	Id      int32
	Name    string
	Age     int32
	Address *Address
}
type Address struct {
	Id       int32
	Province string
	City     string
	Street   string
}

var _ Biz = &userBiz{}

type userBiz struct {
	repo data.Repo
}

func NewBiz(dsn string) (Biz, error) {
	repo, err := data.NewRepo(dsn)
	if err != nil {
		return nil, err
	}
	return &userBiz{
		repo: repo,
	}, nil
}

func (b *userBiz) GetUser(id int32) *User {
	user := b.repo.GetUser(id)
	address := b.repo.GetAddress(user.AddressId)
	ret := &User{
		Id:   user.Id,
		Name: user.Name,
		Age:  user.Age,
		Address: &Address{
			Id:       address.Id,
			Province: address.Province,
			City:     address.City,
			Street:   address.Street,
		},
	}
	return ret
}

func (b *userBiz) UpdateUser(updateUser *User) (*User, error) {
	user := b.repo.GetUser(updateUser.Id)
	address := b.repo.GetAddress(user.AddressId)
	var g errgroup.Group
	g.Go(func() error {
		_, err := b.repo.UpdateUser(&data.User{
			Id:   user.Id,
			Name: user.Name,
			Age:  user.Age,
		})
		return err
	})
	g.Go(func() error {
		_, err := b.repo.UpdateAddress(&data.Address{
			Id:       address.Id,
			Province: address.Province,
			City:     address.City,
			Street:   address.Street,
		})
		return err
	})
	if err := g.Wait(); err != nil {
		return nil, err
	} else {
		return b.GetUser(updateUser.Id), nil
	}
}
