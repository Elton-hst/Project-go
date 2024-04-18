package server

import (
	"github.com/Elton-hst/internal/api/dependency"
	"github.com/Elton-hst/internal/infrastructure/database/config"
	"github.com/Elton-hst/internal/infrastructure/redis"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server struct {
	server *echo.Echo
	Port   int
}

func NewServer() *Server {
	return &Server{
		server: echo.New(),
	}
}

func (s *Server) Start() *echo.Echo {
	dependency.Dependency(config.GetDB(), s.server)
	redis.NewRedis()

	s.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "URI: ${uri} | METHOD: ${method} | STATUS: ${status} | ERROR: ${error} | LATENCY ${latency} \n",
	}))
	s.server.Use(middleware.Recover())
	s.server.Use(middleware.CORS())

	s.server.GET("/swagger/*", echoSwagger.WrapHandler)
	s.server.GET("/ping", healthCheck)

	return s.server
}

func (s *Server) GetServer() *echo.Echo {
	return s.server
}

func healthCheck(c echo.Context) error {
	return c.JSON(200, "pong")
}
