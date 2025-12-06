package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type RoleHandler struct {
	roleService *services.RoleService
}

func NewRoleHandler(roleService *services.RoleService) *RoleHandler {
	return &RoleHandler{
		roleService: roleService,
	}
}

// GetAll handles GET /roles
func (h *RoleHandler) GetAll(c *gin.Context) {
	roles, err := h.roleService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch roles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": roles})
}

// GetByID handles GET /roles/:id
func (h *RoleHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid role id"})
		return
	}

	role, err := h.roleService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrRoleNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": role})
}

// Create handles POST /roles
func (h *RoleHandler) Create(c *gin.Context) {
	var req dtos.CreateRole
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	role, err := h.roleService.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to create role"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": role})
}

// Update handles PUT /roles/:id
func (h *RoleHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid role id"})
		return
	}

	var req dtos.UpdateRole
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	role, err := h.roleService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrRoleNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to update role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": role})
}

// Delete handles DELETE /roles/:id
func (h *RoleHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid role id"})
		return
	}

	err = h.roleService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrRoleNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to delete role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": nil})
}

// parseID extracts an int64 ID from the URL path
func (h *RoleHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
