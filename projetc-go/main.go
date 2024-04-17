package main

import (
	"os"
	"path/filepath"
	"runtime"

	_ "github.com/Elton-hst/docs"
	"github.com/Elton-hst/internal/api/server"
	"github.com/Elton-hst/internal/application/logger"
	"github.com/Elton-hst/internal/infrastructure/database/config"
	"github.com/joho/godotenv"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	if err := godotenv.Load(basepath + "/.env"); err != nil {
		logger.Error.Println("Failed to fetch environment variables")
	}
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for using Swagger with Echo.
// @host localhost:8081
// @BasePath /api/v1
func main() {
	config.DataBaseInit()
	defer config.GetDB().DB()

	serve := server.NewServer()
	seila := serve.Start()

	if err := seila.Start(os.Getenv("SERVER")); err != nil {
		logger.Error.Println("Failed to initialize the server")
	}

}
