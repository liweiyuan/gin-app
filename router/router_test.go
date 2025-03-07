package router

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewGinRouter(t *testing.T) {
	router := NewGinRouter()
	assert.NotNil(t, router)
	assert.IsType(t, &GinRouter{}, router)
}

func TestGinRouterRegister(t *testing.T) {
	router := NewGinRouter()
	router.register("GET", "/test", func(c *gin.Context) {
		c.String(200, "test")
	})

	assert.Equal(t, 1, len(router.routes))
	assert.Equal(t, "GET", router.routes[0].Method)
	assert.Equal(t, "/test", router.routes[0].Path)
}

func TestGinRouterSetup(t *testing.T) {
	router := NewGinRouter()
	engine := router.setup()
	assert.NotNil(t, engine)
	assert.IsType(t, &gin.Engine{}, engine)
}

func TestGinRouterRegisterMiddleware(t *testing.T) {
	router := NewGinRouter()
	router.registerMiddleware(func(c *gin.Context) {
		c.Next()
	})

	assert.Equal(t, 1, len(router.middlewares))
}
