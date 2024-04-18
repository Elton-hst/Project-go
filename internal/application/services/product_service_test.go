package services

import (
	"fmt"
	"testing"
	"time"

	"github.com/Elton-hst/internal/domain/aggregate"
	"github.com/Elton-hst/internal/domain/factory"
	"github.com/Elton-hst/internal/infrastructure/mocks"
	"go.uber.org/mock/gomock"
)

func TestProductService_CreateUser(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	name := "Pão"
	value := 50.0

	product, err := factory.CreatedProduct(name, value, time.Now())
	if err != nil {
		fmt.Println(err)
	}

	repo := mocks.NewMockProductRepository(control)
	service := NewProductService(repo)

	repo.EXPECT().Create(gomock.Any()).Return(aggregate.NewProduct("Pão", 50.0, time.Now()))

	seila, err := service.Create(product)

	if err != nil {
		t.FailNow()
		return
	}

	if seila.Name != name {
		t.FailNow()
		return
	}
}
