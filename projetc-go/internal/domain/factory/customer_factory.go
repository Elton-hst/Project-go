package factory

import "github.com/Elton-hst/internal/domain/aggregate"

func CreatedCustomer(name string, password string, email string, phone string) (*aggregate.Customer, error) {
	customer, err := aggregate.NewCustomer(name, password, email, phone)
	if err != nil {
		return nil, err
	}

	return customer, nil
}