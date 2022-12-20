package routes

import (
	"gin01/src/DataBase"
	"gin01/src/service"
	"github.com/gin-gonic/gin"
)

func UserRouter(e *gin.Engine) {
	e.POST("/addUser", AddUser)
	e.POST("/deleteUser", DeleteUser)
	e.POST("/editUser", EditUser)
	e.POST("/selectUser", SelectUser)
}
func AddUser(c *gin.Context) {
	var userInfo DataBase.UserInfo

	if getUserInfoErr := c.ShouldBindJSON(&userInfo); getUserInfoErr != nil {
		c.JSON(400, gin.H{"message": "缺乏必要参数"})
		return
	}
	_, addErr := service.AddUser(userInfo)
	if addErr != nil {
		c.JSON(500, gin.H{"message": "添加失败"})
		return
	} else {
		c.JSON(200, gin.H{"message": "添加成功"})
	}
}
func EditUser(c *gin.Context) {

}
func DeleteUser(c *gin.Context) {

}
func SelectUser(c *gin.Context) {

}
