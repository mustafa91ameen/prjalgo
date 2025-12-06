package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetAll handles GET /users
func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.userService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": users})
}

// GetByID handles GET /users/:id
func (h *UserHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid user id"})
		return
	}

	user, err := h.userService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to fetch user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": user})
}

// Create handles POST /users
func (h *UserHandler) Create(c *gin.Context) {
	var req dtos.CreateUser
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	user, err := h.userService.Create(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, services.ErrUserExists) {
			c.JSON(http.StatusConflict, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": user})
}

// Update handles PUT /users/:id
func (h *UserHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid user id"})
		return
	}

	var req dtos.UpdateUser
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	user, err := h.userService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": user})
}

// Delete handles DELETE /users/:id
func (h *UserHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid user id"})
		return
	}

	err = h.userService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": nil})
}

// parseID extracts an int64 ID from the URL path
func (h *UserHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
