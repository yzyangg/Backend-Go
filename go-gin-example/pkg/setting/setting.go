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

func init() {
	var err error
	Cfg, err = ini.Load("./conf/app.ini")
	if err != nil {
		log.Println("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()

}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}
func LoadServer() {
	section, err := Cfg.GetSection("app")
	if err != nil {
		log.Println("Fail to get section 'server': %v", err)
	}

	HTTPPort = section.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(section.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(section.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

}
func LoadApp() {
	section, err := Cfg.GetSection("app")
	if err != nil {
		log.Println("Fail to get section 'app': %v", err)
	}

	JwtSecret = section.Key("JWT_SECRET").MustString("yzy")
	PageSize = section.Key("PAGE_SIZE").MustInt(10)

}
