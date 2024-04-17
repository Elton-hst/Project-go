package aggregate

import (
	"time"

	"github.com/Elton-hst/internal/domain/entities/product"
	"github.com/Elton-hst/internal/domain/generics/contract"
)

type Product struct {
	Name       string
	CategoryID string
	Value      float64
	Validity   time.Time
	contract.BaseEntity
}

type Category struct {
	Name string
	contract.BaseEntity
}

func NewProduct(name string, categoryId string, value float64, validity time.Time) (*Product, error) {
	product, err := product.NewProduct(name, value, validity)
	if err != nil {
		return nil, err
	}

	aggregateProduct := &Product{
		Name:       product.Name,
		CategoryID: categoryId,
		Value:      product.Value,
		Validity:   product.Validity,
		BaseEntity: *contract.NewBaseEntity(),
	}

	return aggregateProduct, nil
}

func NewCategory(name string) (*Category, error) {
	category, err := product.NewCategory(name)
	if err != nil {
		return nil, err
	}

	return &Category{
		Name:       category.Name,
		BaseEntity: *contract.NewBaseEntity(),
	}, nil
}
