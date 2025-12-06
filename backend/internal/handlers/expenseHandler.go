package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type ExpenseHandler struct {
	expenseService *services.ExpenseService
}

func NewExpenseHandler(expenseService *services.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{
		expenseService: expenseService,
	}
}

// GetAll handles GET /expenses
func (h *ExpenseHandler) GetAll(c *gin.Context) {
	expenses, err := h.expenseService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch expenses"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": expenses})
}

// GetByID handles GET /expenses/:id
func (h *ExpenseHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid expense id"})
		return
	}

	expense, err := h.expenseService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrExpenseNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch expense"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": expense})
}

// GetByProjectID handles GET /projects/:id/expenses
func (h *ExpenseHandler) GetByProjectID(c *gin.Context) {
	projectID, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid project id"})
		return
	}

	expenses, err := h.expenseService.GetByProjectID(c.Request.Context(), projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch expenses"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": expenses})
}

// Create handles POST /expenses
func (h *ExpenseHandler) Create(c *gin.Context) {
	var req dtos.CreateExpense
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	expense, err := h.expenseService.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to create expense"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": expense})
}

// Update handles PUT /expenses/:id
func (h *ExpenseHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid expense id"})
		return
	}

	var req dtos.UpdateExpense
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	expense, err := h.expenseService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrExpenseNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to update expense"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": expense})
}

// Delete handles DELETE /expenses/:id
func (h *ExpenseHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid expense id"})
		return
	}

	err = h.expenseService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrExpenseNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to delete expense"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": nil})
}

// parseID extracts an int64 ID from the URL path
func (h *ExpenseHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
