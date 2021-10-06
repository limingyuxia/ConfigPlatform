package api

import (
	"errors"
)

func GetDbError(err error) error {
	if err.Error() == "Error 1062: Duplicate entry 'superuser' for key 'user.username'" {
		return errors.New("用户名重复")
	} else {
		return errors.New("未知错误")
	}
}

/*
func GoSqlDrive(ctx context.Context) {
	querySql := `select * from project`

	rows, err := mysql.Conn.QueryContext(ctx, querySql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var projectInfo = new(model.ProjectInfo)
		columns, _ := rows.Columns()

		Member := projectInfo.GetMemberByJsonTag(columns)
		if err := rows.Scan(Member...); err != nil {
			log.Print(err)
			continue
		}
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

}
*/
