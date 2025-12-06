package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type WorkDayLaborHandler struct {
	laborService *services.WorkDayLaborService
}

func NewWorkDayLaborHandler(laborService *services.WorkDayLaborService) *WorkDayLaborHandler {
	return &WorkDayLaborHandler{
		laborService: laborService,
	}
}

// GetAll handles GET /workday-labor
func (h *WorkDayLaborHandler) GetAll(c *gin.Context) {
	labors, err := h.laborService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch labors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": labors})
}

// GetByWorkDayID handles GET /workdays/:workDayId/labors
func (h *WorkDayLaborHandler) GetByWorkDayID(c *gin.Context) {
	workDayID, err := h.parseID(c, "workDayId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid work day id"})
		return
	}

	labors, err := h.laborService.GetByWorkDayID(c.Request.Context(), workDayID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch labors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": labors})
}

// GetByID handles GET /workdays/:workDayId/labors/:id
func (h *WorkDayLaborHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid labor id"})
		return
	}

	labor, err := h.laborService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayLaborNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch labor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": labor})
}

// Create handles POST /workdays/:workDayId/labors
func (h *WorkDayLaborHandler) Create(c *gin.Context) {
	workDayID, err := h.parseID(c, "workDayId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid work day id"})
		return
	}

	var req dtos.CreateWorkDayLabor
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Set workDayId from URL path
	req.WorkDayID = workDayID

	labor, err := h.laborService.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to create labor"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": labor})
}

// Update handles PUT /workdays/:workDayId/labors/:id
func (h *WorkDayLaborHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid labor id"})
		return
	}

	var req dtos.UpdateWorkDayLabor
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	labor, err := h.laborService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayLaborNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to update labor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": labor})
}

// Delete handles DELETE /workdays/:workDayId/labors/:id
func (h *WorkDayLaborHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid labor id"})
		return
	}

	err = h.laborService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayLaborNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to delete labor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": nil})
}

// parseID extracts an int64 ID from the URL path
func (h *WorkDayLaborHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
