package routes

import (
	"ConfigPlatform/conf"
	"ConfigPlatform/routes/middleware/captcha"
	"ConfigPlatform/routes/middleware/cors"
	"ConfigPlatform/routes/middleware/jwts"
	"ConfigPlatform/routes/middleware/secure"
	"ConfigPlatform/services"
	"os"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
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
	// r.StaticFile("/", "web/index.html")
	// r.StaticFile("project.js", "web/project.js")
	// r.StaticFile("project.css", "web/project.css")

	// swagger文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authMiddleware := jwts.JwtInit()

	// 登录
	r.POST("/login", authMiddleware.LoginHandler)

	// 登出
	r.GET("/logout", authMiddleware.LogoutHandler)

	// 刷新token
	r.GET("/auth/refreshToken", authMiddleware.RefreshHandler)

	// 注册
	r.POST("/register", services.Register)

	// api 接口
	project := r.Group("/project")

	captchas := r.Group("/captcha")

	emails := r.Group("/email")

	addProjectRoute(project, authMiddleware)

	addCaptchasRoute(captchas)

	addEmailsRoute(emails)

	// 处理图片验证码
	r.Use(captcha.RefeshCaptcha())

	if os.Getenv("Proto") == "Https" {
		r.Use(secure.TlsHandler())
		r.RunTLS(":"+strconv.Itoa(conf.ServerSetting.HttpsPort),
			"./conf/nginx/limingyuxia.asia.crt", "./conf/nginx/limingyuxia.asia.key")

		return
	}

	// 运行服务
	r.Run(":" + strconv.Itoa(conf.ServerSetting.HttpPort))
}

func addProjectRoute(g *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	g.Use(authMiddleware.MiddlewareFunc())
	{
		g.GET("list", services.GetProjectList)
		g.GET("detail", services.GetProjectDetail)
		g.POST("add", services.AddProject)
		g.POST("edit", services.EditProject)
		g.DELETE("delete", services.DeleteProject)
	}
}

func addCaptchasRoute(g *gin.RouterGroup) {
	// 获取图片验证码id
	g.GET("/get", services.GetCaptcha)

	// 验证图片验证码
	g.POST("/confirm", services.ConfirmCaptcha)
}

func addEmailsRoute(g *gin.RouterGroup) {
	g.POST("/send", services.SendEmailCode)

	g.POST("/confirm", services.ConfirmEmailCode)
}
