package qq

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"ConfigPlatform/services/auth2"

	"github.com/gin-gonic/gin"
)

type authMe struct {
	ClientId string `json:"client_id"` // 网站的appid
	Openid   string `json:"openid"`    // 用户的openid
}

type userInfo struct {
	Ret             int    `json:"ret,omitempty"`
	Msg             string `json:"msg,omitempty"`
	Nickname        string `json:"nickname"`                // QQ昵称
	Gender          string `json:"gender"`                  // 性别
	Province        string `json:"province"`                // 省份
	City            string `json:"city"`                    // 城市
	Year            string `json:"year"`                    // 出生年
	Constellation   string `json:"constellation,omitempty"` // 星座
	Figureurl       string `json:"figureurl"`               // 大小为30×30像素的QQ空间头像URL
	Figureurl1      string `json:"figureurl_1"`             // 大小为50×50像素的QQ空间头像URL
	Figureurl2      string `json:"figureurl_2"`             // 大小为100×100像素的QQ空间头像URL
	FigureurlQq     string `json:"figureurl_qq"`            // 大小为640×640像素的QQ头像URL
	FigureurlQq1    string `json:"figureurl_qq_1"`          // 大小为40×40像素的QQ头像URL
	FigureurlQq2    string `json:"figureurl_qq_2"`          // 大小为100×100像素的QQ头像URL。需要注意，不是所有的用户都拥有QQ的100x100的头像，但40x40像素则是一定会有
	FigureurlType   string `json:"figureurl_type"`          // 文档未说明
	IsYellowVip     string `json:"is_yellow_vip"`           // 标识用户是否为黄钻用户（0：不是；1：是）
	Vip             string `json:"vip"`                     // 标识用户是否为QQ会员用户（0：不是；1：是）
	YellowVipLevel  string `json:"yellow_vip_level"`        // 黄钻等级
	Level           string `json:"level"`                   // QQ会员等级
	IsYellowYearVip string `json:"is_yellow_year_vip"`      // 标识是否为年费黄钻用户（0：不是； 1：是）
}

type accessToken struct {
	AccessToken      string `json:"access_token"`      // access token
	ExpiresIn        int64  `json:"expires_in"`        // access token 过期时间，单位是秒
	Error            int32  `json:"error"`             // 错误码
	ErrorDescription string `json:"error_description"` // 错误信息
}

// 获取返回结构中json数据
func getRespJson(originResp string) string {
	jsonDataStart := strings.Index(originResp, "{")
	jsonDataEnd := strings.Index(originResp, "}")
	if jsonDataStart > 0 && jsonDataEnd > 0 {
		jsonData := originResp[jsonDataStart : jsonDataEnd+1]

		return jsonData
	}

	return ""
}

// 获取accessToken
func getAccessToken(ctx context.Context, authorizationCode string) (string, error) {

	var baseUrl = "https://graph.qq.com/oauth2.0/token"
	var appId = "101490224"
	var appSecret = ""
	var redirectUri = "http%3A%2F%2Flocalhost%3A3000%2Fproxy"

	var urlParam = map[string]string{
		"grant_type":    "authorization_code",
		"client_id":     appId,
		"client_secret": appSecret,
		"code":          authorizationCode, // 10分钟有效
		"redirect_uri":  redirectUri,
	}
	var reqHeader = map[string]string{}

	accessTokenResp, err := auth2.HttpGet(ctx, baseUrl, urlParam, reqHeader)
	if err != nil {
		return "", err
	}

	jsonResp := getRespJson(accessTokenResp)
	if jsonResp == "" {
		log.Print("origin data: ", accessTokenResp)
		return "", err
	}

	var accessToken accessToken
	err = json.Unmarshal([]byte(jsonResp), &accessToken)
	if err != nil {
		log.Print("parse qq access token json data error: ", err)
		return "", err
	}

	if accessToken.Error != 0 {
		log.Print(accessToken.ErrorDescription)
		// return "", errors.New(accessToken.ErrorDescription)
	}

	return accessToken.AccessToken, nil
}

// 获取用户的openid
func getUserOpenid(ctx context.Context, accessToken string) (*authMe, error) {

	var baseUrl = "https://graph.qq.com/oauth2.0/me"

	var urlParam = map[string]string{
		"access_token": accessToken,
	}
	var reqHeader = map[string]string{}

	userOpenid, err := auth2.HttpGet(ctx, baseUrl, urlParam, reqHeader)
	if err != nil {
		return nil, err
	}

	jsonResp := getRespJson(userOpenid)
	if jsonResp == "" {
		log.Print("origin data: ", userOpenid)
		return nil, err
	}

	// 解析openid数据
	var authMe authMe
	err = json.Unmarshal([]byte(jsonResp), &authMe)
	if err != nil {
		log.Print("parse auth me json data error: ", err)
		return nil, err
	}

	return &authMe, nil
}

// 获取用户信息
func getUserInfo(ctx context.Context, accessToken string, authMe *authMe) (*userInfo, error) {
	log.Printf("get %v info", authMe.Openid)

	var baseUrl = "https://graph.qq.com/user/get_user_info"

	var urlParam = map[string]string{
		"access_token":       accessToken,
		"oauth_consumer_key": authMe.ClientId,
		"openid":             authMe.Openid,
		"format":             "json",
	}
	var reqHeader = map[string]string{}

	userInfoResp, err := auth2.HttpGet(ctx, baseUrl, urlParam, reqHeader)
	if err != nil {
		return nil, err
	}

	// 解析用户信息
	var userInfo userInfo
	err = json.Unmarshal([]byte(userInfoResp), &userInfo)
	if err != nil {
		log.Print("parse user info json data error: ", err)
		return nil, err
	}

	// 判断返回
	if userInfo.Ret != 0 {
		return nil, errors.New(userInfo.Msg)
	}

	return &userInfo, nil
}

func QQLogin(c *gin.Context) {
	// 获取鉴权码
	authorizationCode := c.Query("code")

	// 	client端的状态值。用于第三方应用防止CSRF攻击
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

	// 这里先模拟token
	accessToken := "14B954C46E3D0A66D0FEB572B9AACB8B"
	authMe, err := getUserOpenid(c, accessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "get user openid failed")
	}

	userinfo, err := getUserInfo(c, accessToken, authMe)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "get user info failed")
	}

	c.JSON(http.StatusOK, userinfo)
}
