package request

type BuyersLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
