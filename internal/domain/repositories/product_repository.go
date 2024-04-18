package repositories

import (
	"github.com/Elton-hst/internal/domain/aggregate"
	"github.com/Elton-hst/internal/domain/generics/contract"
)

type ProductRepository interface {
	contract.ContractRepository[aggregate.Product]
}
