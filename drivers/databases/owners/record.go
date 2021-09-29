package owners

import (
	"final_project/belee/business/owners"
	"time"
)

//khusus bagian database
//butuh method untuk convert ke object domain
type Owners struct {
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

func (owner *Owners) ToDomain() owners.Domain {
	return owners.Domain{
		Id:        owner.Id,
		Name:      owner.Name,
		Age:       owner.Age,
		NoHp:      owner.NoHp,
		Dob:       owner.Dob,
		Address:   owner.Address,
		Email:     owner.Email,
		Password:  owner.Password,
		CreatedAt: owner.CreatedAt,
		UpdatedAt: owner.UpdatedAt,
	}
}

func FromDomain(domain owners.Domain) Owners {
	return Owners{
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
