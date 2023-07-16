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
	//c.Query可用于获取?name=test&state=1这类 URL 参数，而c.DefaultQuery则支持设置一个默认值
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
