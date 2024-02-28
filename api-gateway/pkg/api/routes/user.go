package routes

import (
	handler "github.com/karthikkalarikal/api-gateway/pkg/api/handlers/interfaces"
	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(e *echo.Group, authHandler handler.AuthHandler) {
	// auth := e.Group("/user")

	signup := e.Group("/signup")
	{
		signup.POST("/", authHandler.UserSignUp)
	}

}
