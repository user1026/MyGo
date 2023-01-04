package routes

import (
	"fmt"
	"gin01/src/DataBase"
	"gin01/src/utils"
	"github.com/gin-gonic/gin"
)

func UserRouter(e *gin.Engine) {
	user := e.Group("/user")
	{
		user.POST("/add", AddUser)
		user.POST("/delete", DeleteUser)
		user.POST("/edit", EditUser)
		user.POST("/select", SelectUser)
		user.POST("/selectById", SelectUserById)
	}
}

func AddUser(c *gin.Context) {
	var userInfo DataBase.UserInfo
	if getUserInfoErr := c.ShouldBindJSON(&userInfo); getUserInfoErr != nil {
		fmt.Println(userInfo)
		c.JSON(400, gin.H{"message": "缺乏必要参数"})
		return
	}
	res := serv.AddUser(userInfo)
	if res != true {
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
func SelectUserById(c *gin.Context) {
	var user DataBase.UserInfo
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.FailWithMsg("未获取到uid", c)
		return
	}

	userinfo, err := serv.SelectUserById(user.Uid)
	if err != nil {
		utils.FailWithMsg("查询失败", c)
		return
	}
	utils.OKWithData(map[string]interface{}{"userinfo": userinfo}, c)
}
