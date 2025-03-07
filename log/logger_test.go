package log

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLoggerInit(t *testing.T) {
	// Assuming the init function sets up the logger
	Init()
	// Add assertions to verify the logger is initialized correctly
	assert.Equal(t, logrus.InfoLevel, Logger.GetLevel())
}
