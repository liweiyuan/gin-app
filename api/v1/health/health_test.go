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
	expectedResponse := `{"code":200,"message":"Application status","data":{"status":"ok","version":"1.0.0","go_version":"go1.24.0","memory":{"alloc":473608,"total_alloc":473608,"sys":8344840,"num_gc":0},"uptime":"2.3535ms"}}`
	var actualResponse map[string]interface{}
	json.Unmarshal([]byte(w.Body.String()), &actualResponse)
	delete(actualResponse["data"].(map[string]interface{}), "timestamp")
	assert.Equal(t, expectedResponse, w.Body.String())
}
