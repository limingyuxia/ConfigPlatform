package conf

import (
	"github.com/go-ini/ini"
)

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
}

var DatabaseSetting = &Database{}

func LoadConf() {
	cfg, err := ini.Load("conf/conf.ini")
	if err != nil {
		panic(err)
	}

	cfg.Section("database").MapTo(DatabaseSetting)
}
