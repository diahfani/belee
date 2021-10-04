package transactions

import (
	"time"
)

type Transactions struct {
	Id int `gorm:"primaryKey" json:"id"`
	// BuyerID int `json:"buyerId"`
	// Buyer        *buyer.Buyers                 `json:"buyer"`
	// WarungID int `json:"warungId"`
	// Warung       *warung.Warungs               `json:"warung"`
	// BarangID int `json:"barangId"`
	// Barang       *products.Products            `json:"barang"`
	// PaymentID int `json:"paymentId"`
	// Payment      *paymentMethod.PaymentMethods `json:"payment"`
	ProductsName string    `json:"productsname" gorm:"not null"`
	TotalQty     int       `json:"totalqty" gorm:"not null"`
	TotalPrice   float32   `json:"totalprice" gorm:"not null"`
	CreatedAt    time.Time `json:"createdAt"`
	// Status       string    `json:"status"`
	// DetailTransaction []DetailTransaction
}
