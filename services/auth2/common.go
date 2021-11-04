package auth2

import (
	"ConfigPlatform/api/users"
	"ConfigPlatform/model"
	"ConfigPlatform/routes/middleware/jwts"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 发送http get请求
func HttpGet(ctx context.Context, url string, urlParam map[string]string,
	reqHeader map[string]string) (string, error) {

	// 新建请求
	var req *http.Request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Print("new http get request failed: ", err)
		return "", err
	}

	// 添加url参数
	params := req.URL.Query()
	for key, value := range urlParam {
		params.Add(key, value)
	}
	req.URL.RawQuery = params.Encode()

	// 这个默认添加，因为是必须的，否则请求会失败
	req.Header.Set("Content-type", "application/json")
	for key, value := range reqHeader {
		req.Header.Set(key, value)
	}

	// 发送请求
	client := http.Client{}
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		log.Print("client do request faile: ", err)
		return "", err
	}

	// 读响应体
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(ctx, "read resp body failed: ", err)
		return "", err
	}

	// 判断http状态码
	if resp.StatusCode != http.StatusOK {
		log.Print(ctx, "http request error: ", string(respBody))
		return "", errors.New(strconv.Itoa(resp.StatusCode))
	}

	return string(respBody), nil
}

// 发送http post请求
func HttpPost(ctx context.Context, url string, reqBody map[string]string,
	reqHeader map[string]string) (string, error) {

	// 转json字符串
	reqBodyByte, err := json.Marshal(reqBody)
	if err != nil {
		log.Print("package request body failed: ", err)
		return "", err
	}

	bodyReader := bytes.NewBuffer(reqBodyByte)
	var req *http.Request

	// 新建请求
	req, err = http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		log.Print("new http post request failed: ", err)
		return "", err
	}

	// 添加请求头参数
	req.Header.Set("Content-type", "application/json")
	for key, value := range reqHeader {
		req.Header.Set(key, value)
	}

	// 发送请求
	client := http.Client{}
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		log.Print("client do request faile: ", err)
		return "", err
	}

	// 读响应体
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("read resp body failed: ", err)
		return "", err
	}

	// 判断http状态码
	if resp.StatusCode != http.StatusOK {
		log.Print("http request error: ", string(respBody))
		return "", errors.New(strconv.Itoa(resp.StatusCode))
	}

	return string(respBody), nil
}

func GetJwtToken(c *gin.Context, auth2User *model.Auth2User) (*model.LoginResp, error) {
	// 手动注册用户
	auth2UserInfo, err := users.CreateUser(c, auth2User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "db error")
		return nil, err
	}

	// 检查第三方登录信息是否改变
	if err := users.UpdateUserAuth2Info(auth2UserInfo, auth2User); err != nil {
		return nil, err
	}

	// 生成jwt token
	mw := jwts.AuthMiddleware
	token := jwt.New(jwt.GetSigningMethod(mw.SigningAlgorithm))
	claims := token.Claims.(jwt.MapClaims)

	if mw.PayloadFunc != nil {
		data := &model.JwtUser{
			UserName: auth2User.UserName,
		}
		for key, value := range mw.PayloadFunc(data) {
			claims[key] = value
		}
	}

	expire := mw.TimeFunc().Add(mw.Timeout)
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = mw.TimeFunc().Unix()
	tokenString, err := token.SignedString(mw.Key)

	return &model.LoginResp{
		Token:  tokenString,
		Expire: expire.Format("2006-01-02 15:04:05"),
	}, nil
}

// 生成用户名和密码
func GenerateNamePwd() (string, string) {
	userName := uuid.New().String()
	passWord := base64.StdEncoding.EncodeToString([]byte(userName))

	return userName, passWord
}
