package data

import (
	"fmt"
	cErros "github.com/walrusyu/gocamp007/demo/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repo interface {
	GetUser(id int32) *User
	GetAddress(id int32) *Address
	UpdateUser(*User) (*User, error)
	UpdateAddress(*Address) (*Address, error)
}

type User struct {
	Id        int32
	Name      string
	Age       int32
	AddressId int32
}
type Address struct {
	Id       int32
	Province string
	City     string
	Street   string
}

var _ Repo = &userRepo{}

type userRepo struct {
	db *gorm.DB
}

func NewRepo(dsn string) (Repo, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, cErros.New(500, "incorrect config", "the mysql connection string is invalid")
	}
	return &userRepo{
		db: db,
	}, nil
}

func (r *userRepo) GetUser(id int32) *User {
	var user *User
	r.db.First(user, "id = ?", fmt.Sprintf("%d", id))
	return user
}

func (r *userRepo) GetAddress(id int32) *Address {
	var address *Address
	r.db.First(address, "id = ?", fmt.Sprintf("%d", id))
	return address
}

func (r *userRepo) UpdateUser(updateUser *User) (*User, error) {
	var user *User
	r.db.First(user, "id = ?", fmt.Sprintf("%d", updateUser.Id))
	r.db.Model(user).Updates(updateUser)
	return user, nil
}

func (r *userRepo) UpdateAddress(updateAddress *Address) (*Address, error) {
	var address *Address
	r.db.First(address, "id = ?", fmt.Sprintf("%d", address.Id))
	r.db.Model(address).Updates(updateAddress)
	return address, nil
}
