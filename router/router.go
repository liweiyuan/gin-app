package router

import (
	"fmt"
	"gin-app/config"
	"gin-app/handler"

	"github.com/gin-gonic/gin"
)

// GinRouter
type GinRouter struct {
	engine *gin.Engine
}

// newGinRouter 创建新的GinRouter实例
func newGinRouter() *GinRouter {
	return &GinRouter{
		engine: gin.Default(),
	}
}

// register 注册路由处理器
func (r *GinRouter) register(method string, path string, h handler.Handler) {
	r.engine.Handle(method, path, h.Handle)
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
	r.register("GET", "/ping", &handler.PingHandler{})
	return r.setup()
}

func Serve() {
	router := register()
	router.Run(fmt.Sprintf(":%d", config.GlobalConfig.App.Port))
}
