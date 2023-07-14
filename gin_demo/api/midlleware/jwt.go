package midlleware

import (
	"Backend-Go/gin_demo/model"
	"Backend-Go/gin_demo/utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

var Secret = []byte("yzy")

func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			utils.ResponseFail(c, "no authorization", 401)
			c.Abort()
			return
		}

		// 从 Authorization 中解析出 token
		parts := strings.SplitN(authHeader, " ", 2)

		if !(len(parts) == 2 && parts[0] == "Bearer") {
			utils.ResponseFail(c, "authorization format error", 401)
			c.Abort()
			return
		}
		token, err := ParseToken(parts[1])
		if err != nil {
			utils.ResponseFail(c, "error !!", 401)
			c.Abort()
			return
		}
		c.Set("username", token.Username)
		c.Next()
	}
}

func ParseToken(tokenString string) (*model.MyClaims, error) {
	// 解析token
	// tokenString 是客户端传过来的token,MyClaims 是一个结构体，用来接收解析出来的值,token 是解析出来的token对象,err 是解析出来的错误信息
	// 解析tokenString，返回token对象，如果有错误就返回错误信息
	token, err := jwt.ParseWithClaims(tokenString, &model.MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
