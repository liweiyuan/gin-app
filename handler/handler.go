package handler

import (
	"gin-app/log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// PingHandler 实现ping路由处理器
type PingHandler struct{}

// Handle 实现ping请求处理
func (h *PingHandler) Handle(c *gin.Context) {
	// 记录日志
	log.Logger.WithFields(logrus.Fields{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
	}).Info("request")

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
