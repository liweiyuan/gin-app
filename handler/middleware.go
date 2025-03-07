package handler

import (
	"context"
	stderrors "errors"
	"gin-app/errors"
	"gin-app/log"
	"gin-app/responses"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ErrorHandlerMiddleware catches AppError types and returns standardized responses
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if there are any errors
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				// Check if it's our custom AppError type
				var appErr *errors.AppError
				if stderrors.As(e.Err, &appErr) {
					// Use the status code and message from our custom error
					responses.Error(c, appErr.StatusCode, appErr.Error(), nil)
					return
				}
			}
			// If it's not our custom error, return a generic 500 error
			responses.InternalServerError(c, http.StatusText(http.StatusInternalServerError), nil)
		}
	}
}

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// Record detailed logs
				log.Logger.WithFields(logrus.Fields{
					"path":       c.Request.URL.Path,
					"method":     c.Request.Method,
					"error":      r,
					"client_ip":  c.ClientIP(),
					"user_agent": c.Request.UserAgent(),
				}).Error("panic recovered")

				// Use our standardized response
				responses.InternalServerError(c, http.StatusText(http.StatusInternalServerError), nil)
				c.Abort()
			}
		}()
		c.Next()
	}
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()

		// Process request
		c.Next()

		// Calculate request duration
		duration := time.Since(start)
		statusCode := c.Writer.Status()

		// Create structured log entry with common fields
		logEntry := log.Logger.WithFields(logrus.Fields{
			"status":     statusCode,
			"duration":   duration.String(),
			"path":       path,
			"method":     method,
			"client_ip":  clientIP,
			"user_agent": userAgent,
			"request_id": c.GetHeader("X-Request-ID"),
		})

		// Log based on status code
		if len(c.Errors) > 0 {
			// Log request with error details
			logEntry.WithField("errors", c.Errors.String()).Error("Request failed with errors")
		} else if statusCode >= http.StatusBadRequest && statusCode < http.StatusInternalServerError {
			// Client errors (4xx)
			logEntry.Warn("Client error")
		} else if statusCode >= http.StatusInternalServerError {
			// Server errors (5xx)
			logEntry.Error("Server error")
		} else {
			// Success (2xx) and redirection (3xx)
			logEntry.Info("Request processed successfully")
		}
	}
}

// CORSMiddleware sets up Cross-Origin Resource Sharing
func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins, should be restricted in production
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

// TimeoutMiddleware aborts requests that take too long to process
func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a context with timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// Update the request with the timeout context
		c.Request = c.Request.WithContext(ctx)

		// Create a channel to track when the request is complete
		done := make(chan struct{})

		go func() {
			c.Next()
			close(done)
		}()

		select {
		case <-done:
			// Request completed before timeout
			return
		case <-ctx.Done():
			// Request timed out
			if ctx.Err() == context.DeadlineExceeded {
				log.Logger.WithFields(logrus.Fields{
					"path":      c.Request.URL.Path,
					"method":    c.Request.Method,
					"client_ip": c.ClientIP(),
					"timeout":   timeout.String(),
				}).Warn("Request timed out")

				responses.RequestTimeout(c, "Request timeout", nil) // Explicitly passing nil as data parameter
				c.Abort()
			}
		}
	}
}
