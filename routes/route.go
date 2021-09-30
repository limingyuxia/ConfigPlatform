package routes

import (
	"ConfigServer/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 初始化路由配置
func InitRouter() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFile("/", "web/index.html")
	r.StaticFile("project.js", "web/project.js")
	r.StaticFile("project.css", "web/project.css")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := r.Group("/config")

	addConfigRoute(api)

	r.Run(":8000")
}

func addConfigRoute(g *gin.RouterGroup) {
	g.GET("list", services.GetProjectList)
}
