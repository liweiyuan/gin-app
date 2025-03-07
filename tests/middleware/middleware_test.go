package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"gin-app/handler"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	router := gin.Default()
	router.Use(handler.ErrorHandlerMiddleware())
	router.Use(handler.RecoveryMiddleware())
	router.Use(handler.LoggerMiddleware())
	router.Use(handler.CORSMiddleware())
	router.Use(handler.TimeoutMiddleware(5 * time.Second))
	return router
}

type ErrorResponse struct {
	Error string `json:"error"`
}

/* func TestRecoveryMiddleware(t *testing.T) {
	router := setupTestRouter()
	router.GET("/panic", func(c *gin.Context) {
		panic("test panic")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/panic", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Contains(t, w.Body.String(), "test panic")
} */

func TestLoggerMiddleware(t *testing.T) {
	router := setupTestRouter()
	router.GET("/log", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/log", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "test", w.Body.String())
}

func TestErrorHandlerMiddleware(t *testing.T) {
	router := setupTestRouter()
	router.GET("/error", func(c *gin.Context) {
		c.Error(errors.New("test error"))
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/error", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Internal Server Error")
}

func TestTimeoutMiddleware(t *testing.T) {
	router := setupTestRouter()
	router.GET("/timeout", func(c *gin.Context) {
		time.Sleep(6 * time.Second)
		c.String(http.StatusOK, "test")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/timeout", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusRequestTimeout, w.Code)
	assert.Contains(t, w.Body.String(), "Request timeout")
}

func TestCORSMiddleware(t *testing.T) {
	router := setupTestRouter()
	router.GET("/cors", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/cors", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "test", w.Body.String())
	assert.Equal(t, "", w.Header().Get("Access-Control-Allow-Origin"))
}
