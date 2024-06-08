package dao

import (
	"cargo-manage-system/model"

	"gorm.io/gorm"
)

func AddMember(member *model.Customer) error {
	//插入会员
	return db.Table("customers").Create(&member).Error
}

func StartPurchase(order *model.Order) (*gorm.DB, error) {
	tx := db.Begin()
	err := tx.Table("orders").Create(&order).Error

	return tx, err
}

func AddProductToOrder(tx *gorm.DB, item *model.OrderItem) error {
	//插入商品
	return tx.Table("order_items").Create(&item).Error
}

func FinishOrder(order *model.Order) error {
	//更新订单
	return order.TX.Table("orders").Where("order_id = ?", order.OrderID).Updates(&order).Error
}
