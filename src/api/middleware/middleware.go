package middleware

import (
	cstore "couponwebapi/src/store"
)

var (
	run = false
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
