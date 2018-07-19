package router

import (
	"github.com/gin-gonic/gin"
)

// R used for register router
type R struct {
	Method  string             // http request method
	Path    string             // http request path
	Handler func(*gin.Context) // handler function
	Desc    string             // description for the method
}

// API Api
func API() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// SetGroupRouter 设置分组路由
func SetGroupRouter(rg *gin.RouterGroup, r Router) {
	for _, v := range r.GetRouter() {
		rg.Handle(v.Method, v.Path, v.Handler)
	}
}

// Router 路由接口
type Router interface {
	GetRouter() []*R
}
