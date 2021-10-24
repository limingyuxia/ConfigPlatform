package wechat

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"ConfigPlatform/services/auth2"

	"github.com/gin-gonic/gin"
)

type accessAuth struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid"`
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
}

// 获取accessToken
func getAccessToken(ctx context.Context, authorizationCode string) (string, error) {

	var baseUrl = "https://api.weixin.qq.com/sns/oauth2/access_token"
	var appId = "wxbdc5610cc59c1631"
	var appSecret = ""

	var urlParam = map[string]string{
		"appid":      appId,
		"secret":     appSecret,
		"code":       authorizationCode,
		"grant_type": "authorization_code",
	}
	var reqHeader = map[string]string{}

	accessTokenResp, err := auth2.HttpGet(ctx, baseUrl, urlParam, reqHeader)
	if err != nil {
		return "", err
	}

	var accessAuth accessAuth
	err = json.Unmarshal([]byte(accessTokenResp), &accessAuth)
	if err != nil {
		log.Print("parse wechat access token json data error: ", err)
		return "", err
	}

	if accessAuth.Errcode != 0 {
		log.Print(accessAuth.Errmsg)
		// return "", errors.New(accessAuth.Errmsg)
	}

	return "", nil
}

func WechatLogin(c *gin.Context) {

	// 获取鉴权码
	authorizationCode := c.Query("code")

	// 	该参数可用于防止csrf攻击（跨站请求伪造攻击）
	state := c.Query("state")

	if authorizationCode == "" || state == "" {
		c.JSON(http.StatusUnauthorized, "failed")
		return
	}

	// 校验state
	if state != "3d6be0a4035d839573b04816624a415e" {
		c.JSON(http.StatusUnauthorized, "failed")
		return
	}

	_, err := getAccessToken(c, authorizationCode)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "get access token failed")
		return
	}

	c.JSON(http.StatusOK, "ok")
}
