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

func (repo *MysqlBuyerRepository) Login(email string, password string, ctx context.Context) (buyers.Domain, error) {
	var buyer Buyers
	result := repo.Conn.First(&buyer, "email = ? AND password = ?", email, password)

	if result.Error != nil {
		return buyers.Domain{}, result.Error
	}

	return buyer.ToDomain(), nil
}
