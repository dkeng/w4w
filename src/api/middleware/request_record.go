package middleware

import (
	"time"

	"github.com/dkeng/w4w/src/entity"
	"github.com/gin-gonic/gin"
)

var (
	//requestRecord 请求记录管道 缓存通道
	requestRecord = make(chan *entity.RequestRecord, 1024)
)

// RequestRecord 请求记录
func RequestRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		// Path
		path := c.Request.URL.Path
		// Get参数
		raw := c.Request.URL.RawQuery
		c.Next()
		// 结束时间
		end := time.Now()
		writerSize := c.Writer.Size()
		if writerSize < 0 {
			writerSize = 0
		}
		rr := &entity.RequestRecord{
			StartTime:   start,
			EndTime:     end,
			Path:        path,
			RawQuery:    raw,
			Latency:     end.Sub(start),
			Referer:     c.Request.Referer(),
			ClientIP:    c.ClientIP(),
			Method:      c.Request.Method,
			StatusCode:  c.Writer.Status(),
			UserAgent:   c.Request.UserAgent(),
			RequestSize: c.Request.ContentLength,
			WriterSize:  int64(writerSize),
		}
		// 请求消息入队
		requestRecord <- rr
	}
}
