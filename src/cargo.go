package src

import (
	"cargo-manage-system/dao"
	"cargo-manage-system/model"
)

func AddCargoStock(cargo model.RestockingRecord) error {
	//插入进货记录
	return dao.AddCargoStock(cargo)
}

func AddProduct(product model.Product) error {
	//插入商品
	return dao.AddProduct(product)
}

func AddBrand(brand *model.Brand) error {
	//插入品牌
	return dao.AddBrand(brand)
}

func AddProductType(t *model.ProductType) error {
	//插入类型
	return dao.AddProductType(t)
}

func AddVendor(vendor *model.Vendor) error {
	//插入供应商
	return dao.AddVendor(vendor)
}

func GetProductPrice(productID string) (float64, error) {
	//获取商品价格
	return dao.GetProductPrice(productID)
}
