package buyer

import (
	"time"
)

type Buyers struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name" gorm:"not null"`
	Age       string    `json:"age" gorm:"not null"`
	NoHp      string    `json:"nohp" gorm:"not null"`
	Dob       string    `json:"dob" gorm:"not null"`
	Address   string    `json:"address" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
