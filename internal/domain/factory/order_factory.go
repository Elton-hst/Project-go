package factory

import "github.com/Elton-hst/internal/domain/aggregate"

func CreatedOrder(orderFactory *aggregate.Order) (*aggregate.Order, error) {
	order, err := aggregate.NewOrder(orderFactory)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func CreatedOrderItem(orderItemFactory *aggregate.OrderItem) (*aggregate.OrderItem, error) {
	item, err := aggregate.NewOrderItem(orderItemFactory)
	if err != nil {
		return nil, err
	}

	return item, nil
}  