package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index 首页
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// Short 短链接
func Short(c *gin.Context) {
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
