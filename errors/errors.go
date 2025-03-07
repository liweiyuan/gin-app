package errors

import (
	"fmt"
	"net/http"
)

// AppError 是应用程序的自定义错误类型
type AppError struct {
	StatusCode int               // HTTP状态码
	Message    string            // 错误消息
	Details    map[string]string // 可选的详细信息
}

// Error 实现错误接口
func (e *AppError) Error() string {
	return e.Message
}

// 创建新的应用程序错误
func NewAppError(statusCode int, message string, details map[string]string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Message:    message,
		Details:    details,
	}
}

// NotFound 创建404错误
func NotFound(message string, details map[string]string) *AppError {
	if message == "" {
		message = "Resource not found"
	}
	return NewAppError(http.StatusNotFound, message, details)
}

// BadRequest 创建400错误
func BadRequest(message string, details map[string]string) *AppError {
	if message == "" {
		message = "Bad request"
	}
	return NewAppError(http.StatusBadRequest, message, details)
}

// Internal 创建500错误
func Internal(message string, details map[string]string) *AppError {
	if message == "" {
		message = "Internal server error"
	}
	return NewAppError(http.StatusInternalServerError, message, details)
}

// Unauthorized 创建401错误
func Unauthorized(message string, details map[string]string) *AppError {
	if message == "" {
		message = "Unauthorized"
	}
	return NewAppError(http.StatusUnauthorized, message, details)
}

// ValidationError 创建带有字段验证错误的400错误
func ValidationError(field, message string) *AppError {
	details := map[string]string{
		field: message,
	}
	return BadRequest(fmt.Sprintf("Validation error on field '%s'", field), details)
}
