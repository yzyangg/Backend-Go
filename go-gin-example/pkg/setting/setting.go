package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg          *ini.File
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PageSize     int
	JwtSecret    string
)

// 读取配置文件
func init() {
	var err error
	Cfg, err = ini.Load("E:\\codes\\gocode\\Backend-Go\\go-gin-example\\conf\\app.ini")
	if err != nil {
		log.Printf("Fail to parse 'conf/app.ini': %v\n", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()

}

func LoadBase() {
	// 读取配置文件（顶级）
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}
func LoadServer() {
	// 读取配置文件（次级）
	section, err := Cfg.GetSection("app")
	if err != nil {
		log.Printf("Fail to get section 'server': %v\n", err)
	}

	HTTPPort = section.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(section.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(section.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

}
func LoadApp() {
	section, err := Cfg.GetSection("app")
	if err != nil {
		log.Printf("Fail to get section 'app': %v\n", err)
	}

	JwtSecret = section.Key("JWT_SECRET").MustString("yzy")
	PageSize = section.Key("PAGE_SIZE").MustInt(10)

}
