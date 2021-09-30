package main

import (
	"ConfigServer/routes"

	"ConfigServer/conf"
	"ConfigServer/conf/mysql"
	_ "ConfigServer/docs"
)

// @title 配置平台
// @contact.name 李明
// @contact.email 2572597150@qq.com
// @version 1.0
// @description  配置平台 HTTP接口文档
func main() {
	// 加载配置文件
	conf.LoadConf()

	// 初始化db
	mysql.InitDb()
	defer mysql.Conn.Close()

	// 初始化路由
	routes.InitRouter()
}
