package routes

import (
	handler "github.com/karthikkalarikal/api-gateway/pkg/api/handlers/interfaces"
	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(e *echo.Group, authHandler handler.AuthHandler, userHandler handler.UserHandler) {
	// auth := e.Group("/user")

	e.POST("/signup", authHandler.UserSignUp)
	e.POST("/signin", authHandler.UserSignIn)
	e.GET("/view", userHandler.ViewAllProblems)

}
