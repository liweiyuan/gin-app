package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents a standardized API response structure
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success sends a successful response with 200 status code
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}

// Created sends a successful response with 201 status code
func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Code:    http.StatusCreated,
		Message: message,
		Data:    data,
	})
}

// NoContent sends a successful response with 204 status code
func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// BadRequest sends an error response with 400 status code
func BadRequest(c *gin.Context, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	}
	
	c.JSON(http.StatusBadRequest, Response{
		Code:    http.StatusBadRequest,
		Message: message,
		Data:    responseData,
	})
}

// Unauthorized sends an error response with 401 status code
func Unauthorized(c *gin.Context, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	}
	
	c.JSON(http.StatusUnauthorized, Response{
		Code:    http.StatusUnauthorized,
		Message: message,
		Data:    responseData,
	})
}

// Forbidden sends an error response with 403 status code
func Forbidden(c *gin.Context, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	}
	
	c.JSON(http.StatusForbidden, Response{
		Code:    http.StatusForbidden,
		Message: message,
		Data:    responseData,
	})
}

// NotFound sends an error response with 404 status code
func NotFound(c *gin.Context, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	}
	
	c.JSON(http.StatusNotFound, Response{
		Code:    http.StatusNotFound,
		Message: message,
		Data:    responseData,
	})
}

// RequestTimeout sends an error response with 408 status code
func RequestTimeout(c *gin.Context, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	}
	
	c.JSON(http.StatusRequestTimeout, Response{
		Code:    http.StatusRequestTimeout,
		Message: message,
		Data:    responseData,
	})
}

// Conflict sends an error response with 409 status code
func Conflict(c *gin.Context, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	}
	
	c.JSON(http.StatusConflict, Response{
		Code:    http.StatusConflict,
		Message: message,
		Data:    responseData,
	})
}

// UnprocessableEntity sends an error response with 422 status code
func UnprocessableEntity(c *gin.Context, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	}
	
	c.JSON(http.StatusUnprocessableEntity, Response{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
		Data:    responseData,
	})
}

// InternalServerError sends an error response with 500 status code
func InternalServerError(c *gin.Context, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	}
	
	c.JSON(http.StatusInternalServerError, Response{
		Code:    http.StatusInternalServerError,
		Message: message,
		Data:    responseData,
	})
}

// ServiceUnavailable sends an error response with 503 status code
func ServiceUnavailable(c *gin.Context, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	}
	
	c.JSON(http.StatusServiceUnavailable, Response{
		Code:    http.StatusServiceUnavailable,
		Message: message,
		Data:    responseData,
	})
}

// Error is a generic function that sends an error response with the specified status code
func Error(c *gin.Context, statusCode int, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	}
	
	c.JSON(statusCode, Response{
		Code:    statusCode,
		Message: message,
		Data:    responseData,
	})
}

// WithStatusCode sends a response with a custom status code
func WithStatusCode(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Code:    statusCode,
		Message: message,
		Data:    data,
	})
}

// Helper functions for ResponseWithData and ResponseWithoutData

// ResponseWithData creates a response with the specified code, message, and data
func ResponseWithData(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// ResponseWithoutData creates a response with just the specified code and message
func ResponseWithoutData(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}

// MethodNotAllowed returns a method not allowed error response (HTTP 405)
func MethodNotAllowed(c *gin.Context, message string) {
	if message == "" {
		message = "Method not allowed"
	}
	ResponseWithoutData(c, http.StatusMethodNotAllowed, message)
}
