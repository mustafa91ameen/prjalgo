package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/container"
)

// RegisterRoleRoutes sets up /roles, /pages, and /role-pages routes
func RegisterRoleRoutes(rg *gin.RouterGroup, c *container.Container) {
	// Roles
	roles := rg.Group("/roles")
	{
		roles.GET("", c.RoleHandler.GetAll)
		roles.GET("/:id", c.RoleHandler.GetByID)
		roles.POST("", c.RoleHandler.Create)
		roles.PUT("/:id", c.RoleHandler.Update)
		roles.DELETE("/:id", c.RoleHandler.Delete)
	}

	// Pages
	pages := rg.Group("/pages")
	{
		pages.GET("", c.PageHandler.GetAll)
		pages.GET("/:id", c.PageHandler.GetByID)
		pages.POST("", c.PageHandler.Create)
		pages.PUT("/:id", c.PageHandler.Update)
		pages.DELETE("/:id", c.PageHandler.Delete)
	}

	// Role Pages
	rolePages := rg.Group("/role-pages")
	{
		rolePages.GET("", c.RolePageHandler.GetAll)
		rolePages.GET("/:id", c.RolePageHandler.GetByID)
		rolePages.POST("", c.RolePageHandler.Create)
		rolePages.PUT("/:id", c.RolePageHandler.Update)
		rolePages.DELETE("/:id", c.RolePageHandler.Delete)
	}
}
