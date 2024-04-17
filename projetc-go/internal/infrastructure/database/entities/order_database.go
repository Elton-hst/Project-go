package entities

import (
	"github.com/Elton-hst/internal/domain/aggregate"
	"github.com/Elton-hst/internal/domain/generics/contract"
)

type OrderDatabase struct {
	CustomerId string
	OrderItem  []OrderItemDatabase `gorm:"foreignKey:OrderId;"`
	BaseEntityDatabase
}

type OrderItemDatabase struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
	OrderId     string
	ProductId   string
	BaseEntityDatabase
}

func (order *OrderDatabase) ToOderAggregate() *aggregate.Order {
	var newOrderItem []aggregate.OrderItem

	for _, item := range order.OrderItem {
		orderItem := &aggregate.OrderItem{
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
			Quantity:    item.Quantity,
			OrderId:     item.OrderId,
			ProductId:   item.ProductId,
		}

		newOrderItem = append(newOrderItem, *orderItem)
	}

	return &aggregate.Order{
		CustumerId: order.CustomerId,
		Item:       newOrderItem,
		BaseEntity: contract.BaseEntity{
			PK:        order.BaseEntityDatabase.PK,
			IsActive:  order.BaseEntityDatabase.IsActive,
			CreatedAt: order.BaseEntityDatabase.CreatedAt,
			UpdatedAt: order.BaseEntityDatabase.UpdatedAt,
		},
	}
}

func (item *OrderItemDatabase) ToOderItemAggregate() *aggregate.OrderItem {
	return &aggregate.OrderItem{
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		Quantity:    item.Quantity,
		OrderId:     item.OrderId,
		ProductId:   item.ProductId,
		BaseEntity: contract.BaseEntity{
			PK:        item.BaseEntityDatabase.PK,
			IsActive:  item.BaseEntityDatabase.IsActive,
			CreatedAt: item.BaseEntityDatabase.CreatedAt,
			UpdatedAt: item.BaseEntityDatabase.UpdatedAt,
		},
	}
}
