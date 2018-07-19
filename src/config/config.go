package config

import (
	"log"
	"os"

	"github.com/dkeng/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Init 配置文件初始化
func Init(in string) {
	// 初始化配置文件
	viper.SetConfigFile(in)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件错误：%s", err.Error())
		os.Exit(1)
	}
	// 日志初始化
	logger.Init()
	// logger.RegisterSentry()

	if viper.GetString("system.mode") == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}
