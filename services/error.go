package services

type RETCODE int32

// 错误码从4000开始
const (
	SUCCESS             RETCODE = 2000
	PARAMS_ERROR        RETCODE = 4000
	DB_ERROR            RETCODE = 4001
	NOAUTH_ERROR        RETCODE = 4002
	EXPIRE_ERROR        RETCODE = 4003
	TOKEN_ERROR         RETCODE = 4004
	REGISTER_ERROR      RETCODE = 4005
	NO_TOKEN_ERROR      RETCODE = 4006
	PROJ_TOTAL_ERROR    RETCODE = 4007
	PROJ_LIST_ERROR     RETCODE = 4008
	PROJ_DETAIL_ERROR   RETCODE = 4009
	PROJ_ADD_ERROR      RETCODE = 4010
	PROJ_EDIT_ERROR     RETCODE = 4011
	PROJ_DEL_ERROR      RETCODE = 4012
	CAPTCHA_ERROR       RETCODE = 4013
	CAPTCHA_TOKEN_ERROR RETCODE = 4014
	REDIS_ERROR         RETCODE = 4015
	SEND_EMAIL_ERROR    RETCODE = 4016
	READ_FAILED_ERROR   RETCODE = 4017
	EMAIL_CODE_ERROR    RETCODE = 4018
	FILE_UPLOAD_H_ERROR RETCODE = 4019
	FILE_UPLOAD_ERROR   RETCODE = 4020
	TOKEN_FORMAT_ERROR  RETCODE = 4021
	GET_USER_INFO_ERROR RETCODE = 4022
	AVATAR_UPDATE_ERROR RETCODE = 4023
	USER_AVATAR_ERROR   RETCODE = 4024
	GET_AUTH_INFO_ERROR RETCODE = 4025
	UPDATE_USER_ERROR   RETCODE = 4026
	GROUP_LIST_ERROR    RETCODE = 4027
	GROUP_ADD_ERROR     RETCODE = 4028
	GROUP_EDIT_ERROR    RETCODE = 4029
	GROUP_DELETE_ERROR  RETCODE = 4030
)

// 错误码和错误信息的关系
var (
	RETCODE_MSG = map[RETCODE]string{
		2000: "成功",
		4000: "必要参数确实或参数错误",
		4001: "数据库错误",
		4002: "token错误",
		4003: "token过期",
		4004: "token生成失败",
		4005: "用户注册失败",
		4006: "没有token",
		4007: "获取项目总数失败",
		4008: "获取项目列表失败",
		4009: "获取项目详情失败",
		4010: "项目添加失败",
		4011: "项目编辑失败",
		4012: "项目删除失败",
		4013: "验证码有误",
		4014: "图片验证码过期",
		4015: "redis错误",
		4016: "操作太频繁，1分钟以后在试",
		4017: "读文件出错",
		4018: "邮箱验证码错误或已过期",
		4019: "上传文件的头部参数出错",
		4020: "文件保存失败",
		4021: "token格式错误",
		4022: "获取用户信息失败",
		4023: "更新头像失败",
		4024: "获取用户头像失败",
		4025: "获取用户AUTH应用信息失败",
		4026: "更新用户信息失败",
		4027: "获取分组列表失败",
		4028: "新增分组列表失败",
		4029: "编辑分组列表失败",
		4030: "删除分组列表失败",
	}
)
