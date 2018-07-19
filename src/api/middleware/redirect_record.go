package middleware

import (
	"github.com/dkeng/w4w/src/entity"
)

var (
	//RedirectRecord 请求记录管道 缓存通道
	RedirectRecord = make(chan *entity.RedirectRecord, 1024)
)
