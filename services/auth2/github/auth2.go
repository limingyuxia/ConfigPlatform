package github

import (
	"ConfigPlatform/api/users"
	"ConfigPlatform/model"
	"ConfigPlatform/routes/middleware/jwts"
	"ConfigPlatform/services/auth2"
	"context"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type accessToken struct {
	AccessToken string `json:"access_token"` // access token
	Scope       string `json:"scope"`        // 作用域
	TokenType   string `json:"token_type"`   // token 类型
}

type userInfo struct {
	Login           string `json:"login"`            // 登录用户名
	Id              int64  `json:"id"`               // 用户id
	AvatarUrl       string `json:"avatar_url"`       // 用户头像
	NodeId          string `json:"node_id"`          //
	Gravatar        string `json:"gravatar_id"`      //
	HtmlUrl         string `json:"html_url"`         // 用户主页
	Company         string `json:"company"`          // 公司
	Blog            string `json:"blog"`             // 博客
	Location        string `json:"location"`         // 地址
	Email           string `json:"email"`            // 邮箱
	TwitterUsername string `json:"twitter_username"` // 推特用户名
	PublicRepos     int    `json:"public_repos"`     // 公共代码仓库数量
}

// 获取accessToken
func getAccessToken(ctx context.Context, authorizationCode string) (string, error) {
	url := "https://github.com/login/oauth/access_token"
	clientId := "b0f4b22bfa884640f030"
	redirectUri := "http://config-platform.top:8000/githubLogin"

	// 获取github app secret
	clientSecret, err := ioutil.ReadFile("./conf/auth2Secret/github.secret")
	if err != nil {
		log.Print("read github secret faile! ", err)
		return "", err
	}

	var reqHeader = map[string]string{
		"Accept": "application/json",
	}

	var reqBody = map[string]string{
		"client_id":     clientId,
		"client_secret": string(clientSecret),
		"code":          authorizationCode,
		"redirect_uri":  redirectUri,
	}

	accessTokenResp, err := auth2.HttpPost(ctx, url, reqBody, reqHeader)
	if err != nil {
		return "", err
	}

	var accessToken accessToken
	err = json.Unmarshal([]byte(accessTokenResp), &accessToken)
	if err != nil {
		log.Print("parse github access token json data error: ", err)
		return "", err
	}

	return accessToken.AccessToken, nil
}

// 获取用户信息
func getUserInfo(ctx context.Context, accessToken string) (*userInfo, error) {
	url := "https://api.github.com/user"

	urlParam := map[string]string{}
	reqHeader := map[string]string{
		"Authorization": "token " + accessToken,
	}

	userInfoResp, err := auth2.HttpGet(ctx, url, urlParam, reqHeader)
	if err != nil {
		log.Print("get user info failed error: ", err)
		return nil, err
	}

	// 解析用户信息
	var userInfo userInfo
	err = json.Unmarshal([]byte(userInfoResp), &userInfo)
	if err != nil {
		log.Print("parse user info json data error: ", err)
		return nil, err
	}

	return &userInfo, nil
}

func GithubLogin(c *gin.Context) {
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

	accessToken, err := getAccessToken(c, authorizationCode)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "get access token failed")
		return
	}

	userInfo, err := getUserInfo(c, accessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "get user info failed")
		return
	}

	// 手动注册用户
	userName := "User" + "_" + time.Now().Format("2006-01-02 15:04:05")
	passWord := base64.StdEncoding.EncodeToString([]byte(userName))

	var auth2User = &model.Auth2User{
		Type:     "github",
		Nickname: userInfo.Login,
		Avatar:   userInfo.AvatarUrl,
		UniqueId: strconv.FormatInt(userInfo.Id, 10),
		UserName: userName,
		Password: passWord,
	}

	err = users.CreateUser(c, auth2User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "db error")
		return
	}

	// 生成token
	mw := jwts.AuthMiddleware
	token := jwt.New(jwt.GetSigningMethod(mw.SigningAlgorithm))
	claims := token.Claims.(jwt.MapClaims)

	if mw.PayloadFunc != nil {
		data := &model.JwtUser{
			UserName: userName,
		}
		for key, value := range mw.PayloadFunc(data) {
			claims[key] = value
		}
	}

	expire := mw.TimeFunc().Add(mw.Timeout)
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = mw.TimeFunc().Unix()
	tokenString, err := token.SignedString(mw.Key)

	c.JSON(http.StatusOK, model.LoginResp{
		Token:  tokenString,
		Expire: expire.Format("2006-01-02 15:04:05"),
	})
}
