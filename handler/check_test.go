package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)                     // 设置测试模式
	recorder := httptest.NewRecorder()            // 创建一个记录器
	context, _ := gin.CreateTestContext(recorder) // 创建一个测试上下文

	// 创建一个请求
	request, _ := http.NewRequest("GET", "/check", nil)
	context.Request = request

	//执行测试
	handler := &CheckEndpoint{}
	handler.Check(context)

	// 断言验证
	assert.Equal(t, http.StatusOK, recorder.Code)                   // 验证状态码是否为200
	assert.JSONEq(t, `{"message":"check"}`, recorder.Body.String()) // 验证返回的JSON是否正确

}
