package api

import (
	"Backend-Go/gin_demo/api/midlleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.Use(midlleware.CORS())

	r.POST("/register", Register) // 注册
	r.POST("/login", Login)       // 登录

	UserRouter := r.Group("/user")
	{
		UserRouter.Use(midlleware.JWTAuth())
		UserRouter.GET("/get", Info) // 获取用户信息
	}

	err := r.Run(":8080")
	if err != nil {
		return
	} // 跑在 8088 端口上

}
