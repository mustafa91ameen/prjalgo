package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type WorkSubCategoryHandler struct {
	subCategoryService *services.WorkSubCategoryService
}

func NewWorkSubCategoryHandler(subCategoryService *services.WorkSubCategoryService) *WorkSubCategoryHandler {
	return &WorkSubCategoryHandler{
		subCategoryService: subCategoryService,
	}
}

// GetAll handles GET /work-subcategories
func (h *WorkSubCategoryHandler) GetAll(c *gin.Context) {
	subCategories, err := h.subCategoryService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch work subcategories"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": subCategories})
}

// GetByID handles GET /work-subcategories/:id
func (h *WorkSubCategoryHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid work subcategory id"})
		return
	}

	subCategory, err := h.subCategoryService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkSubCategoryNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch work subcategory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": subCategory})
}

// GetByCategoryID handles GET /work-categories/:categoryId/subcategories
func (h *WorkSubCategoryHandler) GetByCategoryID(c *gin.Context) {
	categoryID, err := h.parseID(c, "categoryId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid category id"})
		return
	}

	subCategories, err := h.subCategoryService.GetByCategoryID(c.Request.Context(), categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch work subcategories"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": subCategories})
}

// Create handles POST /work-subcategories
func (h *WorkSubCategoryHandler) Create(c *gin.Context) {
	var req dtos.CreateWorkSubCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	subCategory, err := h.subCategoryService.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to create work subcategory"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": subCategory})
}

// Update handles PUT /work-subcategories/:id
func (h *WorkSubCategoryHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid work subcategory id"})
		return
	}

	var req dtos.UpdateWorkSubCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	subCategory, err := h.subCategoryService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrWorkSubCategoryNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to update work subcategory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": subCategory})
}

// Delete handles DELETE /work-subcategories/:id
func (h *WorkSubCategoryHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid work subcategory id"})
		return
	}

	err = h.subCategoryService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkSubCategoryNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to delete work subcategory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": nil})
}

// parseID extracts an int64 ID from the URL path
func (h *WorkSubCategoryHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
