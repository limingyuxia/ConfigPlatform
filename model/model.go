package model

type GetProjectListReq struct {
	ID          int    `json:"id" form:"id"`                                        // 项目id
	Name        string `json:"name" form:"name"`                                    // 项目名称
	Department  string `json:"department" form:"department"`                        // 项目所属部门
	Admin       string `json:"admin" form:"admin"`                                  // 管理员
	ProjectUser string `json:"project_user" form:"project_user" binding:"required"` // 项目创建者
	DevelopUser string `json:"develop_user" form:"develop_user"`                    // 项目的开发人员
	PageIndex   int    `json:"page_index" form:"page_index" binding:"required"`     // 页码
	PageSize    int    `json:"page_size"  form:"page_size"  binding:"required"`     // 页大小
}

type ProjectInfo struct {
	ID          int      `json:"id"`                     // 项目id
	Name        string   `json:"name"`                   // 项目名称
	Description string   `json:"description,omitempty"`  // 项目描述
	Department  []string `json:"department,omitempty"`   // 项目所属部门
	Admin       []string `json:"admin,omitempty"`        // 管理员
	ProjectUser string   `json:"project_user"`           // 项目创建者
	DevelopUser []string `json:"develop_user,omitempty"` // 项目的开发人员
	CreateTime  string   `json:"create_time"`            // 项目创建时间
	UpdateTime  string   `json:"update_time"`            // 更新时间
}

type GetProjectListResp struct {
	Total       int64          `json:"total"`        // 总数
	ProjectList []*ProjectInfo `json:"project_list"` // 项目列表
}

type GetProjectDetailReq struct {
	ID   int    `json:"id" form:"id" binding:"required"` // 项目id
	Name string `json:"name" form:"name"`                // 项目名称
}

type AddProject struct {
	Name        string   `json:"name" binding:"required"`         // 项目名称
	Description string   `json:"description"`                     // 项目描述
	Department  []string `json:"department"`                      // 项目所属部门
	Admin       []string `json:"admin"`                           // 管理员
	ProjectUser string   `json:"project_user" binding:"required"` // 项目创建者
	DevelopUser []string `json:"develop_user"`                    // 项目的开发人员
}

type EditProject struct {
	ID          int      `json:"id" form:"id" binding:"required"` // 项目id
	Description string   `json:"description"`                     // 项目描述
	Department  []string `json:"department"`                      // 项目所属部门
	Admin       []string `json:"admin"`                           // 管理员
	DevelopUser []string `json:"develop_user"`                    // 项目的开发人员
}

type DeleteProject struct {
	ID          int    `json:"id" form:"id" binding:"required"` // 项目id
	ProjectUser string `json:"project_user" form:"id"`          // 项目创建者
}

type GetCaptchaResp struct {
	CaptchaId string `json:"captcha_id"` // 验证码id
}

type ConfirmCaptcha struct {
	CaptchaId       string `json:"captcha_id" binding:"required"`       // 验证码id
	CaptchaSolution string `json:"captcha_solution" binding:"required"` // 验证码的值
}

type RefeshCaptcha struct {
	Reload int64 `json:"reload"` // 时间戳或随机数
}

type LoginReq struct {
	Username string `json:"username" form:"username" binding:"required"` // 用户名
	Password string `json:"password" form:"password" binding:"required"` // 密码
}

type LoginResp struct {
	Token  string `json:"token"`  // token
	Expire string `json:"expire"` // 过期时间
}

type Register struct {
	Username     string `json:"username" binding:"required"`      // 用户名
	Password     string `json:"password" binding:"required"`      // 密码
	EmailAddress string `json:"email_address" binding:"required"` // 邮箱地址
}

type AuthError struct {
	Code    int    `json:"code"`    // http 状态码
	Message string `json:"message"` // 错误信息
}

// 嵌入到jwt中的user信息
type JwtUser struct {
	UserName string `json:"username"` // jwt token第二部分的值
}

type CaptchaToken struct {
	CaptchaToken string `json:"captcha_token"` // 图片验证码验证成功返回的token
}

type SendEmailCode struct {
	CaptchaToken string `json:"captcha_token" binding:"required"` // 图片验证码验证成功返回的token
	EmailAddress string `json:"email_address" binding:"required"` // 邮箱地址
}

type ConfirmEmail struct {
	EmailCode    string `json:"email_code" binding:"required"`    // 邮箱验证码
	CaptchaToken string `json:"captcha_token" binding:"required"` // 图片验证码验证成功返回的token
}

type JwtToken struct {
	Exp                 int64  `json:"exp"`                   // token 过期时间戳
	OrigIat             int64  `json:"orig_iat"`              // token 开始有效时间戳
	TokenRefeshValidity int64  `json:"token_refesh_validity"` // token 最后有效刷新时间
	UserName            string `json:"username"`              // token 所属的用户名
}

type UserInfo struct {
	UserName       string   `json:"username"`        // 用户名
	Nickname       string   `json:"nickname"`        // 用户昵称
	Gender         int      `json:"gender"`          // 性别 1 男 2 女
	Region         string   `json:"region"`          // 地区
	Phone          string   `json:"phone"`           // 电话号码
	Email          string   `json:"email"`           // 邮箱
	Photo          string   `json:"photo"`           // 头像
	CreateTime     string   `json:"create_time"`     // 注册时间
	CreateProject  []string `json:"create_project"`  // 创建的项目
	ManageProject  []string `json:"manage_project"`  // 管理的项目
	DevelopProject []string `json:"develop_project"` // 参与的项目
}

type Auth2User struct {
	Type     string `json:"type"`      // 第三方应用类型: qq wechat weibo github
	Nickname string `json:"nickname"`  // 第三方应用昵称
	UserName string `json:"username"`  // 随机生成的平台用户名
	UniqueId string `json:"unique_id"` // 第三方平台的唯一id
	Password string `json:"password"`  // 随机生成的平台密码
	Avatar   string `json:"avatar"`    // 第三方应用头像
}
