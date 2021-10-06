package project

import (
	"ConfigPlatform/conf/mysql"
	"ConfigPlatform/model"
	"strconv"

	"ConfigPlatform/models"
	"context"
	"log"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// 获取用户创建的项目总数
func GetProjectTotal(ctx context.Context, userName string) (int64, error) {
	total, err := models.Projects(
		qm.Where("project_user = ?", userName),
	).Count(ctx, mysql.Conn)

	if err != nil {
		log.Print("GetProjectTotal error: ", err)
		return 0, err
	}

	return total, nil
}

func GetProjectList(ctx context.Context, req *model.GetProjectListReq) (
	[]*model.ProjectInfo, error) {

	queryMod := []qm.QueryMod{}
	if req.ID != 0 {
		queryMod = append(queryMod, qm.Select("id = ?", strconv.Itoa(int(req.ID))))
	}
	if req.Name != "" {
		queryMod = append(queryMod, qm.Select("name = ?", req.Name))
	}
	queryMod = append(queryMod, qm.Offset(req.PageIndex), qm.Limit(req.PageSize))

	list, err := models.Projects(
		queryMod...,
	).All(ctx, mysql.Conn)
	if err != nil {
		log.Print("GetProjectList error: ", err)
		return nil, err
	}

	var projectList []*model.ProjectInfo
	for _, project := range list {
		projectList = append(projectList, &model.ProjectInfo{
			ID: int64(project.ID),
		})
	}

	return projectList, nil
}

func InsertProject() {

}

func UpdateProject() {

}

func DeleteProject() {

}
