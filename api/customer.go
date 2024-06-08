package api

import (
	"cargo-manage-system/model"
	"cargo-manage-system/src"
	"cargo-manage-system/util/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func newMembership(c *gin.Context) {
	name := c.PostForm("name")
	statusString := c.PostForm("status")
	phone := c.PostForm("phone")

	status, err := strconv.Atoi(statusString)
	if err != nil {
		response.NewMembershipResponse(c, false, "状态格式错误", 0)
		return
	}

	//生成会员
	var membership model.Customer
	membership.Name = name
	membership.MembershipStatus = status
	membership.ContactInfo = phone

	//将会员写入数据库
	err = src.AddMember(&membership)
	if err != nil {
		response.NewMembershipResponse(c, false, err.Error(), 0)
		return
	}

	response.NewMembershipResponse(c, true, "添加会员成功", membership.CustomerID)
}

func StartPurchase(c *gin.Context) {
	StoreIDString := c.PostForm("storeID")
	CustomerIDString := c.PostForm("customerID")

	StoreID, err := strconv.Atoi(StoreIDString)
	if err != nil {
		response.StartPurchaseResponse(c, false, err.Error(), "")
		return
	}

	var CustomerID int

	if CustomerIDString != "" {
		CustomerID, err = strconv.Atoi(CustomerIDString)
		if err != nil {
			response.StartPurchaseResponse(c, false, err.Error(), "")
			return
		}
	} else {
		CustomerID = 0
	}

	uuid, err := src.StartPurchase(StoreID, CustomerID)
	if err != nil {
		response.StartPurchaseResponse(c, false, err.Error(), "")
		return
	}

	response.StartPurchaseResponse(c, true, "创建订单成功", uuid)
}

func AddItem(c *gin.Context) {
	orderID := c.PostForm("orderID")
	UPCCode := c.PostForm("UPCCode")
	quantityString := c.PostForm("quantity")

	quantity, err := strconv.Atoi(quantityString)
	if err != nil {
		response.AddItemResponse(c, false, "数量格式错误")
		return
	}

	itemPrice, err := src.GetProductPrice(UPCCode)
	if err != nil {
		response.AddItemResponse(c, false, err.Error())
		return
	}

	//生成商品
	var product model.OrderItem
	product.OrderID = orderID
	product.UPCCode = UPCCode
	product.Quantity = quantity
	product.Price = itemPrice

	//将商品写入数据库
	err = src.AddItemsToOrder(orderID, &product)
	if err != nil {
		response.AddItemResponse(c, false, err.Error())
		return
	}

	response.AddItemResponse(c, true, "订单商品添加")
}

func FinishPurchase(c *gin.Context) {
	orderID := c.PostForm("orderID")

	totalPrice, err := src.FinishPurchase(orderID)
	if err != nil {
		response.FinishPurchaseResponse(c, false, err.Error(), 0)
		return
	}

	response.FinishPurchaseResponse(c, true, "订单完成", totalPrice)
}
