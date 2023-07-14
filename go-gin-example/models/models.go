package models

import (
	"Backend-Go/go-gin-example/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	//"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

// 初始化数据库
func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	// 读取配置文件
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Printf("Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Printf("Fail to open database: %v", err)
	}

	// 设置表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	// 设置表名为单数
	db.SingularTable(true)
	// 设置数据库连接池
	db.DB().SetMaxIdleConns(10)
	// 设置最大打开连接数
	db.DB().SetMaxOpenConns(100)
}
