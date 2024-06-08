package dao

import "cargo-manage-system/model"

func BrandQueryByName(name string) (id int, err error) {
	var brand model.Brand
	err = db.Where("brand_name = ?", name).First(&brand).Error

	return brand.BrandID, err
}

func TypeQueryByName(name string) (id int, err error) {
	var Type model.ProductType
	err = db.Where("type_name = ?", name).First(&Type).Error

	return Type.TypeID, err
}

func VendorQueryByName(name string) (id int, err error) {
	var vendor model.Vendor
	err = db.Where("vendor_name = ?", name).First(&vendor).Error

	return vendor.VendorID, err
}

func StoreQueryByName(name string) (store model.Store, err error) {
	err = db.Where("store_name = ?", name).First(&store).Error

	return store, err
}

func ProductQueryByName(name string) (product model.Product, err error) {
	err = db.Where("name = ?", name).First(&product).Error

	return product, err
}

func GetStores() ([]model.Store, error) {
	var stores []model.Store
	err := db.Find(&stores).Error
	return stores, err
}

func GetProductsByStoreID(storeID int) ([]model.Product, error) {
	var products []model.Product
	err := db.Table("products").Joins("JOIN product_store ON products.upc_code = product_store.upc_code").Where("product_store.store_id = ?", storeID).Find(&products).Error
	return products, err
}

func GetProductsByUPC(upc string) (product model.Product, err error) {
	err = db.Where("upc_code = ?", upc).First(&product).Error
	return product, err
}

func GetStoreByStoreID(id int) (store model.Store, err error) {
	err = db.Where("store_id = ?", id).First(&store).Error
	return store, err
}

func GetStockByUPCAndStoreID(upc string, storeID int) (stock model.Stock, err error) {
	err = db.Table("product_store").Where("upc_code = ? AND store_id = ?", upc, storeID).First(&stock).Error
	return stock, err
}

func GetTypeByTypeID(id int) (productType model.ProductType, err error) {
	err = db.Where("type_id = ?", id).First(&productType).Error
	return productType, err
}

func GetBrandByBrandID(id int) (brand model.Brand, err error) {
	err = db.Where("brand_id = ?", id).First(&brand).Error
	return brand, err
}
