package main

import (
	"log"

	"ConfigPlatform/conf"
	"ConfigPlatform/conf/mysql"
	"ConfigPlatform/conf/redis"
	_ "ConfigPlatform/docs"
	"ConfigPlatform/routes"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// @title 配置平台
// @contact.name 李明
// @contact.email 2572597150@qq.com
// @version 1.0
// @description  配置管理平台 HTTP接口文档
func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Print(nil, "got_a_panic: ", err)
		}
	}()

	// 加载配置文件
	if err := conf.LoadConf(); err != nil {
		return
	}

	// 初始化db
	if err := mysql.InitDb(); err != nil {
		return
	}
	defer mysql.Conn.Close()

	// 初始化redis
	if err := redis.InitRedis(); err != nil {
		return
	}
	defer redis.RedisConn.Close()

	// MySQL 设置Debug模式
	boil.DebugMode = true

	// 初始化路由
	routes.InitRouter()
}
