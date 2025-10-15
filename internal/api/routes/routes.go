package routes

import (
	"go-api-backend/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes creates a gin router, registers routes and returns it.
func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/health", handlers.HealthCheck)

	// User routes
	router.GET("/users", handlers.GetUsers)
	router.POST("/users", handlers.CreateUser)

	return router
}
