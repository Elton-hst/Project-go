package services

import (
	"github.com/Elton-hst/internal/domain/aggregate"
	"github.com/Elton-hst/internal/domain/generics/contract"
)

type ProductService interface {
	contract.ContractService[aggregate.Product]
}
