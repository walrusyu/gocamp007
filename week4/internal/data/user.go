package data

import "github.com/walrusyu/gocamp007/week4/internal/po"

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (repo *UserRepo) GetUser(id int) (*po.User, error) {
	return &po.User{Id: 1, Name: "ywf"}, nil
}

func (repo *UserRepo) SaveUser(user *po.User) error {
	return nil
}

func (repo *UserRepo) SaveUsers(user []*po.User) error {
	return nil
}
