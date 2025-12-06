package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/container"
)

// RegisterFinanceRoutes sets up /expenses and /income routes
func RegisterFinanceRoutes(rg *gin.RouterGroup, c *container.Container) {
	// Expenses
	expenses := rg.Group("/expenses")
	{
		expenses.GET("", c.ExpenseHandler.GetAll)
		expenses.GET("/:id", c.ExpenseHandler.GetByID)
		expenses.POST("", c.ExpenseHandler.Create)
		expenses.PUT("/:id", c.ExpenseHandler.Update)
		expenses.DELETE("/:id", c.ExpenseHandler.Delete)
	}

	// Income
	income := rg.Group("/income")
	{
		income.GET("", c.IncomeHandler.GetAll)
		income.GET("/:id", c.IncomeHandler.GetByID)
		income.POST("", c.IncomeHandler.Create)
		income.PUT("/:id", c.IncomeHandler.Update)
		income.DELETE("/:id", c.IncomeHandler.Delete)
	}
}
