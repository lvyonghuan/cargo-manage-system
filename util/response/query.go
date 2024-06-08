package response

import (
	"cargo-manage-system/model"

	"github.com/gin-gonic/gin"
)

func BrandResponse(c *gin.Context, state bool, msg string, id int) {
	c.JSON(200, gin.H{
		"state": state,
		"msg":   msg,
		"id":    id,
	})
}

func TypeResponse(c *gin.Context, state bool, msg string, id int) {
	c.JSON(200, gin.H{
		"state": state,
		"msg":   msg,
		"id":    id,
	})
}

func VendorResponse(c *gin.Context, state bool, msg string, id int) {
	c.JSON(200, gin.H{
		"state": state,
		"msg":   msg,
		"id":    id,
	})
}

func StoreResponse(c *gin.Context, state bool, msg string, id int) {
	c.JSON(200, gin.H{
		"state": state,
		"msg":   msg,
		"id":    id,
	})
}

func UPCResponse(c *gin.Context, state bool, msg string, upc string) {
	c.JSON(200, gin.H{
		"state": state,
		"msg":   msg,
		"upc":   upc,
	})
}

func GetStoresResponse(c *gin.Context, state bool, msg string, stores []model.Store) {
	c.JSON(200, gin.H{
		"state":  state,
		"msg":    msg,
		"stores": stores,
	})
}

func GetProductsByStoreIDResponse(c *gin.Context, state bool, msg string, products []model.ResponseStock) {
	c.JSON(200, gin.H{
		"state":    state,
		"msg":      msg,
		"products": products,
	})
}
