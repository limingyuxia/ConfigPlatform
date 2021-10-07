package jwt

import (
	"ConfigPlatform/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code services.RETCODE = services.SUCCESS

		// 取token
		token := c.GetHeader("Authorization")
		if token == "" {
			code = services.NO_TOKEN_ERROR
		} else {
			_, err := services.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = services.EXPIRE_ERROR
				default:
					code = services.NOAUTH_ERROR
				}
			}
		}

		if code != services.SUCCESS {
			services.ResponseError(code, services.RETCODE_MSG[code], c)
			c.Abort()
			return
		}

		c.Next()
	}
}
