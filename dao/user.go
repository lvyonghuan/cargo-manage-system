package dao

import (
	"cargo-manage-system/model"
	"errors"

	"gorm.io/gorm"
)

func CheckAdminIsLegal(user model.Admin) error {
	var result model.Admin
	err := db.First(&result, "username = ? ", user.Username).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("用户不存在")
		}
		return err
	}

	if result.Password != user.Password {
		return errors.New("密码错误")
	}

	return nil
}
