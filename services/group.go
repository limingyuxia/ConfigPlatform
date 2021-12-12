package services

import (
	"ConfigPlatform/api/group"
	"ConfigPlatform/model"
	"log"

	"github.com/gin-gonic/gin"
)

// @Tags 配置分组
// @Summary 获取配置分组列表
// @Accept  json
// @Produce  json
// @Param GetConfigGroupList query model.GetConfigGroupReq true "请求列表参数"
// @Success 200 {object} model.GetConfigGroupResp
// @Failure 400 {object} WebResponse
// @Router /group/list [get]
func GetConfigGroupList(c *gin.Context) {
	var req model.GetConfigGroupReq
	if err := c.ShouldBindQuery(&req); err == nil {

		configGroups, err := group.GetConfigGroupList(c, &req)
		if err != nil {
			ResponseError(GROUP_LIST_ERROR, RETCODE_MSG[GROUP_LIST_ERROR], c)
			return
		}

		ResponseData(configGroups, c)
	} else {
		log.Print("GetConfigGroupList param error: ", err)
		ResponseError(PARAMS_ERROR, RETCODE_MSG[PARAMS_ERROR], c)
	}
}

// @Tags 配置分组
// @Summary 添加配置分组
// @Accept  json
// @Produce  json
// @Param AddConfigGroup body model.AddConfigGroupReq true "分组数据"
// @Success 200 {string} string
// @Failure 400 {object} WebResponse
// @Router /group/add [post]
func AddConfigGroup(c *gin.Context) {
	var req model.AddConfigGroupReq

	if err := c.ShouldBindJSON(&req); err == nil {
		Resp, err := group.AddConfigGroup(c, &req)
		if err != nil {
			log.Print(err)
			ResponseError(GROUP_ADD_ERROR, RETCODE_MSG[GROUP_ADD_ERROR], c)
			return
		}

		ResponseData(Resp, c)
	} else {
		log.Print("AddConfigGroup param error: ", err)
		ResponseError(PARAMS_ERROR, RETCODE_MSG[PARAMS_ERROR], c)
	}
}

// @Tags 配置分组
// @Summary 编辑配置分组
// @Accept  json
// @Produce  json
// @Param EditConfigGroup body model.EditConfigGroupReq true "配置分组修改项"
// @Success 200 {string} string
// @Failure 400 {object} WebResponse
// @Router /group/edit [post]
func EditConfigGroup(c *gin.Context) {
	var req model.EditConfigGroupReq

	if err := c.ShouldBindJSON(&req); err == nil {
		Resp, err := group.EditConfigGroup(c, &req)
		if err != nil {
			ResponseError(GROUP_EDIT_ERROR, RETCODE_MSG[GROUP_EDIT_ERROR], c)
			return
		}

		ResponseData(Resp, c)

	} else {
		log.Print("EditConfigGroup param error: ", err)
		ResponseError(PARAMS_ERROR, RETCODE_MSG[PARAMS_ERROR], c)
	}
}

// @Tags 配置分组
// @Summary 删除配置分组
// @Accept  json
// @Produce  json
// @Param DeleteConfigGroup query model.DeleteConfigGroupReq true "删除配置分组信息"
// @Success 200 {string} string
// @Failure 400 {object} WebResponse
// @Router /group/delete [delete]
func DeleteConfigGroup(c *gin.Context) {
	var req model.DeleteConfigGroupReq

	if err := c.ShouldBindQuery(&req); err == nil {
		Resp, err := group.DeleteConfigGroup(c, &req)
		if err != nil {
			ResponseError(GROUP_DELETE_ERROR, RETCODE_MSG[GROUP_DELETE_ERROR], c)
			return
		}

		ResponseData(Resp, c)
	} else {
		log.Print("DeleteConfigGroup param error: ", err)
		ResponseError(PARAMS_ERROR, RETCODE_MSG[PARAMS_ERROR], c)
	}
}
