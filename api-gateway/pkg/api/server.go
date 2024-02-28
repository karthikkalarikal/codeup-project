package api

import (
	"fmt"
	"net/http"

	handler "github.com/karthikkalarikal/api-gateway/pkg/api/handlers/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/api/routes"
	"github.com/karthikkalarikal/api-gateway/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	port   string
	engine *echo.Echo
}

func NewServerHTTP(cfg *config.Config, authHandler handler.AuthHandler) *Server {
	fmt.Println("here")
	e := echo.New()

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
	routes.SetupUserRoutes(e.Group("/user"), authHandler)

	return &Server{
		engine: e,
		port:   cfg.Port,
	}
}

func (c *Server) Start() {
	// c.engine.Run(c.port)
	c.engine.Start(c.port)
}
