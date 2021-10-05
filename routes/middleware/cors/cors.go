package cors

import (
	"github.com/gin-gonic/gin"
)

// 处理跨域
func HandleCors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://limingyuxia.asia:9528")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTION")
	c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type")
	c.Header("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	} else {
		c.Next()
	}
}
