package routes

import (
	handler "github.com/karthikkalarikal/api-gateway/pkg/api/handlers/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/api/middleware"
	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(e *echo.Group, authHandler handler.AuthHandler, userHandler handler.UserHandler) {
	// auth := e.Group("/user")

	e.POST("/signup", authHandler.UserSignUp)
	e.POST("/signin", authHandler.UserSignIn)
	e.GET("/view", userHandler.ViewAllProblems)
	// e.POST("/logout", authHandler.UserLogout)
	// userManagent := e.Group("/logout")
	problem := e.Group("/problem")
	problem.Use(middleware.UserMiddleware)
	{
		problem.POST("/{id}", userHandler.GetOneProblemById)
	}
}
