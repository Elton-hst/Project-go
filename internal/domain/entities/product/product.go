package product

import (
	"errors"
	"time"
)

type Product struct {
	Name       string
	CategoryId string
	Value      float64
	Validity   time.Time
}

func NewProduct(name string, value float64, validity time.Time) (*Product, error) {
	product := &Product{
		Name:     name,
		Value:    value,
		Validity: validity,
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {

	if p.Name == "" || p.Value == 0 || p.Validity.IsZero() {
		return errors.New("Product invalide")
	}

	return nil
}
