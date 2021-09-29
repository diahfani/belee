package request

import "final_project/belee/business/buyers"

type BuyersLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *BuyersLogin) ToDomain() *buyers.Domain {
	return &buyers.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}
