package biz

import "github.com/walrusyu/gocamp007/demo/cmd/order/internal/data"

type Biz interface {
	GetOrder(int32) *Order
}

type Order struct {
	Id          int32
	Description string
	OrderItems  []*OrderItem
}
type OrderItem struct {
	Id       int32
	Offer    string
	Quantity int32
}

var _ Biz = &orderBiz{}

type orderBiz struct {
	repo data.Repo
}

func NewBiz(dsn string) (Biz, error) {
	repo, err := data.NewRepo(dsn)
	if err != nil {
		return nil, err
	}
	return &orderBiz{
		repo: repo,
	}, nil
}

func (b *orderBiz) GetOrder(id int32) *Order {
	order := b.repo.GetOrder(id)
	items := b.repo.GetOrderItems(id)
	ret := &Order{
		Id:          order.Id,
		Description: order.Description,
		OrderItems:  []*OrderItem{},
	}
	for _, item := range items {
		ret.OrderItems = append(ret.OrderItems, &OrderItem{
			Id:       item.Id,
			Offer:    item.Offer,
			Quantity: item.Quantity,
		})
	}
	return ret
}
