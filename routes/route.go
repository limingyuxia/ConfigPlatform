package routes

import (
	"ConfigPlatform/routes/middleware/cors"
	"ConfigPlatform/routes/middleware/jwt"
	"ConfigPlatform/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 初始化路由配置
func InitRouter() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 处理跨域
	r.Use(cors.HandleCors)

	// 加载静态文件
	r.StaticFile("/", "web/index.html")
	r.StaticFile("project.js", "web/project.js")
	r.StaticFile("project.css", "web/project.css")

	// swagger文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 登录
	r.POST("/login", services.Login)

	// 注册
	r.POST("/register", services.Register)

	// api 接口
	api := r.Group("/project")

	addProjectRoute(api)

	r.Run(":8000")
}

func addProjectRoute(g *gin.RouterGroup) {
	g.Use(jwt.JWT())
	{
		g.GET("list", services.GetProjectList)
		g.GET("detail", services.GetProjectDetail)
		g.POST("add", services.AddProject)
		g.POST("edit", services.EditProject)
		g.DELETE("delete", services.DeleteProject)
	}
}
