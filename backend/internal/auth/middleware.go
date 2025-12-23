package auth

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mustafa91ameen/prjalgo/backend/internal/response"
)

const (
	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer "
	UserIDKey           = "user_id"
	UsernameKey         = "username"
	TargetIDKey         = "target_id"
)

// Context key type for type-safe context values
type contextKey string

// ContextKeyUserID is the key for storing userId in standard context
const ContextKeyUserID contextKey = "userId"

// GetUserIDFromContext extracts userId from standard context (for use in services)
func GetUserIDFromContext(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(ContextKeyUserID).(int64)
	return userID, ok
}

// GetUserIDPtrFromContext extracts userId as pointer from standard context
func GetUserIDPtrFromContext(ctx context.Context) *int64 {
	if userID, ok := ctx.Value(ContextKeyUserID).(int64); ok {
		return &userID
	}
	return nil
}

// PermissionChecker interface for checking user permissions
type PermissionChecker interface {
	HasPermission(ctx context.Context, userID int64, route string, permission string) (bool, error)
}

func JWTMiddleware(jwtManager *JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader == "" {
			response.Unauthorized(c, "authorization header required")
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, BearerPrefix) {
			response.Unauthorized(c, "invalid authorization format")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, BearerPrefix)
		claims, err := jwtManager.ValidateToken(tokenString)
		if err != nil {
			response.Unauthorized(c, err.Error())
			c.Abort()
			return
		}

		c.Set(UserIDKey, claims.UserID)
		c.Set(UsernameKey, claims.Username)

		// Add userID to standard context so it's accessible from services
		ctx := context.WithValue(c.Request.Context(), ContextKeyUserID, claims.UserID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// AuthorizationMiddleware checks if user has permission to access a route
func AuthorizationMiddleware(checker PermissionChecker, route string, permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get(UserIDKey)
		if !exists {
			response.Unauthorized(c, "user not authenticated")
			c.Abort()
			return
		}

		hasPermission, err := checker.HasPermission(c.Request.Context(), userID.(int64), route, permission)
		if err != nil {
			response.InternalError(c, "failed to check permissions")
			c.Abort()
			return
		}

		if !hasPermission {
			response.Error(c, http.StatusForbidden, "access denied")
			c.Abort()
			return
		}

		c.Next()
	}
}

// GetUserID helper to get user ID from context
func GetUserID(c *gin.Context) (int64, bool) {
	userID, exists := c.Get(UserIDKey)
	if !exists {
		return 0, false
	}
	return userID.(int64), true
}

// AuditLogger interface for audit logging
type AuditLogger interface {
	Log(ctx context.Context, actorID *int64, action, targetType string, targetID *int64, method, status, ipAddress string)
}

// AuditMiddleware logs actions to audit log
func AuditMiddleware(auditService AuditLogger, action, targetType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		status := "success"
		if c.Writer.Status() >= 400 {
			status = "failure"
		}

		var actorID *int64
		if userID, exists := c.Get(UserIDKey); exists {
			id := userID.(int64)
			actorID = &id
		}

		var targetID *int64
		// First check if handler set target ID in context (for create actions)
		if tid, exists := c.Get(TargetIDKey); exists {
			if id, ok := tid.(int64); ok {
				targetID = &id
			}
		} else if idStr := c.Param("id"); idStr != "" {
			// Fall back to URL param for update/delete actions
			if id, err := strconv.ParseInt(idStr, 10, 64); err == nil {
				targetID = &id
			}
		}

		auditService.Log(
			c.Request.Context(),
			actorID,
			action,
			targetType,
			targetID,
			c.Request.Method,
			status,
			c.ClientIP(),
		)
	}
}

// AuditAuthMiddleware for login/logout/refresh
// After handler runs, it tries to get user ID from context (set by auth handler on success)
func AuditAuthMiddleware(auditService AuditLogger, action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		status := "success"
		if c.Writer.Status() >= 400 {
			status = "failure"
		}

		// Try to get user ID from context (set by auth handler after successful login/logout)
		var actorID *int64
		var targetID *int64
		if userID, exists := c.Get(UserIDKey); exists {
			id := userID.(int64)
			actorID = &id
			targetID = &id // For auth actions, the target is the user themselves
		}

		auditService.Log(
			c.Request.Context(),
			actorID,
			action,
			"auth",
			targetID,
			c.Request.Method,
			status,
			c.ClientIP(),
		)
	}
}
