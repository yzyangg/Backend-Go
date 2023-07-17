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
	// 自动迁移
	//这样，通过 db.AutoMigrate() 方法，你可以保持数据库表的结构与定义的 student 结构体模型同步
	//无需手动创建或修改数据库表
	//这对于在开发过程中更改模型结构或确保表结构与代码定义一致非常有用。
	db.AutoMigrate(&student{})

}

func insert(st *student) {
	//插入
	db.Create(st)
}
