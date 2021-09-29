package request

import "final_project/belee/business/owners"

type OwnersLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *OwnersLogin) ToDomain() *owners.Domain {
	return &owners.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}
