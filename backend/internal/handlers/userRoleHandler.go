package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type UserRoleHandler struct {
	userRoleService *services.UserRoleService
}

func NewUserRoleHandler(userRoleService *services.UserRoleService) *UserRoleHandler {
	return &UserRoleHandler{
		userRoleService: userRoleService,
	}
}

// GetAll handles GET /user-roles
func (h *UserRoleHandler) GetAll(c *gin.Context) {
	userRoles, err := h.userRoleService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch user roles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": userRoles})
}

// GetByID handles GET /user-roles/:id
func (h *UserRoleHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid user role id"})
		return
	}

	userRole, err := h.userRoleService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrUserRoleNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch user role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": userRole})
}

// GetByUserID handles GET /users/:id/roles
func (h *UserRoleHandler) GetByUserID(c *gin.Context) {
	userID, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid user id"})
		return
	}

	userRoles, err := h.userRoleService.GetByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch user roles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": userRoles})
}

// Create handles POST /user-roles
func (h *UserRoleHandler) Create(c *gin.Context) {
	var req dtos.CreateUserRole
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	userRole, err := h.userRoleService.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to create user role"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": userRole})
}

// Delete handles DELETE /user-roles/:id
func (h *UserRoleHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid user role id"})
		return
	}

	err = h.userRoleService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrUserRoleNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to delete user role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": nil})
}

// parseID extracts an int64 ID from the URL path
func (h *UserRoleHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
