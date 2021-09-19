package buyers

import (
	"time"
)

type Buyers struct {
	Id         int       `gorm:"primaryKey" json:"id"`
	Name       string    `json:"name"`
	Age        string    `json:"age"`
	NoHp       string    `json:"nohp"`
	Dob        string    `json:"dob"`
	Address    string    `json:"address"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
