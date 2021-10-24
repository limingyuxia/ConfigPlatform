package auth2

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
