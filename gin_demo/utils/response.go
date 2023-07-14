package utils

import "github.com/gin-gonic/gin"

func ResponseSuccess(c *gin.Context, message any, statusCode int) {
	c.JSON(statusCode, gin.H{
		"code": statusCode,
		"msg":  message,
	})
}

func ResponseFail(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, gin.H{
		"code": statusCode,
		"msg":  message,
	})
}
