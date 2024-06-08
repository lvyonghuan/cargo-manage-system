package api

import (
	"cargo-manage-system/model"
	"cargo-manage-system/src"
	"cargo-manage-system/util/response"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var admin model.Admin
	admin.Username = username
	admin.Password = password

	err := src.Login(admin)
	if err != nil {
		response.LoginState(c, false, err.Error())
		return
	}

	response.LoginState(c, true, "登录成功")
}
