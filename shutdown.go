package minutil

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// Run 函数封装了启动和优雅退出的逻辑
func RunServer(router *gin.Engine, addr string) {
	// 创建一个HTTP服务器
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	// 启动HTTP服务器
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 创建一个通道来接收系统信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	Info("Shutting down server......")

	// 创建一个上下文，设置超时时间为5秒
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭HTTP服务器
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	Info("Server Exiting")
}

// func main() {
// 	// 创建一个Gin引擎
// 	router := gin.Default()

// 	// 定义一个简单的路由
// 	router.GET("/", func(c *gin.Context) {
// 		time.Sleep(5 * time.Second) // 模拟一个耗时的请求
// 		c.String(http.StatusOK, "Hello, World!")
// 	})

// 	// 调用封装的Run函数启动服务器
// 	Run(router, ":8080")
// }
