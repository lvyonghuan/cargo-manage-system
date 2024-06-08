package model

import (
	"sync"

	"gorm.io/gorm"
)

// Customer 顾客模型
// 0表示非会员
type Customer struct {
	CustomerID       int `gorm:"primary_key;auto_increment"`
	Name             string
	MembershipStatus int
	ContactInfo      string
}

type Order struct {
	OrderID    string `gorm:"primary_key"`
	StoreID    int
	CustomerID int    //无会员指向0
	OrderDate  string `gorm:"type:datetime"`
	TotalPrice float64

	TX  *gorm.DB   `gorm:"-"`
	Mux sync.Mutex `gorm:"-"`
}

type OrderItem struct {
	OrderItemID int `gorm:"primary_key"`
	OrderID     string
	UPCCode     string
	Quantity    int
	Price       float64 `gorm:"-"`
}
