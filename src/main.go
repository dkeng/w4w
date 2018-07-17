package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/dkeng/w4w/src/config"
	"github.com/dkeng/w4w/src/server"
)

func main() {
	// 配置文件
	config.Init()
	server.Startup()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	defer close()
	log.Println("Server exiting")
}

// 关闭
func close() {
	server.Close()
}
