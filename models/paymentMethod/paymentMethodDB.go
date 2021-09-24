package paymentMethod

import "time"

type PaymentMethods struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	PaymentName string    `json:"paymentname"`
	CreatedAt   time.Time `json:"createdat"`
}
