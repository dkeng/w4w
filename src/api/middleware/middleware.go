package middleware

import (
	cstore "couponwebapi/src/store"
	"time"

	"github.com/dkeng/w4w/src/entity"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

var (
	//requestRecord 请求记录管道 缓存通道
	requestRecord = make(chan *entity.RequestRecord, 1024)
	run           = false
)

// Start 启动
func Start(store *cstore.Store) {
	run = true
	go func(s *cstore.Store) {
		for {
			select {
			case r := <-requestRecord:
				s.DB.Create(r)
			}
			// 用户主动关闭，如果请求记录管道内容为空，关闭管道
			if !run && len(requestRecord) == 0 {
				close(requestRecord)
				break
			}
		}
	}(store)
}

// Close 关闭
func Close() {
	run = false
}

// Cors 跨域
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	// config.AddAllowHeaders("Authorization")
	return cors.New(config)
}

// Header 头处理
func Header() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		c.Writer.Header().Set("Server", "W4WServer")
	}
}

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
