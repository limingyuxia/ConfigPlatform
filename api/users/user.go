package users

import (
	"ConfigPlatform/api"
	"ConfigPlatform/conf/mysql"
	"context"
	"errors"
	"log"
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

func UserRegisted(ctx context.Context, userName, passWord string) (int64, error) {

	insertSql := `insert into user(username, password) VALUES(?, ?)`

	result, err := mysql.Conn.ExecContext(ctx, insertSql, userName, passWord)
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
