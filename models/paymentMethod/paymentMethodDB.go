package paymentmethod

type PaymentMethods struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	PaymentName string `json:"paymentname"`
}

type AddPayment struct {
	PaymentName string `json:"paymentname"`
}
