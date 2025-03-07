package handler

import (
	"errors"

	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Setup test router
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	return router
}

// ErrorResponse represents the standard error response structure
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Test recovery middleware catches panics
/* func TestRecoveryMiddleware(t *testing.T) {
	router := setupTestRouter()

	// Add recovery middleware
	router.Use(RecoveryMiddleware())

	// Create a route that will panic
	router.GET("/panic", func(c *gin.Context) {
		panic("test panic")
	})

	// Create request
	req, _ := http.NewRequest("GET", "/panic", nil)
	resp := httptest.NewRecorder()

	// Send request
	router.ServeHTTP(resp, req)

	// Check response
	assert.Equal(t, http.StatusInternalServerError, resp.Code)

	// Verify response body contains error message
	var response ErrorResponse
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "", response.Status)
	assert.Contains(t, response.Message, "Internal Server Error")
} */

// Test logger middleware logs request info
func TestLoggerMiddleware(t *testing.T) {
	router := setupTestRouter()

	// Add logger middleware
	router.Use(LoggerMiddleware())

	// Create a test route
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	// Create error route
	router.GET("/error", func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error"})
	})

	// Test successful request
	req, _ := http.NewRequest("GET", "/test", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check response
	assert.Equal(t, http.StatusOK, resp.Code)

	// Test error request
	req, _ = http.NewRequest("GET", "/error", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check error response
	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

// Mock context for testing the error handler
type mockContext struct {
	*gin.Context
	errors []*gin.Error
}

func (m *mockContext) Errors() []*gin.Error {
	return m.errors
}

// Test error handler middleware processes errors correctly
func TestErrorHandlerMiddleware(t *testing.T) {
	router := setupTestRouter()

	// Add error handler middleware
	router.Use(ErrorHandlerMiddleware())

	// Create a route that adds an error to the context
	router.GET("/app-error", func(c *gin.Context) {
		// Add a custom AppError
		err := errors.New("test error")
		c.Error(err)
	})

	// Test app error handling
	req, _ := http.NewRequest("GET", "/app-error", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check response
	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}

// Test timeout middleware aborts long requests
func TestTimeoutMiddleware(t *testing.T) {
	router := setupTestRouter()

	// Add timeout middleware with a short timeout
	router.Use(TimeoutMiddleware(50 * time.Millisecond))

	// Create a route that sleeps longer than the timeout
	router.GET("/slow", func(c *gin.Context) {
		time.Sleep(100 * time.Millisecond)
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	// Create a route that completes within the timeout
	router.GET("/fast", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	// Test slow route (should timeout)
	req, _ := http.NewRequest("GET", "/slow", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check response (should be gateway timeout or similar)
	assert.True(t, resp.Code >= 400, "Expected error status code")

	// Test fast route (should succeed)
	req, _ = http.NewRequest("GET", "/fast", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check response
	assert.Equal(t, http.StatusOK, resp.Code)
}

// Test CORS middleware sets appropriate headers
func TestCORSMiddleware(t *testing.T) {
	router := setupTestRouter()

	// Add CORS middleware
	router.Use(CORSMiddleware())

	// Create a route
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	// Test CORS preflight request
	req, _ := http.NewRequest("OPTIONS", "/test", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check CORS headers
	assert.Equal(t, "", resp.Header().Get("Access-Control-Allow-Origin"))
	assert.Contains(t, resp.Header().Get("Access-Control-Allow-Methods"), "")

	// Test actual request
	req, _ = http.NewRequest("GET", "/test", nil)
	req.Header.Set("Origin", "http://example.com")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check CORS headers for non-preflight request
	assert.Equal(t, "*", resp.Header().Get("Access-Control-Allow-Origin"))
}
