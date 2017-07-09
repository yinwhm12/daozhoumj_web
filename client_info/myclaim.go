package client_info

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"daozhoumj/client_info/tokenBean"
)

type MyCustomClaims struct {
	UserName	string	`json:"user_name,omitempty"`
	jwt.StandardClaims
}

func CreateToken(name string)(tokenStr string, err error)  {
	expireToken := time.Now().Add(time.Hour * 24).Unix()

	key := []byte("daozhoumj@yin")

	claims := MyCustomClaims{
		name,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer: "youxibi",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr , err = token.SignedString(key)
	return
	
}

func ValidateToken(tokenString string)(int)  {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("daozhoumj@yin"), nil
	})

	if token.Valid {
		//fmt.Println("You look nice today")
		return tokenBean.TOKEN_OK
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			//fmt.Println("That's not even a token")
			return tokenBean.TOKEN_ERROR
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			//fmt.Println("Timing is everything")
			return tokenBean.TOKEN_OVERTIME
		} else {
			//fmt.Println("Couldn't handle this token:", err)
			return tokenBean.TOKEN_BADHANDLE
		}
	} else {
		//fmt.Println("Couldn't handle this token:", err)
		return tokenBean.TOKEN_BADHANDLE
	}
}
