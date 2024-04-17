package mapper

import (
	"github.com/Elton-hst/internal/domain/aggregate"
	"github.com/Elton-hst/internal/infrastructure/database/entities"
)

func ToProductDatabase(product *aggregate.Product) *entities.ProductDatabase {
	return &entities.ProductDatabase{
		Name:       product.Name,
		CategoryID: product.CategoryID,
		Value:      product.Value,
		Validity:   product.Validity,
		BaseEntityDatabase: entities.BaseEntityDatabase{
			PK:        product.BaseEntity.PK,
			IsActive:  product.BaseEntity.IsActive,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		},
	}
}
