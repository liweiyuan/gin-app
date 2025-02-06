package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingHandler(t *testing.T) {
	// 创建一个新的 Gin 引擎
	r := gin.Default()

	// 创建一个 PingHandler 实例
	ping := &PingEndpoint{}

	// 注册路由和处理器
	r.GET("/ping", ping.Ping)

	// 创建一个 HTTP 请求
	req, _ := http.NewRequest("GET", "/ping", nil)

	// 创建一个响应记录器
	w := httptest.NewRecorder()

	// 发送请求
	r.ServeHTTP(w, req)

	// 断言响应状态码
	assert.Equal(t, http.StatusOK, w.Code)

	// 断言响应体
	expectedBody := `{"message":"pong"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}
