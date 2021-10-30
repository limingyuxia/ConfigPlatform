package users

import (
	"ConfigPlatform/api"
	"ConfigPlatform/conf"
	"ConfigPlatform/conf/mysql"
	"ConfigPlatform/model"
	"ConfigPlatform/models"
	"ConfigPlatform/services/weedo"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"path"
	"strconv"
	"strings"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

func GetUserNameFromJwtToken(c *gin.Context, authMiddleware *jwt.GinJWTMiddleware) (string, error) {
	// 中间层已经校验 err
	token, _ := authMiddleware.ParseToken(c)

	// 分割jwt token
	tokenParts := strings.Split(token.Raw, ".")
	// jwt 第二部分是自定义内容
	jwtTokenContent, err := base64.RawStdEncoding.DecodeString(tokenParts[1])
	if err != nil {
		log.Print("base decode jwt token failed: ", err)
		return "", err
	}

	// 解析jwt token
	var jwtToken model.JwtToken
	err = json.Unmarshal(jwtTokenContent, &jwtToken)
	if err != nil {
		log.Print("parse jwt token error: ", err)
		return "", err
	}

	return jwtToken.UserName, nil
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

func UpdateUserAvatar(ctx *gin.Context, username, fileName string, fileContent []byte) error {
	// 获取user id
	userId, err := models.Users(
		qm.Select("id"),
		qm.Where("username = ?", username),
	).One(ctx, mysql.Conn)
	if err != nil {
		log.Print("get user id failed: ", err)
		return err
	}

	// 生成uuid
	uuid := uuid.New()
	key := uuid.String()

	// 获取文件后缀
	filesuffix := path.Ext(fileName)
	// 生成新的文件名 filename_uuid.xxx
	fileNameUnique := fileName[0:len(fileName)-len(filesuffix)] + "_" + key + filesuffix

	// 查看该用户的文件映射
	seaweedfsinfo, err := models.Seaweedfs(
		qm.Select("id", "seaweedfs_url"),
		qm.Where("user_id = ?", userId.ID),
	).All(ctx, mysql.Conn)
	if err != nil {
		log.Print("get seaweedfs id failed: ", err)
		return err
	}

	// 连接服务
	seaweedfsServer := conf.SeaweedFsSetting.Domain + ":" + strconv.Itoa(conf.SeaweedFsSetting.ServerPort)
	seaweedfsVolume := conf.SeaweedFsSetting.Domain + ":" + strconv.Itoa(conf.SeaweedFsSetting.VolumePort)
	client := weedo.NewClient(seaweedfsServer, seaweedfsVolume)

	// 生成文件流
	bodyReader := bytes.NewBuffer(fileContent)

	// 本地服务文件流
	var serverUrl string = "/user/avatar/get/" + fileNameUnique

	if len(seaweedfsinfo) == 0 {
		fid, _, err := client.AssignUpload(fileNameUnique, "text/plain", bodyReader)
		if err != nil {
			log.Print("upload file error: ", err)
			// ResponseError(FILE_UPLOAD_ERROR, RETCODE_MSG[FILE_UPLOAD_ERROR], c)
			return err
		}

		// 记录该文件映射
		var seaweedfs = models.Seaweedf{
			UserID:       userId.ID,
			ServerURL:    serverUrl,
			SeaweedfsURL: fid,
		}

		err = seaweedfs.Insert(ctx, mysql.Conn, boil.Infer())
		if err != nil {
			log.Print("insert seaweedfs failed: ", err)
			return err
		}

	} else {
		// 更新文件
		err := client.Delete(seaweedfsinfo[0].SeaweedfsURL, 1)
		if err != nil {
			log.Print("delete seaweedfs failed: ", err)
			return err
		}

		fid, _, err := client.AssignUpload(fileNameUnique, "text/plain", bodyReader)
		if err != nil {
			log.Print("upload file error: ", err)
			// ResponseError(FILE_UPLOAD_ERROR, RETCODE_MSG[FILE_UPLOAD_ERROR], c)
			return err
		}

		// 更新 seaseedfs 映射
		seaweedfs, err := models.FindSeaweedf(ctx, mysql.Conn, seaweedfsinfo[0].ID)
		if err != nil {
			log.Print("find seaweedfs failed: ", err)
			return err
		}
		seaweedfs.SeaweedfsURL = fid
		seaweedfs.ServerURL = serverUrl

		_, err = seaweedfs.Update(ctx, mysql.Conn, boil.Infer())
		if err != nil {
			log.Print("update seaweedfs failed: ", err)
			return err
		}
	}

	// 更新用户头像url
	user, err := models.FindUser(ctx, mysql.Conn, userId.ID)
	if err != nil {
		log.Print("find user failed: ", err)
		return err
	}
	user.Photo = null.StringFrom(serverUrl)

	_, err = user.Update(ctx, mysql.Conn, boil.Infer())
	if err != nil {
		log.Print("update user avatar failed: ", err)
		return err
	}

	return nil
}

func GetUserAvatarFid(ctx context.Context, serverUrl string) (string, error) {
	seaweedfs, err := models.Seaweedfs(
		qm.Select("seaweedfs_url"),
		qm.Where("server_url = ?", serverUrl),
	).One(ctx, mysql.Conn)
	if err != nil {
		log.Print("get seaweedfs_url failed: ", err)
		return "", err
	}

	return seaweedfs.SeaweedfsURL, nil
}
