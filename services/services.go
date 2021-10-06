package services

import (
	"ConfigPlatform/api/project"
	"ConfigPlatform/model"

	"github.com/gin-gonic/gin"
)

// @Tags 项目
// @Summary 获取项目列表
// @Accept  json
// @Produce  json
// @Param GetProjectList query model.GetProjectListReq true "获取用户项目列表"
// @Success 200 {object} model.GetProjectListResp
// @Failure 400 {object} WebResponse
// @Router /project/list [get]
func GetProjectList(c *gin.Context) {

	var req model.GetProjectListReq
	if err := c.ShouldBindQuery(&req); err == nil {

		total, err := project.GetProjectTotal(c, &req)
		if err != nil {
			ResponseError(DB_ERROR, err.Error(), c)
			return
		}

		list, err := project.GetProjectList(c, &req)
		if err != nil {
			ResponseError(DB_ERROR, err.Error(), c)
			return
		}

		ResponseData(model.GetProjectListResp{
			Total:       total,
			ProjectList: list,
		}, c)

	} else {
		ResponseError(PARAMS_ERROR, err.Error(), c)
	}
}

// @Tags 项目
// @Summary 获取项目详情
// @Accept  json
// @Produce  json
// @Param project query model.GetProjectDetailReq true "项目"
// @Success 200 {object} model.ProjectInfo
// @Failure 400 {object} WebResponse
// @Router /project/detail [get]
func GetProjectDetail(c *gin.Context) {

	var req model.GetProjectDetailReq
	if err := c.ShouldBindQuery(&req); err == nil {

		projectDetail, err := project.GetProjectDetail(c, req)
		if err != nil {
			ResponseError(DB_ERROR, err.Error(), c)
			return
		}

		ResponseData(projectDetail, c)

	} else {
		ResponseError(PARAMS_ERROR, err.Error(), c)
	}
}
