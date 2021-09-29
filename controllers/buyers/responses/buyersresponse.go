package responses

import (
	"final_project/belee/business/buyers"
	"time"
)

type BuyersResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Age       string    `json:"age"`
	NoHp      string    `json:"nohp"`
	Dob       string    `json:"dob"`
	Address   string    `json:"address"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain buyers.Domain) BuyersResponse {
	return BuyersResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Age:       domain.Age,
		NoHp:      domain.NoHp,
		Dob:       domain.Dob,
		Address:   domain.Address,
		Email:     domain.Email,
		Token:     domain.Token,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
