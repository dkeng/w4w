package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RedirectShort 短链接
func RedirectShort(c *gin.Context) {
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
func AddShort(c *gin.Context) {
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
