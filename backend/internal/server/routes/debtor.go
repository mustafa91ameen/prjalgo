package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/container"
)

// RegisterDebtorRoutes sets up /debtors routes
func RegisterDebtorRoutes(rg *gin.RouterGroup, c *container.Container) {
	debtors := rg.Group("/debtors")
	{
		debtors.GET("", c.DebtorHandler.GetAll)
		debtors.GET("/:id", c.DebtorHandler.GetByID)
		debtors.POST("", c.DebtorHandler.Create)
		debtors.PUT("/:id", c.DebtorHandler.Update)
		debtors.DELETE("/:id", c.DebtorHandler.Delete)
	}
}
