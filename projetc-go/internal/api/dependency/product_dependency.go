package dependency

import (
	"github.com/Elton-hst/internal/api/controller"
	"github.com/Elton-hst/internal/api/router"
	"github.com/Elton-hst/internal/application/services"
	"github.com/Elton-hst/internal/infrastructure/database/repositories"
	"github.com/Elton-hst/internal/infrastructure/migrations"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ProductDependency(db *gorm.DB, echo *echo.Echo) *echo.Echo {
	productMigration := migrations.NewProductMigration(db)
	productRepository := repositories.NewGormProductRepository(productMigration.Database)
	productService := services.NewProductService(productRepository)
	productController := controller.NewProductController(productService)
	route := router.NewRouter(echo, *productController)
	return route.Router()
}