package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"ConfigPlatform/conf"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Conn *sql.DB
)

func InitDb() error {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DatabaseSetting.User,
		conf.DatabaseSetting.Password,
		conf.DatabaseSetting.Host,
		strconv.Itoa(conf.DatabaseSetting.Port),
		conf.DatabaseSetting.Name)

	// 连接数据库
	var err error
	Conn, err = sql.Open(conf.DatabaseSetting.Type, connection)
	if err != nil {
		log.Print("init db failed: ", err)
		return err
	}

	// 设置属性
	Conn.SetConnMaxLifetime(time.Minute * 3)
	Conn.SetMaxOpenConns(10)
	Conn.SetMaxIdleConns(10)

	return nil
}
