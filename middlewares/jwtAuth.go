package middlewares

import (
	"belee/constant"
	"errors"

	// "go/constant"

	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtClaims struct {
	Id int `json:"Id"`
	jwt.StandardClaims
}

func GenerateTokenBuyersJWT(id int) (string, error) {

	claims := JwtClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 1).Unix(),
		},
	}
	// fmt.Println(claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(constant.SECRET_JWT_BUYERS))

	if err != nil {
		return "", err
	}

	return t, nil
}

func GenerateTokenOwnersJWT(id int) (string, error) {
	claims := JwtClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 2).Unix(),
		},
	}
	// fmt.Println(claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(constant.SECRET_JWT_OWNERS))

	if err != nil {
		return "", err
	}

	return t, nil
}

func GetClaimsBuyers(c echo.Context) (int, error) {
	buyers := c.Get("buyers")
	if buyers != nil {
		buyersJwt := buyers.(*jwt.Token)
		if buyersJwt.Valid {
			claims := buyersJwt.Claims.(jwt.MapClaims)
			buyersId := claims["buyersId"].(float64)
			return int(buyersId), nil
		}
	}
	return 0, errors.New("Failed claims JWT")
}

func GetClaimsOwners(c echo.Context) (int, error) {
	owners := c.Get("owners")
	if owners != nil {
		ownersJwt := owners.(*jwt.Token)
		if ownersJwt.Valid {
			claims := ownersJwt.Claims.(jwt.MapClaims)
			ownersId := claims["ownersId"].(float64)
			return int(ownersId), nil
		}
	}
	return 0, errors.New("Failed claims JWT")
}

// func TokenValid(r *http.Request) error {
// 	tokenString := ExtractToken(r)
// 	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header)
// 		}
// 		return []byte(os.Getenv("SECRET_JWT")), nil
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		Pretty(claims)
// 	}
// 	return nil
// }

// func ExtractToken(r *http.Request) string {
// 	keys := r.URL.Query()
// 	token := keys.Get("token")
// 	if token != "" {
// 		return token
// 	}

// 	bearerToken := r.Header.Get("Authorization")
// 	if len(strings.Split(bearerToken, " ")) == 2 {
// 		return strings.Split(bearerToken, " ")[1]
// 	}

// 	return ""

// }

// func ExtractTokenID(r *http.Request) (uint32, error) {
// 	tokenString := ExtractToken(r)
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(os.Getenv("SECRET_JWT")), nil
// 	})
// 	if err != nil {
// 		return 0, err
// 	}
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if ok && token.Valid {
// 		buyerid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["buyersId"]), 10, 32)
// 		if err != nil {
// 			return 0, err
// 		}
// 		return uint32(buyerid), nil
// 	}
// 	return 0, nil

// }

// func Pretty(data interface{}) {
// 	b, err := json.MarshalIndent(data, "", " ")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	fmt.Println(string(b))
// }
