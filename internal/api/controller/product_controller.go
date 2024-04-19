package controller

import (
	"net/http"

	"github.com/Elton-hst/internal/api/rest"
	"github.com/Elton-hst/internal/application/logger"
	"github.com/Elton-hst/internal/domain/services"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	service services.ProductService
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{
		service: service,
	}
}

// CreateProduct godoc
// @Summary      Create product
// @Description  Add a new product
// @Accept       json
// @Produce      json
// @Param        product body rest.CreateProductRequest true "Product data"
// @Success      201 {object} aggregate.Product "Product created successfully"
// @Failure      400 {object} map[string]string "Bad request"
// @Router       /product [post]
func (p *ProductController) CreateProduct(c echo.Context) error {
	var createProduct rest.CreateProductRequest

	if err := c.Bind(&createProduct); err != nil {
		logger.Error.Println("Failure to write the request body")
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Error": err.Error(),
		})
	}

	product := createProduct.ToAggregate()
	logger.Info.Printf("Start create new product: %s", product.PK)

	result, err := p.service.Create(product)
	if err != nil {
		logger.Error.Printf("Failure on create new product %s", product.PK)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Error": err.Error(),
		})
	}

	logger.Info.Print("Success on create new product")
	return c.JSON(http.StatusCreated, result)
}

// GetProduct   godoc
// @Summary     Recupera um produto por ID
// @Description Retorna os detalhes de um produto com base no ID fornecido.
// @Accept      json
// @Produce     json
// @Param       id path string true "ID do produto"
// @Success     201 {object} aggregate.Product "Product created successfully"
// @Failure     400 {object} map[string]string "Bad request"
// @Router      /product/{id} [get]
func (p *ProductController) GetProduct(c echo.Context) error {
	id := c.Param("id")
	result, err := p.service.FindById(id)
	if err != nil {
		logger.Error.Printf("Failure on find product %s", id)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Error": err.Error(),
		})
	}

	logger.Info.Printf("Success on find product %s", result.PK)
	return c.JSON(http.StatusOK, result)
}
