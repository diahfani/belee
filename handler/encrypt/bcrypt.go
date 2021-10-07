package encrypt

import (
	"golang.org/x/crypto/bcrypt"
)

type HashPassword struct {
	Password string
}

func Hash(secret string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.MinCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// func Hash(password string) ([]byte, error) {
// 	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	// result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
// 	// if err != nil {
// 	// 	return "", err
// 	// }
// 	// return string(result), nil
// }

// func VerifyPassword(hashedPassword, password string) bool {
// 	hash := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// 	return hash == nil
// 	// return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// }

// func BeforeSave() error {
// 	var buyer buyer.Buyers
// 	hashedPassword, err := Hash(buyer.Password)
// 	if err != nil {
// 		return err
// 	}

// 	buyer.Password = string(hashedPassword)
// 	return nil
// }
