package repositories

import (
	"github.com/Elton-hst/internal/application/logger"
	"github.com/Elton-hst/internal/domain/aggregate"
	"github.com/Elton-hst/internal/domain/repositories"
	"github.com/Elton-hst/internal/infrastructure/database/entities"
	"github.com/Elton-hst/internal/infrastructure/database/mapper"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) repositories.ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (p *ProductRepositoryImpl) Create(product *aggregate.Product) (*aggregate.Product, error) {
	newProduct := mapper.ToProductDatabase(product)

	if err := p.db.Create(newProduct).Error; err != nil {
		return nil, err
	}

	result, err := p.FindById(newProduct.PK)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *ProductRepositoryImpl) FindById(id string) (*aggregate.Product, error) {
	var product entities.ProductDatabase

	err := p.db.Find(&product, "pk = ?", id).Error
	if err != nil {
		return nil, err
	}

	result := product.ToProductAggregate()
	if result.PK == "" {
		logger.Error.Printf("Error on find product %s", product.PK)
	}

	return result, nil
}

func (p *ProductRepositoryImpl) DeleteById(id string) error {
	panic("unimplemented")
}

func (p *ProductRepositoryImpl) FindAll() ([]*aggregate.Product, error) {
	panic("unimplemented")
}

func (p *ProductRepositoryImpl) Update(product *aggregate.Product) (*aggregate.Product, error) {
	panic("unimplemented")
}
