package services

import (
	"context"

	"github.com/Elton-hst/internal/application/idempotency"
	"github.com/Elton-hst/internal/domain/aggregate"
	"github.com/Elton-hst/internal/domain/repositories"
	"github.com/Elton-hst/internal/domain/services"
)

type ProductServiceImpl struct {
	i    *idempotency.IdempotentHandler[*aggregate.Product]
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) services.ProductService {
	return &ProductServiceImpl{repo: repo}
}

func (p *ProductServiceImpl) Create(product *aggregate.Product) (*aggregate.Product, error) {
	productIdempotenty, has, err := p.i.Startt(context.TODO(), product.PK)
	if err != nil {
		return nil, err
	}

	if has {
		return productIdempotenty, nil
	}

	result, err := p.repo.Create(product)
	if err != nil {
		return nil, err
	}

	if err := p.i.Store(context.TODO(), product.PK, product); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ProductServiceImpl) FindById(id string) (*aggregate.Product, error) {
	result, err := p.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ProductServiceImpl) DeleteById(id string) error {
	panic("unimplemented")
}

func (p *ProductServiceImpl) FindAll() ([]*aggregate.Product, error) {
	panic("unimplemented")
}

func (p *ProductServiceImpl) Update(entity *aggregate.Product) (*aggregate.Product, error) {
	panic("unimplemented")
}
