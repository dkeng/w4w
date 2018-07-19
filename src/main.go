package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/dkeng/w4w/src/store"

	"github.com/dkeng/w4w/src/api/middleware"
	"github.com/dkeng/w4w/src/config"
	"github.com/dkeng/w4w/src/server"
)

var (
	flagConfig string
	flagHelp   bool
)

func init() {
	flag.StringVar(&flagConfig, "config", "config/w4w.toml", "配置文件")
	flag.BoolVar(&flagHelp, "help", false, "帮助")
	flag.Parse()
}
func float() bool {
	if flagHelp {
		flag.PrintDefaults()
		return false
	}
	return true
}
func main() {
	if !float() {
		return
	}
	// 配置文件
	config.Init(flagConfig)
	store := new(store.Store)
	err := store.Open()
	if err != nil {
		log.Fatalf("打开存储错误：%s", err.Error())
		os.Exit(1)
	}
	// 启动中间件
	middleware.Start(store)
	// 启动服务器
	server.Start(store)

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
