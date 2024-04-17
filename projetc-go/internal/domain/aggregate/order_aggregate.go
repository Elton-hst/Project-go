package aggregate

import (
	"github.com/Elton-hst/internal/domain/entities/order"
	"github.com/Elton-hst/internal/domain/generics/contract"
)

type (
	Order struct {
		CustumerId string
		Item       []OrderItem
		contract.BaseEntity
	}

	OrderItem struct {
		OrderId     string
		ProductId   string
		Name        string
		Description string
		Price       float64
		Quantity    int
		contract.BaseEntity
	}
)

func NewOrderItem(orderItem *OrderItem) (*OrderItem, error) {
	item, err := order.NewOrderItem(
		orderItem.OrderId,
		orderItem.ProductId,
		orderItem.Name,
		orderItem.Description,
		orderItem.Price,
		orderItem.Quantity,
	)

	if err != nil {
		return nil, err
	}

	return &OrderItem{
		OrderId:     item.OrderId,
		ProductId:   item.ProductId,
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		Quantity:    item.Quantity,
		BaseEntity:  *contract.NewBaseEntity(),
	}, nil
}

func NewOrder(orderParam *Order) (*Order, error) {
	order, err := order.NewOrder(orderParam.CustumerId, nil)
	if err != nil {
		return nil, err
	}

	return &Order{
		CustumerId: order.CostumerId,
		BaseEntity: *contract.NewBaseEntity(),
	}, nil
}
