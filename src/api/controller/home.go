package controller

import (
	"net/http"

	"github.com/dkeng/w4w/src/server/router"
	"github.com/dkeng/w4w/src/store"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// HomeController Home
type HomeController struct {
	linkStore           store.LinkStore
	redirectRecordStore store.RedirectRecordStore
}

// Init 初始化 HomeController
func (h *HomeController) Init(allStore *store.AllStore) *HomeController {
	h.linkStore = allStore.LinkStore
	h.redirectRecordStore = allStore.RedirectRecordStore
	return h
}

// GetRouter 获取 HomeController 路由
func (h *HomeController) GetRouter() []*router.R {
	return []*router.R{}
}

// Index 首页
func (h *HomeController) Index(c *gin.Context) {
	linkRanks := h.redirectRecordStore.RankTop100()
	var result []map[string]interface{}
	for _, v := range linkRanks {
		linkRank := h.linkStore.QueryByID(v.LinkID)
		line := structs.Map(linkRank)
		line["count"] = v.Count
		result = append(result, line)
	}
	c.HTML(http.StatusOK, "index.html", result)
}
