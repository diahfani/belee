package buyers

import (
	"context"
	"final_project/belee/business/buyers"

	"gorm.io/gorm"
)

type MysqlBuyerRepository struct {
	Conn *gorm.DB
}

func NewMysqlBuyerRepository(conn *gorm.DB) buyers.Repository {
	return &MysqlBuyerRepository{
		Conn: conn,
	}
}

func (repo *MysqlBuyerRepository) Login(ctx context.Context, email string, password string) (buyers.Domain, error) {
	var buyer Buyers
	result := repo.Conn.First(&buyer, "email = ? AND password = ?", email, password)

	if result.Error != nil {
		return buyers.Domain{}, result.Error
	}

	return buyer.ToDomain(), nil
}

// func (repo *MysqlBuyerRepository) Register(ctx context.Context, buyerDomain *buyers.Domain) (buyers.Domain, string, error) {
// 	rec := FromDomain(*buyerDomain)

// 	result := repo.Conn.Create(rec)
// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	return nil
// }

func (repo *MysqlBuyerRepository) GetByEmail(ctx context.Context, email string) (buyers.Domain, error) {
	rec := Buyers{}
	err := repo.Conn.Where("email = ?", email).First(&rec).Error
	if err != nil {
		return buyers.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (repo *MysqlBuyerRepository) Store(ctx context.Context, buyerDomain *buyers.Domain) (buyers.Domain, error) {
	rec := FromDomain(*buyerDomain)

	result := repo.Conn.Create(&rec)
	if result.Error != nil {
		return buyers.Domain{}, result.Error
	}

	return rec.ToDomain(), nil
}
