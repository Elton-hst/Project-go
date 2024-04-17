package entities

import (
	"time"

	"github.com/Elton-hst/internal/domain/aggregate"
	"github.com/Elton-hst/internal/domain/generics/contract"
)

type ProductDatabase struct {
	Name       string
	CategoryID string
	Value      float64
	Validity   time.Time
	BaseEntityDatabase
}

func (p *ProductDatabase) ToProductAggregate() *aggregate.Product {
	return &aggregate.Product{
		Name:       p.Name,
		CategoryID: p.CategoryID,
		Value:      p.Value,
		Validity:   p.Validity,
		BaseEntity: contract.BaseEntity{
			PK:        p.BaseEntityDatabase.PK,
			IsActive:  p.BaseEntityDatabase.IsActive,
			CreatedAt: p.BaseEntityDatabase.CreatedAt,
			UpdatedAt: p.BaseEntityDatabase.UpdatedAt,
		},
	}
}
