package jwts

import (
	"ConfigPlatform/api/users"
	"ConfigPlatform/conf"
	"ConfigPlatform/model"
	"ConfigPlatform/services"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey = "username"
var tokenValidity = 3 * time.Hour    // token有效期
var tokenRefreshValidity = time.Hour // token过期后，刷新的有效期

func JwtInit() *jwt.GinJWTMiddleware {

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(conf.ServerSetting.JwtSecret),
		Timeout:     tokenValidity,
		MaxRefresh:  tokenRefreshValidity,
		IdentityKey: identityKey,

		Authenticator: func(c *gin.Context) (interface{}, error) {

			var loginReq model.LoginReq
			if err := c.ShouldBindJSON(&loginReq); err != nil {
				log.Print("login param error: ", err)
				return "", jwt.ErrMissingLoginValues
			}

			// 校验用户名密码
			if err := users.IsUserRegisted(c, loginReq.Username, loginReq.Password); err == nil {
				return &model.JwtUser{
					UserName: loginReq.Username,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},

		// 把Authenticator返回的信息嵌入到jwt token中
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*model.JwtUser); ok {
				return true
			}

			return false
		},

		// 成功登入后，把从 Authenticator 返回的数据转换到MapClaims总，并嵌入jwt token
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.JwtUser); ok {
				return jwt.MapClaims{
					identityKey:             v.UserName,
					"token_refesh_validity": time.Now().Add(tokenValidity + tokenRefreshValidity).Unix(),
				}
			}

			return jwt.MapClaims{}
		},

		// 登入成功以后的返回返回值
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(code, model.LoginResp{
				Token:  token,
				Expire: expire.Format("2006-01-02 15:04:05"),
			})
		},

		// 鉴权失败以后的返回
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.JwtUser{
				UserName: claims[identityKey].(string),
			}
		},

		// 登出响应
		LogoutResponse: func(c *gin.Context, code int) {
			services.Logout(c)

			c.JSON(code, gin.H{
				"message": "登出成功",
			})
		},

		// 刷新token响应
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			services.RefeshToken(c)

			c.JSON(code, model.LoginResp{
				Token:  token,
				Expire: expire.Format("2006-01-02 15:04:05"),
			})
		},

		TokenLookup: "header: Authorization",

		TokenHeadName: "Bearer",

		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}
