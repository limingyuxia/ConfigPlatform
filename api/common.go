package api

import "errors"

func GetDbError(err error) error {
	if err.Error() == "Error 1062: Duplicate entry 'superuser' for key 'user.username'" {
		return errors.New("用户名重复")
	} else {
		return errors.New("未知错误")
	}
}
