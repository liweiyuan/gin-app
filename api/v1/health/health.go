package health

import (
	"runtime"
	"time"

	"github.com/gin-gonic/gin"

	"gin-app/responses"
)

// Info holds information about the application status
type Info struct {
	Status    string    `json:"status"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
	GoVersion string    `json:"go_version"`
	Memory    Memory    `json:"memory"`
	Uptime    string    `json:"uptime"`
}

// Memory represents runtime memory stats
type Memory struct {
	Alloc      uint64 `json:"alloc"`
	TotalAlloc uint64 `json:"total_alloc"`
	Sys        uint64 `json:"sys"`
	NumGC      uint32 `json:"num_gc"`
}

var startTime = time.Now()

// Constants for application info
const (
	StatusOK = "ok"
	Version  = "1.0.0" // Update this with your app version
)

// RegisterRoutes registers the health check routes
func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/health", Health)
	router.GET("/status", Status)
}

// Health handles the health check endpoint
func Health(c *gin.Context) {
	responses.Success(c, "Service is healthy", nil)
}

// Status provides detailed information about the application status
func Status(c *gin.Context) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	info := Info{
		Status:    StatusOK,
		Version:   Version,
		Timestamp: time.Now(),
		GoVersion: runtime.Version(),
		Memory: Memory{
			Alloc:      m.Alloc,
			TotalAlloc: m.TotalAlloc,
			Sys:        m.Sys,
			NumGC:      m.NumGC,
		},
		Uptime: time.Since(startTime).String(),
	}
	responses.Success(c, "Application status", info)
}
