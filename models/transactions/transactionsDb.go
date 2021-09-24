package transactions

import (
	"time"
)

type status string

const (
	lunas    status = "lunas"
	tdkLunas status = "tidak lunas"
)

type Transactions struct {
	Id           int       `gorm:"primaryKey" json:"id"`
	WarungID     int       `json:"warungId"`
	PaymentID    int       `json:"paymentId"`
	ProductsName int       `json:"productsName"`
	TotalQty     int       `json:"totalqty"`
	TotalPrice   int       `json:"totalprice"`
	CreatedAt    time.Time `json:"createdAt"`
	Status       status    `json:"status"`
}
