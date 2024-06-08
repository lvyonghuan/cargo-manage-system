package response

import "github.com/gin-gonic/gin"

func AddCargoStockResponse(c *gin.Context, state bool, msg string) {
	c.JSON(200, gin.H{
		"state": state,
		"msg":   msg,
	})
}

func AddProductResponse(c *gin.Context, state bool, msg string) {
	c.JSON(200, gin.H{
		"state": state,
		"msg":   msg,
	})
}

func AddBrandResponse(c *gin.Context, state bool, msg string, brandID int) {
	c.JSON(200, gin.H{
		"state": state,
		"msg":   msg,
		"id":    brandID,
	})
}

func AddProductTypeResponse(c *gin.Context, state bool, msg string, typeID int) {
	c.JSON(200, gin.H{
		"state": state,
		"msg":   msg,
		"id":    typeID,
	})
}

func AddVendorResponse(c *gin.Context, state bool, msg string, vendorID int) {
	c.JSON(200, gin.H{
		"state": state,
		"msg":   msg,
		"id":    vendorID,
	})
}
