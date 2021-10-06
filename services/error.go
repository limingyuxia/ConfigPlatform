package services

type RETCODE int32

// 错误码从4000开始
const (
	SUCCESS        RETCODE = 2000
	PARAMS_ERROR   RETCODE = 4000
	DB_ERROR       RETCODE = 4001
	NOAUTH_ERROR   RETCODE = 4002
	EXPIRE_ERROR   RETCODE = 4003
	TOKEN_ERROR    RETCODE = 4004
	REGISTER_ERROR RETCODE = 4005
)

// 错误码和错误信息的关系
var (
	RETCODE_name = map[RETCODE]string{
		2000: "成功",
		4000: "参数错误",
		4001: "数据库错误",
		4002: "未鉴权",
		4003: "token过期",
		4004: "token生成失败",
		4005: "用户注册失败",
	}
)
