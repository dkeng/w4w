package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/dkeng/w4w/src/api/controller"
	"github.com/dkeng/w4w/src/api/middleware"
	"github.com/dkeng/w4w/src/server/router"
	"github.com/dkeng/w4w/src/store"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	httpServer *http.Server
)

// Start 启动
func Start(store *store.Store) {
	handler := gin.Default()
	// 加载中间件
	handler.Use(middleware.Cors(), middleware.Header(), middleware.RequestRecord())
	// 加载模板
	handler.LoadHTMLGlob("templates/*")

	allController := new(controller.AllController).Init(store)
	handler.GET("/", allController.HomeController.Index)
	handler.GET("/:key", allController.ShortController.RedirectShort)
	// api
	api := handler.Group("/api", router.API())
	router.SetGroupRouter(api, allController.ShortController)

	httpServer = &http.Server{
		Addr:    viper.GetString("system.addr"),
		Handler: handler,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

// Close 关闭
func Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("HttpServer Shutdown:", err)
	}
}
