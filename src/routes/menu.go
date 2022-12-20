package routes

import "github.com/gin-gonic/gin"

func Menu(e *gin.Engine) {
	e.POST("/getMenu", getMenu)
	e.POST("/addMenu", addMenu)
	e.POST("editMenu", editMenu)
	e.POST("/deleteMenu", deleteMenu)
}

func deleteMenu(c *gin.Context) {

}

func editMenu(c *gin.Context) {

}

func addMenu(c *gin.Context) {

}

func getMenu(c *gin.Context) {

}
