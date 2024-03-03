package routes

import (
	"github.com/golang-jwt/jwt"
	handler "github.com/karthikkalarikal/api-gateway/pkg/api/handlers/interfaces"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(e *echo.Group, authHandler handler.AuthHandler) {
	// auth := e.Group("/user")

	signup := e.Group("/signup")
	{
		signup.POST("/", authHandler.UserSignUp)
	}

	
	signup.Use(echojwt.WithConfig())

}
