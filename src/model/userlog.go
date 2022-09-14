package model

import "github.com/gin-gonic/gin"

func UserLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
