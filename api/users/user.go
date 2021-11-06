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
	"database/sql"
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

func GetUserAuth2Info(ctx context.Context, username string) (*model.Auth2UserInfo, error) {
	auth2Info, err := models.Auth2s(
		qm.Select("qq_username", "qq_avatar", "wechat_username", "wechat_avatar",
			"weibo_username", "weibo_avatar", "github_username", "github_avatar"),
		qm.InnerJoin("user on user.id = auth2.user_id"),
		qm.Where("username = ?", username),
	).One(ctx, mysql.Conn)
	if err != nil {
		if err == sql.ErrNoRows {
			return &model.Auth2UserInfo{}, nil
		}

		log.Print("get user auth2 info failed: ", err)
		return nil, err
	}

	return &model.Auth2UserInfo{
		QQUsername:     auth2Info.QQUsername.String,
		QQAvatar:       auth2Info.QQAvatar.String,
		WechatUsername: auth2Info.WechatUsername.String,
		WechatAvatar:   auth2Info.WechatAvatar.String,
		WeiboUsername:  auth2Info.WeiboUsername.String,
		WeiboAvatar:    auth2Info.WeiboAvatar.String,
		GithubUsername: auth2Info.GithubUsername.String,
		GithubAvatar:   auth2Info.GithubAvatar.String,
	}, nil
}

func UpdateUserInfo(ctx context.Context, updateInfo model.UpdateUserInfo) error {

	userId, err := models.Users(
		qm.Select("id"),
		qm.Where("username = ?", updateInfo.UserName),
	).One(ctx, mysql.Conn)
	if err != nil {
		log.Print(ctx, "get user id failed: ", err)
		return err
	}

	user, err := models.FindUser(ctx, mysql.Conn, userId.ID)
	if err != nil {
		log.Print(ctx, "get user failed: ", err)
		return err
	}

	user.Username = updateInfo.UserName
	user.Nickname = null.StringFrom(updateInfo.Nickname)
	user.Gender = null.IntFrom(updateInfo.Gender)
	user.Region = null.StringFrom(updateInfo.Region)
	user.Phone = null.StringFrom(updateInfo.Phone)
	user.Email = null.StringFrom(updateInfo.Email)

	_, err = user.Update(ctx, mysql.Conn, boil.Infer())
	if err != nil {
		log.Print(ctx, "update user failed: ", err)
		return err
	}

	return nil
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

// 用户是否是第一次用第三方平台登录
func isAuth2UserRegister(ctx context.Context, user *model.Auth2User) (bool, error) {
	queryMod := []qm.QueryMod{
		qm.Select("id", "user_id"),
	}

	switch user.Type {
	case "qq":
		qqOpenId := null.StringFrom(user.UniqueId)
		queryMod = append(queryMod, qm.Where("qq_openid = ?", qqOpenId))
	case "wechat":
		wechatOpenid := null.StringFrom(user.UniqueId)
		queryMod = append(queryMod, qm.Where("wechat_openid = ?", wechatOpenid))
	case "weibo":
		uId := null.StringFrom(user.UniqueId)
		queryMod = append(queryMod, qm.Where("uid = ?", uId))
	case "github":
		githubId := null.StringFrom(user.UniqueId)
		queryMod = append(queryMod, qm.Where("github_id = ?", githubId))
	default:
		return true, errors.New("unknow type")
	}

	auth2User, err := models.Auth2s(
		queryMod...,
	).One(ctx, mysql.Conn)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return true, err
	}

	// 检查用户的auth2信息是否变更
	err = UpdateUserAuth2Info(ctx, auth2User, user)
	if err != nil {
		return true, err
	}

	// 获取原来的用户名
	userName, err := models.Users(
		qm.Select("username"),
		qm.Where("id = ?", auth2User.UserID),
	).One(ctx, mysql.Conn)
	if err != nil {
		return true, err
	}

	user.UserName = userName.Username
	return true, nil
}

// 更新用户auth2信息
func UpdateUserAuth2Info(ctx context.Context, auth2UserInfo *models.Auth2,
	auth2User *model.Auth2User) error {

	updateAuth2Info, err := models.FindAuth2(ctx, mysql.Conn, auth2UserInfo.ID)
	if err != nil {
		return err
	}

	switch auth2User.Type {
	case "qq":
		if auth2UserInfo.QQUsername.String != auth2User.Nickname {
			updateAuth2Info.QQUsername = null.StringFrom(auth2User.Nickname)
		}
		if auth2UserInfo.QQAvatar.String != auth2User.Avatar {
			updateAuth2Info.QQAvatar = null.StringFrom(auth2User.Avatar)
		}
	case "wechat":
		if auth2UserInfo.WechatUsername.String != auth2User.Nickname {
			updateAuth2Info.WechatUsername = null.StringFrom(auth2User.Nickname)
		}
		if auth2UserInfo.WechatAvatar.String != auth2User.Avatar {
			updateAuth2Info.WechatAvatar = null.StringFrom(auth2User.Avatar)
		}
	case "weibo":
		if auth2UserInfo.WeiboUsername.String != auth2User.Nickname {
			updateAuth2Info.WeiboUsername = null.StringFrom(auth2User.Nickname)
		}
		if auth2UserInfo.WeiboAvatar.String != auth2User.Avatar {
			updateAuth2Info.WeiboAvatar = null.StringFrom(auth2User.Avatar)
		}
	case "github":
		if auth2UserInfo.GithubUsername.String != auth2User.Nickname {
			updateAuth2Info.GithubUsername = null.StringFrom(auth2User.Nickname)
		}
		if auth2UserInfo.GithubAvatar.String != auth2User.Avatar {
			updateAuth2Info.GithubAvatar = null.StringFrom(auth2User.Avatar)
		}
	default:
		return errors.New("unknow type")
	}

	_, err = updateAuth2Info.Update(ctx, mysql.Conn, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

// 新增用户
func CreateUser(ctx context.Context, user *model.Auth2User) error {

	isRegister, err := isAuth2UserRegister(ctx, user)
	if err != nil {
		log.Print(ctx, "check user register failed: ", err)
		return err
	}

	// 第一次第三方登录
	if !isRegister {
		var newUser = &models.User{
			Username: user.UserName,
			Password: user.Password,
		}

		err := newUser.Insert(ctx, mysql.Conn, boil.Infer())
		if err != nil {
			log.Print("create user failed: ", err)
			return err
		}

		var newUserAuth = &models.Auth2{
			UserID: newUser.ID,
		}

		switch user.Type {
		case "qq":
			newUserAuth.QQOpenid = null.StringFrom(user.UniqueId)
			newUserAuth.QQUsername = null.StringFrom(user.Nickname)
			newUserAuth.QQAvatar = null.StringFrom(user.Avatar)
		case "wechat":
			newUserAuth.WechatOpenid = null.StringFrom(user.UniqueId)
			newUserAuth.WechatUsername = null.StringFrom(user.Nickname)
			newUserAuth.WechatAvatar = null.StringFrom(user.Avatar)
		case "weibo":
			newUserAuth.UID = null.StringFrom(user.UniqueId)
			newUserAuth.WeiboUsername = null.StringFrom(user.Nickname)
			newUserAuth.WeiboAvatar = null.StringFrom(user.Avatar)
		case "github":
			newUserAuth.GithubID = null.StringFrom(user.UniqueId)
			newUserAuth.GithubUsername = null.StringFrom(user.Nickname)
			newUserAuth.GithubAvatar = null.StringFrom(user.Avatar)
		}

		err = newUserAuth.Insert(ctx, mysql.Conn, boil.Infer())
		if err != nil {
			log.Print("create user auth2 failed: ", err)
			return err
		}
	}

	return nil
}
