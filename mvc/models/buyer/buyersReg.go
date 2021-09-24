package buyer

type BuyersRegist struct {
	Name     string `json:"name"`
	Age      string `json:"age"`
	NoHp     string `json:"nohp"`
	Dob      string `json:"dob"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
