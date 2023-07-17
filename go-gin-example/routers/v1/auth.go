package v1

import (
	"Backend-Go/go-gin-example/models"
	"Backend-Go/go-gin-example/pkg/err"
	"Backend-Go/go-gin-example/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
)

type auth struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username,
		Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := err.INVALID_PARAMS

	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err2 := util.GenerateToken(username, password)

			if err2 != nil {
				code = err.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = err.SUCCESS
			}

		} else {
			code = err.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	util.Success(c, code, err.GetMsg(code), data)
}
