package router

import (
	"github.com/Elton-hst/internal/api/controller"
	"github.com/labstack/echo/v4"
)

type Router struct {
	Echo              *echo.Echo
	ProductController controller.ProductController
}

func NewRouter(echo *echo.Echo, productController controller.ProductController) *Router {
	return &Router{
		Echo:              echo,
		ProductController: productController,
	}
}

func (r *Router) Router() *echo.Echo {
	v1 := r.Echo.Group("/api/v1")
	{
		product := v1.Group("/product")
		{
			product.POST("", r.ProductController.CreateProduct)
			product.GET(":id", r.ProductController.GetProduct)
		}
	}

	return r.Echo
}
