package aggregate

import (
	"github.com/Elton-hst/internal/domain/entities/person"
	"github.com/Elton-hst/internal/domain/generics/contract"
)

type Customer struct {
	person.Person
	contract.BaseEntity
}

func NewCustomer(name string, password string, email string, phone string) (*Customer, error) {
	person, err := person.NewPerson(name, password, email, phone)
	if err != nil {
		return nil, err
	}

	customer := &Customer{
		Person:     *person,
		BaseEntity: *contract.NewBaseEntity(),
	}

	return customer, nil
}
