package model

type GetProjectListReq struct {
	ID          int    `json:"id" form:"id"`                                    // 项目id
	Name        string `json:"name" form:"name"`                                // 项目名称
	Department  string `json:"department" form:"department"`                    // 项目所属部门
	Admin       string `json:"admin" form:"admin"`                              // 管理员
	ProjectUser string `json:"project_user" form:"project_user"`                // 项目创建者
	DevelopUser string `json:"develop_user" form:"develop_user"`                // 项目的开发人员
	PageIndex   int    `json:"page_index" form:"page_index" binding:"required"` // 页码
	PageSize    int    `json:"page_size"  form:"page_size"  binding:"required"` // 页大小
}

type ProjectInfo struct {
	ID          int    `json:"id"`                     // 项目id
	Name        string `json:"name"`                   // 项目名称
	Description string `json:"description,omitempty"`  // 项目描述
	Department  string `json:"department,omitempty"`   // 项目所属部门
	Admin       string `json:"admin,omitempty"`        // 管理员
	ProjectUser string `json:"project_user"`           // 项目创建者
	DevelopUser string `json:"develop_user,omitempty"` // 项目的开发人员
	CreateTime  string `json:"create_time"`            // 项目创建时间
	UpdateTime  string `json:"update_time"`            // 删除时间
}

type GetProjectListResp struct {
	Total       int64
	ProjectList []*ProjectInfo
}

type GetProjectDetailReq struct {
	ID   int    `json:"id" form:"id" binding:"required"` // 项目id
	Name string `json:"name" form:"name"`                // 项目名称
}

/*
func (p *ProjectInfo) GetMemberByJsonTag(columns []string) []interface{} {
	var Member []interface{}

	for _, column := range columns {
		switch column {
		case "id":
			Member = append(Member, &p.ID)
		case "name":
			Member = append(Member, &p.Name)
		case "description":
			Member = append(Member, &p.Description)
		case "department":
			Member = append(Member, &p.Department)
		case "admin":
			Member = append(Member, &p.Admin)
		case "project_user":
			Member = append(Member, &p.ProjectUser)
		case "develop_user":
			Member = append(Member, &p.DevelopUser)
		case "create_time":
			Member = append(Member, &p.CreateTime)
		case "update_time":
			Member = append(Member, &p.UpdateTime)
		}
	}

	return Member
}
*/
