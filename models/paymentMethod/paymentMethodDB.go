package paymentMethod

import "time"

type PaymentMethods struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time
}
