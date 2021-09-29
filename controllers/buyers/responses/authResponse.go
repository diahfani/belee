package responses

import "final_project/belee/business/buyers"

type AuthResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func GetAuthResp(buyerDomain buyers.Domain, token string) AuthResponse {
	return AuthResponse{
		Email:    buyerDomain.Email,
		Password: buyerDomain.Password,
		Token:    buyerDomain.Token,
	}
}
