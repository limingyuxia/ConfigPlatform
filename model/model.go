package model

type GetProjectListReq struct {
	ID          int64  `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`                                // 项目名称
	Department  string `json:"department" form:"department"`                    // 项目所属部门
	Admin       string `json:"admin,omitempty" form:"admin"`                    // 管理员
	ProjectUser string `json:"project_user" form:"project_user"`                // 项目创建者
	DevelopUser string `json:"develop_user" form:"develop_user"`                // 项目的开发人员
	PageIndex   int    `json:"page_index" form:"page_index" binding:"required"` // 页码
	PageSize    int    `json:"page_size"  form:"page_size"  binding:"required"` // 页大小
}

type ProjectInfo struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`                  // 项目名称
	Description string `json:"description,omitempty"` // 项目描述
	Department  string `json:"department,omitempty"`  // 项目所属部门
	Admin       string `json:"admin,omitempty"`       // 管理员
	ProjectUser string `json:"project_user"`          // 项目创建者
	DevelopUser string `json:"develop_user"`          // 项目的开发人员
	CreateTime  string `json:"create_time"`           // 项目创建时间
	UpdateTime  string `json:"update_time"`           // 删除时间
}

type GetProjectListResp struct {
	Total       int64
	ProjectList []*ProjectInfo
}
