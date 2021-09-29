package paymentMethod

import (
	"belee/business/paymentMethod"
	"context"

	"gorm.io/gorm"
)

type MysqlPaymentRepository struct {
	Conn *gorm.DB
}

func NewMysqlPaymentRepo(conn *gorm.DB) paymentMethod.Repository {
	return &MysqlPaymentRepository{
		Conn: conn,
	}
}

func (pr *MysqlPaymentRepository) Find(ctx context.Context) ([]paymentMethod.Domain, error) {
	rec := []PaymentMethod{}

	err := pr.Conn.Find(&rec).Error
	if err != nil {
		return []paymentMethod.Domain{}, err
	}

	payment := []paymentMethod.Domain{}
	for _, value := range rec {
		payment = append(payment, value.ToDomain())
	}

	return payment, nil
}

func (pr *MysqlPaymentRepository) Add(ctx context.Context, pDomain *paymentMethod.Domain) (paymentMethod.Domain, error) {
	rec := FromDomain(pDomain)
	result := pr.Conn.Create(&rec)
	if result.Error != nil {
		return paymentMethod.Domain{}, result.Error
	}
	return rec.ToDomain(), nil
}

func (pr *MysqlPaymentRepository) FindById(id int) (paymentMethod.Domain, error) {
	rec := PaymentMethod{}
	if err := pr.Conn.Where("id = ?", id).First(&rec).Error; err != nil {
		return paymentMethod.Domain{}, err
	}
	return rec.ToDomain(), nil
}
