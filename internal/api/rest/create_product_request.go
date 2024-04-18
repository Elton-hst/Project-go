package rest

import (
	"time"

	"github.com/Elton-hst/internal/application/logger"
	"github.com/Elton-hst/internal/domain/aggregate"
	"github.com/Elton-hst/internal/domain/factory"
)

type CreateProductRequest struct {
	PK    string  `json:"PK"`
	Name  string  `json:"Name"`
	Value float64 `json:"Value"`
}

func (req *CreateProductRequest) ToAggregate() *aggregate.Product {
	product, err := factory.CreatedProduct(req.Name, req.Value, time.Now())
	if err != nil {
		logger.Error.Printf(err.Error())
	}
	return product
}
