package api

import (
	"cargo-manage-system/src"
	"cargo-manage-system/util/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func brandQuery(c *gin.Context) {
	name := c.PostForm("brandName")

	id, err := src.BrandQuery(name)
	if err != nil {
		response.BrandResponse(c, false, err.Error(), 0)
		return
	}

	response.BrandResponse(c, true, "查询成功", id)
}
func typeQuery(c *gin.Context) {
	name := c.PostForm("typeName")

	id, err := src.TypeQuery(name)
	if err != nil {
		response.TypeResponse(c, false, err.Error(), 0)
		return
	}

	response.TypeResponse(c, true, "查询成功", id)
}

func vendorQuery(c *gin.Context) {
	name := c.PostForm("vendorName")

	id, err := src.VendorQuery(name)
	if err != nil {
		response.VendorResponse(c, false, err.Error(), 0)
		return
	}

	response.VendorResponse(c, true, "查询成功", id)
}

func storeQuery(c *gin.Context) {
	name := c.PostForm("storeName")

	store, err := src.StoreQuery(name)
	if err != nil {
		response.StoreResponse(c, false, err.Error(), 0)
		return
	}

	response.StoreResponse(c, true, "查询成功", store.StoreID)
}

func upcQuery(c *gin.Context) {
	name := c.PostForm("productName")

	product, err := src.UPCQuery(name)
	if err != nil {
		response.UPCResponse(c, false, err.Error(), "")
		return
	}

	response.UPCResponse(c, true, "查询成功", product.UPCCode)
}

func getStores(c *gin.Context) {
	stores, err := src.GetStores()
	if err != nil {
		response.GetStoresResponse(c, false, err.Error(), nil)
		return
	}

	response.GetStoresResponse(c, true, "获取成功", stores)
}

func getProductsByStoreID(c *gin.Context) {
	storeIDString := c.PostForm("storeID")

	storeID, err := strconv.Atoi(storeIDString)
	if err != nil {
		response.GetProductsByStoreIDResponse(c, false, "storeID参数错误", nil)
		return
	}

	products, err := src.GetProductsByStoreID(storeID)
	if err != nil {
		response.GetProductsByStoreIDResponse(c, false, err.Error(), nil)
		return
	}

	response.GetProductsByStoreIDResponse(c, true, "获取成功", products)
}
