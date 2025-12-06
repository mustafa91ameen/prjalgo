package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type PageHandler struct {
	pageService *services.PageService
}

func NewPageHandler(pageService *services.PageService) *PageHandler {
	return &PageHandler{
		pageService: pageService,
	}
}

// GetAll handles GET /pages
func (h *PageHandler) GetAll(c *gin.Context) {
	pages, err := h.pageService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch pages"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": pages})
}

// GetByID handles GET /pages/:id
func (h *PageHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid page id"})
		return
	}

	page, err := h.pageService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrPageNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch page"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": page})
}

// GetActivePages handles GET /pages/active
func (h *PageHandler) GetActivePages(c *gin.Context) {
	pages, err := h.pageService.GetActivePages(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch active pages"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": pages})
}

// Create handles POST /pages
func (h *PageHandler) Create(c *gin.Context) {
	var req dtos.CreatePage
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	page, err := h.pageService.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to create page"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": page})
}

// Update handles PUT /pages/:id
func (h *PageHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid page id"})
		return
	}

	var req dtos.UpdatePage
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	page, err := h.pageService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrPageNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to update page"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": page})
}

// Delete handles DELETE /pages/:id
func (h *PageHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid page id"})
		return
	}

	err = h.pageService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrPageNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to delete page"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": nil})
}

// parseID extracts an int64 ID from the URL path
func (h *PageHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
