package model

import (
	"fmt"
	"gin01/src/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.Index(path, "login") < 0 {
			tokenString := c.GetHeader("Token")
			_, claims, err := utils.ParseToken(tokenString)
			if err != nil {
				c.JSON(400, gin.H{"message": "登陆过期，请重新登陆!"})
				return
			}
			fmt.Println(claims)

		}
		c.Next()
	}
}
