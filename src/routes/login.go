package routes

import (
	"gin01/src/service"
	"github.com/gin-gonic/gin"
)

func LoginRoute(e *gin.Engine) {
	e.POST("/login", service.Login)
	e.POST("/getuserinfo", service.Getuserinfo)
	e.POST("/testUrl", service.TestUrl)
	e.POST("/addUser", service.AddUser)
	e.POST("/editUser", service.EditUser)
	e.POST("/delUser", service.DelUser)
	e.POST("/selectUser", service.SelectUser)
}
