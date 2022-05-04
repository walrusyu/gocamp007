package data

import (
	"fmt"
	cErros "github.com/walrusyu/gocamp007/demo/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repo interface {
	GetOrder(id int32) *Order
	GetOrderItems(id int32) []*OrderItem
	UpdateOrder(*Order) (*Order, error)
}

type Order struct {
	Id          int32
	Description string
}

type OrderItem struct {
	Id       int32
	OrderId  int32
	Offer    string
	Quantity int32
}

var _ Repo = &orderRepo{}

type orderRepo struct {
	db *gorm.DB
}

func NewRepo(dsn string) (Repo, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, cErros.New(500, "incorrect config", "the mysql connection string is invalid")
	}
	return &orderRepo{
		db: db,
	}, nil
}

func (r *orderRepo) GetOrder(id int32) *Order {
	var order *Order
	r.db.First(order, "id = ?", fmt.Sprintf("%d", id))
	return order
}

func (r *orderRepo) GetOrderItems(id int32) []*OrderItem {
	var items []*OrderItem
	r.db.Where("orderId = ?", fmt.Sprintf("%d", id)).Find(&items)
	return items
}

func (r *orderRepo) UpdateOrder(updateOrder *Order) (*Order, error) {
	var order *Order
	r.db.First(order, "id = ?", fmt.Sprintf("%d", updateOrder.Id))
	r.db.Model(order).Updates(updateOrder)
	return order, nil
}
