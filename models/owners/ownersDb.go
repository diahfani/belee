package owners

import (
	"time"
)

type Owners struct {
	Id         int       `gorm:"primaryKey" json:"id"`
	Name       string    `json:"name"`
	Age        int       `json:"age"`
	NoHp       int       `json:"nohp"`
	Dob        time.Time `json:"dob"`
	Address    string    `json:"address"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
