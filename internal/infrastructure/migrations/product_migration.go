package migrations

import (
	"github.com/Elton-hst/internal/infrastructure/database/entities"
	"gorm.io/gorm"
)

type ProductMigration struct {
	Database *gorm.DB
	Product  *entities.ProductDatabase
}

func NewProductMigration(db *gorm.DB) *ProductMigration {
	return &ProductMigration{Database: db}
}


func (pm *ProductMigration) CreateTable() error {
	if err := pm.Database.Migrator().AutoMigrate(pm.Product); err != nil {
		return err
	}
	return nil
}

func (pm *ProductMigration) DropTable() error {
	if err := pm.Database.Migrator().DropTable(pm.Product); err != nil {
		return err
	}
	return nil
}

func (pm *ProductMigration) TableExist() bool {
	return pm.Database.Migrator().HasTable(pm.Product)
}
