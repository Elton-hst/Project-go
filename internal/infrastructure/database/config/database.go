package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Elton-hst/internal/application/logger"
	"github.com/Elton-hst/internal/infrastructure/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	log "gorm.io/gorm/logger"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password int
}

func buildDatabase() *DBConfig {
	dbConfig := DBConfig{
		User:     os.Getenv("DB_USERNAME"),
		Password: envInt(os.Getenv("DB_PASSWORD")),
		Host:     os.Getenv("DB_HOST"),
		Port:     envInt(os.Getenv("DB_PORT")),
		DBName:   os.Getenv("DB_NAME"),
	}
	return &dbConfig
}

func envInt(envString string) int {
	portInt, err := strconv.Atoi(envString)
	if err != nil {
		logger.Error.Printf("Error to read variable: %s", envString)
	}
	return portInt
}

func dbUrl(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%d dbname=%s port=%d sslmode=disable",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
	)
}

func DataBaseInit() {
	dbConfig := buildDatabase()
	dbUrl := dbUrl(dbConfig)

	var err error
	DB, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		Logger: log.Default.LogMode(log.Info),
	})

	if err != nil {
		logger.Error.Println("Failed to initialize database")
	}

	initializeMigration(DB)
}

func initializeMigration(DB *gorm.DB) {
	order := migrations.NewOrderMigration(DB)
	orderTable := order.TableExist()

	if orderTable {
		if err := order.DropTable(); err != nil {
			logger.Error.Println("Failed to drop table: order")
		}

		if err := order.CreateTable(); err != nil {
			logger.Error.Println("Error creating order")
		}
	} else {
		if err := order.CreateTable(); err != nil {
			logger.Error.Println("Error creating order")
		}
	}

	customer := migrations.NewCustomerMigration(DB)
	customerTable := customer.TableExist()

	if customerTable {
		if err := customer.DropTable(); err != nil {
			logger.Error.Println("Failed to drop table: customer")
		}

		if err := customer.CreateTable(); err != nil {
			logger.Error.Println("Error creating customer")
		}
	} else {
		if err := customer.CreateTable(); err != nil {
			logger.Error.Println("Error creating customer")
		}
	}

	product := migrations.NewProductMigration(DB)
	productTable := product.TableExist()

	if productTable {
		if err := product.DropTable(); err != nil {
			logger.Error.Println("Failed to drop table: product")
		}

		if err := product.CreateTable(); err != nil {
			logger.Error.Println("Error creating product")
		}
	} else {
		if err := product.CreateTable(); err != nil {
			logger.Error.Println("Error creating product")
		}
	}

}

func GetDB() *gorm.DB {
	return DB
}
