package router

import (
	"context"
	"fmt"
	"gin-app/api/v1/health"
	"gin-app/api/v1/user"
	"gin-app/config"
	"gin-app/handler"
	"gin-app/log"
	"gin-app/models"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Method string
	Path   string
}

type GinRouter struct {
	routes      []Route
	middlewares []gin.HandlerFunc
	engine      *gin.Engine
}

// NewGinRouter 创建GinRouter实例
func NewGinRouter() *GinRouter {
	return &GinRouter{
		engine: gin.New(), // 使用gin.New()以便手动添加中间件
	}
}

// register 注册路由处理器
func (r *GinRouter) register(method string, path string, h gin.HandlerFunc) {
	r.engine.Handle(method, path, h)
	r.routes = append(r.routes, Route{Method: method, Path: path})
}

// setup 实现路由设置
func (r *GinRouter) setup() *gin.Engine {
	return r.engine
}

// 注册中间件
func (r *GinRouter) registerMiddleware(middleware gin.HandlerFunc) {
	r.engine.Use(middleware)
	r.middlewares = append(r.middlewares, middleware)
}

// register 配置并返回Gin引擎实例
func Register() *gin.Engine {
	r := NewGinRouter()

	// 注册全局中间件
	// 顺序很重要 - 请求首先经过Logger、CORS，然后是超时检测，最后是错误处理和恢复
	r.registerMiddleware(handler.LoggerMiddleware())                  // 记录请求日志
	r.registerMiddleware(handler.CORSMiddleware())                    // 处理跨域请求
	r.registerMiddleware(handler.TimeoutMiddleware(10 * time.Second)) // 10秒请求超时
	r.registerMiddleware(handler.ErrorHandlerMiddleware())            // 统一错误处理
	r.registerMiddleware(handler.RecoveryMiddleware())                // 从panic中恢复

	// 创建处理器
	userRepo := models.NewInMemoryUserRepository()
	userHandler := user.NewUserHandler(userRepo)

	// 注册API路由 - v1版本 (RESTful API设计)
	// 所有新功能应该添加在此版本下，保持向后兼容性
	v1 := r.engine.Group("/api/v1")
	{
		// 健康检查和系统状态
		health.RegisterRoutes(v1)

		// 用户管理 - RESTful设计
		userHandler.RegisterRoutes(v1)
	}

	// 兼容旧版API（保持向后兼容性）
	// 这些路由将继续支持，但新客户端应使用v1 API
	// 这些路由将继续支持，但新客户端应使用v1 API
	// 未来版本可能会考虑弃用这些端点
	r.register("GET", "/ping", func(c *gin.Context) { health.Health(c) })
	r.register("GET", "/status", func(c *gin.Context) { health.Status(c) })

	// Legacy user routes
	r.register("GET", "/user/:id", userHandler.GetUser)       // 旧版本
	r.register("POST", "/user", userHandler.CreateUser)       // 旧版本
	r.register("DELETE", "/user/:id", userHandler.DeleteUser) // 旧版本
	r.register("PUT", "/user/:id", userHandler.UpdateUser)    // 旧版本
	return r.setup()
}

// Serve 启动HTTP服务器并处理优雅关闭
func Serve() {
	router := Register()

	// 配置HTTP服务器
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.GlobalConfig.App.Port),
		Handler: router,
		// 添加合理的超时设置
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// 在一个新的goroutine中启动服务器
	go func() {
		appName := config.GlobalConfig.App.Name
		appVersion := config.GlobalConfig.App.Version
		appPort := config.GlobalConfig.App.Port

		log.Logger.Infof("%s v%s starting on port %d", appName, appVersion, appPort)
		log.Logger.Infof("API v1 available at: http://localhost:%d/api/v1", appPort)

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
