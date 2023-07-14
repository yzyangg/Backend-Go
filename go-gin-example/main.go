package main

import (
	"Backend-Go/go-gin-example/pkg/setting"
	"Backend-Go/go-gin-example/routers"
	"fmt"
	"net/http"
)

func main() {
	//router := gin.Default() // 默认的路由实例
	//router.GET("/get", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "Hello, World!",
	//	})
	//})

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		MaxHeaderBytes: 1 << 20, // 请求头最大字节数
	}
	// 启动http服务，并且开始监听请求
	err := s.ListenAndServe()

	if err != nil {
		return
	}
}
