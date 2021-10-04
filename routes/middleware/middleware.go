package middleware

import "github.com/gin-gonic/gin"

// 处理跨域
func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:7000")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTION")
	c.Header("Access-Control-Allow-Headers", "authorization, Access-Token,origin, content-type, accept")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
}
