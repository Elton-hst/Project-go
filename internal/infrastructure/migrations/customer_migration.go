package migrations

import (
	"github.com/Elton-hst/internal/infrastructure/database/entities"
	"gorm.io/gorm"
)

type CustomerMigration struct {
	Database *gorm.DB
	Customer  *entities.CustomerDatabase
}

func NewCustomerMigration(db *gorm.DB) *CustomerMigration {
	return &CustomerMigration{Database: db}
}


func (c *CustomerMigration) CreateTable() error {
	if err := c.Database.Migrator().AutoMigrate(c.Customer); err != nil {
		return err
	}
	return nil
}

func (c *CustomerMigration) DropTable() error {
	if err := c.Database.Migrator().DropTable(c.Customer); err != nil {
		return err
	}
	return nil
}

func (c *CustomerMigration) TableExist() bool {
	return c.Database.Migrator().HasTable(c.Customer)
}
