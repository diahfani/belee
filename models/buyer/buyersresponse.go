package buyer

import "time"

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

// type AuthenticationBuyer struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }
