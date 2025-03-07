package handler

import (
	"gin-app/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				//记录日志
				log.Logger.WithFields(logrus.Fields{
					"path":   c.Request.URL.Path,
					"method": c.Request.Method,
					"error":  r,
				}).Error("panic")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				c.Abort()
			}
		}()
		c.Next()
	}
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // 处理请求

		// 检查是否有错误
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				log.Logger.WithFields(logrus.Fields{
					"status": c.Writer.Status(),
					"error":  e.Error(),
					"path":   c.Request.URL.Path,
					"method": c.Request.Method,
				}).Error("HTTP error occurred")
			}
		} else {
			status := c.Writer.Status()
			if status >= http.StatusBadRequest {
				// Log non-2xx responses
				log.Logger.WithFields(logrus.Fields{
					"status": status,
					"path":   c.Request.URL.Path,
					"method": c.Request.Method,
				}).Warn("Non-2xx response")
			}
		}
	}
}
