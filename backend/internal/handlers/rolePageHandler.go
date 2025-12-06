package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type RolePageHandler struct {
	rolePageService *services.RolePageService
}

func NewRolePageHandler(rolePageService *services.RolePageService) *RolePageHandler {
	return &RolePageHandler{
		rolePageService: rolePageService,
	}
}

// GetAll handles GET /role-pages
func (h *RolePageHandler) GetAll(c *gin.Context) {
	rolePages, err := h.rolePageService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch role pages"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": rolePages})
}

// GetByID handles GET /role-pages/:id
func (h *RolePageHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid role page id"})
		return
	}

	rolePage, err := h.rolePageService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrRolePageNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch role page"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": rolePage})
}

// GetByRoleID handles GET /roles/:roleId/pages
func (h *RolePageHandler) GetByRoleID(c *gin.Context) {
	roleID, err := h.parseID(c, "roleId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid role id"})
		return
	}

	rolePages, err := h.rolePageService.GetByRoleID(c.Request.Context(), roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch role pages"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": rolePages})
}

// Create handles POST /role-pages
func (h *RolePageHandler) Create(c *gin.Context) {
	var req dtos.CreateRolePage
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	rolePage, err := h.rolePageService.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to create role page"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": rolePage})
}

// Update handles PUT /role-pages/:id
func (h *RolePageHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid role page id"})
		return
	}

	var req dtos.UpdateRolePage
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	rolePage, err := h.rolePageService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrRolePageNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to update role page"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": rolePage})
}

// Delete handles DELETE /role-pages/:id
func (h *RolePageHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid role page id"})
		return
	}

	err = h.rolePageService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrRolePageNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to delete role page"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": nil})
}

// parseID extracts an int64 ID from the URL path
func (h *RolePageHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
