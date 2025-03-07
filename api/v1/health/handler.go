package health

import (
	"gin-app/config"
	"gin-app/responses"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

// Handler 健康检查处理器
type Handler struct {
	startTime time.Time
}

// NewHandler 创建健康检查处理器
func NewHandler() *Handler {
	return &Handler{
		startTime: time.Now(),
	}
}

// CheckResponse 健康检查响应
type CheckResponse struct {
	Status    string    `json:"status"`
	AppName   string    `json:"app_name"`
	Version   string    `json:"version"`
	Uptime    string    `json:"uptime"`
	GoVersion string    `json:"go_version"`
	Timestamp time.Time `json:"timestamp"`
}

// Check 健康检查
func (h *Handler) Check(c *gin.Context) {
	response := CheckResponse{
		Status:    "ok",
		AppName:   config.GlobalConfig.App.Name,
		Version:   config.GlobalConfig.App.Version,
		Uptime:    time.Since(h.startTime).String(),
		GoVersion: runtime.Version(),
		Timestamp: time.Now(),
	}

	responses.Success(c, "Health check successful", response)
}

// Ping 简单的ping测试
func (h *Handler) Ping(c *gin.Context) {
	responses.Success(c, "Ping successful", map[string]string{"message": "pong"})
}
