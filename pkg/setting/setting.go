package setting

import (
	"log"
	"os"

	"github.com/go-ini/ini"
)

var AppSetting *App
var MysqlSetting *Mysql
var SrsSetting *Srs

type App struct {
	AppRoot   string `ini:"root"`
	Port      int    `ini:"port"`
	JwtSecret string `ini:"jwt_secret"`
	JwtIssuer string `ini:"jwt_issuer"`
}

type Mysql struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Database string `ini:"database"`
}

type Srs struct {
	RtmpPort       string `ini:"rtmp_port"`
	NginxHttpPort  string `ini:"nginx_http_port"`
	NginxHttpsPort string `ini:"nginx_https_port"`
	HttpApiPort    string `ini:"http_api_port"`
	RtcServerPort  string `ini:"rtc_server_port"`
}

func init() {
	confFile := os.Args[1]
	cfg, err := ini.Load(confFile)
	if err != nil {
		log.Fatalf("setting.init, fail to parse '%s': %v", confFile, err)
	}
	AppSetting = &App{}
	MysqlSetting = &Mysql{}
	SrsSetting = &Srs{}

	err = cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo app err: %v", err)
	}

	err = cfg.Section("mysql").MapTo(MysqlSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo datebase err: %v", err)
	}

	err = cfg.Section("srs").MapTo(SrsSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo srs err: %v", err)
	}
}
