package responses

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestSuccessResponse tests the Success function
func TestSuccessResponse(t *testing.T) {
	// Setup Gin context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Test data
	message := "Operation successful"
	data := map[string]string{"key": "value"}

	// Call function
	Success(c, message, data)

	// Verify response
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	assert.Equal(t, float64(http.StatusOK), response["code"])
	assert.Equal(t, message, response["message"])
	assert.Equal(t, "value", response["data"].(map[string]interface{})["key"])
}

// TestCreatedResponse tests the Created function
func TestCreatedResponse(t *testing.T) {
	// Setup
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Test data
	message := "Resource created"
	data := map[string]int{"id": 123}

	// Call function
	Created(c, message, data)

	// Verify response
	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	assert.Equal(t, float64(http.StatusCreated), response["code"])
	assert.Equal(t, message, response["message"])
	assert.Equal(t, float64(123), response["data"].(map[string]interface{})["id"])
}

// TestNoContentResponse tests the NoContent function
func TestNoContentResponse(t *testing.T) {
	// Setup
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Call function
	NoContent(c)

	// Verify response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Empty(t, w.Body.String())
}

// TestErrorResponse tests the Error function
func TestErrorResponse(t *testing.T) {
	// Setup
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Test data
	statusCode := http.StatusBadRequest
	message := "Invalid input"
	data := map[string]string{"field": "username", "error": "required"}

	// Call function
	Error(c, statusCode, message, data)

	// Verify response
	assert.Equal(t, statusCode, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	assert.Equal(t, float64(statusCode), response["code"])
	assert.Equal(t, message, response["message"])

	respData := response["data"].(map[string]interface{})
	assert.Equal(t, "username", respData["field"])
	assert.Equal(t, "required", respData["error"])
}

// TestBadRequestResponse tests the BadRequest function
func TestBadRequestResponse(t *testing.T) {
	// Setup
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Call function
	BadRequest(c, "Invalid input")

	// Verify response
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	assert.Equal(t, float64(http.StatusBadRequest), response["code"])
	assert.Equal(t, "Invalid input", response["message"])
}

// TestInternalServerErrorResponse tests the InternalServerError function
func TestInternalServerErrorResponse(t *testing.T) {
	// Setup
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Call function
	InternalServerError(c, "Server error occurred")

	// Verify response
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	assert.Equal(t, float64(http.StatusInternalServerError), response["code"])
	assert.Equal(t, "Server error occurred", response["message"])
}

// TestRequestTimeoutResponse tests the RequestTimeout function
func TestRequestTimeoutResponse(t *testing.T) {
	// Setup
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Call function
	RequestTimeout(c, "Request timed out")

	// Verify response
	assert.Equal(t, http.StatusRequestTimeout, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	assert.Equal(t, float64(http.StatusRequestTimeout), response["code"])
	assert.Equal(t, "Request timed out", response["message"])
}
