package router

import (
	"context"
	"fmt"
	"gin-app/config"
	"gin-app/handler"
	"gin-app/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	engine *gin.Engine
}

// newGinRouter creates a new GinRouter instance
func newGinRouter() *GinRouter {
	return &GinRouter{
		engine: gin.New(), // Use gin.New() so you can manually add middlewares in order
	}
}

// register 注册路由处理器
func (r *GinRouter) register(method string, path string, h gin.HandlerFunc) {
	r.engine.Handle(method, path, h)
}

// setup 实现路由设置
func (r *GinRouter) setup() *gin.Engine {
	return r.engine
}

// 注册中间件
func (r *GinRouter) registerMiddleware(middleware gin.HandlerFunc) {
	r.engine.Use(middleware)
}

func register() *gin.Engine {
	r := newGinRouter()
	r.registerMiddleware(handler.RecoveryMiddleware())
	r.registerMiddleware(handler.LoggerMiddleware())

	// 注册路由
	ping := &handler.PingEndpoint{}
	r.register("GET", "/ping", ping.Ping)

	check := &handler.CheckEndpoint{}
	r.register("GET", "/check", check.Check)

	user := &handler.UserEndpoint{}
	r.register("GET", "/user/get", user.GetUser)
	r.register("POST", "/user/create", user.CreateUser)
	r.register("DELETE", "/user/delete", user.DeleteUser)
	r.register("PUT", "/user/update", user.UpdateUser)

	return r.setup()
}

func Serve() {
	router := register()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.GlobalConfig.App.Port),
		Handler: router,
	}
	// 在一个新的goroutine中启动服务器
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Logger.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Logger.Fatal("Server forced to shutdown:", err)
	}
	log.Logger.Info("Server exiting")
}
