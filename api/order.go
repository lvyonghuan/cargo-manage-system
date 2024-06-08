package api

import (
	"cargo-manage-system/src"
	"cargo-manage-system/util/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func addAutoOrder(c *gin.Context) {
	upcCode := c.PostForm("UPCCode")
	storeIDString := c.PostForm("storeID")
	limitString := c.PostForm("limit")
	restockAmountString := c.PostForm("restockAmount")

	storeID, err := strconv.Atoi(storeIDString)
	if err != nil {
		response.AddOrderResponse(c, false, "storeID参数错误")
		return
	}

	limit, err := strconv.Atoi(limitString)
	if err != nil {
		response.AddOrderResponse(c, false, "limit参数错误")
		return
	}

	restockAmount, err := strconv.Atoi(restockAmountString)
	if err != nil {
		response.AddOrderResponse(c, false, "restockAmount参数错误")
		return
	}

	err = src.AddAutoOrder(upcCode, storeID, limit, restockAmount)
	if err != nil {
		response.AddOrderResponse(c, false, err.Error())
		return
	}

	response.AddOrderResponse(c, true, "添加成功")
}

func stopAutoOrder(c *gin.Context) {
	upcCode := c.PostForm("UPCCode")
	storeIDString := c.PostForm("storeID")

	storeID, err := strconv.Atoi(storeIDString)
	if err != nil {
		response.StopOrderResponse(c, false, "storeID参数错误")
		return
	}

	err = src.StopAutoOrder(upcCode, storeID)
	if err != nil {
		response.StopOrderResponse(c, false, err.Error())
		return
	}

	response.StopOrderResponse(c, true, "停止成功")
}

func getAutoOrderList(c *gin.Context) {
	list, err := src.GetAutoOrderList()
	if err != nil {
		response.OrderListResponse(c, false, err.Error(), nil)
		return
	}

	response.OrderListResponse(c, true, "获取成功", list)
}
