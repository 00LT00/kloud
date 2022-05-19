package rest

import (
	"context"
	"encoding/gob"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"kloud/model"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var router *gin.Engine

func init() {
	gob.Register(model.User{})
	router = gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	initRouter()
}

func Run(addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//创建一个信号监听通道
	quit := make(chan os.Signal, 1)
	//监听 syscall.SIGINT 跟 syscall.SIGTERM信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	si := <-quit
	log.Println("Shutting down server...", si)

	//shutdown方法需要传入一个上下文参数，这里就设计到两种用法
	//1.WithCancel带时间，表示接收到信号之后，过完该断时间不管当前请求是否完成，强制断开
	ctx, cancel := context.WithTimeout(context.Background(), 9*time.Second)
	//2.不带时间，表示等待当前请求全部完成再断开
	//ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		//当请求还在的时候强制断开了连接将产生错误，err不为空
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
