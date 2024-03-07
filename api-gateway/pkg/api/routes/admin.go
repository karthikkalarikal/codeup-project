package routes

import (
	handler "github.com/karthikkalarikal/api-gateway/pkg/api/handlers/interfaces"
	"github.com/labstack/echo/v4"
)

func SetupAdminRoutes(e *echo.Group, adminHandler handler.AdminHandler) {
	// auth := e.Group("/user")
	problem := e.Group("/problem")

	{
		problem.POST("/", adminHandler.CreateProblem)
		// problem.GET("/")
		// problem.GET("/:id")
		// problem.PATCH("/")
		// problem.DELETE("/:id")
		// problem.PUT("/:id")
	}
}
