package factory

import (
	"time"

	"github.com/Elton-hst/internal/domain/aggregate"
)

func CreatedProduct(name string, value float64, validity time.Time) (*aggregate.Product, error) {
	product, err := aggregate.NewProduct(name, value, validity)
	if err != nil {
		return nil, err
	}
	return product, nil
}