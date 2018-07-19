package middleware

import wstore "github.com/dkeng/w4w/src/store"

var (
	run = false
)

// Start 启动
func Start(store *wstore.Store) {
	run = true
	go func(s *wstore.Store) {
		for {
			select {
			case r := <-RedirectRecord:
				s.DB.Create(r)
			}
			// 用户主动关闭，如果请求记录管道内容为空，关闭管道
			if !run && len(RedirectRecord) == 0 {
				close(RedirectRecord)
				break
			}
		}
	}(store)
}

// Close 关闭
func Close() {
	run = false
}
