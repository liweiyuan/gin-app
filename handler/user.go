// UserEndpoint 实现user路由处理器, 实现增删改查

package handler

import (
	"gin-app/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserEndpoint struct{}

// User 模拟的用户结构体
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 模拟的用户数据库
var users = map[int]User{
	1: {ID: 1, Name: "张三", Age: 20},
}

// logRequest 记录请求日志
func logRequest(c *gin.Context, action string) {
	log.Logger.WithFields(logrus.Fields{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
		"ip":     c.ClientIP(),
	}).Info(action)
}

func (h *UserEndpoint) CreateUser(c *gin.Context) {
	logRequest(c, "Handling /user/create request")

	// 模拟创建用户
	newUser := User{ID: 2, Name: "李四", Age: 25}
	users[newUser.ID] = newUser

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   newUser,
	})
}

func (h *UserEndpoint) DeleteUser(c *gin.Context) {
	logRequest(c, "Handling /user/delete request")

	// 模拟删除用户
	delete(users, 1)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "user deleted",
	})
}

func (h *UserEndpoint) UpdateUser(c *gin.Context) {
	logRequest(c, "Handling /user/update request")

	// 模拟更新用户
	if user, exists := users[1]; exists {
		user.Age = 30
		users[user.ID] = user
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   user,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "user not found",
		})
	}
}

func (h *UserEndpoint) GetUser(c *gin.Context) {
	logRequest(c, "Handling /user/get request")

	// 模拟获取用户
	if user, exists := users[1]; exists {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   user,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "user not found",
		})
	}
}
