package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/container"
)

// RegisterRoutes sets up all API routes
func RegisterRoutes(router *gin.Engine, c *container.Container) {
	// Health check
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	// API v1 group
	api := router.Group("/api/v1")

	// Register all domain routes
	RegisterProjectRoutes(api, c)
	RegisterWorkDayRoutes(api, c)
	RegisterWorkCategoryRoutes(api, c)
	RegisterFinanceRoutes(api, c)
	RegisterDebtorRoutes(api, c)
	RegisterUserRoutes(api, c)
	RegisterRoleRoutes(api, c)
}
