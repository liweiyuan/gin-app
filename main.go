// main.go
package main

import (
	"gin-app/router"
)

func main() {
	// 启动服务器，默认监听在 0.0.0.0:8080
	router.Serve()
}
