package routes

import (
	handler "github.com/karthikkalarikal/api-gateway/pkg/api/handlers/interfaces"
	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(e *echo.Group, authHandler handler.AuthHandler) {
	// auth := e.Group("/user")

	e.POST("/signup", authHandler.UserSignUp)
	e.POST("/signin", authHandler.UserSignIn)
	// signup.Use(middleware.UserMiddleware)
	// {
	// 	signup.POST("/", authHandler.UserSignUp)
	// }

}
