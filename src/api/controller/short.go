package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dkeng/w4w/src/api/middleware"
	"github.com/dkeng/w4w/src/core"
	"github.com/dkeng/w4w/src/entity"
	"github.com/dkeng/w4w/src/server/router"
	"github.com/dkeng/w4w/src/store"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
	key := c.Param("key")
	key = strings.Trim(key, " ")
	_, flag := c.GetQuery("302")
	link := s.linkStore.QueryByShort(key)
	if link == nil {
		c.String(http.StatusNotFound, "短链接不存在")
		return
	}
	refererType := 1
	if flag {
		refererType = 2
		// 页面重定向
		c.HTML(http.StatusOK, "redirect.html", gin.H{
			"url": link.URL,
		})
	} else {
		// 302 临时重定向
		c.Redirect(http.StatusFound, link.URL)
	}

	defer func() {
		rr := &entity.RedirectRecord{
			LinkID:      link.ID,
			Referer:     c.Request.Referer(),
			RefererType: refererType,
			Path:        c.Request.URL.Path,
			RawQuery:    c.Request.URL.RawQuery,
			ClientIP:    c.ClientIP(),
			UserAgent:   c.Request.UserAgent(),
		}
		// 请求消息入队
		middleware.RedirectRecord <- rr
	}()
}

type shortModel struct {
	URL string `form:"url" json:"url" binding:"required"`
}

// AddShort 添加短链接
func (s *ShortController) AddShort(c *gin.Context) {
	var model shortModel
	err := c.Bind(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请检查参数是否输入有效",
		})
		return
	}
	// 格式化检查URL
	url, err := core.FormatURL(model.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	// 生成短链接
	shortURLs, err := core.ShortURL(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	add := false
	shortURL := ""
	// 添加不存在的
	for i := 0; i < 4; i++ {
		shortURL = shortURLs[i]
		exit := s.linkStore.ExistShort(shortURL)
		if exit {
			add = true
			break
		}
		link := &entity.Link{
			URL:   url,
			Short: shortURL,
		}
		add = s.linkStore.Add(link)
		if add {
			break
		}
	}
	if !add {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "短链接转换失败",
		})
		return
	}
	// 生成短链接
	shortLink1 := fmt.Sprintf("%s/%s", viper.GetString("system.website_url"), shortURL)
	shortLink2 := fmt.Sprintf("%s/%s?302", viper.GetString("system.website_url"), shortURL)
	// 生成短链接二维码
	shortLinkQrcode1, err := core.CreateQrcodeBase64(shortLink1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "短链接转换失败",
		})
		return
	}
	shortLinkQrcode2, err := core.CreateQrcodeBase64(shortLink2)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "短链接转换失败",
		})
		return
	}
	result := resultShort{
		ShortLink1:       shortLink1,
		ShortLinkQrcode1: shortLinkQrcode1,
		ShortLink2:       shortLink2,
		ShortLinkQrcode2: shortLinkQrcode2,
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
