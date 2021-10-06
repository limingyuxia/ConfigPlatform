package services

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"ConfigPlatform/api/users"
	"ConfigPlatform/conf"
)

type LoginReq struct {
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

func encodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

var jwtSecret []byte

func init() {
	jwtSecret = []byte(conf.JwtSecret)
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
// @Param Login body LoginReq true "登录用户的账号密码"
// @Success 200 {object} Token
// @Failure 400 {object} WebResponse
// @Router /login [post]
func Login(c *gin.Context) {
	var login LoginReq

	if err := c.ShouldBindJSON(&login); err == nil {

		// 校验用户名密码
		if err := users.IsUserRegisted(c, login.Username, login.Password); err == nil {

			// 生成token
			token, err := generateToken(login.Username, login.Password)
			if err != nil {
				ResponseError(TOKEN_ERROR, err.Error(), c)
				return
			}

			ResponseData(Token{
				Token:  token,
				Expire: time.Now().Add(3 * time.Hour).Unix(),
			}, c)

		} else {
			ResponseError(DB_ERROR, err.Error(), c)
		}

	} else {
		ResponseError(PARAMS_ERROR, err.Error(), c)
	}
}

// @Tags 鉴权
// @Summary 注册账号
// @Accept  json
// @Produce  json
// @Param Register body LoginReq true "注册用户的账号密码"
// @Success 200 {string} string
// @Failure 400 {object} WebResponse
// @Router /register [post]
func Register(c *gin.Context) {
	var register LoginReq

	if err := c.ShouldBindJSON(&register); err == nil {

		userId, err := users.UserRegisted(c, register.Username, register.Password)
		if err != nil {
			ResponseError(REGISTER_ERROR, err.Error(), c)
			return
		}

		ResponseData("注册成功, id: "+strconv.Itoa(int(userId)), c)
	} else {
		ResponseError(PARAMS_ERROR, err.Error(), c)
	}
}
