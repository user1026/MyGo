package routes

import (
	"gin01/src/service"
	"github.com/gin-gonic/gin"
)

func LoginRoute(e *gin.Engine) {
	e.POST("/login", service.Login)
	e.POST("/getuserinfo", service.Getuserinfo)
	e.POST("/testUrl", service.TestUrl)
}
