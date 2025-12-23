package handlers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/mustafa91ameen/prjalgo/backend/internal/auth"
	"github.com/mustafa91ameen/prjalgo/backend/internal/response"
	"github.com/mustafa91ameen/prjalgo/backend/internal/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	result, err := h.authService.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			response.Unauthorized(c, "invalid username or password")
			return
		}
		if errors.Is(err, services.ErrUserInactive) {
			response.Unauthorized(c, "user account is inactive")
			return
		}
		response.InternalError(c, "login failed")
		return
	}

	// Set user ID in context for audit middleware
	c.Set(auth.UserIDKey, result.UserID)

	response.Success(c, result)
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	var req RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	result, err := h.authService.RefreshToken(c.Request.Context(), req.RefreshToken)
	if err != nil {
		if errors.Is(err, services.ErrInvalidRefreshToken) ||
			errors.Is(err, services.ErrExpiredRefreshToken) ||
			errors.Is(err, services.ErrRevokedRefreshToken) {
			response.Unauthorized(c, "invalid or expired refresh token")
			return
		}
		response.InternalError(c, "refresh failed")
		return
	}

	// Set user ID in context for audit middleware
	c.Set(auth.UserIDKey, result.UserID)

	response.Success(c, result)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	var req LogoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	result, err := h.authService.Logout(c.Request.Context(), req.RefreshToken)
	if err != nil {
		response.InternalError(c, "logout failed")
		return
	}

	// Set user ID in context for audit middleware
	if result != nil {
		c.Set(auth.UserIDKey, result.UserID)
	}

	response.Success(c, nil)
}

// GetMyPages returns the current user's accessible pages with permissions
func (h *AuthHandler) GetMyPages(c *gin.Context) {
	userID, ok := auth.GetUserID(c)
	if !ok {
		response.Unauthorized(c, "user not authenticated")
		return
	}

	pages, err := h.authService.GetUserPages(c.Request.Context(), userID)
	if err != nil {
		response.InternalError(c, "failed to get user pages")
		return
	}

	response.Success(c, pages)
}
