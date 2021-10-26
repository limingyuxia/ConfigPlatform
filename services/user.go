package services

import (
	"ConfigPlatform/api/users"
	"ConfigPlatform/conf"
	"ConfigPlatform/model"
	"ConfigPlatform/services/weedo"
	"bytes"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"path"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	// 生成uuid
	uuid := uuid.New()
	key := uuid.String()

	// 获取文件后缀
	filesuffix := path.Ext(fileName)
	// 生成新的文件名
	fileNameUnique := fileName[0:len(fileName)-len(filesuffix)] + "_" + key + filesuffix

	// 连接服务
	seaweedfsServer := conf.SeaweedFsSetting.Domain + ":" + strconv.Itoa(conf.SeaweedFsSetting.ServerPort)
	seaweedfsVolume := conf.SeaweedFsSetting.Domain + ":" + strconv.Itoa(conf.SeaweedFsSetting.VolumePort)

	client := weedo.NewClient(seaweedfsServer, seaweedfsVolume)

	bodyReader := bytes.NewBuffer(fileContent)
	fid, _, err := client.AssignUpload(fileNameUnique, "text/plain", bodyReader)
	if err != nil {
		log.Print("upload file error: ", err)
		ResponseError(FILE_UPLOAD_ERROR, RETCODE_MSG[FILE_UPLOAD_ERROR], c)
		return
	}

	fid = strings.Replace(fid, ",", "/", 1)

	ResponseData("http://"+seaweedfsVolume+"/"+fid, c)
}
