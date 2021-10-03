package services

import (
	"ConfigPlatform/conf/mysql"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type ProjectInfo struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`                  // 项目名称
	Description string    `json:"description,omitempty"` // 项目描述
	Department  string    `json:"department,omitempty"`  // 项目所属部门
	Admin       string    `json:"admin,omitempty"`       // 管理员
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
	CTime       string    `json:"c_time"`
	UTime       string    `json:"u_time"`
}

type ProjectList struct {
	Total int32          `json:"total"`
	List  []*ProjectInfo `json:"list"`
}

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
		case "create_time":
			Member = append(Member, &p.CreateTime)
		case "update_time":
			Member = append(Member, &p.UpdateTime)
		}
	}

	return Member
}

// @Tags 配置
// @Summary 获取项目列表
// @Accept  json
// @Produce  json
// @Success 200 {object} ProjectList
// @Failure 400 {object} WebResponse
// @Router /config/list [get]
func GetProjectList(c *gin.Context) {

	querySql := `select id, name, description,
		department, admin, create_time, update_time from project`

	projectList := &ProjectList{
		Total: 2,
	}
	rows, err := mysql.GetDb().QueryContext(c, querySql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var projectInfo = new(ProjectInfo)
		columns, _ := rows.Columns()

		Member := projectInfo.GetMemberByJsonTag(columns)
		if err := rows.Scan(Member...); err != nil {
			log.Print(err)
			continue
		}

		projectInfo.CTime = projectInfo.CreateTime.Format("2006-01-02 15:04:05")
		projectInfo.UTime = projectInfo.UpdateTime.Format("2006-01-02 15:04:05")
		projectList.List = append(projectList.List, projectInfo)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	c.JSON(200, projectList)
}
