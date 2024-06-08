package response

import "github.com/gin-gonic/gin"

func OrderListResponse(c *gin.Context, state bool, msg string, list []string) {
	c.JSON(200, gin.H{
		"state": state,
		"msg":   msg,
		"list":  list,
	})
}

func AddOrderResponse(c *gin.Context, state bool, msg string) {
	c.JSON(200, gin.H{
		"state": state,
		"msg":   msg,
	})
}

func StopOrderResponse(c *gin.Context, state bool, msg string) {
	c.JSON(200, gin.H{
		"state": state,
		"msg":   msg,
	})
}
