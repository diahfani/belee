package responses

import "final_project/belee/business/owners"

type AuthResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func GetAuthResp(ownerDomain owners.Domain, token string) AuthResponse {
	return AuthResponse{
		Email:    ownerDomain.Email,
		Password: ownerDomain.Password,
		Token:    ownerDomain.Token,
	}
}
