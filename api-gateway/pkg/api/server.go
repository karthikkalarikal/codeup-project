package api

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/karthikkalarikal/api-gateway/pkg/api/handlers/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/api/routes"
	"github.com/karthikkalarikal/api-gateway/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server struct {
	port   string
	engine *echo.Echo
}

func NewServerHTTP(cfg *config.Config, authHandler handler.AuthHandler, userHandler handler.UserHandler, adminHandler handler.AdminHandler) *Server {
	fmt.Println("here in server")
	e := echo.New()
	// to do - use custom logger to make the echo's logging more appealing , zap package.
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPost, http.MethodConnect},
		AllowHeaders:     []string{echo.HeaderAccept, echo.HeaderContentType, echo.HeaderAuthorization, echo.HeaderXCSRFToken},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})) // to allow front end to connect

	routes.SetupUserRoutes(e.Group("/user"), authHandler, userHandler)
	routes.SetupAdminRoutes(e.Group("/admin"), adminHandler)

	return &Server{
		engine: e,
		port:   cfg.Port,
	}
}

func (c *Server) Start() {
	// c.engine.Run(c.port)
	fmt.Println("port", c.port)
	if err := c.engine.Start(fmt.Sprintf(":%s", c.port)); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
