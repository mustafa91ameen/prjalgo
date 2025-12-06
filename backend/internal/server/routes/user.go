package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/container"
)

// RegisterUserRoutes sets up /users and /user-roles routes
func RegisterUserRoutes(rg *gin.RouterGroup, c *container.Container) {
	// Users
	users := rg.Group("/users")
	{
		users.GET("", c.UserHandler.GetAll)
		users.GET("/:id", c.UserHandler.GetByID)
		users.POST("", c.UserHandler.Create)
		users.PUT("/:id", c.UserHandler.Update)
		users.DELETE("/:id", c.UserHandler.Delete)

		// Nested route under user
		users.GET("/:id/roles", c.UserRoleHandler.GetByUserID)
	}

	// User Roles
	userRoles := rg.Group("/user-roles")
	{
		userRoles.GET("", c.UserRoleHandler.GetAll)
		userRoles.GET("/:id", c.UserRoleHandler.GetByID)
		userRoles.POST("", c.UserRoleHandler.Create)
		userRoles.DELETE("/:id", c.UserRoleHandler.Delete)
	}
}
