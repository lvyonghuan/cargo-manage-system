package src

import (
	"cargo-manage-system/dao"
	"cargo-manage-system/model"
)

func Login(admin model.Admin) error {
	//检查用户名密码是否正确
	err := dao.CheckAdminIsLegal(admin)
	if err != nil {
		return err
	}

	return nil
}
