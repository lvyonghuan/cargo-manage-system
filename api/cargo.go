package api

import (
	"cargo-manage-system/model"
	"cargo-manage-system/src"
	"cargo-manage-system/util/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func addProduct(c *gin.Context) {
	//获取商品信息
	UPCCode := c.PostForm("UPCCode")
	name := c.PostForm("name")
	priceString := c.PostForm("price")
	brandIDString := c.PostForm("brand")
	typeIDString := c.PostForm("type")

	price, err := strconv.ParseFloat(priceString, 64)
	if err != nil {
		response.AddProductResponse(c, false, "价格格式错误")
		return
	}

	brandID, err := strconv.Atoi(brandIDString)
	if err != nil {
		response.AddProductResponse(c, false, "品牌ID格式错误")
		return
	}

	typeID, err := strconv.Atoi(typeIDString)
	if err != nil {
		response.AddProductResponse(c, false, "类型ID格式错误")
		return
	}

	//生成商品
	var product model.Product
	product.UPCCode = UPCCode
	product.Name = name
	product.Price = price
	product.BrandID = brandID
	product.TypeID = typeID

	//将商品写入数据库
	err = src.AddProduct(product)
	if err != nil {
		response.AddProductResponse(c, false, err.Error())
		return
	}

	response.AddProductResponse(c, true, "添加商品成功")
}

func addBrand(c *gin.Context) {
	//获取品牌信息
	name := c.PostForm("name")

	//生成品牌
	var brand model.Brand
	brand.BrandName = name

	//将品牌写入数据库
	err := src.AddBrand(&brand)
	if err != nil {
		response.AddBrandResponse(c, false, err.Error(), 0)
		return
	}

	response.AddBrandResponse(c, true, "添加品牌成功", brand.BrandID)
}

func addProductType(c *gin.Context) {
	//获取类型信息
	name := c.PostForm("name")
	parentTypeIDString := c.PostForm("parentTypeID")

	var (
		parentTypeID *int
		err          error
	)

	if parentTypeIDString != "" {
		tempID, err := strconv.Atoi(parentTypeIDString)
		if err != nil {
			response.AddProductTypeResponse(c, false, "父类型ID格式错误", 0)
			return
		}
		parentTypeID = &tempID
	}

	//生成类型
	var productType model.ProductType
	productType.TypeName = name
	productType.ParentTypeID = parentTypeID

	//将类型写入数据库
	err = src.AddProductType(&productType)
	if err != nil {
		response.AddProductTypeResponse(c, false, err.Error(), 0)
		return
	}

	response.AddProductTypeResponse(c, true, "添加类型成功", productType.TypeID)
}

func addVendor(c *gin.Context) {
	//获取供应商信息
	Name := c.PostForm("name")

	//生成供应商
	var vendor model.Vendor
	vendor.VendorName = Name

	//将供应商写入数据库
	err := src.AddVendor(&vendor)
	if err != nil {
		response.AddVendorResponse(c, false, err.Error(), 0)
		return
	}

	response.AddVendorResponse(c, true, "添加供应商成功", vendor.VendorID)
}

// 进货
func addCargoStock(c *gin.Context) {
	//生成进货记录，获取商品UPCCode、供应商ID、进货数量、总价、商店ID
	upcCode := c.PostForm("UPCCode")
	vendorIDString := c.PostForm("vendorID")
	purchaseAmountString := c.PostForm("purchaseAmount")
	totalPriceString := c.PostForm("totalPrice")
	storeIDString := c.PostForm("storeID")

	//将供应商ID、进货数量、商店ID转换为int,将总价转换为float64
	vendorID, err := strconv.Atoi(vendorIDString)
	if err != nil {
		response.AddCargoStockResponse(c, false, "供应商ID格式错误")
		return
	}

	purchaseAmount, err := strconv.Atoi(purchaseAmountString)
	if err != nil {
		response.AddCargoStockResponse(c, false, "进货数量格式错误")
		return
	}

	storeID, err := strconv.Atoi(storeIDString)
	if err != nil {
		response.AddCargoStockResponse(c, false, "商店ID格式错误")
		return
	}

	totalPrice, err := strconv.ParseFloat(totalPriceString, 64)
	if err != nil {
		response.AddCargoStockResponse(c, false, "总价格式错误")
		return
	}

	//生成进货记录
	var record model.RestockingRecord
	record.UPCCode = upcCode
	record.VendorID = vendorID
	record.PurchaseAmount = purchaseAmount
	record.TotalPrice = totalPrice
	record.StoreID = storeID

	//将进货记录写入数据库
	err = src.AddCargoStock(record)
	if err != nil {
		response.AddCargoStockResponse(c, false, err.Error())
		return
	}

	response.AddCargoStockResponse(c, true, "进货成功")
}
