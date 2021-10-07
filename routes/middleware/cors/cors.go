package cors

import (
	"ConfigPlatform/conf"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 处理跨域
func HandleCors(c *gin.Context) {
	var httpsReq string = "https://" + conf.NginxSetting.Domain + ":" + strconv.Itoa(conf.NginxSetting.HttpsPort)
	var httpReq string = "http://" + conf.NginxSetting.Domain + ":" + strconv.Itoa(conf.NginxSetting.HttpPort)

	if c.GetHeader("Origin") == httpsReq {
		c.Header("Access-Control-Allow-Origin", httpsReq)
	} else if c.GetHeader("Origin") == httpReq {
		c.Header("Access-Control-Allow-Origin", httpReq)
	}

	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTION")
	c.Header("Access-Control-Allow-Headers", "Authorization, Origin, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
		return
	} else {
		c.Next()
	}
}
