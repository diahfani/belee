package transactions

import (
	"belee/models/buyer"
	"belee/models/paymentMethod"
	"belee/models/products"
	"belee/models/warung"
	"time"
)

type TransactionResponse struct {
	Id           int                           `gorm:"primaryKey" json:"id"`
	Buyer        *buyer.Buyers                 `json:"buyer"`
	Warung       *warung.Warungs               `json:"warung"`
	Barang       *products.Products            `json:"barang"`
	PaymentID    int                           `json:"paymentId"`
	Payment      *paymentMethod.PaymentMethods `json:"payment"`
	ProductsName int                           `json:"productsName"`
	TotalQty     int                           `json:"totalqty"`
	TotalPrice   float32                       `json:"totalprice"`
	CreatedAt    time.Time                     `json:"createdAt"`
	Status       string                        `json:"status"`
}
