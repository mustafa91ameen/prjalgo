package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/container"
)

// RegisterWorkDayRoutes sets up /workdays and related routes
func RegisterWorkDayRoutes(rg *gin.RouterGroup, c *container.Container) {
	// WorkDay
	workDays := rg.Group("/workdays")
	{
		workDays.GET("", c.WorkDayHandler.GetAll)
		workDays.GET("/:id", c.WorkDayHandler.GetByID)
		workDays.POST("", c.WorkDayHandler.Create)
		workDays.PUT("/:id", c.WorkDayHandler.Update)
		workDays.DELETE("/:id", c.WorkDayHandler.Delete)
	}

	// WorkDay Materials
	materials := rg.Group("/workday-materials")
	{
		materials.GET("", c.WorkDayMaterialHandler.GetAll)
		materials.GET("/:id", c.WorkDayMaterialHandler.GetByID)
		materials.POST("", c.WorkDayMaterialHandler.Create)
		materials.PUT("/:id", c.WorkDayMaterialHandler.Update)
		materials.DELETE("/:id", c.WorkDayMaterialHandler.Delete)
	}

	// WorkDay Labor
	labor := rg.Group("/workday-labor")
	{
		labor.GET("", c.WorkDayLaborHandler.GetAll)
		labor.GET("/:id", c.WorkDayLaborHandler.GetByID)
		labor.POST("", c.WorkDayLaborHandler.Create)
		labor.PUT("/:id", c.WorkDayLaborHandler.Update)
		labor.DELETE("/:id", c.WorkDayLaborHandler.Delete)
	}

	// WorkDay Equipment
	equipment := rg.Group("/workday-equipment")
	{
		equipment.GET("", c.WorkDayEquipmentHandler.GetAll)
		equipment.GET("/:id", c.WorkDayEquipmentHandler.GetByID)
		equipment.POST("", c.WorkDayEquipmentHandler.Create)
		equipment.PUT("/:id", c.WorkDayEquipmentHandler.Update)
		equipment.DELETE("/:id", c.WorkDayEquipmentHandler.Delete)
	}
}
