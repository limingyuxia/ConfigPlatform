package main

import (
	"log"

	"ConfigPlatform/conf"
	"ConfigPlatform/conf/mysql"
	_ "ConfigPlatform/docs"
	"ConfigPlatform/routes"
)

// @title 配置平台
// @contact.name 李明
// @contact.email 2572597150@qq.com
// @version 1.0
// @description  配置平台 HTTP接口文档
func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Print(nil, "got_a_panic: ", err)
		}
	}()

	// 加载配置文件
	conf.LoadConf()

	// 初始化db
	mysql.InitDb()
	defer mysql.Conn.Close()

	// 初始化路由
	routes.InitRouter()
}
