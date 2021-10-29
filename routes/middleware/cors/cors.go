package cors

import (
	"ConfigPlatform/conf"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 处理跨域
func HandleCors(c *gin.Context) {
	var httpReq string
	var httpPort string = strconv.Itoa(conf.NginxSetting.HttpsPort)

	var httpsReq string
	var httpsPort string = strconv.Itoa(conf.NginxSetting.HttpPort)

	if httpPort == "80" {
		httpReq = "http://" + conf.NginxSetting.Domain
	} else {
		httpReq = "http://" + conf.NginxSetting.Domain + ":" + httpPort
	}

	if httpsPort == "443" {
		httpsReq = "https://" + conf.NginxSetting.Domain
	} else {
		httpsReq = "https://" + conf.NginxSetting.Domain + httpsPort
	}

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
