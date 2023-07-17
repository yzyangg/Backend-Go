package jwt

import (
	"Backend-Go/go-gin-example/pkg/err"
	"Backend-Go/go-gin-example/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = err.SUCCESS

		token := c.Query("token")

		if token == "" {
			code = err.INVALID_PARAMS

		} else {
			claims, e := util.ParseToken(token)
			if e != nil {
				// 验证失败
				code = err.ERROR_AUTH_CHECK_TOKEN_FAIL

			} else if time.Now().Unix() > claims.ExpiresAt {
				// 超时
				code = err.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != err.SUCCESS {
			util.Fail(c, http.StatusUnauthorized, err.GetMsg(code), data)
			c.Abort()
			return
		}
		// 放行
		c.Next()
	}
}
