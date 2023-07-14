package api

import (
	"Backend-Go/gin_demo/dao"
	"Backend-Go/gin_demo/model"
	"Backend-Go/gin_demo/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	var (
		username string
		password string
	)

	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.ResponseFail(c, "username or password is empty", http.StatusBadRequest)
		return
	}

	username = c.PostForm("username")
	password = c.PostForm("password")
	if username == "" || password == "" {
		utils.ResponseSuccess(c, "username or password is empty", http.StatusBadRequest)
		return
	}
	if dao.AddUser(username, password) {
		utils.ResponseSuccess(c, "register success", http.StatusOK)
	}
}

func Login(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.ResponseFail(c, "username or password is empty", http.StatusBadRequest)
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")

	flag := dao.FindUser(username, password)
	if !flag {
		utils.ResponseFail(c, "login fail", http.StatusBadRequest)
		return
	}

	c.SetCookie("username", username, 3600, "/", "localhost", false, true)
	claims := model.MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
			Issuer:    "yzy",
		},
	}

	// 通过claims创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名
	signedString, _ := token.SignedString([]byte("yzy"))

	utils.ResponseSuccess(c, signedString, http.StatusOK)

}
func Info(c *gin.Context) {
	username, _ := c.Get("username")
	utils.ResponseSuccess(c, gin.H{"username": username}, http.StatusOK)
}
