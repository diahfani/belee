package owners

import "time"

type BuyersRegist struct {
	Name     string    `json:"name"`
	Age      string    `json:"age"`
	NoHp     int       `json:"nohp"`
	Dob      time.Time `json:"dob"`
	Address  string    `json:"address"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
