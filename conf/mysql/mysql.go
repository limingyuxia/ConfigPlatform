package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"ConfigPlatform/conf"
	"ConfigPlatform/util"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/boil"
)

var (
	Conn *sql.DB
)

func InitDb() {
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DatabaseSetting.User,
		conf.DatabaseSetting.Password,
		conf.DatabaseSetting.Host,
		conf.DatabaseSetting.Name)

	// 连接数据库
	var err error
	Conn, err = sql.Open(conf.DatabaseSetting.Type, connection)
	if err != nil {
		panic(err)
	}

	// 设置属性
	Conn.SetConnMaxLifetime(time.Minute * 3)
	Conn.SetMaxOpenConns(10)
	Conn.SetMaxIdleConns(10)

	// 设置sql debug
	boil.DebugMode = true
	boil.DebugWriter = &util.LogWriter{}
}

func GetDb() *sql.DB {
	return Conn
}
