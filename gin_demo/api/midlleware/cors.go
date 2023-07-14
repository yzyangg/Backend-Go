package midlleware

import "github.com/gin-gonic/gin"

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token, x-user-id")
		c.Writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Add("Access-Control-Allow-Credentials", "true")

		//在这段代码中，如果请求的方法为 OPTIONS，就会调用 AbortWithStatus 方法将响应状态码设置为 204（No Content），并终止请求的处理。
		//返回 204 状态码表示预检请求成功，告诉浏览器服务器允许跨域访问。
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		//c.Writer.He
		c.Next()
	}
}
