package request

import "belee/business/paymentMethod"

type PaymentMethod struct {
	Name string `json:"name"`
}

func (req *PaymentMethod) ToDomain() *paymentMethod.Domain {
	return &paymentMethod.Domain{
		Name: req.Name,
	}
}
