package controller

import (
	"fmt"
	"net/http"

	"github.com/dkeng/pkg/time"
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
	s, e := time.TodayStartEndTime()
	rankList := h.redirectRecordStore.RankByStartTimeAndEndTime(s, e)
	var ranks = make(map[int64]interface{}, len(rankList))
	var ids = make([]int64, len(rankList))
	for i, v := range rankList {
		ids[i] = v["link_id"].(int64)
		ranks[v["link_id"].(int64)] = v["count"].(int64)
	}
	var result []map[string]interface{}
	rank := h.linkStore.QueryInID(ids...)
	for _, v := range rank {
		line := structs.Map(v)
		line["count"] = ranks[v.ID]
		result = append(result, line)
	}
	fmt.Println(result)
	c.HTML(http.StatusOK, "index.html", result)
}
