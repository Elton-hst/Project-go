package order

import "errors"

type Order struct {
	CostumerId string
	Item       []*OrderItem
}

func NewOrder(custumerId string, items ...*OrderItem) (*Order, error) {
	order := &Order{
		CostumerId: custumerId,
		Item:       items,
	}

	if err := order.Validate(); err != nil {
		return nil, err
	}

	return order, nil
}

func (o *Order) Validate() error {
	if o.CostumerId == "" {
		return errors.New("Order cannot be empty")
	}

	for _, item := range o.Item {
		if err := item.Validate(); err != nil {
			return err
		}
	}

	return nil
}
