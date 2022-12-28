package setting

import (
	"log"
	"os"

	"github.com/go-ini/ini"
)

type App struct {
	AppRoot   string `ini:"root"`
	Port      int    `ini:"port"`
	JwtSecret string `ini:"jwt_secret"`
}

var AppSetting *App

type Mysql struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Database string `ini:"database"`
}

var MysqlSetting *Mysql

func init() {
	confFile := os.Args[1]
	cfg, err := ini.Load(confFile)
	if err != nil {
		log.Fatalf("setting.init, fail to parse '%s': %v", confFile, err)
	}
	AppSetting = &App{}
	MysqlSetting = &Mysql{}

	err = cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo app err: %v", err)
	}

	err = cfg.Section("mysql").MapTo(MysqlSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo datebase err: %v", err)
	}
}
