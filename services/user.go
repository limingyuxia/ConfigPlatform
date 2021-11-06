package services

import (
	"ConfigPlatform/api/users"
	"ConfigPlatform/conf"
	"ConfigPlatform/model"
	"ConfigPlatform/routes/middleware/jwts"
	"ConfigPlatform/services/auth2"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Tags 用户管理
// @Summary 注册账号
// @Accept  json
// @Produce  json
// @Param Register body model.Register true "注册用户的信息"
// @Success 200 {string} string
// @Failure 400 {object} WebResponse
// @Router /user/register [post]
func Register(c *gin.Context) {
	var register model.Register

	if err := c.ShouldBindJSON(&register); err == nil {

		userId, err := users.UserRegisted(c, &register)
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

// @Tags 用户管理
// @Summary 上传用户头像
// @Accept  json
// @Produce  json
// @Param UpdateAvatar body string true "multipart/form-data加密后的文件内容"
// @Success 200 {string} string
// @Failure 400 {object} WebResponse
// @Router /user/avatar/update [post]
func UpdateAvatar(c *gin.Context) {
	contentType := c.GetHeader("Content-Type")
	if contentType == "" {
		ResponseError(FILE_UPLOAD_H_ERROR, RETCODE_MSG[FILE_UPLOAD_H_ERROR], c)
		return
	}

	contentTypeValue, params, err := mime.ParseMediaType(contentType)
	if err != nil || contentTypeValue != "multipart/form-data" {
		ResponseError(FILE_UPLOAD_H_ERROR, RETCODE_MSG[FILE_UPLOAD_H_ERROR], c)
		return
	}

	boundary, ok := params["boundary"]
	if !ok {
		ResponseError(FILE_UPLOAD_H_ERROR, RETCODE_MSG[FILE_UPLOAD_H_ERROR], c)
		return
	}

	reader := multipart.NewReader(c.Request.Body, boundary)
	// 文件内容
	var fileContent = []byte{}
	// 文件名称
	var fileName string

	for {
		part, err := reader.NextPart()
		if err != nil {
			log.Print(err)
			break
		}

		fileContentPart, err := ioutil.ReadAll(part)
		if err != nil {
			log.Print(err)
		}
		fileContent = append(fileContent, fileContentPart...)

		fileName = part.FileName()
	}

	// 获取token中的用户名
	username, err := users.GetUserNameFromJwtToken(c, jwts.AuthMiddleware)
	if err != nil {
		ResponseError(TOKEN_FORMAT_ERROR, RETCODE_MSG[TOKEN_FORMAT_ERROR], c)
		return
	}

	// 更新用户头像
	err = users.UpdateUserAvatar(c, username, fileName, fileContent)
	if err != nil {
		ResponseError(AVATAR_UPDATE_ERROR, RETCODE_MSG[AVATAR_UPDATE_ERROR], c)
		return
	}

	ResponseData("更新成功", c)
}

// @Tags 用户管理
// @Summary 获取用户信息
// @Produce  json
// @Success 200 {object} model.UserInfo
// @Failure 400 {object} WebResponse
// @Router /user/info [get]
func GetUserInfo(c *gin.Context) {
	// 获取token中的用户名
	username, err := users.GetUserNameFromJwtToken(c, jwts.AuthMiddleware)
	if err != nil {
		ResponseError(TOKEN_FORMAT_ERROR, RETCODE_MSG[TOKEN_FORMAT_ERROR], c)
		return
	}

	// 获取用户信息
	userInfo, err := users.GetUserInfo(c, username)
	if err != nil {
		ResponseError(GET_USER_INFO_ERROR, RETCODE_MSG[GET_USER_INFO_ERROR], c)
		return
	}

	ResponseData(userInfo, c)
}

func GetUserAvatar(c *gin.Context) {
	fid, err := users.GetUserAvatarFid(c, c.Request.URL.Path)
	if err != nil {
		ResponseError(GET_USER_INFO_ERROR, RETCODE_MSG[GET_USER_INFO_ERROR], c)
		return
	}

	host := conf.SeaweedFsSetting.Domain + ":" + strconv.Itoa(conf.SeaweedFsSetting.VolumePort)

	data, err := auth2.HttpGet(c, "http://"+host+"/"+fid, nil, nil)
	if err != nil {
		ResponseError(USER_AVATAR_ERROR, RETCODE_MSG[USER_AVATAR_ERROR], c)
		return
	}

	c.Data(200, "string", []byte(data))
}

// @Tags 用户管理
// @Summary 获取用户第三方应用信息
// @Produce  json
// @Success 200 {object} model.Auth2UserInfo
// @Failure 400 {object} WebResponse
// @Router /user/auth2Info [get]
func GetUserAuth2Info(c *gin.Context) {
	// 获取token中的用户名
	username, err := users.GetUserNameFromJwtToken(c, jwts.AuthMiddleware)
	if err != nil {
		ResponseError(TOKEN_FORMAT_ERROR, RETCODE_MSG[TOKEN_FORMAT_ERROR], c)
		return
	}

	// 获取用户第三方应用信息
	auth2Info, err := users.GetUserAuth2Info(c, username)
	if err != nil {
		ResponseError(GET_USER_INFO_ERROR, RETCODE_MSG[GET_USER_INFO_ERROR], c)
		return
	}

	ResponseData(auth2Info, c)
}
