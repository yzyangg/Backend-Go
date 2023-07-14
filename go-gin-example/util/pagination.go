package util

import (
	"Backend-Go/go-gin-example/pkg/setting"
	"github.com/gin-gonic/gin"
	// 类型转换库
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) int {
	result := 0
	// 类型转换
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
