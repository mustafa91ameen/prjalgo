package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type WorkDayHandler struct {
	workDayService *services.WorkDayService
}

func NewWorkDayHandler(workDayService *services.WorkDayService) *WorkDayHandler {
	return &WorkDayHandler{
		workDayService: workDayService,
	}
}

// GetAll handles GET /workdays
func (h *WorkDayHandler) GetAll(c *gin.Context) {
	workDays, err := h.workDayService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch work days"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": workDays})
}

// GetByID handles GET /workdays/:id
func (h *WorkDayHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid work day id"})
		return
	}

	workDay, err := h.workDayService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch work day"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": workDay})
}

// GetByProjectID handles GET /projects/:id/workdays
func (h *WorkDayHandler) GetByProjectID(c *gin.Context) {
	projectID, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid project id"})
		return
	}

	workDays, err := h.workDayService.GetByProjectID(c.Request.Context(), projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch work days"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": workDays})
}

// Create handles POST /workdays
func (h *WorkDayHandler) Create(c *gin.Context) {
	var req dtos.CreateWorkDay
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	workDay, err := h.workDayService.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to create work day"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": workDay})
}

// Update handles PUT /workdays/:id
func (h *WorkDayHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid work day id"})
		return
	}

	var req dtos.UpdateWorkDay
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	workDay, err := h.workDayService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to update work day"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": workDay})
}

// Delete handles DELETE /workdays/:id
func (h *WorkDayHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid work day id"})
		return
	}

	err = h.workDayService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to delete work day"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": nil})
}

// parseID extracts an int64 ID from the URL path
func (h *WorkDayHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
