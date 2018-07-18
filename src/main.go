package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/dkeng/w4w/src/store"

	"github.com/dkeng/w4w/src/api/middleware"
	"github.com/dkeng/w4w/src/config"
	"github.com/dkeng/w4w/src/server"
)

func main() {
	// 配置文件
	config.Init()
	store := new(store.Store)
	err := store.Open()
	if err != nil {
		log.Fatalf("打开存储错误：%s", err.Error())
		os.Exit(1)
	}
	// middleware.Start(nil)
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
	// 关闭服务器
	server.Close()
}
