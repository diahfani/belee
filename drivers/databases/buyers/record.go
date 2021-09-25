package buyers

import (
	"final_project/belee/belee/business/buyers"
	"time"
)

//khusus bagian database
//butuh method untuk convert ke object domain
type Buyers struct {
	Id        int    `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Age       string `gorm:"not null"`
	NoHp      string `gorm:"not null"`
	Dob       string `gorm:"not null"`
	Address   string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (buyer *Buyers) ToDomain() buyers.Domain {
	return buyers.Domain{
		Id:        buyer.Id,
		Name:      buyer.Name,
		Age:       buyer.Age,
		NoHp:      buyer.NoHp,
		Dob:       buyer.Dob,
		Address:   buyer.Address,
		Email:     buyer.Email,
		Password:  buyer.Password,
		CreatedAt: buyer.CreatedAt,
		UpdatedAt: buyer.UpdatedAt,
	}
}

func FromDomain(domain buyers.Domain) Buyers {
	return Buyers{
		Id:        domain.Id,
		Name:      domain.Name,
		Age:       domain.Age,
		NoHp:      domain.NoHp,
		Dob:       domain.Dob,
		Address:   domain.Address,
		Email:     domain.Email,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
