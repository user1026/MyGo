package utils

import (
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> 8305f918e55a1711e3d120a7f69d55a17b0d9af5
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Allow() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

var jwtkey = []byte("Token")

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.Index(path, "login") < 0 {
<<<<<<< HEAD
			tokenString := c.GetHeader("Token")
			_, claims, err := parseToken(tokenString)
			if err != nil {
				c.JSON(400, gin.H{"message": "登陆过期，请重新登陆!"})
				return
			}
			fmt.Println(claims)
=======
			tokenString := c.GetHeader("Authorization")
			_, _, err := parseToken(tokenString)
			if err != nil {
				c.JSON(400, gin.H{"message": "登陆过期，请重新登陆"})
				return
			}
>>>>>>> 8305f918e55a1711e3d120a7f69d55a17b0d9af5
		}
		c.Next()
	}
}

type Claims struct {
	UserId     string
	ExpiryTime int
	jwt.StandardClaims
}

func parseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}
