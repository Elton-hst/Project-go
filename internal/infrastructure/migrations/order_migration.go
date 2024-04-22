package migrations

import (
	"github.com/Elton-hst/internal/infrastructure/database/entities"
	"gorm.io/gorm"
)

type OrderMigration struct {
	Database  *gorm.DB
	Order     *entities.OrderDatabase
	OrderItem *entities.OrderItemDatabase
}

func NewOrderMigration(db *gorm.DB) *OrderMigration {
	return &OrderMigration{Database: db}
}

func (or *OrderMigration) CreateTable() error {
	if err := or.Database.Migrator().AutoMigrate(or.OrderItem, or.Order); err != nil {
		return err
	}
	return nil
}

func (or *OrderMigration) DropTable() error {
	if err := or.Database.Migrator().DropTable(or.OrderItem, or.Order); err != nil {
		return err
	}
	return nil
}

func (or *OrderMigration) TableExist() bool {
	return or.Database.Migrator().HasTable(or.Order)
}
