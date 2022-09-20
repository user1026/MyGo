package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

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

func GetNowTime() string {
	t := time.Now().Format("2006-01-02 15:04:05")
	return t
}

func GetNowTimeByOptions(HasYMD bool, JoinYMD string, HasHMS bool, JoinHMS string) string {
	str := fmt.Sprintf("2006%s01%s02 15%s04%s05", JoinYMD, JoinYMD, JoinHMS, JoinHMS)
	t := time.Now().Format(str)
	YMD, HMS := "", ""
	if HasYMD {
		YMD = strings.Split(t, " ")[0]
	}
	if HasHMS {
		HMS = strings.Split(t, " ")[1]
	}
	return YMD + HMS
}
