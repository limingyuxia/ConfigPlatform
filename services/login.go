package services

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
}

type Token struct {
	Token  string `json:"token"`  // token
	Expire int64  `json:"expire"` // 过期时间
}

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var jwtSecret []byte

func encodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func generateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		encodeMD5(username),
		encodeMD5(password),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// @Tags 鉴权
// @Summary 登录
// @Accept  json
// @Produce  json
// @Param user body User true "用户信息"
// @Success 200 {object} Token
// @Failure 400 {object} WebResponse
// @Router /login [post]
func Login(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err == nil {
		if user.Username == "superuser" && user.Password == "supertoken" {
			token, err := generateToken(user.Username, user.Password)
			if err != nil {
				ResponseError(NOAUTH_ERROR, RETCODE_name[NOAUTH_ERROR], c)
			}

			ResponseData(Token{
				Token:  token,
				Expire: 60,
			}, c)
		} else {
			ResponseError(NOAUTH_ERROR, RETCODE_name[NOAUTH_ERROR], c)
		}
	} else {
		ResponseError(PARAMS_ERROR, err.Error(), c)
	}
}
