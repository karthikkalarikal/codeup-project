package routes

import (
	handler "github.com/karthikkalarikal/api-gateway/pkg/api/handlers/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/api/middleware"
	"github.com/labstack/echo/v4"
)

func SetupAdminRoutes(e *echo.Group, adminHandler handler.AdminHandler) {
	// auth := e.Group("/user")
	adminProblem := e.Group("/problem")
	adminProblem.Use(middleware.UserMiddleware)
	{
		adminProblem.POST("/", adminHandler.CreateProblem)
		// problem.GET("/")
		// problem.GET("/:id")
		// problem.PATCH("/")
		// problem.DELETE("/:id")
		// problem.PUT("/:id")
	}
}