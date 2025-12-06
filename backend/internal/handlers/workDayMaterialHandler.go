package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type WorkDayMaterialHandler struct {
	materialService *services.WorkDayMaterialService
}

func NewWorkDayMaterialHandler(materialService *services.WorkDayMaterialService) *WorkDayMaterialHandler {
	return &WorkDayMaterialHandler{
		materialService: materialService,
	}
}

// GetAll handles GET /workday-materials
func (h *WorkDayMaterialHandler) GetAll(c *gin.Context) {
	materials, err := h.materialService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch materials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": materials})
}

// GetByWorkDayID handles GET /workdays/:workDayId/materials
func (h *WorkDayMaterialHandler) GetByWorkDayID(c *gin.Context) {
	workDayID, err := h.parseID(c, "workDayId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid work day id"})
		return
	}

	materials, err := h.materialService.GetByWorkDayID(c.Request.Context(), workDayID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch materials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": materials})
}

// GetByID handles GET /workdays/:workDayId/materials/:id
func (h *WorkDayMaterialHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid material id"})
		return
	}

	material, err := h.materialService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayMaterialNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch material"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": material})
}

// Create handles POST /workdays/:workDayId/materials
func (h *WorkDayMaterialHandler) Create(c *gin.Context) {
	workDayID, err := h.parseID(c, "workDayId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid work day id"})
		return
	}

	var req dtos.CreateWorkDayMaterial
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Set workDayId from URL path
	req.WorkDayID = workDayID

	material, err := h.materialService.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to create material"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": material})
}

// Update handles PUT /workdays/:workDayId/materials/:id
func (h *WorkDayMaterialHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid material id"})
		return
	}

	var req dtos.UpdateWorkDayMaterial
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	material, err := h.materialService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayMaterialNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to update material"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": material})
}

// Delete handles DELETE /workdays/:workDayId/materials/:id
func (h *WorkDayMaterialHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid material id"})
		return
	}

	err = h.materialService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayMaterialNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to delete material"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": nil})
}

// parseID extracts an int64 ID from the URL path
func (h *WorkDayMaterialHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
