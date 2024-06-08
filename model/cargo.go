package model

type RestockingRecord struct {
	UPCCode        string
	VendorID       int
	PurchaseAmount int
	TotalPrice     float64
	StoreID        int

	RestockingDate string `gorm:"type:date"`
}

type Product struct {
	UPCCode string `gorm:"primary_key"`
	Name    string
	Price   float64
	BrandID int
	TypeID  int
}

type Brand struct {
	BrandID   int `gorm:"primary_key type:AUTO_INCREMENT"`
	BrandName string
}

type ProductType struct {
	TypeID       int `gorm:"primary_key type:AUTO_INCREMENT"`
	TypeName     string
	ParentTypeID *int
}

type Vendor struct {
	VendorID   int `gorm:"primary_key type:AUTO_INCREMENT"`
	VendorName string
}
