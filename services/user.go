package services

import (
	"ConfigPlatform/api/users"
	"ConfigPlatform/model"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Tags 用户管理
// @Summary 注册账号
// @Accept  json
// @Produce  json
// @Param Register body model.LoginReq true "注册用户的账号密码"
// @Success 200 {string} model.LoginResp
// @Failure 400 {object} WebResponse
// @Router /register [post]
func Register(c *gin.Context) {
	var register model.LoginReq

	if err := c.ShouldBindJSON(&register); err == nil {

		userId, err := users.UserRegisted(c, register.Username, register.Password)
		if err != nil {
			ResponseError(REGISTER_ERROR, RETCODE_MSG[REGISTER_ERROR], c)
			return
		}

		ResponseData("注册成功, id: "+strconv.Itoa(int(userId)), c)

	} else {
		log.Print("register param error: ", err)
		ResponseError(PARAMS_ERROR, RETCODE_MSG[PARAMS_ERROR], c)
	}
}
