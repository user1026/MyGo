package model

import (
	"fmt"
	"gin01/src/DataBase"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type Claims struct {
	UserId     string
	ExpiryTime int
	jwt.StandardClaims
}

var jwtkey = []byte("Token")

func createToken(form *DataBase.UserInfo) string {
	expireTime := time.Now().Add(time.Hour * 10)
	claims := Claims{
		UserId: form.Uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1", // 签名颁发者
			Subject:   "userToken", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	// fmt.Println(token)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err)
		return " "
	}
	return tokenString
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}
func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.Index(path, "login") < 0 {
			tokenString := c.GetHeader("Token")
			_, claims, err := ParseToken(tokenString)
			if err != nil {
				c.JSON(400, gin.H{"message": "登陆过期，请重新登陆!"})
				return
			}
			fmt.Println(claims)

		}
		c.Next()
	}
}
