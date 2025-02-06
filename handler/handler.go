package handler

import (
	"gin-app/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// CheckHandler 实现ping路由处理器
type CheckHandler struct{}

// 实现/ping请求处理
func (h *CheckHandler) Ping(c *gin.Context) {
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

// 实现/check请求处理
func (h *CheckHandler) Check(c *gin.Context) {
	// 记录日志
	log.Logger.WithFields(logrus.Fields{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
		"ip":     c.ClientIP(),
	}).Info("Handling /check request")
	c.JSON(http.StatusOK, gin.H{
		"message": "check",
	})
}
