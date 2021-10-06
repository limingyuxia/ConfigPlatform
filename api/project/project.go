package project

import (
	"ConfigPlatform/conf/mysql"
	"ConfigPlatform/model"
	"strconv"
	"strings"
	"time"

	"ConfigPlatform/models"
	"context"
	"log"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func getQueryCond(req *model.GetProjectListReq) []qm.QueryMod {
	queryMod := []qm.QueryMod{
		qm.Where("project_user = ?", req.ProjectUser),
	}

	if req.ID != 0 {
		queryMod = append(queryMod, qm.Where("id = ?", strconv.Itoa(req.ID)))
	}
	if req.Name != "" {
		queryMod = append(queryMod, qm.Where("name = ?", req.Name))
	}
	if req.Department != "" {
		queryMod = append(queryMod, qm.Where("department = ?", req.Department))
	}
	if req.Admin != "" {
		queryMod = append(queryMod, qm.Where("admin = ?", req.Admin))
	}
	if req.DevelopUser != "" {
		queryMod = append(queryMod, qm.Where("develop_user = ?", req.DevelopUser))
	}

	return queryMod
}

// 获取用户创建的项目总数
func GetProjectTotal(ctx context.Context, req *model.GetProjectListReq) (int64, error) {

	queryMod := getQueryCond(req)

	total, err := models.Projects(
		queryMod...,
	).Count(ctx, mysql.Conn)
	if err != nil {
		log.Print("GetProjectTotal error: ", err)
		return 0, err
	}

	return total, nil
}

func GetProjectList(ctx context.Context, req *model.GetProjectListReq) (
	[]*model.ProjectInfo, error) {

	queryMod := getQueryCond(req)
	offset := (req.PageIndex - 1) * req.PageSize
	queryMod = append(queryMod, qm.Limit(req.PageSize), qm.Offset(offset))

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
			ID:          project.ID,
			Name:        project.Name,
			Description: project.Description.String,
			Department:  strings.Split(project.Department.String, "/"),
			Admin:       strings.Split(project.Admin.String, ","),
			ProjectUser: project.ProjectUser,
			DevelopUser: strings.Split(project.DevelopUser.String, ","),
			CreateTime:  project.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:  project.UpdateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return projectList, nil
}

func GetProjectDetail(ctx context.Context, req *model.GetProjectDetailReq) (*model.ProjectInfo, error) {

	queryMod := []qm.QueryMod{
		qm.Where("id = ?", req.ID),
	}
	if req.Name != "" {
		queryMod = append(queryMod, qm.Where("name = ?", req.Name))
	}

	project, err := models.Projects(
		queryMod...,
	).One(ctx, mysql.Conn)
	if err != nil {
		log.Print("GetProjectDetail error: ", err)
		return nil, err
	}

	return &model.ProjectInfo{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description.String,
		Department:  strings.Split(project.Department.String, "/"),
		Admin:       strings.Split(project.Admin.String, ","),
		ProjectUser: project.ProjectUser,
		DevelopUser: strings.Split(project.DevelopUser.String, ","),
		CreateTime:  project.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:  project.UpdateTime.Format("2006-01-02 15:04:05"),
	}, nil
}

func InsertProject(ctx context.Context, req *model.AddProject) error {

	insertData := models.Project{
		Name:        req.Name,
		Description: null.StringFrom(req.Description),
		Department:  null.StringFrom(strings.Join(req.Department, "/")),
		Admin:       null.StringFrom(strings.Join(req.Admin, ",")),
		ProjectUser: req.ProjectUser,
		DevelopUser: null.StringFrom(strings.Join(req.DevelopUser, ",")),
	}

	err := insertData.Insert(ctx, mysql.Conn, boil.Infer())
	if err != nil {
		log.Print("InsertProject error: ", err)
		return err
	}

	return nil
}

func UpdateProject(ctx context.Context, req *model.EditProject) error {
	project, err := models.FindProject(ctx, mysql.Conn, req.ID)
	if err != nil {
		log.Print("FindProject error: ", err)
		return err
	}

	project.Description = null.StringFrom(req.Description)
	project.Department = null.StringFrom(strings.Join(req.Department, "/"))
	project.Admin = null.StringFrom(strings.Join(req.Admin, ","))
	project.DevelopUser = null.StringFrom(strings.Join(req.DevelopUser, ","))
	project.UpdateTime = time.Now()

	_, err = project.Update(ctx, mysql.Conn, boil.Infer())
	if err != nil {
		log.Print("UpdateProject error: ", err)
		return err
	}

	return nil
}

func DeleteProject(ctx context.Context, req *model.DeleteProject) error {
	project, err := models.FindProject(ctx, mysql.Conn, req.ID)
	if err != nil {
		log.Print("FindProject error: ", err)
		return err
	}

	_, err = project.Delete(ctx, mysql.Conn)
	if err != nil {
		log.Print("DeleteProject error: ", err)
		return err
	}

	return nil
}
