package response

import "github.com/gin-gonic/gin"

func NewMembershipResponse(c *gin.Context, state bool, msg string, id int) {
	c.JSON(200, gin.H{
		"state": state,
		"msg":   msg,
		"id":    id,
	})
}

func StartPurchaseResponse(c *gin.Context, state bool, msg string, uuid string) {
	c.JSON(200, gin.H{
		"state": state,
		"msg":   msg,
		"uuid":  uuid,
	})
}

func AddItemResponse(c *gin.Context, state bool, msg string) {
	c.JSON(200, gin.H{
		"state": state,
		"msg":   msg,
	})
}

func FinishPurchaseResponse(c *gin.Context, state bool, msg string, price float64) {
	c.JSON(200, gin.H{
		"state": state,
		"msg":   msg,
		"price": price,
	})
}
