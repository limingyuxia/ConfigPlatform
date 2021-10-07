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
	Port     int    // 数据库port
	Name     string // 要连接的database名称
}

type Nginx struct {
	Domain    string // 域名或者ip
	HttpsPort int    // https port
	HttpPort  int    // http port
}

type Server struct {
	Port int // 服务端的端口
}

type App struct {
	JwtSecret string // jwt秘钥
}

var DatabaseSetting = &Database{}

var NginxSetting = &Nginx{}

var ServerSetting = &Server{}

var AppSetting = &App{}

func LoadConf() error {

	cfg, err := ini.Load("conf/conf.ini")
	if err != nil {
		log.Print("load configure file failed: ", err)
		return err
	}

	// 加载数据库配置
	err = cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Print("load database config failed: ", err)
		return err
	}

	// 加载nginx配置
	err = cfg.Section("nginx").MapTo(NginxSetting)
	if err != nil {
		log.Print("load nginx config failed: ", err)
		return err
	}

	// 加载jwt秘钥
	err = cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Print("load jwt config failed: ", err)
		return err
	}

	// 加载服务端配置
	err = cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Print("load server config failed: ", err)
		return err
	}

	return nil
}
