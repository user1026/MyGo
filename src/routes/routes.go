package routes

import (
	"gin01/src/middleware"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Allow())

	//r.Use(model.CheckToken())
	LoginRoute(r)
	UserRouter(r)
	return r
}
