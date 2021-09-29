package owners

import (
	"context"
	"final_project/belee/business/owners"

	// "final_project/belee/drivers/databases/buyers"

	"gorm.io/gorm"
)

type MysqlOwnerRepository struct {
	Conn *gorm.DB
}

func NewMysqlOwnerRepository(conn *gorm.DB) owners.Repository {
	return &MysqlOwnerRepository{
		Conn: conn,
	}
}

func (repo *MysqlOwnerRepository) Login(ctx context.Context, email string, password string) (owners.Domain, error) {
	rec := Owners{}
	err := repo.Conn.Where("email = ?", email).First(&rec).Error
	if err != nil {
		return owners.Domain{}, err
	}
	// result := repo.Conn.First(&buyer, "email = ? AND password = ?", email, password)

	// if result.Error != nil {
	// 	return buyers.Domain{}, result.Error
	// }

	return rec.ToDomain(), nil
}

// func (repo *MysqlBuyerRepository) Register(ctx context.Context, buyerDomain *buyers.Domain) (buyers.Domain, string, error) {
// 	rec := FromDomain(*buyerDomain)

// 	result := repo.Conn.Create(rec)
// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	return nil
// }

func (repo *MysqlOwnerRepository) GetByEmail(ctx context.Context, email string) (owners.Domain, error) {
	rec := Owners{}
	err := repo.Conn.Where("email = ?", email).First(&rec).Error
	if err != nil {
		return owners.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (repo *MysqlOwnerRepository) Store(ctx context.Context, ownerDomain owners.Domain) (owners.Domain, error) {
	rec := FromDomain(ownerDomain)

	result := repo.Conn.Create(&rec)
	if result.Error != nil {
		return owners.Domain{}, result.Error
	}

	return rec.ToDomain(), nil
}
