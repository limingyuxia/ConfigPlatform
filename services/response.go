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
func ResponseError(errcode int, errmsg string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"errcode": errcode,
		"errmsg":  errmsg,
	})
}
