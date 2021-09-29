package paymentMethod

import (
	"belee/business/paymentMethod"
	"time"
)

type PaymentMethod struct {
	Id        int       `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func FromDomain(domain *paymentMethod.Domain) *PaymentMethod {
	return &PaymentMethod{
		Id:   domain.Id,
		Name: domain.Name,
	}
}

func (rec *PaymentMethod) ToDomain() paymentMethod.Domain {
	return paymentMethod.Domain{
		Id:        rec.Id,
		Name:      rec.Name,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
