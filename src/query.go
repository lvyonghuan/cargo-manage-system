package src

import (
	"cargo-manage-system/dao"
	"cargo-manage-system/model"
	"errors"
)

func BrandQuery(name string) (id int, err error) {
	return dao.BrandQueryByName(name)
}

func TypeQuery(name string) (id int, err error) {
	return dao.TypeQueryByName(name)
}

func VendorQuery(name string) (id int, err error) {
	return dao.VendorQueryByName(name)
}

func StoreQuery(name string) (store model.Store, err error) {
	return dao.StoreQueryByName(name)
}

func UPCQuery(name string) (product model.Product, err error) {
	return dao.ProductQueryByName(name)
}

func GetStores() ([]model.Store, error) {
	return dao.GetStores()
}

func GetProductsByStoreID(storeID int) ([]model.ResponseStock, error) {
	products, err := dao.GetProductsByStoreID(storeID)
	if err != nil {
		return nil, err
	}

	var responses []model.ResponseStock
	for _, product := range products {
		stock, newErr := dao.GetStockByUPCAndStoreID(product.UPCCode, storeID)
		if newErr != nil {
			if err != nil {
				err = errors.New(err.Error() + "\n" + newErr.Error())
			} else {
				err = newErr
			}
			continue
		}

		var response model.ResponseStock
		response.ProductName = product.Name
		response.Inventory = stock.InventoryAmount
		response.Price = product.Price

		brand, newErr := dao.GetBrandByBrandID(product.BrandID)
		if newErr != nil {
			if err != nil {
				err = errors.New(err.Error() + "\n" + newErr.Error())
			} else {
				err = newErr
			}
			continue
		}

		response.BrandName = brand.BrandName

		productType, newErr := dao.GetTypeByTypeID(product.TypeID)
		if newErr != nil {
			if err != nil {
				err = errors.New(err.Error() + "\n" + newErr.Error())
			} else {
				err = newErr
			}
			continue
		}

		response.ProductType = productType.TypeName

		responses = append(responses, response)
	}

	return responses, err
}
