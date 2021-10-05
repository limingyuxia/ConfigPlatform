package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebResponse struct {
	ErrCode int    `json:"errcode"` // 错误码
	ErrMsg  string `json:"errmsg"`  // 错误信息
}

// 错误返回数据
func ResponseError(errcode RETCODE, errmsg string, c *gin.Context) {
	// http 返回的状态码，成功200，错误400
	c.JSON(http.StatusBadRequest, gin.H{
		"errcode": errcode,
		"errmsg":  errmsg,
	})
}

// 正常返回数据
func ResponseData(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, data)
}
