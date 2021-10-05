package conf

import (
	"log"

	"github.com/go-ini/ini"
)

type Database struct {
	Type     string // 数据库类型
	User     string // 数据库用户名
	Password string // 数据库密码
	Host     string // 数据库host
	Name     string // 要连接的database名称
}

type Nginx struct {
	Domain    string // 域名或者ip
	HttpsPort int    // https port
	HttpPort  int    // http port
}

var DatabaseSetting = &Database{}
var NginxSetting = &Nginx{}

func LoadConf() {
	cfg, err := ini.Load("conf/conf.ini")
	if err != nil {
		panic(err)
	}

	// 加载数据库配置
	err = cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Print(err)
	}

	// 加载nginx配置
	err = cfg.Section("nginx").MapTo(NginxSetting)
	if err != nil {
		log.Print(err)
	}
}
