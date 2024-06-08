package response

import "github.com/gin-gonic/gin"

// LoginState 返回登录状态
func LoginState(c *gin.Context, state bool, msg string) {
	c.JSON(200, gin.H{
		"state": state,
		"msg":   msg,
	})
}
