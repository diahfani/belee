package paymentMethod

import "time"

type PaymentMethods struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"unique;not null"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time
}
