package users

import (
	"ConfigPlatform/api"
	"ConfigPlatform/conf/mysql"
	"ConfigPlatform/model"
	"ConfigPlatform/models"
	"context"
	"errors"
	"log"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func IsUserRegisted(ctx context.Context, userName, passWord string) error {
	querySql := "select id from user where `username` = ? and `password` = ?"

	user, err := mysql.Conn.QueryContext(ctx, querySql, userName, passWord)
	if err != nil {
		log.Print("IsUserRegisted error: ", err)
		return err
	}
	defer user.Close()

	if !user.Next() {
		log.Print("username or password error")
		return errors.New("username or password error")
	}

	log.Printf("%s login in", userName)

	return nil
}

func UserRegisted(ctx context.Context, register *model.Register) (int64, error) {

	insertSql := `insert into user(username, password, email) VALUES(?, ?, ?)`

	result, err := mysql.Conn.ExecContext(ctx, insertSql, register.Username,
		register.Password, register.EmailAddress)
	if err != nil {
		log.Print("UserRegisted error: ", err)
		return 0, api.GetDbError(err)
	}

	userId, err := result.LastInsertId()
	if err != nil {
		log.Print(err)
		return 0, api.GetDbError(err)
	}

	return userId, nil
}

func GetUserInfo(ctx context.Context, username string) (*model.UserInfo, error) {
	// 用户基本信息
	userInfo, err := models.Users(
		qm.Select("nickname", "gender", "region", "phone", "email", "photo", "create_time"),
		qm.Where("username = ?", username),
	).One(ctx, mysql.Conn)
	if err != nil {
		log.Print("get user info failed: ", err)
		return nil, err
	}

	// 用户创建的项目
	createProject, err := models.Projects(
		qm.Select("name"),
		qm.Where("project_user = ?", username),
	).All(ctx, mysql.Conn)
	if err != nil {
		log.Print("get user create project failed: ", err)
		return nil, err
	}

	var CProjectName = []string{}
	for _, project := range createProject {
		CProjectName = append(CProjectName, project.Name)
	}

	// 用户参与开发的项目
	developProject, err := models.Projects(
		qm.Select("name"),
		qm.Where("develop_user like ?", username),
	).All(ctx, mysql.Conn)
	if err != nil {
		log.Print("get user develop project failed: ", err)
		return nil, err
	}

	var DProjectName = []string{}
	for _, project := range developProject {
		DProjectName = append(DProjectName, project.Name)
	}

	// 用户管理的项目
	manageProject, err := models.Projects(
		qm.Select("name"),
		qm.Where("admin like ?", username),
	).All(ctx, mysql.Conn)
	if err != nil {
		log.Print("get user manage project failed: ", err)
		return nil, err
	}

	var MProjectName = []string{}
	for _, project := range manageProject {
		MProjectName = append(MProjectName, project.Name)
	}

	var userInfos = &model.UserInfo{
		UserName:       username,
		Nickname:       userInfo.Nickname.String,
		Gender:         userInfo.Gender.Int,
		Region:         userInfo.Region.String,
		Phone:          userInfo.Phone.String,
		Email:          userInfo.Email.String,
		Photo:          userInfo.Photo.String,
		CreateTime:     userInfo.CreateTime.Format("2006-01-02 15:04:05"),
		CreateProject:  CProjectName,
		ManageProject:  MProjectName,
		DevelopProject: DProjectName,
	}

	return userInfos, nil
}
