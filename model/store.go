package model

type Store struct {
	StoreID   int    `gorm:"primary_key" json:"storeID"`
	StoreName string `json:"storeName"`
	Location  string `json:"location"`
}

type Stock struct {
	UPCCode         string
	StoreID         int
	InventoryAmount int
}

type ResponseStock struct {
	ProductName string  `json:"ProductName"`
	Inventory   int     `json:"Inventory"`
	ProductType string  `json:"ProductType"`
	BrandName   string  `json:"BrandName"`
	Price       float64 `json:"Price"`
}
