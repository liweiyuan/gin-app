package handler

import (
	"gin-app/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// PingEndpoint 实现ping路由处理器
type PingEndpoint struct{}

// 实现/ping请求处理
func (h *PingEndpoint) Ping(c *gin.Context) {
	// 记录日志
	log.Logger.WithFields(logrus.Fields{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
		"ip":     c.ClientIP(),
	}).Info("Handling /ping request")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
