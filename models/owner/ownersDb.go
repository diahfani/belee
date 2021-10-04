package owner

import (
	"belee/models/warung"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Owners struct {
	Id        int              `gorm:"primaryKey" json:"id"`
	WarungID  []warung.Warungs `json:",omitempty" gorm:"foreignKey:OwnersID"`
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

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	// if err != nil {
	// 	return "", err
	// }
	// return string(result), nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (o *Owners) BeforeSave(*gorm.DB) error {
	hashedPassword, err := Hash(o.Password)
	if err != nil {
		return err
	}

	o.Password = string(hashedPassword)
	return nil
}
