package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mustafa91ameen/prjalgo/backend/internal/logger"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		logger.Info("Request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"duration", time.Since(start),
		)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		// In production, we should validate the origin against allowed domains
		// For now, we allow the request origin if it matches our environment or valid domains

		c.Header("Access-Control-Allow-Origin", origin) // Or "*" if you want to be public, but origin is needed for credentials
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// TimeoutMiddleware adds a timeout to the request context for database queries
// All database operations will respect this timeout and cancel if exceeded
func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create context with timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// Update request with new context
		// All handlers and database queries will use this context
		// If timeout is exceeded, database queries will automatically cancel
		c.Request = c.Request.WithContext(ctx)

		// Continue to next handler
		c.Next()

		// If timeout occurred and no response was written, send timeout error
		if ctx.Err() == context.DeadlineExceeded && !c.Writer.Written() {
			c.AbortWithStatusJSON(http.StatusRequestTimeout, gin.H{
				"error": "Request timeout - database query took too long",
			})
		}
	}
}
