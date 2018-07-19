package controller

import (
	"net/http"

	"github.com/dkeng/w4w/src/server/router"
	"github.com/dkeng/w4w/src/store"
	"github.com/gin-gonic/gin"
)

// ShortController Short
type ShortController struct {
	linkStore store.LinkStore
}

// Init 初始化 ShortController
func (s *ShortController) Init(allStore *store.AllStore) *ShortController {
	s.linkStore = allStore.LinkStore
	return s
}

// GetRouter 获取 ShortController 路由
func (s *ShortController) GetRouter() []*router.R {
	return []*router.R{
		&router.R{
			Method:  http.MethodPost,
			Path:    "/short",
			Handler: s.AddShort,
			Desc:    "添加短链接",
		},
	}
}

// RedirectShort 短链接
func (s *ShortController) RedirectShort(c *gin.Context) {
	// key := c.Param("key")
	_, flag := c.GetQuery("302")
	url := "http://www.baidu.com"
	if flag {
		c.HTML(http.StatusOK, "redirect.html", gin.H{
			"url": url,
		})
	} else {
		c.Redirect(http.StatusMovedPermanently, url)
	}
}

// AddShort 添加短链接
func (s *ShortController) AddShort(c *gin.Context) {

	result := resultShort{
		ShortLink1:       "",
		ShortLinkQrcode1: "",
		ShortLink2:       "",
		ShortLinkQrcode2: "",
	}
	c.JSON(http.StatusCreated, result)
}

// resultShort 返回短链接
type resultShort struct {
	ShortLink1       string `json:"short_link1"`
	ShortLinkQrcode1 string `json:"short_link_qrcode1"`
	ShortLink2       string `json:"short_link2"`
	ShortLinkQrcode2 string `json:"short_link_qrcode2"`
}
