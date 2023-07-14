package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func main() {
	initDB()
}

// student类
type student struct {
	id       int
	username string
	password string
	gender   string
	img      string
}

func initDB() {
	// 设置一下dns charset:编码方式 parseTime:是否解析time类型 loc:时区
	dsn := "root:2002@tcp(127.0.0.1:3306)/my_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("open mysql failed, err:%v\n", err)
	}
	log.Println("DB connect success")
	db.AutoMigrate(&student{})

}

func insert(st *student) {
	//插入
	db.Create(st)
}
