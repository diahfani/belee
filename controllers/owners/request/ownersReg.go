package request

import "final_project/belee/business/owners"

type OwnersRegist struct {
	Name     string `json:"name"`
	Age      string `json:"age"`
	NoHp     string `json:"nohp"`
	Dob      string `json:"dob"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req OwnersRegist) ToDomain() *owners.Domain {
	return &owners.Domain{
		Name:     req.Name,
		Age:      req.Age,
		NoHp:     req.NoHp,
		Dob:      req.Dob,
		Address:  req.Address,
		Email:    req.Email,
		Password: req.Password,
	}
}
