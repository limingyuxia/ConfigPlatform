package secure

import (
	"ConfigPlatform/conf"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// gtin 处理 https
func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     conf.NginxSetting.Domain + ":" + strconv.Itoa(conf.ServerSetting.HttpsPort),
		})

		if err := secureMiddleware.Process(c.Writer, c.Request); err != nil {
			return
		}

		c.Next()
	}
}
