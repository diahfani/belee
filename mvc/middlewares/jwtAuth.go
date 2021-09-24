package middlewares

import (
	"belee/constants"
	"errors"
	"fmt"

	// "go/constant"

	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtClaims struct {
	BuyersId int `json:"uyersId"`
	jwt.StandardClaims
}

func CreateToken(id int) (string, error) {
	claims := JwtClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 1).Unix(),
		},
	}
	fmt.Println(claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(constants.SECRET_JWT))

	if err != nil {
		return "", err
	}

	return t, nil
}

func GetClaims(c echo.Context) (int, error) {
	buyers := c.Get("buyers")
	if buyers != nil {
		buyersJwt := buyers.(*jwt.Token)
		if buyersJwt.Valid {
			claims := buyersJwt.Claims.(jwt.MapClaims)
			buyersId := claims["buyersId"].(float64)
			return int(buyersId), nil
		}
	}
	return 0, errors.New("Failed Create JWT")
}
