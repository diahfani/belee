package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

type ConfigJwt struct {
	SecretJwt       string
	ExpiresDuration int
}

func (jwtconf *ConfigJwt) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtconf.SecretJwt),
	}
}

func (jwtconf *ConfigJwt) GenerateToken(BuyerId int) (string, error) {
	claims := JwtCustomClaims{
		BuyerId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtconf.ExpiresDuration))).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(jwtconf.SecretJwt))
	return token, err
}

func GetUser(c echo.Context) *JwtCustomClaims {
	buyer := c.Get("buyers").(*jwt.Token)
	claims := buyer.Claims.(*JwtCustomClaims)
	return claims
}

// func RoleValidation(role string) echo.MiddlewareFunc {
// 	return func(hf echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			claims := GetUser(c)

// 			if claims.Role == role {
// 				return hf(c)
// 			} else {
// 				return controllers.NewErrorResponse(c, business.ErrForbiddenRoles)
// 			}
// 		}
// 	}
// }
