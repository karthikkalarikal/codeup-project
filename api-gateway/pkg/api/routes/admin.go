package routes

import (
	handler "github.com/karthikkalarikal/api-gateway/pkg/api/handlers/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/api/middleware"
	"github.com/labstack/echo/v4"
)

func SetupAdminRoutes(e *echo.Group, adminHandler handler.AdminHandler) {
	// auth := e.Group("/user")
	adminProblem := e.Group("/problem")
	adminProblem.Use(middleware.AdminMiddleware)
	{
		adminProblem.POST("/", adminHandler.CreateProblem)
		adminProblem.PUT("/first/:id", adminHandler.InsertFirstHalfProblem)
		adminProblem.PUT("/second/:id", adminHandler.InsertSecondHalfProblem)
		// problem.GET("/")
		// problem.GET("/:id")
		// problem.PATCH("/")
		// problem.DELETE("/:id")
		// problem.PUT("/:id")
	}
	adminUser := e.Group("/user")
	adminUser.Use(middleware.AdminMiddleware)
	{
		adminUser.GET("/", adminHandler.ViewUsers)
	}

}
