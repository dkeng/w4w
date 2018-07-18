package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/dkeng/w4w/src/api/middleware"
	"github.com/dkeng/w4w/src/config"
	"github.com/dkeng/w4w/src/server"
)

func main() {
	// 配置文件
	config.Init()
	middleware.Start(nil)
	server.Start()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	defer close()
	log.Println("Server exiting")
}

// 关闭
func close() {
	// 关闭中间件
	middleware.Close()
	server.Close()
}
