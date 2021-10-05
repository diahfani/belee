package transactions

import (
	"time"

	"gorm.io/gorm"
)

type Transactions struct {
	Id      int `gorm:"primaryKey" json:"id"`
	BuyerID int `json:"buyerId" gorm:"not null"`
	// Buyer        buyer.Buyers                 `json:"buyer" gorm:"foreignKey:BuyerID"`
	WarungID int `json:"warungId" gorm:"not null"`
	// Warung       warung.Warungs               `json:"warung" gorm:"foreignKey:WarungID"`
	BarangID int `json:"barangId" gorm:"not null"`
	// Barang       products.Products            `json:"barang" gorm:"foreignKey:BarangID"`
	PaymentID int `json:"paymentId" gorm:"not null"`
	// Payment      paymentMethod.PaymentMethods `json:"payment" gorm:"foreignKey:PaymentID"`
	ProductsName string    `json:"productsname" gorm:"not null"`
	TotalQty     int       `json:"totalqty" gorm:"not null"`
	TotalPrice   float32   `json:"totalprice" gorm:"not null"`
	CreatedAt    time.Time `json:"createdAt"`
	DeletedAt    gorm.DeletedAt
	// Status       string    `json:"status"`
	// DetailTransaction []DetailTransaction
}
