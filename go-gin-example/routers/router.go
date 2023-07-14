package routers

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	// 新的路由实例
	r := gin.New()

	// 使用中间件
	r.Use(gin.Logger())

	// 使用中间件
	r.Use(gin.Recovery())

	// 设置gin框架的运行模式
	gin.SetMode("debug")

	r.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	return r
}
