package dao

import (
	"cargo-manage-system/model"
	"time"
)

func AddCargoStock(cargo model.RestockingRecord) error {
	t := time.Unix(time.Now().Unix(), 0)
	cargo.RestockingDate = t.Format("2006-01-02")

	//插入进货记录
	return db.Table("restocking_record").Create(&cargo).Error
}

func AddProduct(product model.Product) error {
	//插入商品
	return db.Table("products").Create(&product).Error
}

func AddBrand(brand *model.Brand) error {
	//插入品牌
	return db.Table("brands").Create(&brand).Error
}

func AddProductType(t *model.ProductType) error {
	//插入类型
	return db.Table("product_types").Create(&t).Error
}

func AddVendor(vendor *model.Vendor) error {
	//插入供应商
	return db.Table("vendors").Create(&vendor).Error
}

func GetProductPrice(productID string) (float64, error) {
	//获取商品价格
	var product model.Product
	err := db.Table("products").Where("upc_code = ?", productID).First(&product).Error
	if err != nil {
		return 0, err
	}

	return product.Price, nil
}

// QueryCargoStock 查询对应门店对应货物的数量,返回数量和错误
func QueryCargoStock(storeID int, upcCode string) (int, error) {
	var stock model.Stock
	err := db.Table("product_store").Where("store_id = ? AND upc_code = ?", storeID, upcCode).First(&stock).Error
	if err != nil {
		return 0, err
	}

	return stock.InventoryAmount, nil
}
