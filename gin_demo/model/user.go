package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
