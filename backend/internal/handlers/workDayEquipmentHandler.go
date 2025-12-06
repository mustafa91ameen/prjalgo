package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type WorkDayEquipmentHandler struct {
	equipmentService *services.WorkDayEquipmentService
}

func NewWorkDayEquipmentHandler(equipmentService *services.WorkDayEquipmentService) *WorkDayEquipmentHandler {
	return &WorkDayEquipmentHandler{
		equipmentService: equipmentService,
	}
}

// GetAll handles GET /workday-equipment
func (h *WorkDayEquipmentHandler) GetAll(c *gin.Context) {
	equipments, err := h.equipmentService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch equipments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": equipments})
}

// GetByWorkDayID handles GET /workdays/:workDayId/equipments
func (h *WorkDayEquipmentHandler) GetByWorkDayID(c *gin.Context) {
	workDayID, err := h.parseID(c, "workDayId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid work day id"})
		return
	}

	equipments, err := h.equipmentService.GetByWorkDayID(c.Request.Context(), workDayID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch equipments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": equipments})
}

// GetByID handles GET /workdays/:workDayId/equipments/:id
func (h *WorkDayEquipmentHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid equipment id"})
		return
	}

	equipment, err := h.equipmentService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayEquipmentNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch equipment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": equipment})
}

// Create handles POST /workdays/:workDayId/equipments
func (h *WorkDayEquipmentHandler) Create(c *gin.Context) {
	workDayID, err := h.parseID(c, "workDayId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid work day id"})
		return
	}

	var req dtos.CreateWorkDayEquipment
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Set workDayId from URL path
	req.WorkDayID = workDayID

	equipment, err := h.equipmentService.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to create equipment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": equipment})
}

// Update handles PUT /workdays/:workDayId/equipments/:id
func (h *WorkDayEquipmentHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid equipment id"})
		return
	}

	var req dtos.UpdateWorkDayEquipment
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	equipment, err := h.equipmentService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayEquipmentNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to update equipment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": equipment})
}

// Delete handles DELETE /workdays/:workDayId/equipments/:id
func (h *WorkDayEquipmentHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid equipment id"})
		return
	}

	err = h.equipmentService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayEquipmentNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to delete equipment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": nil})
}

// parseID extracts an int64 ID from the URL path
func (h *WorkDayEquipmentHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
