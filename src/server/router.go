package server

import (
	"net/http"

	"github.com/dkeng/w4w/src/api/controller"
	"github.com/gin-gonic/gin"
)

// R used for register router
type R struct {
	Method  string             // http request method
	Path    string             // http request path
	Handler func(*gin.Context) // handler function
	Desc    string             // description for the method
}

var (
	apiRoute = []*R{
		&R{
			Method:  http.MethodPost,
			Path:    "/short",
			Handler: controller.AddShort,
			Desc:    "添加链接",
		},
	}
)

// API Api
func API() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func setGroupRouter(rg *gin.RouterGroup, routes []*R) {
	for _, v := range routes {
		rg.Handle(v.Method, v.Path, v.Handler)
	}
}
