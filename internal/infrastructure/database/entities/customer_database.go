package entities

import (
	"github.com/Elton-hst/internal/domain/aggregate"
	"github.com/Elton-hst/internal/domain/entities/person"
	"github.com/Elton-hst/internal/domain/generics/contract"
)

type CustomerDatabase struct {
	BaseEntityDatabase
	PersonDatabase
}

type PersonDatabase struct {
	Name     string
	Email    string
	Password string
	Phone    string
}

func (c *CustomerDatabase) ToCustomerAggregate() *aggregate.Customer {
	return &aggregate.Customer{
		Person: person.Person{
			Name:     c.PersonDatabase.Name,
			Password: c.PersonDatabase.Password,
			Email:    c.PersonDatabase.Email,
			Phone:    c.PersonDatabase.Phone,
		},
		BaseEntity: contract.BaseEntity{
			PK:        c.BaseEntityDatabase.PK,
			IsActive:  c.BaseEntityDatabase.IsActive,
			CreatedAt: c.BaseEntityDatabase.CreatedAt,
			UpdatedAt: c.BaseEntityDatabase.UpdatedAt,
		},
	}
}
