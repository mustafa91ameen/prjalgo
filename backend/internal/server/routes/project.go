package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/container"
)

// RegisterProjectRoutes sets up /projects routes
func RegisterProjectRoutes(rg *gin.RouterGroup, c *container.Container) {
	projects := rg.Group("/projects")
	{
		projects.GET("", c.ProjectHandler.GetAll)
		projects.GET("/:id", c.ProjectHandler.GetByID)
		projects.POST("", c.ProjectHandler.Create)
		projects.PUT("/:id", c.ProjectHandler.Update)
		projects.DELETE("/:id", c.ProjectHandler.Delete)

		// Nested resources under project
		projects.GET("/:id/workdays", c.WorkDayHandler.GetByProjectID)
		projects.GET("/:id/expenses", c.ExpenseHandler.GetByProjectID)
	}
}
