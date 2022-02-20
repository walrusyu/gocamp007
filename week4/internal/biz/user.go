package biz

import (
	"github.com/walrusyu/gocamp007/week4/internal/data"
	"github.com/walrusyu/gocamp007/week4/internal/po"
)

type UserBiz struct {
	repo *data.UserRepo
}

func NewUserBiz(repo *data.UserRepo) *UserBiz {
	return &UserBiz{
		repo: repo,
	}
}

func (biz *UserBiz) SaveAddressBook(people []*Person) error {
	var newUsers []*po.User
	for _, p := range people {
		newUsers = append(newUsers, &po.User{Id: p.Id, Name: p.Name})
	}
	err := biz.repo.SaveUsers(newUsers)
	return err
}

type Person struct {
	Id           int32
	Name         string
	Email        string
	PhoneNumbers []*PhoneNumber
}

type PhoneNumber struct {
	Number string
	Type   int32
}
