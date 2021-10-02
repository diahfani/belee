package encrypt

import (
	"belee/models/buyer"

	"golang.org/x/crypto/bcrypt"
)

type HashPassword struct {
	Password string
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

func BeforeSave() error {
	var buyer buyer.Buyers
	hashedPassword, err := Hash(buyer.Password)
	if err != nil {
		return err
	}

	buyer.Password = string(hashedPassword)
	return nil
}
