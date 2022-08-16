package routes

import (
	"gin01/src/model"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()
	r.Use(model.Allow())
	r.Use(model.CheckToken())
	LoginRoute(r)
	return r
}
