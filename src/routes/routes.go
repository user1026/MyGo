package routes

import (
	"gin01/src/utils"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()
	r.Use(utils.Allow())
	r.Use(utils.CheckToken())
	LoginRoute(r)
	return r
}
