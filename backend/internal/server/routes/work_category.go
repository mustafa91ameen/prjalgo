package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/container"
)

// RegisterWorkCategoryRoutes sets up /work-categories and /work-subcategories routes
func RegisterWorkCategoryRoutes(rg *gin.RouterGroup, c *container.Container) {
	// Work Categories
	categories := rg.Group("/work-categories")
	{
		categories.GET("", c.WorkCategoryHandler.GetAll)
		categories.GET("/:id", c.WorkCategoryHandler.GetByID)
		categories.POST("", c.WorkCategoryHandler.Create)
		categories.PUT("/:id", c.WorkCategoryHandler.Update)
		categories.DELETE("/:id", c.WorkCategoryHandler.Delete)
	}

	// Work Sub-Categories
	subCategories := rg.Group("/work-subcategories")
	{
		subCategories.GET("", c.WorkSubCategoryHandler.GetAll)
		subCategories.GET("/:id", c.WorkSubCategoryHandler.GetByID)
		subCategories.POST("", c.WorkSubCategoryHandler.Create)
		subCategories.PUT("/:id", c.WorkSubCategoryHandler.Update)
		subCategories.DELETE("/:id", c.WorkSubCategoryHandler.Delete)
	}
}
