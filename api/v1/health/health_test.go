package health

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	router := gin.Default()
	router.GET("/health", Health)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"code\":200,\"message\":\"Service is healthy\"}", w.Body.String())
}

func TestStatus(t *testing.T) {
	router := gin.Default()
	router.GET("/status", Status)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/status", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var actualResponse map[string]interface{}
	json.Unmarshal([]byte(w.Body.String()), &actualResponse)
	delete(actualResponse["data"].(map[string]interface{}), "timestamp")
	assert.Contains(t, w.Body.String(), "code")
	assert.Contains(t, w.Body.String(), "message")
	assert.Contains(t, w.Body.String(), "data")
}
