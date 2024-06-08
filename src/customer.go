package src

import (
	"cargo-manage-system/dao"
	"cargo-manage-system/model"
	"cargo-manage-system/util"
	"errors"
	"time"
)

// 保持上下文
var ctx = make(map[string]*model.Order)

func AddMember(member *model.Customer) error {
	//插入会员
	return dao.AddMember(member)
}

func StartPurchase(storeID, customerID int) (string, error) {
	uuid := util.GenerateUUID()
	var order model.Order
	order.OrderID = uuid
	order.StoreID = storeID
	order.CustomerID = customerID
	order.TotalPrice = 0
	order.OrderDate = time.Now().Format(time.RFC3339)

	tx, err := dao.StartPurchase(&order)
	if err != nil {
		return "", err
	}

	order.TX = tx

	ctx[uuid] = &order

	return uuid, nil
}

func AddItemsToOrder(orderID string, product *model.OrderItem) error {
	order, ok := ctx[orderID]
	if !ok {
		return errors.New("order not found")
	}

	//查找商店库存，大于0才可交易
	stock, err := dao.QueryCargoStock(order.StoreID, product.UPCCode)
	if err != nil {
		return err
	} else if stock < product.Quantity {
		return errors.New("库存不足")
	}

	order.Mux.Lock()
	order.TotalPrice += product.Price * float64(product.Quantity)

	ctx[orderID] = order
	order.Mux.Unlock()

	return dao.AddProductToOrder(order.TX, product)
}

func FinishPurchase(orderID string) (float64, error) {
	order, ok := ctx[orderID]
	if !ok {
		return 0, errors.New("order not found")
	}

	err := dao.FinishOrder(order)
	if err != nil {
		return 0, err
	}

	err = order.TX.Commit().Error
	if err != nil {
		return 0, err
	}

	price := order.TotalPrice

	delete(ctx, orderID)

	return price, nil
}
