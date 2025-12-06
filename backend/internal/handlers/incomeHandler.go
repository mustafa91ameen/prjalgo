package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type IncomeHandler struct {
	incomeService *services.IncomeService
}

func NewIncomeHandler(incomeService *services.IncomeService) *IncomeHandler {
	return &IncomeHandler{
		incomeService: incomeService,
	}
}

// GetAll handles GET /incomes
func (h *IncomeHandler) GetAll(c *gin.Context) {
	incomes, err := h.incomeService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch incomes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": incomes})
}

// GetByID handles GET /incomes/:id
func (h *IncomeHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid income id"})
		return
	}

	income, err := h.incomeService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrIncomeNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch income"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": income})
}

// Create handles POST /incomes
func (h *IncomeHandler) Create(c *gin.Context) {
	var req dtos.CreateIncome
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	income, err := h.incomeService.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to create income"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": income})
}

// Update handles PUT /incomes/:id
func (h *IncomeHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid income id"})
		return
	}

	var req dtos.UpdateIncome
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	income, err := h.incomeService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrIncomeNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to update income"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": income})
}

// Delete handles DELETE /incomes/:id
func (h *IncomeHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid income id"})
		return
	}

	err = h.incomeService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrIncomeNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to delete income"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": nil})
}

// parseID extracts an int64 ID from the URL path
func (h *IncomeHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
