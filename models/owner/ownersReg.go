package owner

import (
	"belee/models/warung"
	"time"
)

type OwnersRegist struct {
	// Id        int             `gorm:"primaryKey" json:"id"`
	WarungID  *warung.Warungs `json:"warungId,omitempty"`
	Name      string          `json:"name"`
	Age       string          `json:"age"`
	NoHp      string          `json:"nohp"`
	Dob       string          `json:"dob"`
	Address   string          `json:"address"`
	Email     string          `json:"email"`
	Password  string          `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
