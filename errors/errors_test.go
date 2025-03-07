package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAppError(t *testing.T) {
	err := NewAppError(400, "test error", nil)
	assert.Equal(t, 400, err.StatusCode)
	assert.Equal(t, "test error", err.Message)
	assert.Nil(t, err.Details)
}

func TestNotFound(t *testing.T) {
	err := NotFound("test not found", nil)
	assert.Equal(t, 404, err.StatusCode)
	assert.Equal(t, "test not found", err.Message)
	assert.Nil(t, err.Details)
}

func TestBadRequest(t *testing.T) {
	err := BadRequest("test bad request", nil)
	assert.Equal(t, 400, err.StatusCode)
	assert.Equal(t, "test bad request", err.Message)
	assert.Nil(t, err.Details)
}

func TestInternal(t *testing.T) {
	err := Internal("test internal error", nil)
	assert.Equal(t, 500, err.StatusCode)
	assert.Equal(t, "test internal error", err.Message)
	assert.Nil(t, err.Details)
}

func TestUnauthorized(t *testing.T) {
	err := Unauthorized("test unauthorized", nil)
	assert.Equal(t, 401, err.StatusCode)
	assert.Equal(t, "test unauthorized", err.Message)
	assert.Nil(t, err.Details)
}

func TestValidationError(t *testing.T) {
	err := ValidationError("field", "test validation error")
	assert.Equal(t, 400, err.StatusCode)
	assert.Equal(t, "Validation error on field 'field'", err.Message)
	assert.Equal(t, map[string]string{"field": "test validation error"}, err.Details)
}
