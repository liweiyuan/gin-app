package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// resetUsers 函数用于在每个测试之前重置 users map 到初始状态
func resetUsers() {
	users = map[int]User{
		1: {ID: 1, Name: "张三", Age: 20},
	}
}

func TestCreateUser(t *testing.T) {
	resetUsers() // 在测试前重置状态
	// 创建一个新的 Gin 引擎
	r := gin.Default()

	// 创建一个 UserEndpoint 实例
	userEndpoint := &UserEndpoint{}

	// 注册路由和处理器
	r.POST("/user/create", userEndpoint.CreateUser)

	// 创建一个 HTTP 请求
	req, _ := http.NewRequest("POST", "/user/create", nil)

	// 创建一个响应记录器
	w := httptest.NewRecorder()

	// 发送请求
	r.ServeHTTP(w, req)

	// 断言响应状态码
	assert.Equal(t, http.StatusOK, w.Code)

	// 断言响应体
	expectedBody := `{"status":"success","data":{"id":2,"name":"李四","age":25}}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestDeleteUser(t *testing.T) {
	resetUsers() // 在测试前重置状态

	// 创建一个新的 Gin 引擎
	r := gin.Default()

	// 创建一个 UserEndpoint 实例
	userEndpoint := &UserEndpoint{}

	// 注册路由和处理器
	r.DELETE("/user/delete", userEndpoint.DeleteUser)

	// 创建一个 HTTP 请求
	req, _ := http.NewRequest("DELETE", "/user/delete", nil)

	// 创建一个响应记录器
	w := httptest.NewRecorder()

	// 发送请求
	r.ServeHTTP(w, req)

	// 断言响应状态码
	assert.Equal(t, http.StatusOK, w.Code)

	// 断言响应体
	expectedBody := `{"status":"success","message":"user deleted"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestUpdateUser(t *testing.T) {
	resetUsers() // 在测试前重置状态

	// 创建一个新的 Gin 引擎
	r := gin.Default()

	// 创建一个 UserEndpoint 实例
	userEndpoint := &UserEndpoint{}

	// 注册路由和处理器
	r.PUT("/user/update", userEndpoint.UpdateUser)

	// 创建一个 HTTP 请求
	req, _ := http.NewRequest("PUT", "/user/update", nil)

	// 创建一个响应记录器
	w := httptest.NewRecorder()

	// 发送请求
	r.ServeHTTP(w, req)

	// 断言响应状态码
	assert.Equal(t, http.StatusOK, w.Code)

	// 断言响应体
	expectedBody := `{"status":"success","data":{"id":1,"name":"张三","age":30}}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestGetUser(t *testing.T) {
	resetUsers() // 在测试前重置状态
	// 创建一个新的 Gin 引擎
	r := gin.Default()

	// 创建一个 UserEndpoint 实例
	userEndpoint := &UserEndpoint{}

	// 注册路由和处理器
	r.GET("/user/get", userEndpoint.GetUser)

	// 创建一个 HTTP 请求
	req, _ := http.NewRequest("GET", "/user/get", nil)

	// 创建一个响应记录器
	w := httptest.NewRecorder()

	// 发送请求
	r.ServeHTTP(w, req)

	// 断言响应状态码
	assert.Equal(t, http.StatusOK, w.Code)

	// 断言响应体
	expectedBody := `{"status":"success","data":{"id":1,"name":"张三","age":20}}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}
