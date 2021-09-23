package owners

import (
	"final_project/belee/models/warung"
	"time"
)

type Owners struct {
	Id        int              `gorm:"primaryKey" json:"id"`
	WarungID  []warung.Warungs `gorm:"foreignKey:OwnersID; OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name      string           `json:"name" gorm:"not null"`
	Age       string           `json:"age" gorm:"not null"`
	NoHp      string           `json:"nohp" gorm:"not null"`
	Dob       string           `json:"dob" gorm:"not null"`
	Address   string           `json:"address" gorm:"not null"`
	Email     string           `json:"email" gorm:"unique"`
	Password  string           `json:"password" gorm:"not null"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}
