package group

import (
	"ConfigPlatform/api/users"
	"ConfigPlatform/conf/mysql"
	"ConfigPlatform/model"
	"ConfigPlatform/models"
	"ConfigPlatform/routes/middleware/jwts"
	"context"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// 查看用户是否有权限编辑分组
func userConfigGroupAndProject(ctx context.Context, projectId int, username string) bool {
	exist, err := models.Projects(
		qm.Where("id = ?", projectId),
		qm.And("(admin like ?", username),
		qm.Or("project_user like ?", username),
		qm.Or("develop_user like ?)", username),
	).Exists(ctx, mysql.Conn)
	if err != nil {
		log.Print("query project error")
		return false
	}

	return exist
}

// 获取配置分组列表
func GetConfigGroupList(ctx *gin.Context, req *model.GetConfigGroupReq) (*model.GetConfigGroupResp, error) {
	// 获取token中的用户名
	username, err := users.GetUserNameFromJwtToken(ctx, jwts.AuthMiddleware)
	if err != nil {
		log.Print("get username from jwt token failed")
		return nil, errors.New("jwt token error")
	}

	if !userConfigGroupAndProject(ctx, req.ProjectId, username) {
		return nil, errors.New("user don't have the project permission")
	}

	configGroups, err := models.ConfigGroups(
		qm.Select(`id`, `group_name`, `comment`),
		qm.Where("project_id = ?", req.ProjectId),
		qm.Where("env = ?", req.Env),
	).All(ctx, mysql.Conn)
	if err != nil {
		log.Print("query config group failed: ", err.Error())
		return nil, err
	}

	var respData []model.ConfigGroup
	for _, configGroup := range configGroups {
		respData = append(respData, model.ConfigGroup{
			Id:        configGroup.ID,
			GroupName: configGroup.GroupName,
			Comment:   configGroup.Comment.String,
		})
	}

	return &model.GetConfigGroupResp{
		ConfigGroupList: respData,
	}, nil
}

// 新增分组列表
func AddConfigGroup(ctx *gin.Context, req *model.AddConfigGroupReq) (*model.AddConfigGroupResp, error) {
	// 获取token中的用户名
	username, err := users.GetUserNameFromJwtToken(ctx, jwts.AuthMiddleware)
	if err != nil {
		log.Print("get username from jwt token failed")
		return nil, errors.New("jwt token error")
	}

	if !userConfigGroupAndProject(ctx, req.ProjectId, username) {
		return nil, errors.New("user don't have the project permission")
	}

	if req.Env != "Dev" && req.Env != "Test" &&
		req.Env != "Pre" && req.Env != "Prod" {
		return nil, errors.New("env is unknown")
	}

	var configGroup = models.ConfigGroup{
		ProjectID:  req.ProjectId,
		Env:        req.Env,
		GroupName:  req.GroupName,
		Comment:    null.StringFrom(req.Comment),
		CreateUser: username,
		UpdateUser: username,
	}

	err = configGroup.Insert(ctx, mysql.Conn, boil.Infer())
	if err != nil {
		log.Print("insert config_platform error")
		return nil, err
	}

	return &model.AddConfigGroupResp{
		Message: "成功",
	}, nil
}

func EditConfigGroup(ctx *gin.Context, req *model.EditConfigGroupReq) (*model.EditConfigGroupResp, error) {
	// 获取token中的用户名
	username, err := users.GetUserNameFromJwtToken(ctx, jwts.AuthMiddleware)
	if err != nil {
		log.Print("get username from jwt token failed")
		return nil, errors.New("jwt token error")
	}

	if !userConfigGroupAndProject(ctx, req.ProjectId, username) {
		return nil, errors.New("user don't have the project permission")
	}

	configGroup, err := models.FindConfigGroup(ctx, mysql.Conn, req.Id)
	if err != nil {
		log.Print("find config_platform error")
		return nil, err
	}
	configGroup.GroupName = req.GroupName
	configGroup.Comment = null.StringFrom(req.Comment)
	configGroup.UpdateUser = username

	_, err = configGroup.Update(ctx, mysql.Conn, boil.Infer())
	if err != nil {
		log.Print("update config_platform error")
		return nil, err
	}

	return &model.EditConfigGroupResp{
		Message: "成功",
	}, nil
}

func DeleteConfigGroup(ctx *gin.Context, req *model.DeleteConfigGroupReq) (*model.DeleteConfigGroupResp, error) {
	// 获取token中的用户名
	username, err := users.GetUserNameFromJwtToken(ctx, jwts.AuthMiddleware)
	if err != nil {
		log.Print("get username from jwt token failed")
		return nil, errors.New("jwt token error")
	}

	if !userConfigGroupAndProject(ctx, req.ProjectId, username) {
		return nil, errors.New("user don't have the project permission")
	}

	configGroup, err := models.FindConfigGroup(ctx, mysql.Conn, req.Id)
	if err != nil {
		log.Print("find config_platform error")
		return nil, err
	}

	_, err = configGroup.Delete(ctx, mysql.Conn)
	if err != nil {
		log.Print("delete config_platform error")
		return nil, err
	}

	return &model.DeleteConfigGroupResp{
		Message: "成功",
	}, nil
}
