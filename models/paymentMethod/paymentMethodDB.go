package paymentMethod

import "time"

type PaymentMethods struct {
	Id             int    `json:"id" gorm:"primaryKey"`
	Name           string `json:"name" gorm:"unique;not null"`
	TransactionsId int
	CreatedAt      time.Time `json:"createdat"`
	UpdatedAt      time.Time
}
