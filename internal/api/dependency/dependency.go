package dependency

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Dependency(db *gorm.DB, echo *echo.Echo) *echo.Echo {
	ProductDependency(db, echo)
	return echo
}