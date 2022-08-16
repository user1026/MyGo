package utils

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserId     string
	ExpiryTime int
	jwt.StandardClaims
}

var jwtkey = []byte("Token")

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}
