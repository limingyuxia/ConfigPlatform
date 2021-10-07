package services

import (
	"ConfigPlatform/api/project"
	"ConfigPlatform/model"
	"log"

	"github.com/gin-gonic/gin"
)

// @Tags 项目
// @Summary 获取项目列表
// @Accept  json
// @Produce  json
// @Param GetProjectList query model.GetProjectListReq true "请求列表参数"
// @Success 200 {object} model.GetProjectListResp
// @Failure 400 {object} WebResponse
// @Router /project/list [get]
func GetProjectList(c *gin.Context) {

	var req model.GetProjectListReq
	if err := c.ShouldBindQuery(&req); err == nil {

		total, err := project.GetProjectTotal(c, &req)
		if err != nil {
			ResponseError(PROJ_TOTAL_ERROR, RETCODE_MSG[PROJ_TOTAL_ERROR], c)
			return
		}

		list, err := project.GetProjectList(c, &req)
		if err != nil {
			ResponseError(PROJ_LIST_ERROR, RETCODE_MSG[PROJ_LIST_ERROR], c)
			return
		}

		ResponseData(model.GetProjectListResp{
			Total:       total,
			ProjectList: list,
		}, c)

	} else {
		log.Print("GetProjectList param error: ", err)
		ResponseError(PARAMS_ERROR, RETCODE_MSG[PARAMS_ERROR], c)
	}
}

// @Tags 项目
// @Summary 获取项目详情
// @Accept  json
// @Produce  json
// @Param project query model.GetProjectDetailReq true "项目id或名称"
// @Success 200 {object} model.ProjectInfo
// @Failure 400 {object} WebResponse
// @Router /project/detail [get]
func GetProjectDetail(c *gin.Context) {

	var req model.GetProjectDetailReq
	if err := c.ShouldBindQuery(&req); err == nil {

		projectDetail, err := project.GetProjectDetail(c, &req)
		if err != nil {
			ResponseError(PROJ_DETAIL_ERROR, RETCODE_MSG[PROJ_DETAIL_ERROR], c)
			return
		}

		ResponseData(projectDetail, c)

	} else {
		log.Print("GetProjectDetail param failed: ", err)
		ResponseError(PARAMS_ERROR, RETCODE_MSG[PARAMS_ERROR], c)
	}
}

// @Tags 项目
// @Summary 添加项目
// @Accept  json
// @Produce  json
// @Param project body model.AddProject true "添加项目数据"
// @Success 200 {string} string
// @Failure 400 {object} WebResponse
// @Router /project/add [post]
func AddProject(c *gin.Context) {

	var req model.AddProject
	if err := c.ShouldBindJSON(&req); err == nil {

		if err := project.InsertProject(c, &req); err != nil {
			ResponseError(PROJ_ADD_ERROR, RETCODE_MSG[PROJ_ADD_ERROR], c)
			return
		}

		ResponseData("新建成功", c)

	} else {
		log.Print("AddProject param error: ", err)
		ResponseError(PARAMS_ERROR, RETCODE_MSG[PARAMS_ERROR], c)
	}
}

// @Tags 项目
// @Summary 编辑项目
// @Accept  json
// @Produce  json
// @Param project body model.EditProject true "编辑项目数据"
// @Success 200 {string} string
// @Failure 400 {object} WebResponse
// @Router /project/edit [post]
func EditProject(c *gin.Context) {

	var req model.EditProject
	if err := c.ShouldBindJSON(&req); err == nil {

		if err := project.UpdateProject(c, &req); err != nil {
			ResponseError(PROJ_EDIT_ERROR, RETCODE_MSG[PROJ_EDIT_ERROR], c)
			return
		}

		ResponseData("更新成功", c)

	} else {
		log.Print("EditProject param error: ", err)
		ResponseError(PARAMS_ERROR, RETCODE_MSG[PARAMS_ERROR], c)
	}
}

// @Tags 项目
// @Summary 删除项目
// @Accept  json
// @Produce  json
// @Param project query model.DeleteProject true "项目id和创建人"
// @Success 200 {string} string
// @Failure 400 {object} WebResponse
// @Router /project/delete [delete]
func DeleteProject(c *gin.Context) {

	var req model.DeleteProject
	if err := c.ShouldBindQuery(&req); err == nil {

		if err := project.DeleteProject(c, &req); err != nil {
			ResponseError(PROJ_DEL_ERROR, RETCODE_MSG[PROJ_DEL_ERROR], c)
			return
		}

		ResponseData("删除成功", c)

	} else {
		ResponseError(PARAMS_ERROR, RETCODE_MSG[PARAMS_ERROR], c)
	}
}
