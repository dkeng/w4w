package controller

import (
	"net/http"

	"github.com/dkeng/w4w/src/server/router"
	"github.com/dkeng/w4w/src/store"
	"github.com/gin-gonic/gin"
)

// HomeController Home
type HomeController struct {
}

// Init 初始化 HomeController
func (h *HomeController) Init(allStore *store.AllStore) *HomeController {

	return h
}

// GetRouter 获取 HomeController 路由
func (h *HomeController) GetRouter() []*router.R {
	return []*router.R{}
}

// Index 首页
func (h *HomeController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
