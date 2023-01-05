package routes

import (
	"gin01/src/middleware"
	"gin01/src/service"
	"github.com/gin-gonic/gin"
)

var serv = service.ServiceApi

func Routes() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Allow())

	//r.Use(model.CheckToken())

	LoginRoute(r)
	UserRouter(r)
	return r
}
