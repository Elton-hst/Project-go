package order

import (
	"errors"
)

type OrderItem struct {
	OrderId     string
	ProductId   string
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func NewOrderItem(ordeId string, productId string, name string, description string, price float64, quantity int) (*OrderItem, error) {
	orderItem := &OrderItem{
		OrderId: ordeId,
		ProductId: productId,
		Name: name,
		Description: description,
		Price: price,
		Quantity: quantity,
	}

	err := orderItem.Validate()
	if err != nil {
		return nil, err
	}

	return orderItem, nil
}

func (o *OrderItem) Validate() error {
	if o.OrderId == "" || o.ProductId == "" || o.Price == 0 {
		return errors.New("OrderItem cannot be empty")
	}
	
	return nil
}
